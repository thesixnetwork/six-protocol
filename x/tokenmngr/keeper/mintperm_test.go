package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMintperm(keeper keeper.Keeper, ctx context.Context, n int) []types.Mintperm {
	items := make([]types.Mintperm, n)
	for i := range items {
		items[i].Token = strconv.Itoa(i)
		items[i].Address = strconv.Itoa(i)

		keeper.SetMintperm(ctx, items[i])
	}
	return items
}

func TestMintpermGet(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	items := createNMintperm(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMintperm(ctx,
			item.Token,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestMintpermRemove(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	items := createNMintperm(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMintperm(ctx,
			item.Token,
			item.Address,
		)
		_, found := keeper.GetMintperm(ctx,
			item.Token,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestMintpermGetAll(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	items := createNMintperm(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMintperm(ctx)),
	)
}
