package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/x/nft/keeper"
	"github.com/thesixnetwork/six-protocol/x/nft/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NftKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
