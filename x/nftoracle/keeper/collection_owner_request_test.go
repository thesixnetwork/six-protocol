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

func createNCollectionOwnerRequest(keeper keeper.Keeper, ctx sdk.Context, n int) []types.CollectionOwnerRequest {
	items := make([]types.CollectionOwnerRequest, n)
	for i := range items {
		items[i].CreatedAt = time.Unix(int64(i), 0).UTC()
		items[i].ValidUntil = time.Unix(int64(i+100), 0).UTC()
		items[i].Id = keeper.AppendCollectionOwnerRequest(ctx, items[i])
	}
	return items
}

func TestCollectionOwnerRequestGet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNCollectionOwnerRequest(k, ctx, 10)
	for _, item := range items {
		got, found := k.GetCollectionOwnerRequest(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCollectionOwnerRequestRemove(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNCollectionOwnerRequest(k, ctx, 10)
	for _, item := range items {
		k.RemoveCollectionOwnerRequest(ctx, item.Id)
		_, found := k.GetCollectionOwnerRequest(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCollectionOwnerRequestGetAll(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNCollectionOwnerRequest(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllCollectionOwnerRequest(ctx)),
	)
}

func TestCollectionOwnerRequestCount(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNCollectionOwnerRequest(k, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, k.GetCollectionOwnerRequestCount(ctx))
}

func TestCollectionOwnerRequestSet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNCollectionOwnerRequest(k, ctx, 1)
	item := items[0]
	item.NftSchemaCode = "schema.code"
	k.SetCollectionOwnerRequest(ctx, item)
	got, found := k.GetCollectionOwnerRequest(ctx, item.Id)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&got),
	)
}

func TestCollectionOwnerRequestIDBytes(t *testing.T) {
	id := uint64(42)
	require.Equal(t, id, keeper.GetCollectionOwnerRequestIDFromBytes(keeper.GetCollectionOwnerRequestIDBytes(id)))
}

func TestActiveVerifyCollectionOwnerQueue(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNCollectionOwnerRequest(k, ctx, 3)
	endTime := time.Unix(1000, 0).UTC()
	for _, item := range items {
		k.InsertActiveVerifyCollectionOwnerRequestQueue(ctx, item.Id, endTime)
	}

	var iterated []types.CollectionOwnerRequest
	k.IterateActiveVerifyCollectionOwnersQueue(ctx, endTime.Add(time.Second), func(verifyRequest types.CollectionOwnerRequest) bool {
		iterated = append(iterated, verifyRequest)
		return false
	})
	require.Len(t, iterated, 3)

	k.RemoveFromActiveVerifyCollectionOwnerQueue(ctx, items[0].Id, endTime)
	iterated = nil
	k.IterateActiveVerifyCollectionOwnersQueue(ctx, endTime.Add(time.Second), func(verifyRequest types.CollectionOwnerRequest) bool {
		iterated = append(iterated, verifyRequest)
		return false
	})
	require.Len(t, iterated, 2)
}
