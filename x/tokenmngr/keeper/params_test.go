package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TokenmngrKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
