package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/evmbind/keeper"
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestBindingMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.EvmbindKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateBinding{Creator: creator,
			EthAddress: strconv.Itoa(i),
		}
		_, err := srv.CreateBinding(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetBinding(ctx,
			expected.EthAddress,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestBindingMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateBinding
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateBinding{Creator: creator,
				EthAddress: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateBinding{Creator: "B",
				EthAddress: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateBinding{Creator: creator,
				EthAddress: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EvmbindKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateBinding{Creator: creator,
				EthAddress: strconv.Itoa(0),
			}
			_, err := srv.CreateBinding(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateBinding(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetBinding(ctx,
					expected.EthAddress,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestBindingMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteBinding
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteBinding{Creator: creator,
				EthAddress: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteBinding{Creator: "B",
				EthAddress: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteBinding{Creator: creator,
				EthAddress: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EvmbindKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateBinding(wctx, &types.MsgCreateBinding{Creator: creator,
				EthAddress: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteBinding(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetBinding(ctx,
					tc.request.EthAddress,
				)
				require.False(t, found)
			}
		})
	}
}
