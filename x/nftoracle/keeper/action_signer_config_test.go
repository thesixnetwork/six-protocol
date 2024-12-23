package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNActionSignerConfig(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActionSignerConfig {
	items := make([]types.ActionSignerConfig, n)
	for i := range items {
		items[i].Chain = strconv.Itoa(i)

		keeper.SetActionSignerConfig(ctx, items[i])
	}
	return items
}

func TestActionSignerConfigGet(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSignerConfig(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetActionSignerConfig(ctx,
			item.Chain,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestActionSignerConfigRemove(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSignerConfig(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActionSignerConfig(ctx,
			item.Chain,
		)
		_, found := keeper.GetActionSignerConfig(ctx,
			item.Chain,
		)
		require.False(t, found)
	}
}

func TestActionSignerConfigGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSignerConfig(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActionSignerConfig(ctx)),
	)
}
