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

func TestVirtualSchemaProposalQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVirtualSchemaProposal(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetVirtualSchemaProposalRequest
		response *types.QueryGetVirtualSchemaProposalResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVirtualSchemaProposalRequest{
				Index: msgs[0].Id,
			},
			response: &types.QueryGetVirtualSchemaProposalResponse{VirtualSchemaProposal: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVirtualSchemaProposalRequest{
				Index: msgs[1].Id,
			},
			response: &types.QueryGetVirtualSchemaProposalResponse{VirtualSchemaProposal: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVirtualSchemaProposalRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VirtualSchemaProposal(wctx, tc.request)
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
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVirtualSchemaProposal(keeper, ctx, 5)

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
			resp, err := keeper.VirtualSchemaProposalAll(wctx, request(nil, uint64(i), uint64(step), false))
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
			resp, err := keeper.VirtualSchemaProposalAll(wctx, request(next, 0, uint64(step), false))
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
		resp, err := keeper.VirtualSchemaProposalAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.VirtualSchemaProposal),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.VirtualSchemaProposalAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

// Prevent strconv unused error
var _ = strconv.IntSize

func TestActiveDislabeVirtualSchemaProposalQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActiveDislabeVirtualSchemaProposal(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetActiveDislabeVirtualSchemaProposalRequest
		response *types.QueryGetActiveDislabeVirtualSchemaProposalResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetActiveDislabeVirtualSchemaProposalRequest{
				Index: msgs[0].Id,
			},
			response: &types.QueryGetActiveDislabeVirtualSchemaProposalResponse{ActiveDislabeVirtualSchemaProposal: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetActiveDislabeVirtualSchemaProposalRequest{
				Index: msgs[1].Id,
			},
			response: &types.QueryGetActiveDislabeVirtualSchemaProposalResponse{ActiveDislabeVirtualSchemaProposal: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetActiveDislabeVirtualSchemaProposalRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ActiveDislabeVirtualSchemaProposal(wctx, tc.request)
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

func TestActiveDislabeVirtualSchemaProposalQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActiveDislabeVirtualSchemaProposal(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActiveDislabeVirtualSchemaProposalRequest {
		return &types.QueryAllActiveDislabeVirtualSchemaProposalRequest{
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
			resp, err := keeper.ActiveDislabeVirtualSchemaProposalAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActiveDislabeVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActiveDislabeVirtualSchemaProposal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActiveDislabeVirtualSchemaProposalAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActiveDislabeVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActiveDislabeVirtualSchemaProposal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActiveDislabeVirtualSchemaProposalAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ActiveDislabeVirtualSchemaProposal),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActiveDislabeVirtualSchemaProposalAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
