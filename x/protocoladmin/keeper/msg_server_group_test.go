package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/keeper"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGroupMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ProtocoladminKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	owner := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateGroup{
			Creator: owner,
			Name:    strconv.Itoa(i),
		}
		_, err := srv.CreateGroup(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetGroup(ctx,
			expected.Name,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Owner)
	}
}

func TestGroupMsgServerUpdate(t *testing.T) {
	owner := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateGroup
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateGroup{
				Creator: owner,
				Name:    strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateGroup{
				Creator: "B",
				Name:    strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateGroup{
				Creator: owner,
				Name:    strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ProtocoladminKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateGroup{
				Creator: owner,
				Name:    strconv.Itoa(0),
			}
			_, err := srv.CreateGroup(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateGroup(ctx, tc.request)
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
	owner := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteGroup
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteGroup{
				Creator: owner,
				Name:    strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteGroup{
				Creator: "B",
				Name:    strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteGroup{
				Creator: owner,
				Name:    strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ProtocoladminKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateGroup(ctx, &types.MsgCreateGroup{
				Creator: owner,
				Name:    strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteGroup(ctx, tc.request)
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
