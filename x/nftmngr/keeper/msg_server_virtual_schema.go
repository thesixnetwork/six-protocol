package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) CreateVirtualSchemaProposal(goCtx context.Context, msg *types.MsgCreateVirtualSchemaProposal) (*types.MsgCreateVirtualSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if schema already
	_, found := k.GetNFTSchema(ctx, msg.VirtualNftSchemaCode)
	if found {
		return nil, sdkerrors.Wrap(types.ErrSchemaAlreadyExists, "Schema name already existed")
	}

	// Check if the value already exists
	_, found = k.GetVirtualSchema(
		ctx,
		msg.VirtualNftSchemaCode,
	)

	if found {
		return nil, sdkerrors.Wrap(types.ErrSchemaAlreadyExists, "Schema name already existed")
	}

	isOwner := false
	// check if creator is path of schema registry owner
	for _, registry := range msg.Registry {
		schema, found := k.GetNFTSchema(ctx, registry.NftSchemaCode)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, registry.NftSchemaCode)
		}
		if schema.Owner == msg.Creator {
			isOwner = true
		}
	}

	if !isOwner {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "Only owner of registry schema can create proposal")
	}

	registry := []*types.VirtualSchemaRegistry{}

	for _, reqRegistry := range msg.Registry {
		registry = append(registry, reqRegistry.ConvertRequestToVirtualRegistry())
	}

	lastProposalId := len(k.GetAllVirtualSchemaProposal(ctx))
	proposalId := lastProposalId + 1
	strProposalId := strconv.FormatInt(int64(proposalId), 10)

	submitTime := ctx.BlockHeader().Time
	votingPeriod := k.govKeeper.GetVotingParams(ctx).VotingPeriod

	endTime := submitTime.Add(votingPeriod)

	k.SetVirtualSchemaProposal(ctx, types.VirtualSchemaProposal{
		Id:                strProposalId,
		VirtualSchemaCode: msg.VirtualNftSchemaCode,
		Registry:          registry,
		SubmitTime:        submitTime,
		VotinStartTime:    submitTime,
		VotingEndTime:     endTime,
	})

	k.SetActiveVirtualSchemaProposal(ctx, types.ActiveVirtualSchemaProposal{
		Id: strProposalId,
	})

	return &types.MsgCreateVirtualSchemaResponse{
		Id:                   strProposalId,
		VirtualNftSchemaCode: msg.VirtualNftSchemaCode,
	}, nil
}

// NOTE:: implement only do not use on production
func (k msgServer) DeleteVirtualSchema(goCtx context.Context, msg *types.MsgDeleteVirtualSchema) (*types.MsgDeleteVirtualSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetVirtualSchema(
		ctx,
		msg.VirtualNftSchemaCode,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	var isOwner bool

	for _, registry := range valFound.Registry {
		nftSchema, _ := k.GetNFTSchema(ctx, registry.NftSchemaCode)

		if msg.Creator == nftSchema.Owner {
			isOwner = true
		}
	}

	if !isOwner {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	k.RemoveVirtualSchema(
		ctx,
		msg.VirtualNftSchemaCode,
	)

	return &types.MsgDeleteVirtualSchemaResponse{}, nil
}

func (k msgServer) DisableVirtualSchemaProposal(goCtx context.Context, msg *types.MsgDisableVirtualSchemaProposal) (*types.MsgDisableVirtualSchemaProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentBlock := ctx.BlockHeight()
	proposalEndBlock, err := strconv.ParseInt(msg.ProposalExpiredBlock, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidHeight, msg.ProposalExpiredBlock)
	}

	if currentBlock >= proposalEndBlock {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidHeight, msg.ProposalExpiredBlock)
	}

	virtualSchema, found := k.GetVirtualSchema(ctx, msg.VirtualNftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.VirtualNftSchemaCode)
	}

	if !virtualSchema.Enable {
		return nil, sdkerrors.Wrap(types.ErrSchemaIsDisable, msg.VirtualNftSchemaCode)
	}

	isOwner := false
	// check if creator is path of schema registry owner
	for _, registry := range virtualSchema.Registry {
		schema, found := k.GetNFTSchema(ctx, registry.NftSchemaCode)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, registry.NftSchemaCode)
		}
		if schema.Owner == msg.Creator {
			isOwner = true
		}
	}

	if !isOwner {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	// iterate to assing id of proposal
	allProposal := k.GetAllDisableVirtualSchema(ctx)
	proposalId := len(allProposal) + 1

	k.SetDisableVirtualSchema(ctx, types.DisableVirtualSchema{
		Id:                   strconv.FormatInt(int64(proposalId), 10),
		VirtualNftSchemaCode: msg.VirtualNftSchemaCode,
		ProposalExpiredBlock: msg.ProposalExpiredBlock,
	})

	return &types.MsgDisableVirtualSchemaProposalResponse{}, nil
}
