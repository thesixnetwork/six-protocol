package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNftData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NftData {
	items := make([]types.NftData, n)
	for i := range items {
		items[i].NftSchemaCode = strconv.Itoa(i)
		items[i].TokenId = strconv.Itoa(i)

		keeper.SetNftData(ctx, items[i])
	}
	return items
}

func TestNftDataGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNftData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNftData(ctx,
			item.NftSchemaCode,
			item.TokenId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestNftDataRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNftData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNftData(ctx,
			item.NftSchemaCode,
			item.TokenId,
		)
		_, found := keeper.GetNftData(ctx,
			item.NftSchemaCode,
			item.TokenId,
		)
		require.False(t, found)
	}
}

func TestNftDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNftData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNftData(ctx)),
	)
}

func TestCreateMetadata(t *testing.T) {
	_, _ = keepertest.InitSchema(t, "../../../resources/nft-schema.json")
	meta := keepertest.InitMetadata(t, "../../../resources/nft-data.json")

	keeper, ctx := keepertest.NftmngrKeeper(t)
	keeper.SetNftData(ctx, meta)

	_, found := keeper.GetNftData(ctx, meta.NftSchemaCode, meta.TokenId)
	if !found {
		fmt.Println("Metadata not found")
	} else {
		require.True(t, found)
	}
}

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

func createNMetadataCreator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MetadataCreator {
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
