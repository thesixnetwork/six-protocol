package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) AdminAll(ctx context.Context, req *types.QueryAllAdminRequest) (*types.QueryAllAdminResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var admins []types.Admin

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	adminStore := prefix.NewStore(store, types.KeyPrefix(types.AdminKeyPrefix))

	pageRes, err := query.Paginate(adminStore, req.Pagination, func(key []byte, value []byte) error {
		var admin types.Admin
		if err := k.cdc.Unmarshal(value, &admin); err != nil {
			return err
		}

		admins = append(admins, admin)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAdminResponse{Admin: admins, Pagination: pageRes}, nil
}

func (k Keeper) Admin(ctx context.Context, req *types.QueryGetAdminRequest) (*types.QueryGetAdminResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetAdmin(
		ctx,
		req.Group,
		req.Admin,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAdminResponse{Admin: val}, nil
}
