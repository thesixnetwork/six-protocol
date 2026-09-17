package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createNActionSignerConfig(keeper keeper.Keeper, ctx sdk.Context, n int) []types.ActionSignerConfig {
	items := make([]types.ActionSignerConfig, n)
	for i := range items {
		items[i].Chain = strconv.Itoa(i)

		keeper.SetActionSignerConfig(ctx, items[i])
	}
	return items
}

func TestActionSignerConfigGet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSignerConfig(k, ctx, 10)
	for _, item := range items {
		got, found := k.GetActionSignerConfig(ctx,
			item.Chain,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestActionSignerConfigRemove(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSignerConfig(k, ctx, 10)
	for _, item := range items {
		k.RemoveActionSignerConfig(ctx,
			item.Chain,
		)
		_, found := k.GetActionSignerConfig(ctx,
			item.Chain,
		)
		require.False(t, found)
	}
}

func TestActionSignerConfigGetAll(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	items := createNActionSignerConfig(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllActionSignerConfig(ctx)),
	)
}
