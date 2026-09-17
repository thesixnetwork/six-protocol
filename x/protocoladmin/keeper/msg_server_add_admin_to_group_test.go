package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"
)

func TestMsgServerAddAdminToGroup(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)

	owner := sample.AccAddress()
	superAdmin := sample.AccAddress()
	outsider := sample.AccAddress()

	k.SetAdmin(ctx, types.Admin{Group: keeper.SUPER_ADMIN, Admin: superAdmin})

	// group does not exist
	_, err := ms.AddAdminToGroup(ctx, &types.MsgAddAdminToGroup{
		Creator: owner, Name: "group1", Address: sample.AccAddress(),
	})
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)

	k.SetGroup(ctx, types.Group{Owner: owner, Name: "group1"})

	// creator is neither owner nor super admin
	_, err = ms.AddAdminToGroup(ctx, &types.MsgAddAdminToGroup{
		Creator: outsider, Name: "group1", Address: sample.AccAddress(),
	})
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)

	// owner can add an admin
	admin1 := sample.AccAddress()
	_, err = ms.AddAdminToGroup(ctx, &types.MsgAddAdminToGroup{
		Creator: owner, Name: "group1", Address: admin1,
	})
	require.NoError(t, err)

	stored, found := k.GetAdmin(ctx, "group1", admin1)
	require.True(t, found)
	require.Equal(t, types.Admin{Group: "group1", Admin: admin1}, stored)

	// admin already exists
	_, err = ms.AddAdminToGroup(ctx, &types.MsgAddAdminToGroup{
		Creator: owner, Name: "group1", Address: admin1,
	})
	require.ErrorIs(t, err, sdkerrors.ErrInvalidRequest)

	// super admin can add an admin as well
	admin2 := sample.AccAddress()
	_, err = ms.AddAdminToGroup(ctx, &types.MsgAddAdminToGroup{
		Creator: superAdmin, Name: "group1", Address: admin2,
	})
	require.NoError(t, err)

	_, found = k.GetAdmin(ctx, "group1", admin2)
	require.True(t, found)
}
