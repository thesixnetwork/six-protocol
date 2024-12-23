package keeper

import (
	"context"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// TODO:: Feat(VirtualSchema)
func (k msgServer) CreateVirtualAction(goCtx context.Context, msg *types.MsgCreateVirtualAction) (*types.MsgCreateVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	virtualSchema, found := k.GetVirtualSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	isOwner := false
	for _, schemaRegistry := range virtualSchema.Registry {
		srcSchema, found := k.GetNFTSchema(ctx, schemaRegistry.NftSchemaCode)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, schemaRegistry.NftSchemaCode)
		}
		if msg.Creator == srcSchema.Owner {
			isOwner = true
		}
	}

	if !isOwner {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	newVirtualAction, err := base64.StdEncoding.DecodeString(msg.Base64VirtualActionStruct)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	virtualAction := types.VirtualAction{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(newVirtualAction, &virtualAction)
	if err != nil {
		return nil, err
	}

	// Check if the virtual action already exists
	_, found = k.GetVirtualAction(
		ctx,
		msg.NftSchemaCode,
		virtualAction.Name,
	)

	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	err = ValidateVirutualAction(virtualAction.ToAction())
	if err != nil {
		return nil, err
	}

	// TODO:: verify more about action
	// 1. attribute in action is found in registry

	k.SetVirtualAction(
		ctx,
		virtualAction,
	)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddAction,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyAddActionName, virtualAction.Name),
		),
	})

	return &types.MsgCreateVirtualActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		VirtualAction: &virtualAction,
	}, nil
}

// TODO:: Feat(VirtualSchema)
func (k msgServer) UpdateVirtualAction(goCtx context.Context, msg *types.MsgUpdateVirtualAction) (*types.MsgUpdateVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	virtualSchema, found := k.GetVirtualSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	isOwner := false
	for _, schemaRegistry := range virtualSchema.Registry {
		srcSchema, found := k.GetNFTSchema(ctx, schemaRegistry.NftSchemaCode)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, schemaRegistry.NftSchemaCode)
		}
		if msg.Creator == srcSchema.Owner {
			isOwner = true
		}
	}

	if !isOwner {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	toUpdateAction, err := base64.StdEncoding.DecodeString(msg.Base64VirtualActionStruct)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	toUpdateVirtualAction := types.VirtualAction{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(toUpdateAction, &toUpdateVirtualAction)
	if err != nil {
		return nil, err
	}

	// Check if the virtual action already exists
	_, found = k.GetVirtualAction(
		ctx,
		msg.NftSchemaCode,
		toUpdateVirtualAction.Name,
	)

	if !found {
		return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, toUpdateVirtualAction.Name)
	}

	k.SetVirtualAction(ctx, toUpdateVirtualAction)

	return &types.MsgUpdateVirtualActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		VirtualAction: &toUpdateVirtualAction,
	}, nil
}

// TODO:: Feat(VirtualSchema)
func (k msgServer) DeleteVirtualAction(goCtx context.Context, msg *types.MsgDeleteVirtualAction) (*types.MsgDeleteVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	virtualSchema, found := k.GetVirtualSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	isOwner := false
	for _, schemaRegistry := range virtualSchema.Registry {
		srcSchema, found := k.GetNFTSchema(ctx, schemaRegistry.NftSchemaCode)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, schemaRegistry.NftSchemaCode)
		}
		if msg.Creator == srcSchema.Owner {
			isOwner = true
		}
	}

	if !isOwner {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	// Check if the virtual action already exists
	_, found = k.GetVirtualAction(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)

	if !found {
		return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, msg.Name)
	}

	k.RemoveVirtualAction(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)

	return &types.MsgDeleteVirtualActionResponse{
		Creator: msg.Creator,
		Status:  "success",
	}, nil
}
