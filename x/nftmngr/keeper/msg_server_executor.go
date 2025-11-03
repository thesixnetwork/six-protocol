package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) AddActionExecutor(ctx context.Context, creator, nftSchemaName, executorAddress string) error {
	schema, foundNftSchema := k.GetNFTSchema(ctx, nftSchemaName)
	_, foundVirtualSchema := k.GetVirtualSchema(ctx, nftSchemaName)

	if !foundNftSchema && !foundVirtualSchema {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if foundNftSchema && creator != schema.Owner {
		return errormod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if foundVirtualSchema && creator != k.GetModuleAddress().String() {
		return errormod.Wrap(sdkerrors.ErrUnauthorized, "Only module account can do this process")
	}

	// Check if the value already exists
	_, isFound := k.GetActionExecutor(
		ctx,
		nftSchemaName,
		executorAddress,
	)

	if isFound {
		return errormod.Wrap(sdkerrors.ErrInvalidRequest, "Action Executor already exists")
	}

	actionExecutor := types.ActionExecutor{
		Creator:         creator,
		NftSchemaCode:   nftSchemaName,
		ExecutorAddress: executorAddress,
	}

	val, found := k.GetExecutorOfSchema(ctx, nftSchemaName)
	if !found {
		val = types.ExecutorOfSchema{
			NftSchemaCode:   nftSchemaName,
			ExecutorAddress: []string{},
		}
	}

	// set executorOfSchema
	val.ExecutorAddress = append(val.ExecutorAddress, executorAddress)

	k.SetExecutorOfSchema(ctx, types.ExecutorOfSchema{
		NftSchemaCode:   val.NftSchemaCode,
		ExecutorAddress: val.ExecutorAddress,
	})

	k.SetActionExecutor(ctx, actionExecutor)

	return nil
}

// RemoveActionExecutor removes a actionExecutor from the store
func (k Keeper) DelActionExecutor(ctx context.Context, creator, nftSchemaName, executorAddress string) error {
	// Retrieve the schema
	schema, foundNftSchema := k.GetNFTSchema(ctx, nftSchemaName)
	_, foundVirtualSchema := k.GetVirtualSchema(ctx, nftSchemaName)

	if !foundNftSchema && !foundVirtualSchema {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if foundNftSchema && creator != schema.Owner {
		return errormod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if foundVirtualSchema && creator != k.GetModuleAddress().String() {
		return errormod.Wrap(sdkerrors.ErrUnauthorized, "Only module account can do this process")
	}

	// Check if the value exists
	_, isFound := k.GetActionExecutor(
		ctx,
		nftSchemaName,
		executorAddress,
	)
	if !isFound {
		return errormod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveActionExecutor(
		ctx,
		nftSchemaName,
		executorAddress,
	)

	val, found := k.GetExecutorOfSchema(ctx, nftSchemaName)
	if !found {
		val = types.ExecutorOfSchema{
			NftSchemaCode:   nftSchemaName,
			ExecutorAddress: []string{},
		}
	}

	// remove executorOfSchema
	for i, executor := range val.ExecutorAddress {
		if executor == executorAddress {
			val.ExecutorAddress = append(val.ExecutorAddress[:i], val.ExecutorAddress[i+1:]...)
			break
		}
	}

	k.SetExecutorOfSchema(ctx, types.ExecutorOfSchema{
		NftSchemaCode:   val.NftSchemaCode,
		ExecutorAddress: val.ExecutorAddress,
	})

	return nil
}

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
