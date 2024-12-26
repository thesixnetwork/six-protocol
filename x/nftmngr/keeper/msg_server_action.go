package keeper

import (
	"context"
	"encoding/base64"
	"strconv"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddAction(goCtx context.Context, msg *types.MsgAddAction) (*types.MsgAddActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// structure for new action
	var new_action types.Action

	// decode base64 string to bytes
	input_action, err := base64.StdEncoding.DecodeString(msg.Base64NewAction)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	// unmarshal bytes to Action structure
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(input_action, &new_action)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	err = k.AddActionKeeper(ctx, msg.Creator, msg.Code, new_action)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddAction,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.Code),
			sdk.NewAttribute(types.AttributeKeyAddActionName, new_action.Name),
			sdk.NewAttribute(types.AttributeKeyAddActionResult, "success"),
		),
	})

	return &types.MsgAddActionResponse{
		Code: msg.GetCode(),
		Name: new_action.Name,
	}, nil
}

func (k msgServer) PerformVirtualAction(goCtx context.Context, msg *types.MsgPerformVirtualAction) (*types.MsgPerformVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Emit events on metadata change
	changeList, err := k.PerformVirtualKeeper(ctx, msg.Creator, msg.NftSchemaName, msg.TokenIdMap, msg.Action, msg.RefId, msg.Parameters)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
			sdk.NewAttribute(types.AttributeKeyRunActionChangeList, string(changeList)),
		),
	)

	return &types.MsgPerformVirtualActionResponse{
		NftSchemaName: msg.NftSchemaName,
	}, nil
}

func (k msgServer) UpdateAction(goCtx context.Context, msg *types.MsgUpdateAction) (*types.MsgUpdateActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	_updateAction, err := base64.StdEncoding.DecodeString(msg.Base64UpdateAction)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	updateAction := types.Action{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(_updateAction, &updateAction)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	err = k.Keeper.UpdateActionKeeper(ctx, msg.Creator, msg.NftSchemaCode, updateAction)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateAction,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyUpdateActionName, updateAction.Name),
		),
	})

	return &types.MsgUpdateActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		Name:          updateAction.Name,
	}, nil
}

func (k msgServer) ToggleAction(goCtx context.Context, msg *types.MsgToggleAction) (*types.MsgToggleActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	err = k.ToggleActionKeeper(ctx, msg.Creator, msg.Code, msg.Action, msg.Status)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeToggleNFTAction,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.Code),
			sdk.NewAttribute(types.AttributeKeyToggleNFTAction, msg.Action),
			sdk.NewAttribute(types.AttributeKeyToggleNFTActionResult, strconv.FormatBool(msg.Status)),
		),
	})

	return &types.MsgToggleActionResponse{
		Code:   msg.Code,
		Name:   msg.Action,
		Status: msg.Status,
	}, nil
}

func (k msgServer) PerformActionByAdmin(goCtx context.Context, msg *types.MsgPerformActionByAdmin) (*types.MsgPerformActionByAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Emit events on metadata change
	changeList, err := k.ActionByAdmin(ctx, msg.Creator, msg.NftSchemaCode, msg.TokenId, msg.Action, msg.RefId, msg.Parameters)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
			sdk.NewAttribute(types.AttributeKeyRunActionChangeList, string(changeList)),
		),
	)

	return &types.MsgPerformActionByAdminResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}
