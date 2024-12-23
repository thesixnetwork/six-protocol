package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVirtualSchemaProposal(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VirtualSchemaProposal {
	items := make([]types.VirtualSchemaProposal, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)

		keeper.SetVirtualSchemaProposal(ctx, items[i])
	}
	return items
}

func TestVirtualSchemaProposalGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualSchemaProposal(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVirtualSchemaProposal(ctx,
			item.Id,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestVirtualSchemaProposalRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualSchemaProposal(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVirtualSchemaProposal(ctx,
			item.Id,
		)
		_, found := keeper.GetVirtualSchemaProposal(ctx,
			item.Id,
		)
		require.False(t, found)
	}
}

func TestVirtualSchemaProposalGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualSchemaProposal(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVirtualSchemaProposal(ctx)),
	)
}
