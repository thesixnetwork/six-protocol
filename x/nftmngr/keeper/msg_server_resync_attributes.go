package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

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
