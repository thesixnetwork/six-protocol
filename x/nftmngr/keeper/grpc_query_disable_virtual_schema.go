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

func (k Keeper) DisableVirtualSchemaAll(c context.Context, req *types.QueryAllDisableVirtualSchemaRequest) (*types.QueryAllDisableVirtualSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var disableVirtualSchemas []types.DisableVirtualSchema
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	disableVirtualSchemaStore := prefix.NewStore(store, types.KeyPrefix(types.DisableVirtualSchemaKeyPrefix))

	pageRes, err := query.Paginate(disableVirtualSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var disableVirtualSchema types.DisableVirtualSchema
		if err := k.cdc.Unmarshal(value, &disableVirtualSchema); err != nil {
			return err
		}

		disableVirtualSchemas = append(disableVirtualSchemas, disableVirtualSchema)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDisableVirtualSchemaResponse{DisableVirtualSchema: disableVirtualSchemas, Pagination: pageRes}, nil
}

func (k Keeper) DisableVirtualSchema(c context.Context, req *types.QueryGetDisableVirtualSchemaRequest) (*types.QueryGetDisableVirtualSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDisableVirtualSchema(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDisableVirtualSchemaResponse{DisableVirtualSchema: val}, nil
}
