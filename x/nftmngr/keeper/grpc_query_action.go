package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k Keeper) ActionByRefIdAll(c context.Context, req *types.QueryAllActionByRefIdRequest) (*types.QueryAllActionByRefIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionByRefIds []types.ActionByRefId
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	actionByRefIdStore := prefix.NewStore(store, types.KeyPrefix(types.ActionByRefIdKeyPrefix))

	pageRes, err := query.Paginate(actionByRefIdStore, req.Pagination, func(key []byte, value []byte) error {
		var actionByRefId types.ActionByRefId
		if err := k.cdc.Unmarshal(value, &actionByRefId); err != nil {
			return err
		}

		actionByRefIds = append(actionByRefIds, actionByRefId)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActionByRefIdResponse{ActionByRefId: actionByRefIds, Pagination: pageRes}, nil
}

func (k Keeper) ActionByRefId(c context.Context, req *types.QueryGetActionByRefIdRequest) (*types.QueryGetActionByRefIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetActionByRefId(
		ctx,
		req.RefId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetActionByRefIdResponse{ActionByRefId: val}, nil
}

func (k Keeper) ActionOfSchemaAll(c context.Context, req *types.QueryAllActionOfSchemaRequest) (*types.QueryAllActionOfSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionOfSchemas []types.ActionOfSchema
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	actionOfSchemaStore := prefix.NewStore(store, types.KeyPrefix(types.ActionOfSchemaKeyPrefix))

	pageRes, err := query.Paginate(actionOfSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var actionOfSchema types.ActionOfSchema
		if err := k.cdc.Unmarshal(value, &actionOfSchema); err != nil {
			return err
		}

		actionOfSchemas = append(actionOfSchemas, actionOfSchema)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActionOfSchemaResponse{ActionOfSchema: actionOfSchemas, Pagination: pageRes}, nil
}

func (k Keeper) ActionOfSchema(c context.Context, req *types.QueryGetActionOfSchemaRequest) (*types.QueryGetActionOfSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetActionOfSchema(
		ctx,
		req.NftSchemaCode,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetActionOfSchemaResponse{ActionOfSchema: val}, nil
}
