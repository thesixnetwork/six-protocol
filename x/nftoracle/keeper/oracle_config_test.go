package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createTestOracleConfig(keeper keeper.Keeper, ctx sdk.Context) types.OracleConfig {
	item := types.OracleConfig{MinimumConfirmation: 4}
	keeper.SetOracleConfig(ctx, item)
	return item
}

func TestOracleConfigGet(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)

	_, found := k.GetOracleConfig(ctx)
	require.False(t, found)

	item := createTestOracleConfig(k, ctx)
	got, found := k.GetOracleConfig(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&got),
	)
}

func TestOracleConfigRemove(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	createTestOracleConfig(k, ctx)
	k.RemoveOracleConfig(ctx)
	_, found := k.GetOracleConfig(ctx)
	require.False(t, found)
}
