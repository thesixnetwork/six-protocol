package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestActionSignerMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateActionSigner{
			Creator:                      creator,
			Base64EncodedSetSignerAction: strconv.Itoa(i),
		}
		_, err := srv.CreateActionSigner(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetActionSigner(ctx,
			creator,
			expected.Base64EncodedSetSignerAction,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.ActorAddress)
	}
}

func TestActionSignerMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateActionSigner
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateActionSigner{
				Creator:                      creator,
				Base64EncodedSetSignerAction: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateActionSigner{
				Creator:                      "B",
				Base64EncodedSetSignerAction: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateActionSigner{
				Creator:                      creator,
				Base64EncodedSetSignerAction: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftoracleKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateActionSigner{
				Creator:                      creator,
				Base64EncodedSetSignerAction: strconv.Itoa(0),
			}
			_, err := srv.CreateActionSigner(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateActionSigner(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetActionSigner(ctx,
					creator,
					expected.Base64EncodedSetSignerAction,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.ActorAddress)
			}
		})
	}
}

func TestActionSignerMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteActionSigner
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteActionSigner{
				Creator:                      creator,
				Base64EncodedSetSignerAction: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteActionSigner{
				Creator:                      "B",
				Base64EncodedSetSignerAction: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteActionSigner{
				Creator:                      creator,
				Base64EncodedSetSignerAction: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftoracleKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateActionSigner(wctx, &types.MsgCreateActionSigner{
				Creator:                      creator,
				Base64EncodedSetSignerAction: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteActionSigner(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetActionSigner(ctx,
					tc.request.Creator,
					tc.request.Base64EncodedSetSignerAction,
				)
				require.False(t, found)
			}
		})
	}
}
