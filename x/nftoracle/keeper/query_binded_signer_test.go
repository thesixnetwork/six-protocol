package keeper_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestBindedSignerQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	msgs := createNBindedSigner(&keeper, ctx, 2)
	for _, tc := range []struct {
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
				OwnerAddress: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.BindedSigner(ctx, tc.request)
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
