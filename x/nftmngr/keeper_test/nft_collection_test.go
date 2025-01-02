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

func createNNftCollection(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NftCollection {
	items := make([]types.NftCollection, n)
	for i := range items {
		items[i].NftSchemaCode = strconv.Itoa(i)

		keeper.SetNftCollection(ctx, items[i])
	}
	return items
}

func TestNftCollectionGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNftCollection(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNftCollection(ctx,
			item.NftSchemaCode,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNftCollectionRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNftCollection(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNftCollection(ctx,
			item.NftSchemaCode,
		)
		_, found := keeper.GetNftCollection(ctx,
			item.NftSchemaCode,
		)
		require.False(t, found)
	}
}

func TestNftCollectionGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNftCollection(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNftCollection(ctx)),
	)
}
