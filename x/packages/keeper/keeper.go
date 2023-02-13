package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/lavanet/lava/x/packages/types"
)

type (
	Keeper struct {
		cdc                codec.BinaryCodec
		storeKey           sdk.StoreKey
		memKey             sdk.StoreKey
		paramstore         paramtypes.Subspace
		epochStorageKeeper types.EpochstorageKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	epochStorageKeeper types.EpochstorageKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:                cdc,
		storeKey:           storeKey,
		memKey:             memKey,
		paramstore:         ps,
		epochStorageKeeper: epochStorageKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// we dont want to do the calculation here too, epochStorage keeper did it
func (k Keeper) IsEpochStart(ctx sdk.Context) (res bool) {
	return k.epochStorageKeeper.GetEpochStart(ctx) == uint64(ctx.BlockHeight())
}