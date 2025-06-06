package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/keeper"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNGroup(keeper keeper.Keeper, ctx context.Context, n int) []types.Group {
	items := make([]types.Group, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetGroup(ctx, items[i])
	}
	return items
}

func TestGroupGet(t *testing.T) {
	keeper, ctx := keepertest.ProtocoladminKeeper(t)
	items := createNGroup(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGroup(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestGroupRemove(t *testing.T) {
	keeper, ctx := keepertest.ProtocoladminKeeper(t)
	items := createNGroup(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGroup(ctx,
			item.Name,
		)
		_, found := keeper.GetGroup(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestGroupGetAll(t *testing.T) {
	keeper, ctx := keepertest.ProtocoladminKeeper(t)
	items := createNGroup(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGroup(ctx)),
	)
}
