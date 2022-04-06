package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func TestOptionsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.TokenmngrKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateOptions{Creator: creator}
	_, err := srv.CreateOptions(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetOptions(ctx)
	require.True(t, found)
	require.Equal(t, expected.DefaultMintee, rst.DefaultMintee)
}

func TestOptionsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateOptions
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateOptions{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateOptions{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TokenmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateOptions{Creator: creator}
			_, err := srv.CreateOptions(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateOptions(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetOptions(ctx)
				require.True(t, found)
				require.Equal(t, expected.DefaultMintee, rst.DefaultMintee)
			}
		})
	}
}

func TestOptionsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteOptions
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteOptions{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteOptions{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TokenmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateOptions(wctx, &types.MsgCreateOptions{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteOptions(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetOptions(ctx)
				require.False(t, found)
			}
		})
	}
}
