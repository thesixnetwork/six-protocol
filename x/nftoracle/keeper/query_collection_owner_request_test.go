package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func TestCollectionOwnerRequestQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	msgs := createNCollectionOwnerRequest(&keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCollectionOwnerRequestRequest
		response *types.QueryGetCollectionOwnerRequestResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCollectionOwnerRequestRequest{Id: msgs[0].Id},
			response: &types.QueryGetCollectionOwnerRequestResponse{CollectionOwnerRequest: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetCollectionOwnerRequestRequest{Id: msgs[1].Id},
			response: &types.QueryGetCollectionOwnerRequestResponse{CollectionOwnerRequest: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetCollectionOwnerRequestRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.CollectionOwnerRequest(ctx, tc.request)
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

func TestCollectionOwnerRequestQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	msgs := createNCollectionOwnerRequest(&keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCollectionOwnerRequestRequest {
		return &types.QueryAllCollectionOwnerRequestRequest{
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
			resp, err := keeper.CollectionOwnerRequestAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CollectionOwnerRequest), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.CollectionOwnerRequest),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CollectionOwnerRequestAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CollectionOwnerRequest), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.CollectionOwnerRequest),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CollectionOwnerRequestAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.CollectionOwnerRequest),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CollectionOwnerRequestAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
