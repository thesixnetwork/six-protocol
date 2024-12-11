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

func TestVirtualSchemaMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateVirtualSchema{Creator: creator,
			VirtualNftSchemaCode: strconv.Itoa(i),
		}
		_, err := srv.CreateVirtualSchema(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetVirtualSchema(ctx,
			expected.VirtualNftSchemaCode,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.VirtualNftSchemaCode)
	}
}

func TestVirtualSchemaMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteVirtualSchema
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteVirtualSchema{Creator: creator,
				VirtualNftSchemaCode: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteVirtualSchema{Creator: "B",
				VirtualNftSchemaCode: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteVirtualSchema{Creator: creator,
				VirtualNftSchemaCode: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateVirtualSchema(wctx, &types.MsgCreateVirtualSchema{Creator: creator,
				VirtualNftSchemaCode: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteVirtualSchema(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetVirtualSchema(ctx,
					tc.request.VirtualNftSchemaCode,
				)
				require.False(t, found)
			}
		})
	}
}
