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
	var (
		registry []*types.VirtualSchemaRegistry
		err      error
	)

	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.ProposalType == types.ProposalType_CREATE {
		registry, err = k.validateCreateVirtualSchemaProposal(ctx, msg.VirtualNftSchemaCode, msg.Registry)
		if err != nil {
			return nil, err
		}
		// **** SCHEMA FEE ****
		feeConfig, found := k.GetNFTFeeConfig(ctx)
		if found {
			// Get Denom
			amount, err := sdk.ParseCoinNormalized(feeConfig.SchemaFee.FeeAmount)
			if err != nil {
				return nil, sdkerrors.Wrap(types.ErrInvalidFeeAmount, err.Error())
			}
			feeBalances, found := k.GetNFTFeeBalance(ctx)
			if !found {
				feeBalances = types.NFTFeeBalance{
					FeeBalances: []string{
						"0" + amount.Denom,
					},
				}
			}

			if len(feeBalances.FeeBalances) > 0 {
				feeBalances.FeeBalances[types.FeeSubject_CREATE_NFT_SCHEMA] = "0" + amount.Denom
			}
			err = k.ProcessFee(ctx, &feeConfig, &feeBalances, types.FeeSubject_CREATE_NFT_SCHEMA, sdk.AccAddress(msg.Creator))
			if err != nil {
				return nil, sdkerrors.Wrap(types.ErrProcessingFee, err.Error())
			}
			// Set Fee Balance
			k.SetNFTFeeBalance(ctx, feeBalances)
		}
	} else {
		registry, err = k.validateUpdateVirtualSchemaProposal(ctx, msg.VirtualNftSchemaCode, msg.Registry)
		if err != nil {
			return nil, err
		}
	}

	err = k.validateOwnerOfRegistry(ctx, msg.Creator, registry)
	if err != nil {
		return nil, err
	}

	actionNameMap := make(map[string]bool)
	// validateAction
	for _, action := range msg.Actions {
		if err := ValidateVirutualAction(action); err != nil {
			return nil, err
		}
		if _, found := actionNameMap[action.Name]; found {
			return nil, sdkerrors.Wrap(types.ErrDuplicateActionName, action.Name)
		}
		actionNameMap[action.Name] = true
	}

	strProposalId := k.getLastVirtualSchemaProposalId(ctx)

	submitTime := ctx.BlockHeader().Time
	votingPeriod := k.govKeeper.GetVotingParams(ctx).VotingPeriod

	endTime := submitTime.Add(votingPeriod)

	k.SetVirtualSchemaProposal(ctx, types.VirtualSchemaProposal{
		Id:           strProposalId,
		ProposalType: msg.ProposalType,
		VirtualSchema: &types.VirtualSchema{
			VirtualNftSchemaCode: msg.VirtualNftSchemaCode,
			Registry:             registry,
			Enable:               msg.Enable,
		},
		Actions:         msg.Actions,
		SubmitTime:      submitTime,
		VotingStartTime: submitTime,
		VotingEndTime:   endTime,
	})

	k.SetActiveVirtualSchemaProposal(ctx, types.ActiveVirtualSchemaProposal{
		Id: strProposalId,
	})

	return &types.MsgProposalVirtualSchemaResponse{
		Id:                   strProposalId,
		VirtualNftSchemaCode: msg.VirtualNftSchemaCode,
		ProposalType:         msg.ProposalType,
	}, nil
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
