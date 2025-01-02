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

func TestActiveDisableVirtualSchemaProposalQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActiveDisableVirtualSchemaProposal(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetActiveDisableVirtualSchemaProposalRequest
		response *types.QueryGetActiveDisableVirtualSchemaProposalResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetActiveDisableVirtualSchemaProposalRequest{
				Index: msgs[0].Id,
			},
			response: &types.QueryGetActiveDisableVirtualSchemaProposalResponse{ActiveDisableVirtualSchemaProposal: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetActiveDisableVirtualSchemaProposalRequest{
				Index: msgs[1].Id,
			},
			response: &types.QueryGetActiveDisableVirtualSchemaProposalResponse{ActiveDisableVirtualSchemaProposal: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetActiveDisableVirtualSchemaProposalRequest{
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
			response, err := keeper.ActiveDisableVirtualSchemaProposal(wctx, tc.request)
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

func TestActiveDisableVirtualSchemaProposalQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNActiveDisableVirtualSchemaProposal(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActiveDisableVirtualSchemaProposalRequest {
		return &types.QueryAllActiveDisableVirtualSchemaProposalRequest{
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
			resp, err := keeper.ActiveDisableVirtualSchemaProposalAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActiveDisableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActiveDisableVirtualSchemaProposal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ActiveDisableVirtualSchemaProposalAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActiveDisableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActiveDisableVirtualSchemaProposal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ActiveDisableVirtualSchemaProposalAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ActiveDisableVirtualSchemaProposal),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ActiveDisableVirtualSchemaProposalAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

// Prevent strconv unused error
var _ = strconv.IntSize

func TestInactiveEnableVirtualSchemaProposalQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNInactiveEnableVirtualSchemaProposal(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetInactiveEnableVirtualSchemaProposalRequest
		response *types.QueryGetInactiveEnableVirtualSchemaProposalResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetInactiveEnableVirtualSchemaProposalRequest{
				Index: msgs[0].Id,
			},
			response: &types.QueryGetInactiveEnableVirtualSchemaProposalResponse{InactiveEnableVirtualSchemaProposal: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetInactiveEnableVirtualSchemaProposalRequest{
				Index: msgs[1].Id,
			},
			response: &types.QueryGetInactiveEnableVirtualSchemaProposalResponse{InactiveEnableVirtualSchemaProposal: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetInactiveEnableVirtualSchemaProposalRequest{
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
			response, err := keeper.InactiveEnableVirtualSchemaProposal(wctx, tc.request)
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

func TestInactiveEnableVirtualSchemaProposalQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNInactiveEnableVirtualSchemaProposal(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllInactiveEnableVirtualSchemaProposalRequest {
		return &types.QueryAllInactiveEnableVirtualSchemaProposalRequest{
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
			resp, err := keeper.InactiveEnableVirtualSchemaProposalAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.InactiveEnableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.InactiveEnableVirtualSchemaProposal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.InactiveEnableVirtualSchemaProposalAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.InactiveEnableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.InactiveEnableVirtualSchemaProposal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.InactiveEnableVirtualSchemaProposalAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.InactiveEnableVirtualSchemaProposal),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.InactiveEnableVirtualSchemaProposalAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

// Prevent strconv unused error
var _ = strconv.IntSize

func TestInactiveDisableVirtualSchemaProposalQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNInactiveDisableVirtualSchemaProposal(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetInactiveDisableVirtualSchemaProposalRequest
		response *types.QueryGetInactiveDisableVirtualSchemaProposalResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetInactiveDisableVirtualSchemaProposalRequest{
				Index: msgs[0].Id,
			},
			response: &types.QueryGetInactiveDisableVirtualSchemaProposalResponse{InactiveDisableVirtualSchemaProposal: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetInactiveDisableVirtualSchemaProposalRequest{
				Index: msgs[1].Id,
			},
			response: &types.QueryGetInactiveDisableVirtualSchemaProposalResponse{InactiveDisableVirtualSchemaProposal: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetInactiveDisableVirtualSchemaProposalRequest{
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
			response, err := keeper.InactiveDisableVirtualSchemaProposal(wctx, tc.request)
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

func TestInactiveDisableVirtualSchemaProposalQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNInactiveDisableVirtualSchemaProposal(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllInactiveDisableVirtualSchemaProposalRequest {
		return &types.QueryAllInactiveDisableVirtualSchemaProposalRequest{
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
			resp, err := keeper.InactiveDisableVirtualSchemaProposalAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.InactiveDisableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.InactiveDisableVirtualSchemaProposal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.InactiveDisableVirtualSchemaProposalAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.InactiveDisableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.InactiveDisableVirtualSchemaProposal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.InactiveDisableVirtualSchemaProposalAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.InactiveDisableVirtualSchemaProposal),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.InactiveDisableVirtualSchemaProposalAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
