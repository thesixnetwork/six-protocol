package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) ActionExecutorAll(ctx context.Context, req *types.QueryAllActionExecutorRequest) (*types.QueryAllActionExecutorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionExecutors []types.ActionExecutor

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
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

func (k Keeper) ActionExecutor(ctx context.Context, req *types.QueryGetActionExecutorRequest) (*types.QueryGetActionExecutorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

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

// ExecutorOfSchema implements types.QueryServer.
func (k Keeper) ExecutorOfSchema(ctx context.Context, req *types.QueryGetExecutorOfSchemaRequest) (*types.QueryGetExecutorOfSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetExecutorOfSchema(
		ctx,
		req.NftSchemaCode,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetExecutorOfSchemaResponse{ExecutorOfSchema: val}, nil
}

// ExecutorOfSchemaAll implements types.QueryServer.
func (k Keeper) ExecutorOfSchemaAll(ctx context.Context, req *types.QueryAllExecutorOfSchemaRequest) (*types.QueryAllExecutorOfSchemaResponse, error) {
	var executorOfSchemas []types.ExecutorOfSchema
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

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
