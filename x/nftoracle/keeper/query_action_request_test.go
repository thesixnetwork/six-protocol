package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func TestActionRequestQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	msgs := createNActionRequest(&keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetActionRequestRequest
		response *types.QueryGetActionRequestResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetActionRequestRequest{Id: msgs[0].Id},
			response: &types.QueryGetActionRequestResponse{ActionOracleRequest: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetActionRequestRequest{Id: msgs[1].Id},
			response: &types.QueryGetActionRequestResponse{ActionOracleRequest: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetActionRequestRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ActionOracleRequest(ctx, tc.request)
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

func TestActionRequestQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	msgs := createNActionRequest(&keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActionRequestRequest {
		return &types.QueryAllActionRequestRequest{
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
			resp, err := keeper.ActionRequestAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionOracleRequest), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionOracleRequest),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActionRequestAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionOracleRequest), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionOracleRequest),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActionRequestAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ActionOracleRequest),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActionRequestAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
