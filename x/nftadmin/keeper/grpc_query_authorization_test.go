package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftadmin/types"
)

func TestAuthorizationQuery(t *testing.T) {
	keeper, ctx := keepertest.NftAdminKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestAuthorization(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAuthorizationRequest
		response *types.QueryGetAuthorizationResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetAuthorizationRequest{},
			response: &types.QueryGetAuthorizationResponse{Authorization: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Authorization(wctx, tc.request)
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
