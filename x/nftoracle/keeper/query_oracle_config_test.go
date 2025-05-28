package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestOracleConfigQuery(t *testing.T) {
	keeper, ctx := keepertest.NftoracleKeeper(t)
	item := createTestOracleConfig(&keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetOracleConfigRequest
		response *types.QueryGetOracleConfigResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetOracleConfigRequest{},
			response: &types.QueryGetOracleConfigResponse{OracleConfig: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.OracleConfig(ctx, tc.request)
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
