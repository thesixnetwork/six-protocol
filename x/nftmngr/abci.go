package nftmngr

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	fmt.Println("###### ENTER ENBLOCKER ########")

	// logger := k.Logger(ctx)

	k.IterateActiveProposal(ctx, ctx.BlockHeader().Time, func(proposal types.VirtualSchemaProposal) (stop bool) {
		fmt.Println("###### ENTER ITERATOR ########")

		pass := k.AfterProposalSuccess(ctx, proposal)
		fmt.Println("###### CHECK PASS OR NOT ########")
		// logger.Info(
		// 	"proposal tallied",
		// 	"proposal", proposal.ProposalId,
		// 	"title", proposal.GetTitle(),
		// 	"result", logMsg,
		// )

		return pass
	})
}
