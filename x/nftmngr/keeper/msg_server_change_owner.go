package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ChangeOrgOwner(goCtx context.Context, msg *types.MsgChangeOrgOwner) (*types.MsgChangeOrgOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	_, err = sdk.AccAddressFromBech32(msg.ToNewOwner)
	if err != nil {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidAddress, msg.ToNewOwner)
	}

	err = k.Keeper.ChangeOrgOwner(ctx, msg.Creator, msg.ToNewOwner, msg.OrgName)
	if err != nil {
		return nil, err
	}

	// emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChangeOrgOwner,
			sdk.NewAttribute(types.AttributeKeyOrgName, msg.OrgName),
			sdk.NewAttribute(types.AttributeKeyOldOwner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyNewOwner, msg.ToNewOwner),
		),
	)

	return &types.MsgChangeOrgOwnerResponse{
		OrgName:  msg.OrgName,
		OldOwner: msg.Creator,
		NewOwner: msg.ToNewOwner,
	}, nil
}

func (k msgServer) ChangeSchemaOwner(goCtx context.Context, msg *types.MsgChangeSchemaOwner) (*types.MsgChangeSchemaOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	_, err = sdk.AccAddressFromBech32(msg.NewOwner)
	if err != nil {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidAddress, msg.NewOwner)
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
