package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// TODO:: Feat(VirtualSchema)
// 1. Check if virtual schema is exist
// 2. Check if virtual schema is enable
// 3. Check if voter(creator) is owner of some src schema
// 4. Check if src schema already vote
func (k msgServer) VoteCreateVirtualSchema(goCtx context.Context, msg *types.MsgVoteCreateVirtualSchema) (*types.MsgVoteCreateVirtualSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	virtualSchemaProposal, found := k.GetVirtualSchemaProposal(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.Id)
	}

	srcSchema, found := k.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	if srcSchema.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
	}

	registryIndex := -1
	for i, registry := range virtualSchemaProposal.Registry {
		if registry.NftSchemaCode == msg.NftSchemaCode {
			registryIndex = i
			// Check if already voted
			if registry.Status != types.RegistryStatus_PENDING {
				return nil, sdkerrors.Wrapf(types.ErrAlreadyVote, "schema %s has already voted", msg.NftSchemaCode)
			}
			break
		}
	}

	if registryIndex == -1 {
		return nil, sdkerrors.Wrapf(types.ErrSchemaNotInRegistry, "schema %s not found in registry", msg.NftSchemaCode)
	}

	// 4. Update vote
	virtualSchemaProposal.Registry[registryIndex].Status = msg.Option

	// 5. Check if voting is complete and process results
	_, totalVotes := countVotes(virtualSchemaProposal.Registry)
	voteTreshold := len(virtualSchemaProposal.Registry)

	// Check if all votes are in
	if totalVotes == voteTreshold {
		k.AfterProposalSuccess(ctx, virtualSchemaProposal.Id)
	}

	// save
	k.SetVirtualSchemaProposal(ctx, virtualSchemaProposal)

	return &types.MsgVoteCreateVirtualSchemaResponse{}, nil
}

// countVotes returns the number of accept votes and total votes cast
func countVotes(registry []*types.VirtualSchemaRegistry) (acceptCount, totalVotes int) {
	for _, reg := range registry {
		if reg.Status != types.RegistryStatus_PENDING {
			totalVotes++
			if reg.Status == types.RegistryStatus_ACCEPT {
				acceptCount++
			}
		}
	}
	return
}
