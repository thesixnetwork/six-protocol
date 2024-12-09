package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO:: Feat(VirtualSchema)
func (k Keeper) VirtualActionAll(c context.Context, req *types.QueryAllVirtualActionRequest) (*types.QueryAllVirtualActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virtuals []types.VirtualAction
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	virtualStore := prefix.NewStore(store, types.KeyPrefix(types.VirtualActionKeyPrefix))

	pageRes, err := query.Paginate(virtualStore, req.Pagination, func(key []byte, value []byte) error {
		var virtual types.VirtualAction
		if err := k.cdc.Unmarshal(value, &virtual); err != nil {
			return err
		}

		virtuals = append(virtuals, virtual)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVirtualActionResponse{VirtualAction: virtuals, Pagination: pageRes}, nil
}

// TODO:: Feat(VirtualSchema)
func (k Keeper) VirtualAction(c context.Context, req *types.QueryGetVirtualActionRequest) (*types.QueryGetVirtualActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVirtualAction(
		ctx,
		req.NftSchemaCode,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVirtualActionResponse{VirtualAction: val}, nil
}
