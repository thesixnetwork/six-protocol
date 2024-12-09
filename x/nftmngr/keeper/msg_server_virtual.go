package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// TODO:: Feat(VirtualSchema)
func (k msgServer) CreateVirtualAction(goCtx context.Context, msg *types.MsgCreateVirtualAction) (*types.MsgCreateVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetVirtualAction(
		ctx,
		msg.NftSchemaCode,
		msg.Base64VirtualActionStruct,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var virtual = types.VirtualAction{
		NftSchemaCode:   "",
		Name:            "",
		Desc:            "",
		Disable:         false,
		When:            "",
		Then:            []string{},
		AllowedActioner: 0,
		Params:          []*types.ActionParams{},
	}

	k.SetVirtualAction(
		ctx,
		virtual,
	)
	return &types.MsgCreateVirtualActionResponse{}, nil
}

// TODO:: Feat(VirtualSchema)
func (k msgServer) UpdateVirtualAction(goCtx context.Context, msg *types.MsgUpdateVirtualAction) (*types.MsgUpdateVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetVirtualAction(
		ctx,
		msg.NftSchemaCode,
		msg.Base64VirtualActionStruct,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	_ = valFound.Desc

	// // Checks if the the msg creator is the same as the current owner
	// if msg.Creator != valFound.Creator {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	var virtual = types.VirtualAction{
		NftSchemaCode:   "",
		Name:            "",
		Desc:            "",
		Disable:         false,
		When:            "",
		Then:            []string{},
		AllowedActioner: 0,
		Params:          []*types.ActionParams{},
	}

	k.SetVirtualAction(ctx, virtual)

	return &types.MsgUpdateVirtualActionResponse{}, nil
}

// TODO:: Feat(VirtualSchema)
func (k msgServer) DeleteVirtualAction(goCtx context.Context, msg *types.MsgDeleteVirtualAction) (*types.MsgDeleteVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetVirtualAction(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	_ = valFound.Desc
	// // Checks if the the msg creator is the same as the current owner
	// if msg.Creator != valFound.Creator {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	k.RemoveVirtualAction(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)

	return &types.MsgDeleteVirtualActionResponse{}, nil
}
