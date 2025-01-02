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
