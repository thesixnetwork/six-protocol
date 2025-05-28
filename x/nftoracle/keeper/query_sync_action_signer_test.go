package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func TestSyncActionSignerQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	msgs := createNSyncActionSigner(&keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSyncActionSignerRequest
		response *types.QueryGetSyncActionSignerResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetSyncActionSignerRequest{Id: msgs[0].Id},
			response: &types.QueryGetSyncActionSignerResponse{SyncActionSigner: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetSyncActionSignerRequest{Id: msgs[1].Id},
			response: &types.QueryGetSyncActionSignerResponse{SyncActionSigner: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetSyncActionSignerRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SyncActionSigner(ctx, tc.request)
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

func TestSyncActionSignerQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSyncActionSigner(&keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSyncActionSignerRequest {
		return &types.QueryAllSyncActionSignerRequest{
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
			resp, err := keeper.SyncActionSignerAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SyncActionSigner), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SyncActionSigner),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SyncActionSignerAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SyncActionSigner), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.SyncActionSigner),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SyncActionSignerAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.SyncActionSigner),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SyncActionSignerAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
