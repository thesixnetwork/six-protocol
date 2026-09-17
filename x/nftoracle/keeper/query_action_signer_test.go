package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	"github.com/cosmos/cosmos-sdk/types/query"
)

func TestActionSignerQuerySingle(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	msgs := createNActionSigner(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetActionSignerRequest
		response *types.QueryGetActionSignerResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetActionSignerRequest{
				ActorAddress: msgs[0].ActorAddress,
				OwnerAddress: msgs[0].OwnerAddress,
			},
			response: &types.QueryGetActionSignerResponse{ActionSigner: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetActionSignerRequest{
				ActorAddress: msgs[1].ActorAddress,
				OwnerAddress: msgs[1].OwnerAddress,
			},
			response: &types.QueryGetActionSignerResponse{ActionSigner: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetActionSignerRequest{
				ActorAddress: "missing",
				OwnerAddress: "missing",
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := k.ActionSigner(ctx, tc.request)
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

func TestActionSignerQueryPaginated(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	msgs := createNActionSigner(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllActionSignerRequest {
		return &types.QueryAllActionSignerRequest{
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
			resp, err := k.ActionSignerAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionSigner), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionSigner),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := k.ActionSignerAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ActionSigner), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ActionSigner),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := k.ActionSignerAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ActionSigner),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.ActionSignerAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
