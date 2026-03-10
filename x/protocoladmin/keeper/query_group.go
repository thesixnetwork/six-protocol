package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) GroupAll(ctx context.Context, req *types.QueryAllGroupRequest) (*types.QueryAllGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var groups []types.Group

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	groupStore := prefix.NewStore(store, types.KeyPrefix(types.GroupKeyPrefix))

	pageRes, err := query.Paginate(groupStore, req.Pagination, func(key []byte, value []byte) error {
		var group types.Group
		if err := k.cdc.Unmarshal(value, &group); err != nil {
			return err
		}

		groups = append(groups, group)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGroupResponse{Group: groups, Pagination: pageRes}, nil
}

func (k Keeper) Group(ctx context.Context, req *types.QueryGetGroupRequest) (*types.QueryGetGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetGroup(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGroupResponse{Group: val}, nil
}
