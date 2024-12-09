package keeper

import (
	"context"
	"encoding/base64"

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
