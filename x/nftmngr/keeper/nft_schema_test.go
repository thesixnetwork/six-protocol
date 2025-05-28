package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNftschema(keeper keeper.Keeper, ctx context.Context, n int) []types.NFTSchema {
	items := make([]types.NFTSchema, n)
	for i := range items {
		items[i].Code = strconv.Itoa(i)

		keeper.SetNftschema(ctx, items[i])
	}
	return items
}

func TestNftschemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNftschema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNftschema(ctx,
			item.Code,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestNftschemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNftschema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNftschema(ctx,
			item.Code,
		)
		_, found := keeper.GetNftschema(ctx,
			item.Code,
		)
		require.False(t, found)
	}
}

func TestNftschemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNftschema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNftschema(ctx)),
	)
}
