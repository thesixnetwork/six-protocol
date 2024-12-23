package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

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
