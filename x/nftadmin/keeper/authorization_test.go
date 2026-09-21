package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/nftadmin/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/nftadmin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func createTestAuthorization(keeper keeper.Keeper, ctx sdk.Context) types.Authorization {
	item := types.Authorization{
		RootAdmin: sample.AccAddress(),
		Permissions: []*types.Permission{
			{
				Name:      "minter",
				Addresses: []string{sample.AccAddress()},
			},
		},
	}
	keeper.SetAuthorization(ctx, item)
	return item
}

func TestAuthorizationGet(t *testing.T) {
	keeper, ctx := keepertest.NftadminKeeper(t)

	_, found := keeper.GetAuthorization(ctx)
	require.False(t, found)

	item := createTestAuthorization(keeper, ctx)
	rst, found := keeper.GetAuthorization(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestAuthorizationRemove(t *testing.T) {
	keeper, ctx := keepertest.NftadminKeeper(t)
	createTestAuthorization(keeper, ctx)
	keeper.RemoveAuthorization(ctx)
	_, found := keeper.GetAuthorization(ctx)
	require.False(t, found)
}

func TestHasPermission(t *testing.T) {
	granted := sample.AccAddressBytes()
	other := sample.AccAddressBytes()

	t.Run("no authorization set", func(t *testing.T) {
		keeper, ctx := keepertest.NftadminKeeper(t)
		require.False(t, keeper.HasPermission(ctx, "minter", granted))
	})

	t.Run("nil permissions", func(t *testing.T) {
		keeper, ctx := keepertest.NftadminKeeper(t)
		keeper.SetAuthorization(ctx, types.Authorization{
			RootAdmin: sample.AccAddress(),
		})
		require.False(t, keeper.HasPermission(ctx, "minter", granted))
	})

	t.Run("permission name not found", func(t *testing.T) {
		keeper, ctx := keepertest.NftadminKeeper(t)
		keeper.SetAuthorization(ctx, types.Authorization{
			RootAdmin: sample.AccAddress(),
			Permissions: []*types.Permission{
				{Name: "minter", Addresses: []string{granted.String()}},
			},
		})
		require.False(t, keeper.HasPermission(ctx, "burner", granted))
	})

	t.Run("address not in list", func(t *testing.T) {
		keeper, ctx := keepertest.NftadminKeeper(t)
		keeper.SetAuthorization(ctx, types.Authorization{
			RootAdmin: sample.AccAddress(),
			Permissions: []*types.Permission{
				{Name: "minter", Addresses: []string{granted.String()}},
			},
		})
		require.False(t, keeper.HasPermission(ctx, "minter", other))
	})

	t.Run("address has permission", func(t *testing.T) {
		keeper, ctx := keepertest.NftadminKeeper(t)
		keeper.SetAuthorization(ctx, types.Authorization{
			RootAdmin: sample.AccAddress(),
			Permissions: []*types.Permission{
				{Name: "minter", Addresses: []string{granted.String()}},
			},
		})
		require.True(t, keeper.HasPermission(ctx, "minter", granted))
	})
}
