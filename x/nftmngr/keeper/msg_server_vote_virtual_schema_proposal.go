package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) VoteVirtualSchemaProposal(goCtx context.Context, msg *types.MsgVoteVirtualSchemaProposal) (*types.MsgVoteVirtualSchemaProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgVoteVirtualSchemaProposalResponse{}, nil
}
