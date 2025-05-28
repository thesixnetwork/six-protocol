package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTokenMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.TokenmngrKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateToken{
			Creator: creator,
			Name:    strconv.Itoa(i),
		}
		_, err := srv.CreateToken(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetToken(ctx,
			expected.Name,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestTokenMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateToken
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateToken{
				Creator: creator,
				Name:    strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateToken{
				Creator: "B",
				Name:    strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateToken{
				Creator: creator,
				Name:    strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TokenmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateToken{
				Creator: creator,
				Name:    strconv.Itoa(0),
			}
			_, err := srv.CreateToken(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateToken(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetToken(ctx,
					expected.Name,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestTokenMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteToken
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteToken{
				Creator: creator,
				Name:    strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteToken{
				Creator: "B",
				Name:    strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteToken{
				Creator: creator,
				Name:    strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TokenmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateToken(ctx, &types.MsgCreateToken{
				Creator: creator,
				Name:    strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteToken(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetToken(ctx,
					tc.request.Name,
				)
				require.False(t, found)
			}
		})
	}
}
