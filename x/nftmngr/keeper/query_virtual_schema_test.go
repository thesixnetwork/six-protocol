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

func TestVirtualSchemaQuerySingle(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNVirtualSchema(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetVirtualSchemaRequest
		response *types.QueryGetVirtualSchemaResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetVirtualSchemaRequest{NftSchemaCode: msgs[0].VirtualNftSchemaCode},
			response: &types.QueryGetVirtualSchemaResponse{VirtualSchema: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetVirtualSchemaRequest{NftSchemaCode: msgs[1].VirtualNftSchemaCode},
			response: &types.QueryGetVirtualSchemaResponse{VirtualSchema: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetVirtualSchemaRequest{NftSchemaCode: strconv.Itoa(100000)},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.VirtualSchema(ctx, tc.request)
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
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNVirtualSchema(k, ctx, 5)

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
			resp, err := k.VirtualSchemaAll(ctx, request(nil, uint64(i), uint64(step), false))
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
			resp, err := k.VirtualSchemaAll(ctx, request(next, 0, uint64(step), false))
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
		resp, err := k.VirtualSchemaAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.VirtualSchema),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.VirtualSchemaAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestVirtualSchemaProposalQuerySingle(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNVirtualSchemaProposal(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetVirtualSchemaProposalRequest
		response *types.QueryGetVirtualSchemaProposalResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetVirtualSchemaProposalRequest{Index: msgs[0].Id},
			response: &types.QueryGetVirtualSchemaProposalResponse{VirtualSchemaProposal: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetVirtualSchemaProposalRequest{Index: msgs[1].Id},
			response: &types.QueryGetVirtualSchemaProposalResponse{VirtualSchemaProposal: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetVirtualSchemaProposalRequest{Index: strconv.Itoa(100000)},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.VirtualSchemaProposal(ctx, tc.request)
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

func TestVirtualSchemaProposalQueryPaginated(t *testing.T) {
	k, ctx := keepertest.NftmngrKeeper(t)
	msgs := createNVirtualSchemaProposal(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVirtualSchemaProposalRequest {
		return &types.QueryAllVirtualSchemaProposalRequest{
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
			resp, err := k.VirtualSchemaProposalAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VirtualSchemaProposal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := k.VirtualSchemaProposalAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VirtualSchemaProposal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.VirtualSchemaProposalAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.VirtualSchemaProposal),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.VirtualSchemaProposalAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
