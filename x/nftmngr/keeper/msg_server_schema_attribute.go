package keeper

import (
	"context"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

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
