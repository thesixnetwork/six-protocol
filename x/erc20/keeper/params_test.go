package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/erc20/types"
)

func TestGetSetParams(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	params := types.NewParams(true, []string{}, []string{})
	require.NoError(t, k.SetParams(ctx, params))

	got := k.GetParams(ctx)
	require.Equal(t, params.EnableErc20, got.EnableErc20)
}

func TestSetParamsDisabled(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	params := types.NewParams(false, []string{}, []string{})
	require.NoError(t, k.SetParams(ctx, params))

	got := k.GetParams(ctx)
	require.False(t, got.EnableErc20)
}

func TestIsERC20Enabled(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	require.NoError(t, k.SetParams(ctx, types.NewParams(true, []string{}, []string{})))
	require.True(t, k.IsERC20Enabled(ctx))

	require.NoError(t, k.SetParams(ctx, types.NewParams(false, []string{}, []string{})))
	require.False(t, k.IsERC20Enabled(ctx))
}

func TestGetParamsDefault(t *testing.T) {
	k, ctx := keepertest.Erc20Keeper(t)

	params := k.GetParams(ctx)
	require.False(t, params.EnableErc20)
	require.Empty(t, params.DynamicPrecompiles)
	require.Empty(t, params.NativePrecompiles)
}
