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

func TestOracleConfigQuery(t *testing.T) {
	k, ctx := keepertest.NftoracleKeeper(t)

	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := k.OracleConfig(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})

	t.Run("NotFound", func(t *testing.T) {
		_, err := k.OracleConfig(ctx, &types.QueryGetOracleConfigRequest{})
		require.ErrorIs(t, err, status.Error(codes.NotFound, "not found"))
	})

	t.Run("Found", func(t *testing.T) {
		item := createTestOracleConfig(k, ctx)
		response, err := k.OracleConfig(ctx, &types.QueryGetOracleConfigRequest{})
		require.NoError(t, err)
		require.Equal(t,
			nullify.Fill(&types.QueryGetOracleConfigResponse{OracleConfig: item}),
			nullify.Fill(response),
		)
	})
}
