package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func createTestOptions(keeper *keeper.Keeper, ctx sdk.Context) types.Options {
	item := types.Options{}
	keeper.SetOptions(ctx, item)
	return item
}

func TestOptionsGet(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	item := createTestOptions(keeper, ctx)
	rst, found := keeper.GetOptions(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestOptionsRemove(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	createTestOptions(keeper, ctx)
	keeper.RemoveOptions(ctx)
	_, found := keeper.GetOptions(ctx)
	require.False(t, found)
}
