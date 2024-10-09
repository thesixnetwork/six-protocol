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

func createNSchemaAttribute(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SchemaAttribute {
	items := make([]types.SchemaAttribute, n)
	for i := range items {
		items[i].NftSchemaCode = strconv.Itoa(i)
		items[i].Name = strconv.Itoa(i)

		keeper.SetSchemaAttribute(ctx, items[i])
	}
	return items
}

func TestSchemaAttributeGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNSchemaAttribute(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSchemaAttribute(ctx,
			item.NftSchemaCode,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSchemaAttributeRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNSchemaAttribute(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSchemaAttribute(ctx,
			item.NftSchemaCode,
			item.Name,
		)
		_, found := keeper.GetSchemaAttribute(ctx,
			item.NftSchemaCode,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestSchemaAttributeGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNSchemaAttribute(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSchemaAttribute(ctx)),
	)
}
