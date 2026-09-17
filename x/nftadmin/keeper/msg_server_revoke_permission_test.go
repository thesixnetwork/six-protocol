package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/nftadmin/types"
)

func TestMsgServerRevokePermission(t *testing.T) {
	rootAdmin := sample.AccAddress()
	revokee := sample.AccAddress()

	t.Run("authorization not found", func(t *testing.T) {
		_, ms, ctx := setupMsgServer(t)
		_, err := ms.RevokePermission(ctx, &types.MsgRevokePermission{
			Creator: rootAdmin,
			Name:    "minter",
			Revokee: revokee,
		})
		require.ErrorIs(t, err, types.ErrAuthorizationNotFound)
	})

	t.Run("unauthorized creator", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetAuthorization(ctx, types.Authorization{RootAdmin: rootAdmin})
		_, err := ms.RevokePermission(ctx, &types.MsgRevokePermission{
			Creator: sample.AccAddress(),
			Name:    "minter",
			Revokee: revokee,
		})
		require.ErrorIs(t, err, types.ErrUnauthorized)
	})

	t.Run("nil permissions", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetAuthorization(ctx, types.Authorization{RootAdmin: rootAdmin})
		_, err := ms.RevokePermission(ctx, &types.MsgRevokePermission{
			Creator: rootAdmin,
			Name:    "minter",
			Revokee: revokee,
		})
		require.ErrorIs(t, err, types.ErrNoPermissions)
	})

	t.Run("invalid revokee", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetAuthorization(ctx, types.Authorization{
			RootAdmin: rootAdmin,
			Permissions: []*types.Permission{
				{Name: "minter", Addresses: []string{revokee}},
			},
		})
		_, err := ms.RevokePermission(ctx, &types.MsgRevokePermission{
			Creator: rootAdmin,
			Name:    "minter",
			Revokee: "invalid_address",
		})
		require.ErrorIs(t, err, types.ErrInvalidGrantee)
	})

	t.Run("no permissions for name", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetAuthorization(ctx, types.Authorization{
			RootAdmin: rootAdmin,
			Permissions: []*types.Permission{
				{Name: "minter", Addresses: []string{revokee}},
			},
		})
		_, err := ms.RevokePermission(ctx, &types.MsgRevokePermission{
			Creator: rootAdmin,
			Name:    "burner",
			Revokee: revokee,
		})
		require.ErrorIs(t, err, types.ErrNoPermissionsForName)
	})

	t.Run("revokee not found for name", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetAuthorization(ctx, types.Authorization{
			RootAdmin: rootAdmin,
			Permissions: []*types.Permission{
				{Name: "minter", Addresses: []string{sample.AccAddress()}},
			},
		})
		_, err := ms.RevokePermission(ctx, &types.MsgRevokePermission{
			Creator: rootAdmin,
			Name:    "minter",
			Revokee: revokee,
		})
		require.ErrorIs(t, err, types.ErrGranteeNotFoundForName)
	})

	t.Run("revoke removes address", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		other := sample.AccAddress()
		k.SetAuthorization(ctx, types.Authorization{
			RootAdmin: rootAdmin,
			Permissions: []*types.Permission{
				{Name: "minter", Addresses: []string{revokee, other}},
			},
		})
		resp, err := ms.RevokePermission(ctx, &types.MsgRevokePermission{
			Creator: rootAdmin,
			Name:    "minter",
			Revokee: revokee,
		})
		require.NoError(t, err)
		require.Equal(t, revokee, resp.Revokee)

		auth, found := k.GetAuthorization(ctx)
		require.True(t, found)
		require.Equal(t, []string{other}, auth.GetPermissionAddressByKey("minter"))
	})

	t.Run("revoke last address removes permission", func(t *testing.T) {
		k, ms, ctx := setupMsgServer(t)
		k.SetAuthorization(ctx, types.Authorization{
			RootAdmin: rootAdmin,
			Permissions: []*types.Permission{
				{Name: "minter", Addresses: []string{revokee}},
			},
		})
		_, err := ms.RevokePermission(ctx, &types.MsgRevokePermission{
			Creator: rootAdmin,
			Name:    "minter",
			Revokee: revokee,
		})
		require.NoError(t, err)

		auth, found := k.GetAuthorization(ctx)
		require.True(t, found)
		require.Nil(t, auth.GetPermissionAddressByKey("minter"))
	})
}
