package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	"github.com/cosmos/cosmos-sdk/types/query"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMintpermQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	msgs := createNMintperm(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetMintpermRequest
		response *types.QueryGetMintpermResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMintpermRequest{
				Token:   msgs[0].Token,
				Address: msgs[0].Address,
			},
			response: &types.QueryGetMintpermResponse{Mintperm: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMintpermRequest{
				Token:   msgs[1].Token,
				Address: msgs[1].Address,
			},
			response: &types.QueryGetMintpermResponse{Mintperm: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMintpermRequest{
				Token:   strconv.Itoa(100000),
				Address: strconv.Itoa(100000),
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
			response, err := keeper.Mintperm(ctx, tc.request)
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

func TestMintpermQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	msgs := createNMintperm(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMintpermRequest {
		return &types.QueryAllMintpermRequest{
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
			resp, err := keeper.MintpermAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Mintperm), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Mintperm),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MintpermAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Mintperm), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Mintperm),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MintpermAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Mintperm),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MintpermAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
