package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/types/query"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGroupQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ProtocoladminKeeper(t)
	msgs := createNGroup(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetGroupRequest
		response *types.QueryGetGroupResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetGroupRequest{
				Name: msgs[0].Name,
			},
			response: &types.QueryGetGroupResponse{Group: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetGroupRequest{
				Name: msgs[1].Name,
			},
			response: &types.QueryGetGroupResponse{Group: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetGroupRequest{
				Name: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Group(ctx, tc.request)
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

func TestGroupQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ProtocoladminKeeper(t)
	msgs := createNGroup(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGroupRequest {
		return &types.QueryAllGroupRequest{
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
			resp, err := keeper.GroupAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Group), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Group),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GroupAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Group), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Group),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GroupAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Group),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GroupAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
