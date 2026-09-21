package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/types/query"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

func TestNFTSchemaQuerySingle(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNNFTSchema(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetNFTSchemaRequest
		response *types.QueryGetNFTSchemaResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetNFTSchemaRequest{Code: msgs[0].Code},
			response: &types.QueryGetNFTSchemaResponse{NFTSchema: msgs[0].ResultWithEmptyVirtualAction()},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetNFTSchemaRequest{Code: msgs[1].Code},
			response: &types.QueryGetNFTSchemaResponse{NFTSchema: msgs[1].ResultWithEmptyVirtualAction()},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetNFTSchemaRequest{Code: "missing"},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.NFTSchema(ctx, tc.request)
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

func TestNFTSchemaQueryPaginated(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNNFTSchema(k, ctx, 5)

	results := make([]types.NFTSchemaQueryResult, len(msgs))
	for i, msg := range msgs {
		results[i] = msg.ResultWithEmptyVirtualAction()
	}

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllNFTSchemaRequest {
		return &types.QueryAllNFTSchemaRequest{
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
		for i := 0; i < len(results); i += step {
			resp, err := k.NFTSchemaAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NFTSchema), step)
			require.Subset(t,
				nullify.Fill(results),
				nullify.Fill(resp.NFTSchema),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(results); i += step {
			resp, err := k.NFTSchemaAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NFTSchema), step)
			require.Subset(t,
				nullify.Fill(results),
				nullify.Fill(resp.NFTSchema),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.NFTSchemaAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(results), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(results),
			nullify.Fill(resp.NFTSchema),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.NFTSchemaAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
