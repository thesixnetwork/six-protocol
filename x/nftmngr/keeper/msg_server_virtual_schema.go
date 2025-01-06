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