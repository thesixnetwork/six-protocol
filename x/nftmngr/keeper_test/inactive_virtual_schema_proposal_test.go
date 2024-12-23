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

func createNInactiveVirtualSchemaProposal(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.InactiveVirtualSchemaProposal {
	items := make([]types.InactiveVirtualSchemaProposal, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)

		keeper.SetInactiveVirtualSchemaProposal(ctx, items[i])
	}
	return items
}

func TestInactiveVirtualSchemaProposalGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNInactiveVirtualSchemaProposal(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetInactiveVirtualSchemaProposal(ctx,
			item.Id,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestInactiveVirtualSchemaProposalRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNInactiveVirtualSchemaProposal(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveInactiveVirtualSchemaProposal(ctx,
			item.Id,
		)
		_, found := keeper.GetInactiveVirtualSchemaProposal(ctx,
			item.Id,
		)
		require.False(t, found)
	}
}

func TestInactiveVirtualSchemaProposalGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNInactiveVirtualSchemaProposal(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllInactiveVirtualSchemaProposal(ctx)),
	)
}
