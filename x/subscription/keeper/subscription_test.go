package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/relayer/sigs"
	keepertest "github.com/lavanet/lava/testutil/keeper"
	"github.com/lavanet/lava/testutil/nullify"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	"github.com/lavanet/lava/x/subscription/keeper"
	"github.com/lavanet/lava/x/subscription/types"
	"github.com/stretchr/testify/require"
)

func createNSubscription(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Subscription {
	items := make([]types.Subscription, n)
	_, creator := sigs.GenerateFloatingKey()

	for i := range items {
		items[i].Creator = creator.String()
		items[i].Consumer = "consumer-" + strconv.Itoa(i)
		items[i].Block = ctx.BlockHeight()
		items[i].PlanIndex = "testplan"
		items[i].PlanBlock = ctx.BlockHeight()

		keeper.SetSubscription(ctx, items[i])
	}
	return items
}

func TestSubscriptionGet(t *testing.T) {
	keeper, ctx := keepertest.SubscriptionKeeper(t)
	items := createNSubscription(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSubscription(ctx,
			item.Consumer,
		)
		require.True(t, found)

		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestSubscriptionRemove(t *testing.T) {
	keeper, ctx := keepertest.SubscriptionKeeper(t)
	items := createNSubscription(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSubscription(ctx,
			item.Creator,
		)
		_, found := keeper.GetSubscription(ctx,
			item.Creator,
		)
		require.False(t, found)
	}
}

func TestSubscriptionGetAll(t *testing.T) {
	keeper, ctx := keepertest.SubscriptionKeeper(t)
	items := createNSubscription(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSubscription(ctx)),
	)
}

func TestCreateSubscription(t *testing. T) {
	_, keepers, _ctx := keepertest.InitAllKeepers(t)
	ctx := sdk.UnwrapSDKContext(_ctx)

	keeper := keepers.Subscription
	bank := keepers.BankKeeper

	creators := []struct {
		address string
		amount  int64
	}{
		{
			address: "FILL",
			amount:  100000,
		},
		{
			address: "FILL",
			amount:  1,
		},
		{
			address: "invalid creator",
			amount:  0,
		},
	}

	for i := range creators {
		if creators[i].address == "FILL" {
			_, addr := sigs.GenerateFloatingKey()
			creators[i].address = addr.String()
			coins := sdk.NewCoins(sdk.NewCoin(
				epochstoragetypes.TokenDenom,
				sdk.NewInt(creators[i].amount),
			))
			bank.SetBalance(ctx, addr, coins)
		}
	}

	consumers := make([]string, 4)
	for i := range consumers {
		_, addr := sigs.GenerateFloatingKey()
		consumers[i] = addr.String()
	}
	consumers[3] = "invalid consumer"

	template := []struct {
		name      string
		index     string
		creator   int
		consumers []int
		success   bool
	}{
		{
			name:      "create subscriptions",
			index:     "testplan",
			creator:   0,
			consumers: []int{0, 1},
			success:   true,
		},
		{
			name:      "invalid creator",
			index:     "testplan",
			creator:   2,
			consumers: []int{2},
			success:   false,
		},
		{
			name:      "invalid consumer",
			index:     "testplan",
			creator:   0,
			consumers: []int{3},
			success:   false,
		},
		{
			name:      "insufficient funds",
			index:     "testplan",
			creator:   1,
			consumers: []int{2},
			success:   false,
		},
		{
			name:      "invalid plan",
			index:     "",
			creator:   0,
			consumers: []int{2},
			success:   false,
		},
		{
			name:      "unknown plan",
			index:     "no-such-plan",
			creator:   0,
			consumers: []int{2},
			success:   false,
		},
		{
			name:      "double subscription",
			index:     "testplan",
			creator:   0,
			consumers: []int{0},
			success:   false,
		},
	}

	for _, tt := range template {
		for _, consumer := range tt.consumers {
			t.Run(tt.name, func(t *testing.T) {
				sub := types.Subscription{
					Creator:   creators[tt.creator].address,
					Consumer:  consumers[consumer],
					PlanIndex: tt.index,
				}

				err := keeper.CreateSubscription(
					ctx, sub.Creator, sub.Consumer, sub.PlanIndex, sub.IsYearly)
				if tt.success {
					require.Nil(t, err, tt.name)
					_, found := keeper.GetSubscription(ctx, sub.Consumer)
					require.True(t, found, tt.name)
				} else {
					require.NotNil(t, err, tt.name)
				}
			})
		}
	}
}
