package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateActionSignerConfig(goCtx context.Context, msg *types.MsgCreateActionSignerConfig) (*types.MsgCreateActionSignerConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate if creator has permission to set fee config
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionAdminSignerConfig, creator)
	if !granted {
		return nil, errormod.Wrap(types.ErrNoPermissionAdminSignerConfig, msg.Creator)
	}

	// Check if the value already exists
	_, isFound := k.GetActionSignerConfig(
		ctx,
		msg.Chain,
	)
	if isFound {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// // validate ContractOwner address
	// if !common.IsHexAddress(msg.ContractOwner) {
	// 	return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "invalid contract owner address")
	// }

	// // validate ContractAddress address
	// if !common.IsHexAddress(msg.ContractAddress) {
	// 	return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "invalid contract address")
	// }

	actionSignerConfig := types.ActionSignerConfig{
		Creator:         msg.Creator,
		Chain:           msg.Chain,
		ContractAddress: msg.ContractAddress,
	}

	k.SetActionSignerConfig(
		ctx,
		actionSignerConfig,
	)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateActionSignerConfig,
			sdk.NewAttribute(types.AttributeKeyChain, msg.Chain),
			sdk.NewAttribute(types.AttributeKeyNewContract, msg.ContractAddress),
		),
	})
	return &types.MsgCreateActionSignerConfigResponse{}, nil
}

func (k msgServer) UpdateActionSignerConfig(goCtx context.Context, msg *types.MsgUpdateActionSignerConfig) (*types.MsgUpdateActionSignerConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate if creator has permission to set fee config
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionAdminSignerConfig, creator)
	if !granted {
		return nil, errormod.Wrap(types.ErrNoPermissionAdminSignerConfig, msg.Creator)
	}
	// Check if the value exists
	valFound, isFound := k.GetActionSignerConfig(
		ctx,
		msg.Chain,
	)
	if !isFound {
		return nil, errormod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errormod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// // validate ContractOwner address
	// if !common.IsHexAddress(msg.ContractOwner) {
	// 	return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "invalid contract owner address")
	// }

	// // validate ContractAddress address
	// if !common.IsHexAddress(msg.ContractAddress) {
	// 	return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "invalid contract address")
	// }

	actionSignerConfig := types.ActionSignerConfig{
		Creator:         msg.Creator,
		Chain:           msg.Chain,
		ContractAddress: msg.ContractAddress,
	}

	k.SetActionSignerConfig(ctx, actionSignerConfig)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetActionSignerConfig,
			sdk.NewAttribute(types.AttributeKeyChain, msg.Chain),
			sdk.NewAttribute(types.AttributeKeyOldContract, valFound.ContractAddress),
			sdk.NewAttribute(types.AttributeKeyNewContract, msg.ContractAddress),
		),
	})
	return &types.MsgUpdateActionSignerConfigResponse{}, nil
}

func (k msgServer) DeleteActionSignerConfig(goCtx context.Context, msg *types.MsgDeleteActionSignerConfig) (*types.MsgDeleteActionSignerConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate if creator has permission to set fee config
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionAdminSignerConfig, creator)
	if !granted {
		return nil, errormod.Wrap(types.ErrNoPermissionAdminSignerConfig, msg.Creator)
	}
	// Check if the value exists
	valFound, isFound := k.GetActionSignerConfig(
		ctx,
		msg.Chain,
	)
	if !isFound {
		return nil, errormod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errormod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveActionSignerConfig(
		ctx,
		msg.Chain,
	)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDeleteActionSignerConfig,
			sdk.NewAttribute(types.AttributeKeyChain, msg.Chain),
		),
	})
	return &types.MsgDeleteActionSignerConfigResponse{}, nil
}
