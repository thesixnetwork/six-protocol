package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"
)

func TestBindedSignerQuerySingle(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)
	msgs := createNBindedSigner(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetBindedSignerRequest
		response *types.QueryGetBindedSignerResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetBindedSignerRequest{
				OwnerAddress: msgs[0].OwnerAddress,
			},
			response: &types.QueryGetBindedSignerResponse{BindedSigner: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetBindedSignerRequest{
				OwnerAddress: msgs[1].OwnerAddress,
			},
			response: &types.QueryGetBindedSignerResponse{BindedSigner: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetBindedSignerRequest{
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
			response, err := k.BindedSigner(ctx, tc.request)
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
