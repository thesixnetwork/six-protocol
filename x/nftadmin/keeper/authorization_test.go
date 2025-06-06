package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftadmin/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftadmin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createTestAuthorization(keeper *keeper.Keeper, ctx sdk.Context) types.Authorization {
	item := types.Authorization{}
	keeper.SetAuthorization(ctx, item)
	return item
}

func TestAuthorizationGet(t *testing.T) {
	keeper, ctx := keepertest.NftadminKeeper(t)
	item := createTestAuthorization(&keeper, ctx)
	rst, found := keeper.GetAuthorization(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestAuthorizationRemove(t *testing.T) {
	keeper, ctx := keepertest.NftadminKeeper(t)
	createTestAuthorization(&keeper, ctx)
	keeper.RemoveAuthorization(ctx)
	_, found := keeper.GetAuthorization(ctx)
	require.False(t, found)
}
