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

func TestSchemaAttributeQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSchemaAttribute(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSchemaAttributeRequest
		response *types.QueryGetSchemaAttributeResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSchemaAttributeRequest{
				NftSchemaCode: msgs[0].NftSchemaCode,
				Name:          msgs[0].Name,
			},
			response: &types.QueryGetSchemaAttributeResponse{SchemaAttribute: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSchemaAttributeRequest{
				NftSchemaCode: msgs[1].NftSchemaCode,
				Name:          msgs[1].Name,
			},
			response: &types.QueryGetSchemaAttributeResponse{SchemaAttribute: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSchemaAttributeRequest{
				NftSchemaCode: strconv.Itoa(100000),
				Name:          strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SchemaAttribute(wctx, tc.request)
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

func TestSchemaAttributeQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSchemaAttribute(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSchemaAttributeRequest {
		return &types.QueryAllSchemaAttributeRequest{
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
			resp, err := keeper.SchemaAttributeAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SchemaAttribute), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SchemaAttribute),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SchemaAttributeAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SchemaAttribute), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SchemaAttribute),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SchemaAttributeAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.SchemaAttribute),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SchemaAttributeAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
