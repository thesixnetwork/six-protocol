package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

func (k msgServer) AddAdminToGroup(goCtx context.Context, msg *types.MsgAddAdminToGroup) (*types.MsgAddAdminToGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	group, foundGroup := k.GetGroup(ctx, msg.Name)
	if !foundGroup {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "group not found")
	}

	_, foundAdmin := k.GetAdmin(ctx, group.Name, msg.Address)
	if foundAdmin {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "admin already existed")
	}

	_, foundSuperAdmin := k.GetAdmin(ctx, SUPER_ADMIN, msg.Creator)
	if group.Owner != msg.Creator && !foundSuperAdmin {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator is not owner or super admin")
	}

	admin := types.Admin{
		Group: group.Name,
		Admin: msg.Address,
	}

	k.SetAdmin(ctx, admin)

	return &types.MsgAddAdminToGroupResponse{}, nil
}
