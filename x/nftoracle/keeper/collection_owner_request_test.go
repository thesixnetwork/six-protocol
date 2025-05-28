package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createNCollectionOwnerRequest(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CollectionOwnerRequest {
	items := make([]types.CollectionOwnerRequest, n)
	for i := range items {
		items[i].Id = keeper.AppendCollectionOwnerRequest(ctx, items[i])
	}
	return items
}

func TestCollectionOwnerRequestGet(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNCollectionOwnerRequest(&keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCollectionOwnerRequest(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCollectionOwnerRequestRemove(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNCollectionOwnerRequest(&keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCollectionOwnerRequest(ctx, item.Id)
		_, found := keeper.GetCollectionOwnerRequest(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCollectionOwnerRequestGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNCollectionOwnerRequest(&keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCollectionOwnerRequest(ctx)),
	)
}

func TestCollectionOwnerRequestCount(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNCollectionOwnerRequest(&keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCollectionOwnerRequestCount(ctx))
}
