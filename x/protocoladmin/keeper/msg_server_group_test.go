package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/keeper"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGroupMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ProtocoladminKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	owner := "B"
	expected := &types.MsgCreateGroup{Creator: owner,
		Name: "test.admin",
	}
	_, err := srv.CreateGroup(wctx, expected)
	require.Error(t, err, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator is not super admin"))
}

func TestGroupMsgServerCreateWithSuperAdmin(t *testing.T) {
	k, ctx := keepertest.ProtocoladminKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	owner := "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateGroup{Creator: owner,
			Name: strconv.Itoa(i),
		}
		_, err := srv.CreateGroup(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetGroup(ctx,
			expected.Name,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Owner)
	}
}

func TestGroupMsgServerUpdate(t *testing.T) {
	owner := "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateGroup
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateGroup{Creator: owner,
				Name: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateGroup{Creator: "B",
				Name: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateGroup{Creator: owner,
				Name: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ProtocoladminKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateGroup{Creator: owner,
				Name: strconv.Itoa(0),
			}
			_, err := srv.CreateGroup(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateGroup(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetGroup(ctx,
					expected.Name,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Owner)
			}
		})
	}
}

func TestGroupMsgServerDelete(t *testing.T) {
	owner := "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteGroup
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteGroup{Creator: owner,
				Name: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteGroup{Creator: "B",
				Name: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteGroup{Creator: owner,
				Name: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ProtocoladminKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateGroup(wctx, &types.MsgCreateGroup{Creator: owner,
				Name: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteGroup(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetGroup(ctx,
					tc.request.Name,
				)
				require.False(t, found)
			}
		})
	}
}
