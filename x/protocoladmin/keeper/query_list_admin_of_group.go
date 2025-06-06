package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) ListAdminOfGroup(goCtx context.Context, req *types.QueryListAdminOfGroupRequest) (*types.QueryListAdminOfGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AdminKeyPrefix))

	var admins []string
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var admin types.Admin
		if err := k.cdc.Unmarshal(value, &admin); err != nil {
			return err
		}

		if admin.Group == req.Group {
			admins = append(admins, admin.Admin)
		}

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListAdminOfGroupResponse{Admin: admins, Pagination: pageRes}, nil
}
