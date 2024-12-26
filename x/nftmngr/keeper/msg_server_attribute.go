package keeper

import (
	"context"
	"encoding/base64"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddAttribute(goCtx context.Context, msg *types.MsgAddAttribute) (*types.MsgAddAttributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	var new_add_attribute types.AttributeDefinition

	input_addribute, err := base64.StdEncoding.DecodeString(msg.Base64NewAttriuteDefenition)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(input_addribute, &new_add_attribute)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	err = k.AddAttributeKeeper(ctx, msg.Creator, msg.Code, new_add_attribute, msg.Location)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.Code),
			sdk.NewAttribute(types.AttributeKeyAddAttributeName, new_add_attribute.Name),
			sdk.NewAttribute(types.AttributeKeyAddAttributeLocation, types.AttributeLocation.String(msg.Location)),
		),
	})

	return &types.MsgAddAttributeResponse{
		Code: msg.Code,
		Name: new_add_attribute.Name,
	}, nil
}

func (k msgServer) ResyncAttributes(goCtx context.Context, msg *types.MsgResyncAttributes) (*types.MsgResyncAttributesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	err = k.ResyncAttibutesKeeper(ctx, msg.Creator, msg.NftSchemaCode, msg.TokenId)
	if err != nil {
		return nil, err
	}

	// Emit Event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeResyncAttributes,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyTokenId, msg.TokenId),
		),
	)
	return &types.MsgResyncAttributesResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}


func (k msgServer) UpdateSchemaAttribute(goCtx context.Context, msg *types.MsgUpdateSchemaAttribute) (*types.MsgUpdateSchemaAttributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	var update_attribute types.AttributeDefinition
	input_addribute, err := base64.StdEncoding.DecodeString(msg.Base64UpdateAttriuteDefenition)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(input_addribute, &update_attribute)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	err = k.UpdateAttributeKeeper(ctx, msg.Creator, msg.NftSchemaCode, update_attribute)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.EventTypeAddAttribute, update_attribute.Name),
		),
	)

	return &types.MsgUpdateSchemaAttributeResponse{
		NftSchemaCode: msg.NftSchemaCode,
		Name:          update_attribute.Name,
	}, nil
}

func (k msgServer) SetAttributeOveriding(goCtx context.Context, msg *types.MsgSetAttributeOveriding) (*types.MsgSetAttributeOveridingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	err = k.Keeper.SetAttributeOveridingKeeper(ctx, msg.Creator, msg.SchemaCode, msg.NewOveridingType)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetAttributeOveriding,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeySetAttributeOverrideResult, "success"),
		),
	})

	return &types.MsgSetAttributeOveridingResponse{}, nil
}

func (k msgServer) ShowAttributes(goCtx context.Context, msg *types.MsgShowAttributes) (*types.MsgShowAttributesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	err = k.ShowAttributeKeeper(ctx, msg.Creator, msg.NftSchemaCode, msg.Show, msg.AttributeNames)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeShowAttribute,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
		),
	)

	return &types.MsgShowAttributesResponse{
		NftSchema: msg.NftSchemaCode,
	}, nil
}
