package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"
)

func TestMsgServerCreateGroup(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	creator := sample.AccAddress()

	// creator is not super admin
	_, err := ms.CreateGroup(ctx, &types.MsgCreateGroup{Creator: creator, Name: "group1"})
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)

	// promote creator to super admin
	k.SetAdmin(ctx, types.Admin{Group: keeper.SUPER_ADMIN, Admin: creator})

	_, err = ms.CreateGroup(ctx, &types.MsgCreateGroup{Creator: creator, Name: "group1"})
	require.NoError(t, err)

	group, found := k.GetGroup(ctx, "group1")
	require.True(t, found)
	require.Equal(t, types.Group{Owner: creator, Name: "group1"}, group)

	// group index already set
	_, err = ms.CreateGroup(ctx, &types.MsgCreateGroup{Creator: creator, Name: "group1"})
	require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
}

func TestMsgServerUpdateGroup(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	creator := sample.AccAddress()

	// operation is disabled, even for super admins
	k.SetAdmin(ctx, types.Admin{Group: keeper.SUPER_ADMIN, Admin: creator})

	_, err := ms.UpdateGroup(ctx, &types.MsgUpdateGroup{Creator: creator, Name: "group1"})
	require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)
	require.Contains(t, err.Error(), "operation not available")
}

func TestMsgServerDeleteGroup(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	creator := sample.AccAddress()

	// creator is not super admin
	_, err := ms.DeleteGroup(ctx, &types.MsgDeleteGroup{Creator: creator, Name: "group1"})
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)

	// promote creator to super admin
	k.SetAdmin(ctx, types.Admin{Group: keeper.SUPER_ADMIN, Admin: creator})

	// group does not exist
	_, err = ms.DeleteGroup(ctx, &types.MsgDeleteGroup{Creator: creator, Name: "group1"})
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)

	k.SetGroup(ctx, types.Group{Owner: creator, Name: "group1"})

	_, err = ms.DeleteGroup(ctx, &types.MsgDeleteGroup{Creator: creator, Name: "group1"})
	require.NoError(t, err)

	_, found := k.GetGroup(ctx, "group1")
	require.False(t, found)
}
