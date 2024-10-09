package keeper_test

import (
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

func createNNFTSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NFTSchema {
	items := make([]types.NFTSchema, n)
	for i := range items {
		items[i].Code = strconv.Itoa(i)

		keeper.SetNFTSchema(ctx, items[i])
	}
	return items
}

func TestNFTSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNFTSchema(ctx,
			item.Code,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNFTSchemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNFTSchema(ctx,
			item.Code,
		)
		_, found := keeper.GetNFTSchema(ctx,
			item.Code,
		)
		require.False(t, found)
	}
}

func TestNFTSchemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNFTSchema(ctx)),
	)
}
