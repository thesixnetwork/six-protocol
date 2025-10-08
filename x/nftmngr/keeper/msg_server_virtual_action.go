package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// deprecated
func (k msgServer) CreateVirtualAction(goCtx context.Context, msg *types.MsgCreateVirtualAction) (*types.MsgCreateVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	listNewVirtualAction := []*types.VirtualAction{}

	if err := msg.CheckDuplicateAction(); err != nil {
		return nil, err
	}

	// check permission of the creator
	err := k.validateVirtualSchemaPermission(ctx, msg.NftSchemaCode, msg.Creator)
	if err != nil {
		return nil, err
	}

	for _, newAction := range msg.NewActions {
		err := k.AddVirtualActionKeeper(ctx, msg.NftSchemaCode, *newAction)
		if err != nil {
			return nil, err
		}

		listNewVirtualAction = append(listNewVirtualAction, &types.VirtualAction{
			VirtualNftSchemaCode: msg.NftSchemaCode,
			Name:                 newAction.Name,
			Desc:                 newAction.Desc,
			When:                 newAction.When,
			Then:                 newAction.Then,
			Disable:              newAction.Disable,
			AllowedActioner:      newAction.AllowedActioner,
			Params:               newAction.Params,
		})
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddAction,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
		),
	})

	return &types.MsgCreateVirtualActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		VirtualAction: listNewVirtualAction,
	}, nil
}

// deprecated
func (k msgServer) UpdateVirtualAction(goCtx context.Context, msg *types.MsgUpdateVirtualAction) (*types.MsgUpdateVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	listNewVirtualAction := []*types.VirtualAction{}

	if err := msg.CheckDuplicateAction(); err != nil {
		return nil, err
	}

	for _, newAction := range msg.NewActions {
		err := k.UpdateVirtualActionKeeper(ctx, msg.NftSchemaCode, *newAction)
		if err != nil {
			return nil, err
		}

		listNewVirtualAction = append(listNewVirtualAction, &types.VirtualAction{
			VirtualNftSchemaCode: msg.NftSchemaCode,
			Name:                 newAction.Name,
			Desc:                 newAction.Desc,
			When:                 newAction.When,
			Then:                 newAction.Then,
			Disable:              newAction.Disable,
			AllowedActioner:      newAction.AllowedActioner,
			Params:               newAction.Params,
		})
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddAction,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
		),
	})

	return &types.MsgUpdateVirtualActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		VirtualAction: listNewVirtualAction,
	}, nil
}

// deprecated
func (k msgServer) DeleteVirtualAction(goCtx context.Context, msg *types.MsgDeleteVirtualAction) (*types.MsgDeleteVirtualActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	virtualSchema, found := k.GetVirtualSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, errormod.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	isOwner := false
	for _, schemaRegistry := range virtualSchema.Registry {
		srcSchema, found := k.GetNFTSchema(ctx, schemaRegistry.NftSchemaCode)
		if !found {
			return nil, errormod.Wrap(types.ErrSchemaDoesNotExists, schemaRegistry.NftSchemaCode)
		}
		if msg.Creator == srcSchema.Owner {
			isOwner = true
		}
	}

	if !isOwner {
		return nil, errormod.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	// Check if the virtual action already exists
	_, found = k.GetVirtualAction(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)

	if !found {
		return nil, errormod.Wrap(types.ErrActionDoesNotExists, msg.Name)
	}

	k.RemoveVirtualAction(
		ctx,
		msg.NftSchemaCode,
		msg.Name,
	)

	return &types.MsgDeleteVirtualActionResponse{
		Creator: msg.Creator,
		Status:  "success",
	}, nil
}
