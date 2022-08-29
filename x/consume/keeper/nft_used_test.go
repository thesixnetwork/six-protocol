package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/consume/keeper"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNftUsed(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NftUsed {
	items := make([]types.NftUsed, n)
	for i := range items {
		items[i].Token = strconv.Itoa(i)

		keeper.SetNftUsed(ctx, items[i])
	}
	return items
}

func TestNftUsedGet(t *testing.T) {
	keeper, ctx := keepertest.ConsumeKeeper(t)
	items := createNNftUsed(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNftUsed(ctx,
			item.Token,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNftUsedRemove(t *testing.T) {
	keeper, ctx := keepertest.ConsumeKeeper(t)
	items := createNNftUsed(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNftUsed(ctx,
			item.Token,
		)
		_, found := keeper.GetNftUsed(ctx,
			item.Token,
		)
		require.False(t, found)
	}
}

func TestNftUsedGetAll(t *testing.T) {
	keeper, ctx := keepertest.ConsumeKeeper(t)
	items := createNNftUsed(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNftUsed(ctx)),
	)
}
