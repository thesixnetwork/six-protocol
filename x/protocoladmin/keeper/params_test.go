package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ProtocoladminKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
