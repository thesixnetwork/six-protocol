package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) SetMetadataFormat(goCtx context.Context, msg *types.MsgSetMetadataFormat) (*types.MsgSetMetadataFormatResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	k.SetMetadataFormatKeeper(ctx, msg.Creator, msg.SchemaCode, msg.NewFormat)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetMetadataFormat,
			sdk.NewAttribute(types.EventTypeSetMetadataFormat, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NewFormat),
		),
	})

	return &types.MsgSetMetadataFormatResponse{
		SchemaCode: msg.SchemaCode,
		NewFormat:  msg.NewFormat,
	}, nil
}
