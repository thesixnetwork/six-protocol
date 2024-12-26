package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) CreateActionExecutor(goCtx context.Context, msg *types.MsgCreateActionExecutor) (*types.MsgCreateActionExecutorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	_, err = sdk.AccAddressFromBech32(msg.ExecutorAddress)
	if err != nil {
		return nil, err
	}

	err = k.AddActionExecutor(ctx, msg.Creator, msg.NftSchemaCode, msg.ExecutorAddress)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddActionExecutor,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyActionExecutor, msg.ExecutorAddress),
		),
	)

	return &types.MsgCreateActionExecutorResponse{
		NftSchemaCode:   msg.NftSchemaCode,
		ExecutorAddress: msg.ExecutorAddress,
	}, nil
}

func (k msgServer) DeleteActionExecutor(goCtx context.Context, msg *types.MsgDeleteActionExecutor) (*types.MsgDeleteActionExecutorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	_, err = sdk.AccAddressFromBech32(msg.ExecutorAddress)
	if err != nil {
		return nil, err
	}

	err = k.DelActionExecutor(ctx, msg.Creator, msg.NftSchemaCode, msg.ExecutorAddress)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemoveActionExecutor,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyActionExecutor, msg.ExecutorAddress),
		),
	)

	return &types.MsgDeleteActionExecutorResponse{}, nil
}
