package keeper

import (
	"context"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

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
