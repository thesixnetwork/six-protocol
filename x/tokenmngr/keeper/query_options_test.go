package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestOptionsQuery(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)
	item := createTestOptions(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetOptionsRequest
		response *types.QueryGetOptionsResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetOptionsRequest{},
			response: &types.QueryGetOptionsResponse{Options: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Options(ctx, tc.request)
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
