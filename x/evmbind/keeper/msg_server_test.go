package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/evmbind/keeper"
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EvmbindKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}