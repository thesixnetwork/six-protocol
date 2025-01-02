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

	k.IterateActiveProposalCreateVirtualSchema(ctx, ctx.BlockHeader().Time, func(proposal types.VirtualSchemaProposal) (stop bool) {
		pass := k.AfterProposalCreateVirtualSchemaSuccess(ctx, proposal)
		return pass
	})

	k.IterateActiveProposalDisableVirtualSchema(ctx, ctx.BlockHeader().Time, func(proposal types.DisableVirtualSchemaProposal) (stop bool) {
		pass := k.AfterProposalDisableVirtualSchemaSuccess(ctx, proposal)
		return pass
	})

	k.IterateActiveProposalEnableVirtualSchema(ctx, ctx.BlockHeader().Time, func(proposal types.EnableVirtualSchemaProposal) (stop bool) {
		pass := k.AfterProposalEnableVirtualSchemaSuccess(ctx, types.EnableVirtualSchemaProposal(proposal))
		return pass
	})
}
