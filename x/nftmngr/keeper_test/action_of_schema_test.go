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

func createNActionOfSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActionOfSchema {
	items := make([]types.ActionOfSchema, n)
	for i := range items {
		items[i].NftSchemaCode = strconv.Itoa(i)
		items[i].Name = strconv.Itoa(i)

		keeper.SetActionOfSchema(ctx, items[i])
	}
	return items
}

func TestActionOfSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActionOfSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetActionOfSchema(ctx,
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

func TestActionOfSchemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActionOfSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActionOfSchema(ctx,
			item.NftSchemaCode,
			item.Name,
		)
		_, found := keeper.GetActionOfSchema(ctx,
			item.NftSchemaCode,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestActionOfSchemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActionOfSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActionOfSchema(ctx)),
	)
}
