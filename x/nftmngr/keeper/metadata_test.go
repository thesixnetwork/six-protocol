package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMetadataCreator(keeper keeper.Keeper, ctx sdk.Context, n int) []types.MetadataCreator {
	items := make([]types.MetadataCreator, n)
	for i := range items {
		items[i].NftSchemaCode = strconv.Itoa(i)
		keeper.SetMetadataCreator(ctx, items[i])
	}
	return items
}

func TestMetadataCreatorGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNMetadataCreator(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMetadataCreator(ctx,
			item.NftSchemaCode,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestMetadataCreatorRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNMetadataCreator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMetadataCreator(ctx,
			item.NftSchemaCode,
		)
		_, found := keeper.GetMetadataCreator(ctx,
			item.NftSchemaCode,
		)
		require.False(t, found)
	}
}

func TestMetadataCreatorGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNMetadataCreator(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMetadataCreator(ctx)),
	)
}

func createNNftCollection(keeper keeper.Keeper, ctx sdk.Context, n int) []types.NftCollection {
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

func TestNftCollectionDataCount(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	require.Equal(t, uint64(0), keeper.GetNftCollectionDataCount(ctx))
	keeper.SetNftCollectionDataCount(ctx, 42)
	require.Equal(t, uint64(42), keeper.GetNftCollectionDataCount(ctx))
}
