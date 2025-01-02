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
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVirtualSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VirtualSchema {
	items := make([]types.VirtualSchema, n)
	for i := range items {
		items[i].VirtualNftSchemaCode = strconv.Itoa(i)

		keeper.SetVirtualSchema(ctx, items[i])
	}
	return items
}

func TestVirtualSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVirtualSchema(ctx,
			item.VirtualNftSchemaCode,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestVirtualSchemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVirtualSchema(ctx,
			item.VirtualNftSchemaCode,
		)
		_, found := keeper.GetVirtualSchema(ctx,
			item.VirtualNftSchemaCode,
		)
		require.False(t, found)
	}
}

func TestVirtualSchemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNVirtualSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVirtualSchema(ctx)),
	)
}

func createNDisableVirtualSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DisableVirtualSchemaProposal {
	items := make([]types.DisableVirtualSchemaProposal, n)
	for i := range items {
		items[i].Id = strconv.Itoa(i)

		keeper.SetDisableVirtualSchemaProposal(ctx, items[i])
	}
	return items
}

func TestDisableVirtualSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNDisableVirtualSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDisableVirtualSchemaProposal(ctx,
			item.Id,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestDisableVirtualSchemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNDisableVirtualSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDisableVirtualSchemaProposal(ctx,
			item.Id,
		)
		_, found := keeper.GetDisableVirtualSchemaProposal(ctx,
			item.Id,
		)
		require.False(t, found)
	}
}

func TestDisableVirtualSchemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNDisableVirtualSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDisableVirtualSchemaProposal(ctx)),
	)
}

func TestDisableVirtualSchemaQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDisableVirtualSchema(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetDisableVirtualSchemaProposalRequest
		response *types.QueryGetDisableVirtualSchemaProposalResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetDisableVirtualSchemaProposalRequest{
				Index: msgs[0].Id,
			},
			response: &types.QueryGetDisableVirtualSchemaProposalResponse{DisableVirtualSchemaProposal: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetDisableVirtualSchemaProposalRequest{
				Index: msgs[1].Id,
			},
			response: &types.QueryGetDisableVirtualSchemaProposalResponse{DisableVirtualSchemaProposal: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetDisableVirtualSchemaProposalRequest{
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
			response, err := keeper.DisableVirtualSchemaProposal(wctx, tc.request)
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

func TestDisableVirtualSchemaQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDisableVirtualSchema(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDisableVirtualSchemaProposalRequest {
		return &types.QueryAllDisableVirtualSchemaProposalRequest{
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
			resp, err := keeper.DisableVirtualSchemaProposalAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.DisableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.DisableVirtualSchemaProposal),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DisableVirtualSchemaProposalAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.DisableVirtualSchemaProposal), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.DisableVirtualSchemaProposal),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.DisableVirtualSchemaProposalAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.DisableVirtualSchemaProposal),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.DisableVirtualSchemaProposalAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
