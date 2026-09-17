package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/nftadmin/types"
)

func TestAuthorizationQuery(t *testing.T) {
	keeper, ctx := keepertest.NftadminKeeper(t)

	t.Run("invalid request", func(t *testing.T) {
		_, err := keeper.Authorization(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})

	t.Run("not found", func(t *testing.T) {
		_, err := keeper.Authorization(ctx, &types.QueryGetAuthorizationRequest{})
		require.ErrorIs(t, err, status.Error(codes.NotFound, "not found"))
	})

	t.Run("found", func(t *testing.T) {
		item := createTestAuthorization(keeper, ctx)
		response, err := keeper.Authorization(ctx, &types.QueryGetAuthorizationRequest{})
		require.NoError(t, err)
		require.Equal(t,
			nullify.Fill(&types.QueryGetAuthorizationResponse{Authorization: item}),
			nullify.Fill(response),
		)
	})
}
