package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) ChangeOrgOwner(goCtx context.Context, msg *types.MsgChangeOrgOwner) (*types.MsgChangeOrgOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	_, err = sdk.AccAddressFromBech32(msg.ToNewOwner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ToNewOwner)
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
