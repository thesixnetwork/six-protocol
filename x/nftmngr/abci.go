package nftmngr

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	k.IterateActiveProposal(ctx, ctx.BlockHeader().Time, func(proposal types.VirtualSchemaProposal) (stop bool) {
		pass := k.AfterProposalSuccess(ctx, proposal)
		return pass
	})
}
