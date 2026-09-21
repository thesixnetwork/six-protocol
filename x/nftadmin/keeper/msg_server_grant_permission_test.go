package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/nftadmin/types"
)

func TestMsgServerGrantPermission(t *testing.T) {
	rootAdmin := sample.AccAddress()
	grantee := sample.AccAddress()

	t.Run("authorization not found", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.GrantPermission(ctx, &types.MsgGrantPermission{
			Creator: rootAdmin,
			Name:    "minter",
			Grantee: grantee,
		})
		require.ErrorIs(t, err, types.ErrAuthorizationNotFound)
	})

	t.Run("unauthorized creator", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetAuthorization(ctx, types.Authorization{RootAdmin: rootAdmin})
		_, err := ms.GrantPermission(ctx, &types.MsgGrantPermission{
			Creator: sample.AccAddress(),
			Name:    "minter",
			Grantee: grantee,
		})
		require.ErrorIs(t, err, types.ErrUnauthorized)
	})

	t.Run("invalid grantee", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetAuthorization(ctx, types.Authorization{RootAdmin: rootAdmin})
		_, err := ms.GrantPermission(ctx, &types.MsgGrantPermission{
			Creator: rootAdmin,
			Name:    "minter",
			Grantee: "invalid_address",
		})
		require.ErrorIs(t, err, types.ErrInvalidGrantee)
	})

	t.Run("grant with nil permissions creates permission", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetAuthorization(ctx, types.Authorization{RootAdmin: rootAdmin})
		resp, err := ms.GrantPermission(ctx, &types.MsgGrantPermission{
			Creator: rootAdmin,
			Name:    "minter",
			Grantee: grantee,
		})
		require.NoError(t, err)
		require.Equal(t, grantee, resp.Grantee)

		auth, found := k.GetAuthorization(ctx)
		require.True(t, found)
		require.Equal(t, []string{grantee}, auth.GetPermissionAddressByKey("minter"))
	})

	t.Run("grant appends to existing permission", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		existing := sample.AccAddress()
		k.SetAuthorization(ctx, types.Authorization{
			RootAdmin: rootAdmin,
			Permissions: []*types.Permission{
				{Name: "minter", Addresses: []string{existing}},
			},
		})
		_, err := ms.GrantPermission(ctx, &types.MsgGrantPermission{
			Creator: rootAdmin,
			Name:    "minter",
			Grantee: grantee,
		})
		require.NoError(t, err)

		auth, found := k.GetAuthorization(ctx)
		require.True(t, found)
		require.Equal(t, []string{existing, grantee}, auth.GetPermissionAddressByKey("minter"))
	})

	t.Run("grant creates new permission name", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		existing := sample.AccAddress()
		k.SetAuthorization(ctx, types.Authorization{
			RootAdmin: rootAdmin,
			Permissions: []*types.Permission{
				{Name: "minter", Addresses: []string{existing}},
			},
		})
		_, err := ms.GrantPermission(ctx, &types.MsgGrantPermission{
			Creator: rootAdmin,
			Name:    "burner",
			Grantee: grantee,
		})
		require.NoError(t, err)

		auth, found := k.GetAuthorization(ctx)
		require.True(t, found)
		require.Equal(t, []string{existing}, auth.GetPermissionAddressByKey("minter"))
		require.Equal(t, []string{grantee}, auth.GetPermissionAddressByKey("burner"))
	})

	t.Run("grantee already exists", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetAuthorization(ctx, types.Authorization{
			RootAdmin: rootAdmin,
			Permissions: []*types.Permission{
				{Name: "minter", Addresses: []string{grantee}},
			},
		})
		_, err := ms.GrantPermission(ctx, &types.MsgGrantPermission{
			Creator: rootAdmin,
			Name:    "minter",
			Grantee: grantee,
		})
		require.ErrorIs(t, err, types.ErrGranteeAlreadyExists)
	})
}
