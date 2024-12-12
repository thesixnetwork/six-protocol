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

func createNDisableVirtualSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DisableVirtualSchema {
	items := make([]types.DisableVirtualSchema, n)
	for i := range items {
		items[i].NftSchemaCode = strconv.Itoa(i)

		keeper.SetDisableVirtualSchema(ctx, items[i])
	}
	return items
}

func TestDisableVirtualSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNDisableVirtualSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDisableVirtualSchema(ctx,
			item.NftSchemaCode,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDisableVirtualSchemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNDisableVirtualSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDisableVirtualSchema(ctx,
			item.NftSchemaCode,
		)
		_, found := keeper.GetDisableVirtualSchema(ctx,
			item.NftSchemaCode,
		)
		require.False(t, found)
	}
}

func TestDisableVirtualSchemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNDisableVirtualSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDisableVirtualSchema(ctx)),
	)
}
