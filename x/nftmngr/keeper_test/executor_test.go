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

func createNActionExecutor(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActionExecutor {
	items := make([]types.ActionExecutor, n)
	for i := range items {
		items[i].NftSchemaCode = strconv.Itoa(i)
		items[i].ExecutorAddress = strconv.Itoa(i)

		keeper.SetActionExecutor(ctx, items[i])
	}
	return items
}

func TestActionExecutorGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActionExecutor(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetActionExecutor(ctx,
			item.NftSchemaCode,
			item.ExecutorAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestActionExecutorRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActionExecutor(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActionExecutor(ctx,
			item.NftSchemaCode,
			item.ExecutorAddress,
		)
		_, found := keeper.GetActionExecutor(ctx,
			item.NftSchemaCode,
			item.ExecutorAddress,
		)
		require.False(t, found)
	}
}

func TestActionExecutorGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActionExecutor(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActionExecutor(ctx)),
	)
}

func createNExecutorOfSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ExecutorOfSchema {
	items := make([]types.ExecutorOfSchema, n)
	for i := range items {
		items[i].NftSchemaCode = strconv.Itoa(i)

		keeper.SetExecutorOfSchema(ctx, items[i])
	}
	return items
}

func TestExecutorOfSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNExecutorOfSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetExecutorOfSchema(ctx,
			item.NftSchemaCode,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestExecutorOfSchemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNExecutorOfSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExecutorOfSchema(ctx,
			item.NftSchemaCode,
		)
		_, found := keeper.GetExecutorOfSchema(ctx,
			item.NftSchemaCode,
		)
		require.False(t, found)
	}
}

func TestExecutorOfSchemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNExecutorOfSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllExecutorOfSchema(ctx)),
	)
}
