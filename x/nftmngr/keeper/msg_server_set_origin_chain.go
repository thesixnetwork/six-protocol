package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) SetOriginChain(goCtx context.Context, msg *types.MsgSetOriginChain) (*types.MsgSetOriginChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	err = k.SetOriginChainKeeper(ctx, msg.Creator, msg.SchemaCode, msg.NewOriginChain)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetBaseURI,
			sdk.NewAttribute(types.EventTypeSetOriginChain, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeySetOriginChain, msg.NewOriginChain),
			sdk.NewAttribute(types.AttributeKeySetOriginChainResult, "success"),
		),
	})

	return &types.MsgSetOriginChainResponse{
		SchemaCode:     msg.SchemaCode,
		NewOriginChain: msg.NewOriginChain,
	}, nil
}
