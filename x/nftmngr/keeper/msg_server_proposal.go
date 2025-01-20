package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// TODO::
// 1. For proposaltype CREATE when start proposal lock amout of value to module account
// 2. when proposal is rejected, unlock the amount and burn some as penalty. The amount left will be refunded to the creator
// 3. when proposal is accepted, and process fee so on.
func (k msgServer) ProposalVirtualSchema(goCtx context.Context, msg *types.MsgProposalVirtualSchema) (*types.MsgProposalVirtualSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	strProposalId, err := k.ProposalVirtualSchemaKeeper(ctx, msg.Creator, msg.VirtualNftSchemaCode, msg.ProposalType, msg.Registry, msg.Actions, msg.Enable)
	if err != nil {
		return nil, err
	}

	return &types.MsgProposalVirtualSchemaResponse{
		Id:                   strProposalId,
		VirtualNftSchemaCode: msg.VirtualNftSchemaCode,
		ProposalType:         msg.ProposalType,
	}, nil
}

func (k Keeper) ProposalVirtualSchemaKeeper(ctx sdk.Context, creator, virtualNftSchemaCode string, proposalType types.ProposalType, registryReq []*types.VirtualSchemaRegistryRequest, actions []*types.Action, enable bool) (string, error) {
	var (
		registry []*types.VirtualSchemaRegistry
		err      error
	)

	if proposalType == types.ProposalType_CREATE {
		registry, err = k.validateCreateVirtualSchemaProposal(ctx, virtualNftSchemaCode, registryReq)
		if err != nil {
			return "", err
		}
	} else {
		registry, err = k.validateUpdateVirtualSchemaProposal(ctx, virtualNftSchemaCode, registryReq)
		if err != nil {
			return "", err
		}
	}

	err = k.validateOwnerOfRegistry(ctx, creator, registry)
	if err != nil {
		return "", err
	}

	actionNameMap := make(map[string]bool)
	// validateAction
	for _, action := range actions {
		if err := ValidateVirutualAction(action); err != nil {
			return "", err
		}
		if _, found := actionNameMap[action.Name]; found {
			return "", sdkerrors.Wrap(types.ErrDuplicateActionName, action.Name)
		}
		actionNameMap[action.Name] = true
	}

	strProposalId := k.getLastVirtualSchemaProposalId(ctx)

	submitTime := ctx.BlockHeader().Time
	votingPeriod := k.govKeeper.GetVotingParams(ctx).VotingPeriod

	endTime := submitTime.Add(votingPeriod)

	k.SetVirtualSchemaProposal(ctx, types.VirtualSchemaProposal{
		Id:           strProposalId,
		ProposalType: proposalType,
		VirtualSchema: &types.VirtualSchema{
			VirtualNftSchemaCode: virtualNftSchemaCode,
			Registry:             registry,
			Enable:               enable,
		},
		Actions:         actions,
		SubmitTime:      submitTime,
		VotingStartTime: submitTime,
		VotingEndTime:   endTime,
	})

	k.SetActiveVirtualSchemaProposal(ctx, types.ActiveVirtualSchemaProposal{
		Id: strProposalId,
	})

	if proposalType == types.ProposalType_CREATE {
		// lock the amount of value to module account
		feeConfig, found := k.GetNFTFeeConfig(ctx)
		// **** SCHEMA FEE ****
		if found {
			// Get Denom
			amount, err := sdk.ParseCoinNormalized(feeConfig.SchemaFee.FeeAmount)
			if err != nil {
				return "", sdkerrors.Wrap(types.ErrInvalidFeeAmount, err.Error())
			}

			creatorAddress, err := sdk.AccAddressFromBech32(creator)
			if err != nil {
				return "", err
			}

			// Lock the amount
			err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddress, types.ModuleName, sdk.NewCoins(amount))
			if err != nil {
				return "", err
			}

			k.SetLockSchemaFee(ctx, types.LockSchemaFee{
				Id:                strProposalId,
				VirtualSchemaCode: virtualNftSchemaCode,
				Amount:            amount,
				Proposer:          creator,
			})
		}
	}

	return strProposalId, nil
}

func (k Keeper) validateVirtualSchemaPermission(ctx sdk.Context, virtualNftSchemaCode, creator string) error {
	virtualSchema, found := k.GetVirtualSchema(ctx, virtualNftSchemaCode)
	if !found {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, virtualNftSchemaCode)
	}

	for _, reg := range virtualSchema.Registry {
		schema, found := k.GetNFTSchema(ctx, reg.NftSchemaCode)
		if !found {
			return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, reg.NftSchemaCode)
		}
		if schema.Owner != creator {
			return sdkerrors.Wrap(types.ErrUnauthorized, "Only owner of registry schema can create proposal: "+creator)
		}
	}

	return nil
}

func (k Keeper) validateOwnerOfRegistry(ctx sdk.Context, creator string, registry []*types.VirtualSchemaRegistry) error {
	isOwner := false
	// check if creator is part of schema registry owner
	for _, reg := range registry {
		schema, found := k.GetNFTSchema(ctx, reg.NftSchemaCode)
		if !found {
			return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, reg.NftSchemaCode)
		}
		if schema.Owner == creator {
			isOwner = true
		}
	}

	if !isOwner {
		return sdkerrors.Wrap(types.ErrUnauthorized, "Only owner of registry schema can create proposal: "+creator)
	}

	return nil
}

func (k Keeper) validateCreateVirtualSchemaProposal(ctx sdk.Context, virtualNftSchemaCode string, registryReq []*types.VirtualSchemaRegistryRequest) ([]*types.VirtualSchemaRegistry, error) {
	registry := []*types.VirtualSchemaRegistry{}
	// Check if schema already
	_, found := k.GetNFTSchema(ctx, virtualNftSchemaCode)
	if found {
		return nil, sdkerrors.Wrap(types.ErrSchemaAlreadyExists, "Schema name already existed")
	}

	// Check if the value already exists
	_, found = k.GetVirtualSchema(
		ctx,
		virtualNftSchemaCode,
	)

	if found {
		return nil, sdkerrors.Wrap(types.ErrSchemaAlreadyExists, "Schema name already existed")
	}

	err := k.checkDuplicateSchemaCodeWithVirtualSchemaProposal(ctx, virtualNftSchemaCode)
	if err != nil {
		return nil, err
	}

	for _, regis := range registryReq {
		registry = append(registry, regis.ConvertRequestToVirtualRegistry())
	}

	return registry, nil
}

func (k Keeper) validateUpdateVirtualSchemaProposal(ctx sdk.Context, virtualNftSchemaCode string, registryReq []*types.VirtualSchemaRegistryRequest) ([]*types.VirtualSchemaRegistry, error) {
	registry := []*types.VirtualSchemaRegistry{}

	// Check if the value already exists
	_, found := k.GetVirtualSchema(
		ctx,
		virtualNftSchemaCode,
	)

	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaNotFound, "Virtual Schema code not found")
	}

	for _, regis := range registryReq {
		registry = append(registry, regis.ConvertRequestToVirtualRegistry())
	}

	return registry, nil
}

func (k Keeper) getLastVirtualSchemaProposalId(ctx sdk.Context) string {
	lastProposalId := len(k.GetAllVirtualSchemaProposal(ctx))
	proposalId := lastProposalId + 1
	strProposalId := strconv.FormatInt(int64(proposalId), 10)
	return strProposalId
}
