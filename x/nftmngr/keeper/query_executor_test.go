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

func TestActionExecutorQuerySingle(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNActionExecutor(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetActionExecutorRequest
		response *types.QueryGetActionExecutorResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetActionExecutorRequest{NftSchemaCode: msgs[0].NftSchemaCode, ExecutorAddress: msgs[0].ExecutorAddress},
			response: &types.QueryGetActionExecutorResponse{ActionExecutor: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetActionExecutorRequest{NftSchemaCode: msgs[1].NftSchemaCode, ExecutorAddress: msgs[1].ExecutorAddress},
			response: &types.QueryGetActionExecutorResponse{ActionExecutor: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetActionExecutorRequest{NftSchemaCode: strconv.Itoa(100000), ExecutorAddress: strconv.Itoa(100000)},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.ActionExecutor(ctx, tc.request)
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

func TestActionExecutorQueryPaginated(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNActionExecutor(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActionExecutorRequest {
		return &types.QueryAllActionExecutorRequest{
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
			resp, err := k.ActionExecutorAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionExecutor), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionExecutor),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := k.ActionExecutorAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionExecutor), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionExecutor),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.ActionExecutorAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ActionExecutor),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.ActionExecutorAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestExecutorOfSchemaQuerySingle(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNExecutorOfSchema(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetExecutorOfSchemaRequest
		response *types.QueryGetExecutorOfSchemaResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetExecutorOfSchemaRequest{NftSchemaCode: msgs[0].NftSchemaCode},
			response: &types.QueryGetExecutorOfSchemaResponse{ExecutorOfSchema: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetExecutorOfSchemaRequest{NftSchemaCode: msgs[1].NftSchemaCode},
			response: &types.QueryGetExecutorOfSchemaResponse{ExecutorOfSchema: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetExecutorOfSchemaRequest{NftSchemaCode: strconv.Itoa(100000)},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.ExecutorOfSchema(ctx, tc.request)
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

func TestExecutorOfSchemaQueryPaginated(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNExecutorOfSchema(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllExecutorOfSchemaRequest {
		return &types.QueryAllExecutorOfSchemaRequest{
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
			resp, err := k.ExecutorOfSchemaAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ExecutorOfSchema), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ExecutorOfSchema),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := k.ExecutorOfSchemaAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ExecutorOfSchema), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ExecutorOfSchema),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.ExecutorOfSchemaAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ExecutorOfSchema),
		)
	})
}
