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

func (k Keeper) ActionByRefIdAll(ctx context.Context, req *types.QueryAllActionByRefIdRequest) (*types.QueryAllActionByRefIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var actionByRefIds []types.ActionByRefId

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
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

func (k Keeper) ActionByRefId(ctx context.Context, req *types.QueryGetActionByRefIdRequest) (*types.QueryGetActionByRefIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetActionByRefId(
		ctx,
		req.RefId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetActionByRefIdResponse{ActionByRefId: val}, nil
}
