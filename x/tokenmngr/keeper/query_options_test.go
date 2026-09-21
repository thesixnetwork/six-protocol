package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/thesixnetwork/six-protocol/v4/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/v4/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
)

func TestOptionsQuery(t *testing.T) {
	keeper, ctx := keepertest.TokenmngrKeeper(t)

	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.Options(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})

	t.Run("NotFound", func(t *testing.T) {
		_, err := keeper.Options(ctx, &types.QueryGetOptionsRequest{})
		require.ErrorIs(t, err, status.Error(codes.NotFound, "not found"))
	})

	t.Run("Found", func(t *testing.T) {
		item := createTestOptions(keeper, ctx)
		response, err := keeper.Options(ctx, &types.QueryGetOptionsRequest{})
		require.NoError(t, err)
		require.Equal(t,
			nullify.Fill(&types.QueryGetOptionsResponse{Options: item}),
			nullify.Fill(response),
		)
	})
}
