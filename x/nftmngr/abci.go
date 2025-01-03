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

	k.CreateVirtualSchemaIterateActiveProposal(ctx, ctx.BlockHeader().Time, func(proposal types.VirtualSchemaProposal) (stop bool) {
		pass := k.CreateVirtualSchemaAfterProposalSuccess(ctx, proposal)
		return pass
	})

	k.DisableVirtualSchemaIterateActiveProposal(ctx, ctx.BlockHeader().Time, func(proposal types.DisableVirtualSchemaProposal) (stop bool) {
		pass := k.DisableVirtualSchemaAfterProposalSuccess(ctx, proposal)
		return pass
	})

	k.EnableVirtualSchemaIterateActiveProposal(ctx, ctx.BlockHeader().Time, func(proposal types.EnableVirtualSchemaProposal) (stop bool) {
		pass := k.EnableVirtualSchemaAfterProposalSuccess(ctx, proposal)
		return pass
	})
}
