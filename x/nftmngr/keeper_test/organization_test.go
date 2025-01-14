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

func createNOrganization(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Organization {
	items := make([]types.Organization, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetOrganization(ctx, items[i])
	}
	return items
}

func TestOrganizationGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNOrganization(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetOrganization(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestOrganizationRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNOrganization(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOrganization(ctx,
			item.Name,
		)
		_, found := keeper.GetOrganization(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestOrganizationGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNOrganization(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOrganization(ctx)),
	)
}
