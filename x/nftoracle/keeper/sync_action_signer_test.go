package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createNSyncActionSigner(keeper keeper.Keeper, ctx sdk.Context, n int) []types.SyncActionSigner {
	items := make([]types.SyncActionSigner, n)
	for i := range items {
		items[i].CreatedAt = time.Unix(int64(i), 0).UTC()
		items[i].ValidUntil = time.Unix(int64(i+100), 0).UTC()
		items[i].Id = keeper.AppendSyncActionSigner(ctx, items[i])
	}
	return items
}

func TestSyncActionSignerGet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNSyncActionSigner(k, ctx, 10)
	for _, item := range items {
		got, found := k.GetSyncActionSigner(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSyncActionSignerRemove(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNSyncActionSigner(k, ctx, 10)
	for _, item := range items {
		k.RemoveSyncActionSigner(ctx, item.Id)
		_, found := k.GetSyncActionSigner(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSyncActionSignerGetAll(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNSyncActionSigner(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllSyncActionSigner(ctx)),
	)
}

func TestSyncActionSignerCount(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNSyncActionSigner(k, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, k.GetSyncActionSignerCount(ctx))
}

func TestSyncActionSignerSet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNSyncActionSigner(k, ctx, 1)
	item := items[0]
	item.Chain = "ethereum"
	item.ActorAddress = "actor"
	k.SetSyncActionSigner(ctx, item)
	got, found := k.GetSyncActionSigner(ctx, item.Id)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&got),
	)
}

func TestSyncActionSignerIDBytes(t *testing.T) {
	id := uint64(42)
	require.Equal(t, id, keeper.GetSyncActionSignerIDFromBytes(keeper.GetSyncActionSignerIDBytes(id)))
}

func TestActiveSyncActionSignerQueue(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNSyncActionSigner(k, ctx, 3)
	endTime := time.Unix(1000, 0).UTC()
	for _, item := range items {
		k.InsertActiveSyncActionSignerQueue(ctx, item.Id, endTime)
	}

	var iterated []types.SyncActionSigner
	k.IterateActiveSyncActionSignerQueue(ctx, endTime.Add(time.Second), func(syncRequest types.SyncActionSigner) bool {
		iterated = append(iterated, syncRequest)
		return false
	})
	require.Len(t, iterated, 3)

	k.RemoveFromActiveSyncActionSignerQueue(ctx, items[0].Id, endTime)
	iterated = nil
	k.IterateActiveSyncActionSignerQueue(ctx, endTime.Add(time.Second), func(syncRequest types.SyncActionSigner) bool {
		iterated = append(iterated, syncRequest)
		return false
	})
	require.Len(t, iterated, 2)
}
