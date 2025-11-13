package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateGroup(goCtx context.Context, msg *types.MsgCreateGroup) (*types.MsgCreateGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, foundAdmin := k.GetAdmin(ctx, SUPER_ADMIN, msg.Creator)
	if !foundAdmin {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "message creator is not super admin")
	}

	// Check if the value already exists
	_, isFound := k.GetGroup(
		ctx,
		msg.Name,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	group := types.Group{
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
	return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "operation not available")
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// _, foundAdmin := k.GetAdmin(ctx, SUPER_ADMIN, msg.Creator)
	// if !foundAdmin {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "message creator is not super admin")
	// }

	// // Check if the value exists
	// valFound, isFound := k.GetGroup(
	// 	ctx,
	// 	msg.Name,
	// )

	// if !isFound {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	// }

	// // Checks if the msg owner is the same as the current owner
	// if msg.Creator != valFound.Owner && !foundAdmin {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	// var group = types.Group{
	// 	Owner: msg.Creator,
	// 	Name:  msg.Name,
	// }

	// k.SetGroup(ctx, group)

	// return &types.MsgUpdateGroupResponse{}, nil
}

func (k msgServer) DeleteGroup(goCtx context.Context, msg *types.MsgDeleteGroup) (*types.MsgDeleteGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, foundAdmin := k.GetAdmin(ctx, SUPER_ADMIN, msg.Creator)
	if !foundAdmin {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "message creator is not super admin")
	}

	// Check if the value exists
	valFound, isFound := k.GetGroup(
		ctx,
		msg.Name,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Creator != valFound.Owner && !foundAdmin {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveGroup(
		ctx,
		msg.Name,
	)

	return &types.MsgDeleteGroupResponse{}, nil
}
