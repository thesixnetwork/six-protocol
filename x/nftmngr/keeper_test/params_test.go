package keeper_test

import (
	"testing"

	testkeeper "github.com/thesixnetwork/six-protocol/testutil/keeper"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NftmngrKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
