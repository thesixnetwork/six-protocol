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

func createNAdmin(keeper keeper.Keeper, ctx context.Context, n int) []types.Admin {
	items := make([]types.Admin, n)
	for i := range items {
		items[i].Group = strconv.Itoa(i)
		items[i].Admin = strconv.Itoa(i)

		keeper.SetAdmin(ctx, items[i])
	}
	return items
}

func TestAdminGet(t *testing.T) {
	keeper, ctx := keepertest.ProtocoladminKeeper(t)
	items := createNAdmin(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAdmin(ctx,
			item.Group,
			item.Admin,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestAdminRemove(t *testing.T) {
	keeper, ctx := keepertest.ProtocoladminKeeper(t)
	items := createNAdmin(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAdmin(ctx,
			item.Group,
			item.Admin,
		)
		_, found := keeper.GetAdmin(ctx,
			item.Group,
			item.Admin,
		)
		require.False(t, found)
	}
}

func TestAdminGetAll(t *testing.T) {
	keeper, ctx := keepertest.ProtocoladminKeeper(t)
	items := createNAdmin(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAdmin(ctx)),
	)
}
