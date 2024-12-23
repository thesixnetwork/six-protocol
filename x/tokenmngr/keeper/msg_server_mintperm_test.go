package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMintpermMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.TokenmngrKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateMintperm{
			Creator: creator,
			Token:   strconv.Itoa(i),
			Address: strconv.Itoa(i),
		}
		_, err := srv.CreateMintperm(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetMintperm(ctx,
			expected.Token,
			expected.Address,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestMintpermMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateMintperm
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateMintperm{
				Creator: creator,
				Token:   strconv.Itoa(0),
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateMintperm{
				Creator: "B",
				Token:   strconv.Itoa(0),
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateMintperm{
				Creator: creator,
				Token:   strconv.Itoa(100000),
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TokenmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateMintperm{
				Creator: creator,
				Token:   strconv.Itoa(0),
				Address: strconv.Itoa(0),
			}
			_, err := srv.CreateMintperm(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateMintperm(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetMintperm(ctx,
					expected.Token,
					expected.Address,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestMintpermMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteMintperm
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteMintperm{
				Creator: creator,
				Token:   strconv.Itoa(0),
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteMintperm{
				Creator: "B",
				Token:   strconv.Itoa(0),
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteMintperm{
				Creator: creator,
				Token:   strconv.Itoa(100000),
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TokenmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateMintperm(wctx, &types.MsgCreateMintperm{
				Creator: creator,
				Token:   strconv.Itoa(0),
				Address: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteMintperm(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetMintperm(ctx,
					tc.request.Token,
					tc.request.Address,
				)
				require.False(t, found)
			}
		})
	}
}
