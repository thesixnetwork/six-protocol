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

func (k Keeper) ActionExecutorAll(c context.Context, req *types.QueryAllActionExecutorRequest) (*types.QueryAllActionExecutorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionExecutors []types.ActionExecutor
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	actionExecutorStore := prefix.NewStore(store, types.KeyPrefix(types.ActionExecutorKeyPrefix))

	pageRes, err := query.Paginate(actionExecutorStore, req.Pagination, func(key []byte, value []byte) error {
		var actionExecutor types.ActionExecutor
		if err := k.cdc.Unmarshal(value, &actionExecutor); err != nil {
			return err
		}

		actionExecutors = append(actionExecutors, actionExecutor)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActionExecutorResponse{ActionExecutor: actionExecutors, Pagination: pageRes}, nil
}

func (k Keeper) ActionExecutor(c context.Context, req *types.QueryGetActionExecutorRequest) (*types.QueryGetActionExecutorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetActionExecutor(
		ctx,
		req.NftSchemaCode,
		req.ExecutorAddress,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetActionExecutorResponse{ActionExecutor: val}, nil
}

func (k Keeper) ExecutorOfSchemaAll(c context.Context, req *types.QueryAllExecutorOfSchemaRequest) (*types.QueryAllExecutorOfSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var executorOfSchemas []types.ExecutorOfSchema
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	executorOfSchemaStore := prefix.NewStore(store, types.KeyPrefix(types.ExecutorOfSchemaKeyPrefix))

	pageRes, err := query.Paginate(executorOfSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var executorOfSchema types.ExecutorOfSchema
		if err := k.cdc.Unmarshal(value, &executorOfSchema); err != nil {
			return err
		}

		executorOfSchemas = append(executorOfSchemas, executorOfSchema)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllExecutorOfSchemaResponse{ExecutorOfSchema: executorOfSchemas, Pagination: pageRes}, nil
}

func (k Keeper) ExecutorOfSchema(c context.Context, req *types.QueryGetExecutorOfSchemaRequest) (*types.QueryGetExecutorOfSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetExecutorOfSchema(
		ctx,
		req.NftSchemaCode,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetExecutorOfSchemaResponse{ExecutorOfSchema: val}, nil
}
