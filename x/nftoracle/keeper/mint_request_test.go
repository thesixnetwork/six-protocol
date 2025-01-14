package keeper_test

import (
	"testing"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func createNMintRequest(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MintRequest {
	items := make([]types.MintRequest, n)
	for i := range items {
		items[i].Id = keeper.AppendMintRequest(ctx, items[i])
	}
	return items
}

func TestMintRequestGet(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNMintRequest(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMintRequest(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMintRequestRemove(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNMintRequest(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMintRequest(ctx, item.Id)
		_, found := keeper.GetMintRequest(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMintRequestGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNMintRequest(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMintRequest(ctx)),
	)
}

func TestMintRequestCount(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNMintRequest(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMintRequestCount(ctx))
}
