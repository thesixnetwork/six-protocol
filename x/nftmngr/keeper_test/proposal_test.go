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
		items[i].Id = strconv.Itoa(i)

		keeper.SetActiveVirtualSchemaProposal(ctx, items[i])
	}
	return items
}

func TestActiveVirtualSchemaProposalGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActiveVirtualSchemaProposal(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetActiveVirtualSchemaProposal(ctx,
			item.Id,
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
			item.Id,
		)
		_, found := keeper.GetActiveVirtualSchemaProposal(ctx,
			item.Id,
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

