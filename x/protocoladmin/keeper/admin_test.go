package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"
)

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

func TestAuthenticate(t *testing.T) {
	k, ctx := keepertest.ProtocoladminKeeper(t)

	superAdmin := sample.AccAddress()
	groupAdmin := sample.AccAddress()
	outsider := sample.AccAddress()

	k.SetAdmin(ctx, types.Admin{Group: keeper.SUPER_ADMIN, Admin: superAdmin})
	k.SetAdmin(ctx, types.Admin{Group: "group1", Admin: groupAdmin})

	// super admin authenticates for any group
	require.True(t, k.Authenticate(ctx, "group1", superAdmin))
	require.True(t, k.Authenticate(ctx, "group2", superAdmin))

	// group admin authenticates only for its own group
	require.True(t, k.Authenticate(ctx, "group1", groupAdmin))
	require.False(t, k.Authenticate(ctx, "group2", groupAdmin))

	// unknown address never authenticates
	require.False(t, k.Authenticate(ctx, "group1", outsider))
}
