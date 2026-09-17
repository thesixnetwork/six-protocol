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

func createNMintRequest(keeper keeper.Keeper, ctx sdk.Context, n int) []types.MintRequest {
	items := make([]types.MintRequest, n)
	for i := range items {
		items[i].CreatedAt = time.Unix(int64(i), 0).UTC()
		items[i].ValidUntil = time.Unix(int64(i+100), 0).UTC()
		items[i].Id = keeper.AppendMintRequest(ctx, items[i])
	}
	return items
}

func TestMintRequestGet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNMintRequest(k, ctx, 10)
	for _, item := range items {
		got, found := k.GetMintRequest(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMintRequestRemove(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNMintRequest(k, ctx, 10)
	for _, item := range items {
		k.RemoveMintRequest(ctx, item.Id)
		_, found := k.GetMintRequest(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMintRequestGetAll(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNMintRequest(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllMintRequest(ctx)),
	)
}

func TestMintRequestCount(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNMintRequest(k, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, k.GetMintRequestCount(ctx))
}

func TestMintRequestSet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNMintRequest(k, ctx, 1)
	item := items[0]
	item.NftSchemaCode = "schema.code"
	item.TokenId = "1"
	k.SetMintRequest(ctx, item)
	got, found := k.GetMintRequest(ctx, item.Id)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&got),
	)
}

func TestMintRequestIDBytes(t *testing.T) {
	id := uint64(42)
	require.Equal(t, id, keeper.GetMintRequestIDFromBytes(keeper.GetMintRequestIDBytes(id)))
}

func TestActiveMintRequestQueue(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNMintRequest(k, ctx, 3)
	endTime := time.Unix(1000, 0).UTC()
	for _, item := range items {
		k.InsertActiveMintRequestQueue(ctx, item.Id, endTime)
	}

	var iterated []types.MintRequest
	k.IterateActiveMintRequestsQueue(ctx, endTime.Add(time.Second), func(mintRequest types.MintRequest) bool {
		iterated = append(iterated, mintRequest)
		return false
	})
	require.Len(t, iterated, 3)

	// nothing is active before the queue end time
	iterated = nil
	k.IterateActiveMintRequestsQueue(ctx, endTime.Add(-time.Second), func(mintRequest types.MintRequest) bool {
		iterated = append(iterated, mintRequest)
		return false
	})
	require.Len(t, iterated, 0)

	k.RemoveFromActiveMintRequestQueue(ctx, items[0].Id, endTime)
	iterated = nil
	k.IterateActiveMintRequestsQueue(ctx, endTime.Add(time.Second), func(mintRequest types.MintRequest) bool {
		iterated = append(iterated, mintRequest)
		return false
	})
	require.Len(t, iterated, 2)
}
