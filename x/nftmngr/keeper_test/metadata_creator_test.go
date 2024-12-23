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
