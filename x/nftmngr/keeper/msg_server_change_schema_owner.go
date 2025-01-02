package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) ChangeSchemaOwner(goCtx context.Context, msg *types.MsgChangeSchemaOwner) (*types.MsgChangeSchemaOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	_, err = sdk.AccAddressFromBech32(msg.NewOwner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.NewOwner)
	}

	err = k.Keeper.ChangeSchemaOwner(ctx, msg.Creator, msg.NewOwner, msg.NftSchemaCode)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSchemaOwnerChanged,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyNewOwner, msg.NewOwner),
		),
	})

	return &types.MsgChangeSchemaOwnerResponse{
		NftSchemaCode: msg.NftSchemaCode,
		NewOwner:      msg.NewOwner,
	}, nil
}
