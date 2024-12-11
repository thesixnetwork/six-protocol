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

	virtualSchema, found := k.GetVirtualSchema(ctx, msg.VirtualNftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.VirtualNftSchemaCode)
	}

	srcSchema, found := k.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	if srcSchema.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
	}

	// Track if the vote has been processed
	voteProcessed := false

	// loop to find then schema registry
	for i, registry := range virtualSchema.Registry {
		if registry.NftSchemaCode == msg.NftSchemaCode {
			// Check if already voted
			if registry.Status != types.RegistryStatus_PENDING {
				return nil, sdkerrors.Wrap(types.ErrAlreadyVote, msg.Creator)
			}

			// Update the status
			virtualSchema.Registry[i].Status = msg.Option
			voteProcessed = true
		}
	}

	// Ensure the vote was processed
	if !voteProcessed {
		return nil, sdkerrors.Wrap(types.ErrSchemaNotInRegistry, msg.NftSchemaCode)
	}

	// Count votes
	var (
		acceptCount int
		totalVotes  int
		voteTreshold = len(virtualSchema.Registry)
	)

	for _, registry := range virtualSchema.Registry {
		if registry.Status == types.RegistryStatus_ACCEPT {
			acceptCount++
		}
		if registry.Status != types.RegistryStatus_PENDING {
			totalVotes++
		}
	}

	// Check if all votes are in
	if totalVotes == voteTreshold {
		if acceptCount == voteTreshold {
			virtualSchema.Enable = true
		} else {
			virtualSchema.Enable = false
		}
	}

	// save
	k.SetVirtualSchema(ctx, virtualSchema)

	return &types.MsgVoteCreateVirtualSchemaResponse{}, nil
}


func (K Keeper) AfterAllVoteAccept(ctx sdk.Context, virtualSchemaCode string) error {
	
	return nil
}


