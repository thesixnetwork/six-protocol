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

func (k Keeper) VirtualSchemaAll(c context.Context, req *types.QueryAllVirtualSchemaRequest) (*types.QueryAllVirtualSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virSchemas []types.VirtualSchema
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	virSchemaStore := prefix.NewStore(store, types.KeyPrefix(types.VirtualSchemaKeyPrefix))

	pageRes, err := query.Paginate(virSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var virSchema types.VirtualSchema
		if err := k.cdc.Unmarshal(value, &virSchema); err != nil {
			return err
		}

		virSchemas = append(virSchemas, virSchema)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVirtualSchemaResponse{VirtualSchema: virSchemas, Pagination: pageRes}, nil
}

func (k Keeper) VirtualSchema(c context.Context, req *types.QueryGetVirtualSchemaRequest) (*types.QueryGetVirtualSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVirtualSchema(
		ctx,
		req.NftSchemaCode,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVirtualSchemaResponse{VirtualSchema: val}, nil
}