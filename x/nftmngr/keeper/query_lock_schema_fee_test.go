package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/types/query"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

func TestLockSchemaFeeQuerySingle(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNLockSchemaFee(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetLockSchemaFeeRequest
		response *types.QueryGetLockSchemaFeeResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLockSchemaFeeRequest{Index: msgs[0].Id},
			response: &types.QueryGetLockSchemaFeeResponse{LockSchemaFee: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetLockSchemaFeeRequest{Index: msgs[1].Id},
			response: &types.QueryGetLockSchemaFeeResponse{LockSchemaFee: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetLockSchemaFeeRequest{Index: strconv.Itoa(100000)},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.LockSchemaFee(ctx, tc.request)
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

func TestLockSchemaFeeQueryPaginated(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNLockSchemaFee(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLockSchemaFeeRequest {
		return &types.QueryAllLockSchemaFeeRequest{
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
			resp, err := k.LockSchemaFeeAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LockSchemaFee), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LockSchemaFee),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := k.LockSchemaFeeAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LockSchemaFee), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LockSchemaFee),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.LockSchemaFeeAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.LockSchemaFee),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.LockSchemaFeeAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
