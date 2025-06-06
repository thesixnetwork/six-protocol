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
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestVirtualSchemaQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVirtualSchema(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetVirtualSchemaRequest
		response *types.QueryGetVirtualSchemaResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVirtualSchemaRequest{
				NftSchemaCode: msgs[0].VirtualNftSchemaCode,
			},
			response: &types.QueryGetVirtualSchemaResponse{VirtualSchema: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVirtualSchemaRequest{
				NftSchemaCode: msgs[1].VirtualNftSchemaCode,
			},
			response: &types.QueryGetVirtualSchemaResponse{VirtualSchema: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVirtualSchemaRequest{
				NftSchemaCode: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VirtualSchema(wctx, tc.request)
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

func TestVirtualSchemaQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVirtualSchema(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVirtualSchemaRequest {
		return &types.QueryAllVirtualSchemaRequest{
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
			resp, err := keeper.VirtualSchemaAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VirtualSchema), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VirtualSchema),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VirtualSchemaAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VirtualSchema), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VirtualSchema),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.VirtualSchemaAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.VirtualSchema),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.VirtualSchemaAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
