package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) ProposalVirtualSchema(goCtx context.Context, msg *types.MsgProposalVirtualSchema) (*types.MsgProposalVirtualSchemaResponse, error) {
	var (
		registry []*types.VirtualSchemaRegistry
		err      error
	)

	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.ProposalType == types.ProposalType_CREATE {
		registry, err = k.validateCreateVirtualSchemaProposal(ctx, msg.VirtualSchema.VirtualNftSchemaCode, msg.VirtualSchema.Registry)
		if err != nil {
			return nil, err
		}
	} else {
		// get virtual schema entity
		existedSchema, found := k.GetVirtualSchema(ctx, msg.VirtualSchema.VirtualNftSchemaCode)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.VirtualSchema.VirtualNftSchemaCode)
		}

		// mark pending to registry
		for _, vsRegistry := range existedSchema.Registry {
			registry = append(registry, &types.VirtualSchemaRegistry{
				NftSchemaCode:    vsRegistry.NftSchemaCode,
				SharedAttributes: vsRegistry.SharedAttributes,
				Decision:         types.RegistryStatus_PENDING,
			})
		}
	}

	err = k.validateOwnerOfRegistry(ctx, msg.Creator, registry)
	if err != nil {
		return nil, err
	}

	strProposalId := k.getLastVirtualSchemaProposalId(ctx)

	submitTime := ctx.BlockHeader().Time
	votingPeriod := k.govKeeper.GetVotingParams(ctx).VotingPeriod

	endTime := submitTime.Add(votingPeriod)

	k.SetVirtualSchemaProposal(ctx, types.VirtualSchemaProposal{
		Id:           strProposalId,
		ProposalType: msg.ProposalType,
		VirtualSchema: &types.VirtualSchema{
			VirtualNftSchemaCode: msg.VirtualSchema.VirtualNftSchemaCode,
			Registry:             registry,
			Enable:               false,
		},
		Actions:         []*types.Action{},
		SubmitTime:      submitTime,
		VotingStartTime: submitTime,
		VotingEndTime:   endTime,
	})

	k.SetActiveVirtualSchemaProposal(ctx, types.ActiveVirtualSchemaProposal{
		Id: strProposalId,
	})

	return &types.MsgProposalVirtualSchemaResponse{
		Id:                   strProposalId,
		VirtualNftSchemaCode: msg.VirtualSchema.VirtualNftSchemaCode,
	}, nil
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
		return sdkerrors.Wrap(types.ErrUnauthorized, "Only owner of registry schema can create proposal")
	}

	return nil
}

func (k Keeper) validateCreateVirtualSchemaProposal(ctx sdk.Context, virtualNftSchemaCode string, registryReq []*types.VirtualSchemaRegistry) ([]*types.VirtualSchemaRegistry, error) {
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

	for _, regis := range registryReq {
		registry = append(registry, regis)
	}

	return registry, nil
}

func (k Keeper) getLastVirtualSchemaProposalId(ctx sdk.Context) string {
	lastProposalId := len(k.GetAllVirtualSchemaProposal(ctx))
	proposalId := lastProposalId + 1
	strProposalId := strconv.FormatInt(int64(proposalId), 10)
	return strProposalId
}
