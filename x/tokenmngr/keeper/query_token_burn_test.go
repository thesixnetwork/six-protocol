package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/types/query"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTokenBurnQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	msgs := createNTokenBurn(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetTokenBurnRequest
		response *types.QueryGetTokenBurnResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTokenBurnRequest{
				Index: msgs[0].Amount.Denom,
			},
			response: &types.QueryGetTokenBurnResponse{TokenBurn: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTokenBurnRequest{
				Index: msgs[1].Amount.Denom,
			},
			response: &types.QueryGetTokenBurnResponse{TokenBurn: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTokenBurnRequest{
				Index: strconv.Itoa(100000),
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
			response, err := keeper.TokenBurn(ctx, tc.request)
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

func TestTokenBurnQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	msgs := createNTokenBurn(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTokenBurnRequest {
		return &types.QueryAllTokenBurnRequest{
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
			resp, err := keeper.TokenBurnAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TokenBurn), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TokenBurn),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TokenBurnAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TokenBurn), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TokenBurn),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TokenBurnAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.TokenBurn),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TokenBurnAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
