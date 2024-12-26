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

// TODO:: Feat(VirtualSchema)
func TestVirtualActionMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateVirtualAction{
			Creator:       creator,
			NftSchemaCode: strconv.Itoa(i),
			NewActions:    []*types.Action{},
		}
		_, err := srv.CreateVirtualAction(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetVirtualAction(ctx,
			expected.NftSchemaCode,
			"name",
		)
		// TODO:
		_ = rst.Name
		require.True(t, found)
		require.Equal(t, expected.Creator, "schemaOwner")
	}
}

// TODO:: Feat(VirtualSchema)
func TestVirtualActionMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateVirtualAction
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateVirtualAction{
				Creator:       creator,
				NftSchemaCode: strconv.Itoa(0),
				NewActions:    []*types.Action{},
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateVirtualAction{
				Creator:       "B",
				NftSchemaCode: strconv.Itoa(0),
				NewActions:    []*types.Action{},
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateVirtualAction{
				Creator:       creator,
				NftSchemaCode: strconv.Itoa(100000),
				NewActions:    []*types.Action{},
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateVirtualAction{
				Creator:       creator,
				NftSchemaCode: strconv.Itoa(0),
				NewActions:    []*types.Action{},
			}
			_, err := srv.CreateVirtualAction(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateVirtualAction(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetVirtualAction(ctx,
					expected.NftSchemaCode,
					"name",
				)
				_ = rst.Name
				require.True(t, found)
				require.Equal(t, expected.Creator, "schemaOwner")
			}
		})
	}
}

// TODO:: Feat(VirtualSchema)
func TestVirtualActionMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteVirtualAction
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteVirtualAction{
				Creator:       creator,
				NftSchemaCode: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteVirtualAction{
				Creator:       "B",
				NftSchemaCode: strconv.Itoa(0),
				Name:          strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteVirtualAction{
				Creator:       creator,
				NftSchemaCode: strconv.Itoa(100000),
				Name:          strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NftmngrKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateVirtualAction(wctx, &types.MsgCreateVirtualAction{
				Creator:       creator,
				NftSchemaCode: strconv.Itoa(0),
				NewActions:    []*types.Action{},
			})
			require.NoError(t, err)
			_, err = srv.DeleteVirtualAction(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetVirtualAction(ctx,
					tc.request.NftSchemaCode,
					tc.request.Name,
				)
				require.False(t, found)
			}
		})
	}
}
