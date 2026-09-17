package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
)

func createTestOptions(keeper keeper.Keeper, ctx context.Context) types.Options {
	item := types.Options{
		DefaultMintee: "default-mintee",
	}
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
