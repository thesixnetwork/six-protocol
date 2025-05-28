package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestActionSignerConfigQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	msgs := createNActionSignerConfig(&keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetActionSignerConfigRequest
		response *types.QueryGetActionSignerConfigResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetActionSignerConfigRequest{
				Chain: msgs[0].Chain,
			},
			response: &types.QueryGetActionSignerConfigResponse{ActionSignerConfig: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetActionSignerConfigRequest{
				Chain: msgs[1].Chain,
			},
			response: &types.QueryGetActionSignerConfigResponse{ActionSignerConfig: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetActionSignerConfigRequest{
				Chain: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ActionSignerConfig(ctx, tc.request)
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

func TestActionSignerConfigQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActionSignerConfig(&keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActionSignerConfigRequest {
		return &types.QueryAllActionSignerConfigRequest{
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
			resp, err := keeper.ActionSignerConfigAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionSignerConfig), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionSignerConfig),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActionSignerConfigAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionSignerConfig), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionSignerConfig),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActionSignerConfigAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ActionSignerConfig),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActionSignerConfigAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
