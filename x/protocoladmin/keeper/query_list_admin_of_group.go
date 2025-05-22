package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

func (k Keeper) ListAdminOfGroup(goCtx context.Context, req *types.QueryListAdminOfGroupRequest) (*types.QueryListAdminOfGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var admins []string
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	adminStore := prefix.NewStore(store, types.KeyPrefix(types.AdminKeyPrefix))

	pageRes, err := query.Paginate(adminStore, req.Pagination, func(key []byte, value []byte) error {
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
