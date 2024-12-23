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

func createNActiveVirtualSchemaProposal(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActiveVirtualSchemaProposal {
	items := make([]types.ActiveVirtualSchemaProposal, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetActiveVirtualSchemaProposal(ctx, items[i])
	}
	return items
}

func TestActiveVirtualSchemaProposalGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActiveVirtualSchemaProposal(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetActiveVirtualSchemaProposal(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestActiveVirtualSchemaProposalRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActiveVirtualSchemaProposal(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActiveVirtualSchemaProposal(ctx,
			item.Index,
		)
		_, found := keeper.GetActiveVirtualSchemaProposal(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestActiveVirtualSchemaProposalGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActiveVirtualSchemaProposal(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActiveVirtualSchemaProposal(ctx)),
	)
}
