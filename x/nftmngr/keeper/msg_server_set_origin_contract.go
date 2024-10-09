package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) SetOriginContract(goCtx context.Context, msg *types.MsgSetOriginContract) (*types.MsgSetOriginContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	err = k.SetOriginContractKeeper(ctx, msg.Creator, msg.SchemaCode, msg.NewContractAddress)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetOriginContract,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.SchemaCode),
			sdk.NewAttribute(types.AttributeKeySetOriginContract, msg.NewContractAddress),
			sdk.NewAttribute(types.AttributeKeySetOriginContractResult, "success"),
		),
	})

	return &types.MsgSetOriginContractResponse{
		SchemaCode:         msg.SchemaCode,
		NewContractAddress: msg.NewContractAddress,
	}, nil
}
