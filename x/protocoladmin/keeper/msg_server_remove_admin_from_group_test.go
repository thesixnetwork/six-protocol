package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"
)

func TestMsgServerRemoveAdminFromGroup(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)

	owner := sample.AccAddress()
	superAdmin := sample.AccAddress()
	outsider := sample.AccAddress()
	admin1 := sample.AccAddress()
	admin2 := sample.AccAddress()

	k.SetAdmin(ctx, types.Admin{Group: keeper.SUPER_ADMIN, Admin: superAdmin})

	// group does not exist
	_, err := ms.RemoveAdminFromGroup(ctx, &types.MsgRemoveAdminFromGroup{
		Creator: owner, Name: "group1", Address: admin1,
	})
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)

	k.SetGroup(ctx, types.Group{Owner: owner, Name: "group1"})

	// admin does not exist in group
	_, err = ms.RemoveAdminFromGroup(ctx, &types.MsgRemoveAdminFromGroup{
		Creator: owner, Name: "group1", Address: admin1,
	})
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)

	k.SetAdmin(ctx, types.Admin{Group: "group1", Admin: admin1})
	k.SetAdmin(ctx, types.Admin{Group: "group1", Admin: admin2})

	// creator is neither owner nor super admin
	_, err = ms.RemoveAdminFromGroup(ctx, &types.MsgRemoveAdminFromGroup{
		Creator: outsider, Name: "group1", Address: admin1,
	})
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)

	// owner can remove an admin
	_, err = ms.RemoveAdminFromGroup(ctx, &types.MsgRemoveAdminFromGroup{
		Creator: owner, Name: "group1", Address: admin1,
	})
	require.NoError(t, err)

	_, found := k.GetAdmin(ctx, "group1", admin1)
	require.False(t, found)

	// super admin can remove an admin as well
	_, err = ms.RemoveAdminFromGroup(ctx, &types.MsgRemoveAdminFromGroup{
		Creator: superAdmin, Name: "group1", Address: admin2,
	})
	require.NoError(t, err)

	_, found = k.GetAdmin(ctx, "group1", admin2)
	require.False(t, found)
}
