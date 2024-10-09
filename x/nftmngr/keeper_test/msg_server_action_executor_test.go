package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestActionExecutorMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateActionExecutor{Creator: creator,
			NftSchemaCode:   strconv.Itoa(i),
			ExecutorAddress: strconv.Itoa(i),
		}
		_, err := srv.CreateActionExecutor(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetActionExecutor(ctx,
			expected.NftSchemaCode,
			expected.ExecutorAddress,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestActionExecutorMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateActionExecutor
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateActionExecutor{Creator: creator,
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateActionExecutor{Creator: "B",
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateActionExecutor{Creator: creator,
				NftSchemaCode:   strconv.Itoa(100000),
				ExecutorAddress: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateActionExecutor{Creator: creator,
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			}
			_, err := srv.CreateActionExecutor(wctx, expected)
			require.NoError(t, err)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetActionExecutor(ctx,
					expected.NftSchemaCode,
					expected.ExecutorAddress,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestActionExecutorMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteActionExecutor
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteActionExecutor{Creator: creator,
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteActionExecutor{Creator: "B",
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteActionExecutor{Creator: creator,
				NftSchemaCode:   strconv.Itoa(100000),
				ExecutorAddress: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateActionExecutor(wctx, &types.MsgCreateActionExecutor{Creator: creator,
				NftSchemaCode:   strconv.Itoa(0),
				ExecutorAddress: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteActionExecutor(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetActionExecutor(ctx,
					tc.request.NftSchemaCode,
					tc.request.ExecutorAddress,
				)
				require.False(t, found)
			}
		})
	}
}
