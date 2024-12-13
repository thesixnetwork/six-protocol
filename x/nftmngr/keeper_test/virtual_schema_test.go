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

func createNVirtualSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VirtualSchema {
	items := make([]types.VirtualSchema, n)
	for i := range items {
		items[i].VirtualNftSchemaCode = strconv.Itoa(i)

		keeper.SetVirtualSchema(ctx, items[i])
	}
	return items
}

func TestVirtualSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVirtualSchema(ctx,
			item.VirtualNftSchemaCode,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestVirtualSchemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVirtualSchema(ctx,
			item.VirtualNftSchemaCode,
		)
		_, found := keeper.GetVirtualSchema(ctx,
			item.VirtualNftSchemaCode,
		)
		require.False(t, found)
	}
}

func TestVirtualSchemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVirtualSchema(ctx)),
	)
}
