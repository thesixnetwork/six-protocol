package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

func (k msgServer) CreateGroup(goCtx context.Context, msg *types.MsgCreateGroup) (*types.MsgCreateGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, foundAdmin := k.GetAdmin(ctx, SUPER_ADMIN, msg.Creator)
	if !foundAdmin {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator is not super admin")
	}

	// Check if the value already exists
	_, isFound := k.GetGroup(
		ctx,
		msg.Name,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var group = types.Group{
		Owner: msg.Creator,
		Name:  msg.Name,
	}

	k.SetGroup(
		ctx,
		group,
	)
	return &types.MsgCreateGroupResponse{}, nil
}

func (k msgServer) UpdateGroup(goCtx context.Context, msg *types.MsgUpdateGroup) (*types.MsgUpdateGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	group, isFound := k.GetGroup(
		ctx,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	_, foundAdmin := k.GetAdmin(ctx, SUPER_ADMIN, msg.Creator)

	// Checks if the the msg owner is the same as the current owner or is super admin
	if msg.Creator != group.Owner && !foundAdmin {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator is not super admin")
	}

	var newGroup = types.Group{
		Owner: msg.Creator,
		Name:  msg.Name,
	}

	k.SetGroup(ctx, newGroup)

	return &types.MsgUpdateGroupResponse{}, nil
}

func (k msgServer) DeleteGroup(goCtx context.Context, msg *types.MsgDeleteGroup) (*types.MsgDeleteGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	group, isFound := k.GetGroup(
		ctx,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	_, foundAdmin := k.GetAdmin(ctx, SUPER_ADMIN, msg.Creator)

	// Checks if the the msg owner is the same as the current owner or is super admin
	if msg.Creator != group.Owner && !foundAdmin {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator is not super admin")
	}

	k.RemoveGroup(
		ctx,
		msg.Name,
	)

	return &types.MsgDeleteGroupResponse{}, nil
}
