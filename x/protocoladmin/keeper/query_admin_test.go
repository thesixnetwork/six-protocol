package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAdminQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ProtocoladminKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAdmin(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAdminRequest
		response *types.QueryGetAdminResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAdminRequest{
				Group: msgs[0].Group,
				Admin: msgs[0].Admin,
			},
			response: &types.QueryGetAdminResponse{Admin: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAdminRequest{
				Group: msgs[1].Group,
				Admin: msgs[1].Admin,
			},
			response: &types.QueryGetAdminResponse{Admin: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAdminRequest{
				Group: strconv.Itoa(100000),
				Admin: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Admin(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAdminQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ProtocoladminKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAdmin(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAdminRequest {
		return &types.QueryAllAdminRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AdminAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Admin), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Admin),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AdminAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Admin), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Admin),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AdminAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Admin),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AdminAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
