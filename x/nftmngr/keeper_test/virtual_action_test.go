package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVirtualAction(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VirtualAction {
	items := make([]types.VirtualAction, n)
	for i := range items {
		items[i].VirtualNftSchemaCode = strconv.Itoa(i)

		keeper.SetVirtualAction(ctx, items[i])
	}
	return items
}

func TestVirtualActionGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualAction(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVirtualAction(ctx,
			item.VirtualNftSchemaCode,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestVirtualActionRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualAction(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVirtualAction(ctx,
			item.VirtualNftSchemaCode,
			item.Name,
		)
		_, found := keeper.GetVirtualAction(ctx,
			item.VirtualNftSchemaCode,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestVirtualActionGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualAction(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVirtualAction(ctx)),
	)
}
