package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/x/packages/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowAllPackages(goCtx context.Context, req *types.QueryShowAllPackagesRequest) (*types.QueryShowAllPackagesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var allPackagesInfo []*types.ShowAllPackagesInfoStruct
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get all packages' storages
	allPackagesStorages := k.GetAllPackageVersionsStorage(ctx)

	// go over all the packages' storages
	for _, packageStorage := range allPackagesStorages {
		packageInfoStruct := types.ShowAllPackagesInfoStruct{}

		// get the package index
		packageInfoStruct.Index = packageStorage.GetPackageIndex()

		// get the latest version package
		latestVersionPackage, err := k.GetPackageLatestVersion(ctx, packageStorage.GetPackageIndex())
		if err != nil {
			return nil, utils.LavaError(ctx, ctx.Logger(), "get_package_latest_version", map[string]string{"err": err.Error(), "packageIndex": packageStorage.GetPackageIndex()}, "could not get the latest version of the package")
		}

		// get the name and price of the package
		packageInfoStruct.Name = latestVersionPackage.GetName()
		packageInfoStruct.Price = latestVersionPackage.GetPrice()

		// append the packageInfoStruct to the allPackagesInfo list
		allPackagesInfo = append(allPackagesInfo, &packageInfoStruct)
	}
	_ = ctx

	return &types.QueryShowAllPackagesResponse{PackagesInfo: allPackagesInfo}, nil
}