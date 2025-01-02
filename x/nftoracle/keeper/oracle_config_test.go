package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func createTestOracleConfig(keeper *keeper.Keeper, ctx sdk.Context) types.OracleConfig {
	item := types.OracleConfig{}
	keeper.SetOracleConfig(ctx, item)
	return item
}

func TestOracleConfigGet(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	item := createTestOracleConfig(keeper, ctx)
	rst, found := keeper.GetOracleConfig(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestOracleConfigRemove(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	createTestOracleConfig(keeper, ctx)
	keeper.RemoveOracleConfig(ctx)
	_, found := keeper.GetOracleConfig(ctx)
	require.False(t, found)
}
