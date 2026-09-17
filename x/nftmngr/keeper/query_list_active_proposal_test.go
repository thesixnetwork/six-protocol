package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

func TestListActiveProposalQuery(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)

	_, err := k.ListActiveProposal(ctx, nil)
	require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))

	// no active proposals
	resp, err := k.ListActiveProposal(ctx, &types.QueryListActiveProposalRequest{})
	require.NoError(t, err)
	require.Empty(t, resp.VirtualSchemaProposal)

	// active entry without the underlying proposal is a not found error
	k.SetActiveVirtualSchemaProposal(ctx, types.ActiveVirtualSchemaProposal{Id: "orphan"})
	_, err = k.ListActiveProposal(ctx, &types.QueryListActiveProposalRequest{})
	require.ErrorIs(t, err, status.Error(codes.NotFound, "not found"))
	k.RemoveActiveVirtualSchemaProposal(ctx, "orphan")

	// active entries backed by stored proposals
	proposals := createNVirtualSchemaProposal(k, ctx, 3)
	for _, proposal := range proposals {
		k.SetActiveVirtualSchemaProposal(ctx, types.ActiveVirtualSchemaProposal{Id: proposal.Id})
	}

	resp, err = k.ListActiveProposal(ctx, &types.QueryListActiveProposalRequest{})
	require.NoError(t, err)
	require.ElementsMatch(t,
		nullify.Fill(proposals),
		nullify.Fill(resp.VirtualSchemaProposal),
	)
}
