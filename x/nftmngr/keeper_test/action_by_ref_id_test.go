package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNActionByRefId(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ActionByRefId {
	items := make([]types.ActionByRefId, n)
	for i := range items {
		items[i].RefId = strconv.Itoa(i)

		keeper.SetActionByRefId(ctx, items[i])
	}
	return items
}

func TestActionByRefIdGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActionByRefId(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetActionByRefId(ctx,
			item.RefId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestActionByRefIdRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActionByRefId(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActionByRefId(ctx,
			item.RefId,
		)
		_, found := keeper.GetActionByRefId(ctx,
			item.RefId,
		)
		require.False(t, found)
	}
}

func TestActionByRefIdGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNActionByRefId(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActionByRefId(ctx)),
	)
}
