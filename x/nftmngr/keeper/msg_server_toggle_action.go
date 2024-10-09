package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

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
