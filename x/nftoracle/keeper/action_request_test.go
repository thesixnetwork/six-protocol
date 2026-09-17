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

func createNActionRequest(keeper keeper.Keeper, ctx sdk.Context, n int) []types.ActionOracleRequest {
	items := make([]types.ActionOracleRequest, n)
	for i := range items {
		items[i].CreatedAt = time.Unix(int64(i), 0).UTC()
		items[i].ValidUntil = time.Unix(int64(i+100), 0).UTC()
		items[i].Id = keeper.AppendActionRequest(ctx, items[i])
	}
	return items
}

func TestActionRequestGet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(k, ctx, 10)
	for _, item := range items {
		got, found := k.GetActionRequest(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestActionRequestRemove(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(k, ctx, 10)
	for _, item := range items {
		k.RemoveActionRequest(ctx, item.Id)
		_, found := k.GetActionRequest(ctx, item.Id)
		require.False(t, found)
	}
}

func TestActionRequestGetAll(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllActionRequest(ctx)),
	)
}

func TestActionRequestCount(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(k, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, k.GetActionRequestCount(ctx))
}

func TestActionRequestSet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(k, ctx, 1)
	item := items[0]
	item.NftSchemaCode = "schema.code"
	item.Action = "some_action"
	k.SetActionRequest(ctx, item)
	got, found := k.GetActionRequest(ctx, item.Id)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&got),
	)
}

func TestActionRequestIDBytes(t *testing.T) {
	id := uint64(42)
	require.Equal(t, id, keeper.GetActionRequestIDFromBytes(keeper.GetActionRequestIDBytes(id)))
}

func TestActiveActionRequestQueue(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionRequest(k, ctx, 3)
	endTime := time.Unix(1000, 0).UTC()
	for _, item := range items {
		k.InsertActiveActionRequestQueue(ctx, item.Id, endTime)
	}

	var iterated []types.ActionOracleRequest
	k.IterateActiveActionRequestsQueue(ctx, endTime.Add(time.Second), func(actionRequest types.ActionOracleRequest) bool {
		iterated = append(iterated, actionRequest)
		return false
	})
	require.Len(t, iterated, 3)

	k.RemoveFromActiveActionRequestQueue(ctx, items[0].Id, endTime)
	iterated = nil
	k.IterateActiveActionRequestsQueue(ctx, endTime.Add(time.Second), func(actionRequest types.ActionOracleRequest) bool {
		iterated = append(iterated, actionRequest)
		return false
	})
	require.Len(t, iterated, 2)
}
