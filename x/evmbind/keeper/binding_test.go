package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/evmbind/keeper"
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBinding(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Binding {
	items := make([]types.Binding, n)
	for i := range items {
		items[i].EthAddress = strconv.Itoa(i)

		keeper.SetBinding(ctx, items[i])
	}
	return items
}

func TestBindingGet(t *testing.T) {
	keeper, ctx := keepertest.EvmbindKeeper(t)
	items := createNBinding(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBinding(ctx,
			item.EthAddress,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBindingRemove(t *testing.T) {
	keeper, ctx := keepertest.EvmbindKeeper(t)
	items := createNBinding(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBinding(ctx,
			item.EthAddress,
		)
		_, found := keeper.GetBinding(ctx,
			item.EthAddress,
		)
		require.False(t, found)
	}
}

func TestBindingGetAll(t *testing.T) {
	keeper, ctx := keepertest.EvmbindKeeper(t)
	items := createNBinding(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBinding(ctx)),
	)
}
