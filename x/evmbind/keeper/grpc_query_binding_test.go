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
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestBindingQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EvmbindKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBinding(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetBindingRequest
		response *types.QueryGetBindingResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetBindingRequest{
				EthAddress: msgs[0].EthAddress,
			},
			response: &types.QueryGetBindingResponse{Binding: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetBindingRequest{
				EthAddress: msgs[1].EthAddress,
			},
			response: &types.QueryGetBindingResponse{Binding: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetBindingRequest{
				EthAddress: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Binding(wctx, tc.request)
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

func TestBindingQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EvmbindKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNBinding(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllBindingRequest {
		return &types.QueryAllBindingRequest{
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
			resp, err := keeper.BindingAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Binding), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Binding),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.BindingAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Binding), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Binding),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.BindingAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Binding),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.BindingAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
