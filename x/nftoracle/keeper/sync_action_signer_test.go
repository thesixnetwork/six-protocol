package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func createNSyncActionSigner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SyncActionSigner {
	items := make([]types.SyncActionSigner, n)
	for i := range items {
		items[i].Id = keeper.AppendSyncActionSigner(ctx, items[i])
	}
	return items
}

func TestSyncActionSignerGet(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNSyncActionSigner(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSyncActionSigner(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSyncActionSignerRemove(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNSyncActionSigner(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSyncActionSigner(ctx, item.Id)
		_, found := keeper.GetSyncActionSigner(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSyncActionSignerGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNSyncActionSigner(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSyncActionSigner(ctx)),
	)
}

func TestSyncActionSignerCount(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNSyncActionSigner(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSyncActionSignerCount(ctx))
}
