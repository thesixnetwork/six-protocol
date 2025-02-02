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

func TestActionByRefIdQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActionByRefId(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetActionByRefIdRequest
		response *types.QueryGetActionByRefIdResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetActionByRefIdRequest{
				RefId: msgs[0].RefId,
			},
			response: &types.QueryGetActionByRefIdResponse{ActionByRefId: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetActionByRefIdRequest{
				RefId: msgs[1].RefId,
			},
			response: &types.QueryGetActionByRefIdResponse{ActionByRefId: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetActionByRefIdRequest{
				RefId: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ActionByRefId(wctx, tc.request)
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

func TestActionByRefIdQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActionByRefId(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActionByRefIdRequest {
		return &types.QueryAllActionByRefIdRequest{
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
			resp, err := keeper.ActionByRefIdAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionByRefId), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionByRefId),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActionByRefIdAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionByRefId), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionByRefId),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActionByRefIdAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ActionByRefId),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActionByRefIdAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestActionOfSchemaQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActionOfSchema(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetActionOfSchemaRequest
		response *types.QueryGetActionOfSchemaResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetActionOfSchemaRequest{
				NftSchemaCode: msgs[0].NftSchemaCode,
				Name:          msgs[0].Name,
			},
			response: &types.QueryGetActionOfSchemaResponse{ActionOfSchema: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetActionOfSchemaRequest{
				NftSchemaCode: msgs[1].NftSchemaCode,
				Name:          msgs[1].Name,
			},
			response: &types.QueryGetActionOfSchemaResponse{ActionOfSchema: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetActionOfSchemaRequest{
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
			response, err := keeper.ActionOfSchema(wctx, tc.request)
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

func TestActionOfSchemaQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActionOfSchema(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActionOfSchemaRequest {
		return &types.QueryAllActionOfSchemaRequest{
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
			resp, err := keeper.ActionOfSchemaAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionOfSchema), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionOfSchema),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActionOfSchemaAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionOfSchema), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionOfSchema),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActionOfSchemaAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ActionOfSchema),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActionOfSchemaAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
