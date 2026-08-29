// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package ante

import (
	"math"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

	nftadminkeeper "github.com/thesixnetwork/six-protocol/v4/x/nftadmin/keeper"
	nftoraclekeeper "github.com/thesixnetwork/six-protocol/v4/x/nftoracle/keeper"
	nftoracletypes "github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"
)

const (
	// OraclePriority is the highest priority for oracle transactions
	OraclePriority = math.MaxInt64 - 100

	// MaxGaslessGasWanted bounds the execution a single fee-exempt oracle vote
	// may consume. Gasless txs skip fee deduction, so they must NOT run under an
	// infinite gas meter — otherwise a permissioned (or compromised) oracle key
	// could drive unbounded computation per block for free. A vote that exceeds
	// this limit simply runs out of gas and fails, same as any other tx.
	MaxGaslessGasWanted uint64 = 10_000_000
)

// GaslessDecorator wraps the fee deduction decorator to conditionally apply gas charges.
// Oracle voting transactions are exempted from gas fees to encourage participation.
type GaslessDecorator struct {
	wrappedDecorators []sdk.AnteDecorator
	nftOracleKeeper   nftoraclekeeper.Keeper
	nftAdminKeeper    nftadminkeeper.Keeper
}

// NewGaslessDecorator creates a new GaslessDecorator instance
func NewGaslessDecorator(
	wrappedDecorators []sdk.AnteDecorator,
	nftOracleKeeper nftoraclekeeper.Keeper,
	nftAdminKeeper nftadminkeeper.Keeper,
) GaslessDecorator {
	return GaslessDecorator{
		wrappedDecorators: wrappedDecorators,
		nftOracleKeeper:   nftOracleKeeper,
		nftAdminKeeper:    nftAdminKeeper,
	}
}

// AnteHandle implements the AnteDecorator interface
func (gd GaslessDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// Determine whether this tx qualifies for the fee-exempt oracle path. This
	// check is side-effect free — it does not mutate any state.
	isGasless, oracle, err := IsTxGasless(tx, ctx, gd.nftOracleKeeper, gd.nftAdminKeeper)
	if err != nil {
		return ctx, err
	}

	if isGasless {
		// Enforce the per-block, per-oracle rate limit against COMMITTED state.
		// Only write the marker in DeliverTx (consensus); in CheckTx do a
		// read-only comparison so the mempool can shed obvious duplicates
		// without mutating durable state (a CheckTx write lands in a throwaway
		// cache and cannot be relied on). If the limit is hit, fall back to the
		// normal fee-paying path rather than granting another free execution.
		if ctx.IsCheckTx() {
			if gd.nftOracleKeeper.GetOracleLastVoteHeight(ctx, oracle) == ctx.BlockHeight() {
				isGasless = false
			}
		} else if spamErr := checkAndSetSpamPreventionCounter(ctx, oracle, gd.nftOracleKeeper); spamErr != nil {
			isGasless = false
		}
	}

	if isGasless {
		// Fee-exempt path: bound the gas so a free tx cannot consume unbounded
		// computation, give it high mempool priority, and skip fee deduction.
		gasCtx := ctx.WithGasMeter(storetypes.NewGasMeter(MaxGaslessGasWanted))
		if ctx.IsCheckTx() {
			gasCtx = gasCtx.WithPriority(OraclePriority)
		}
		return next(gasCtx, tx, simulate)
	}

	// Normal path: run the wrapped decorators (including fee deduction).
	newCtx = ctx
	for _, decorator := range gd.wrappedDecorators {
		newCtx, err = decorator.AnteHandle(newCtx, tx, simulate, func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
			return ctx, nil
		})
		if err != nil {
			return ctx, err
		}
	}

	return next(newCtx, tx, simulate)
}

// IsTxGasless determines if a transaction should be exempted from gas fees. It
// is side-effect free: any condition that disqualifies the tx returns
// (false, nil, nil) so the caller falls through to the normal fee-paying path,
// rather than returning an error that would reject an otherwise-valid tx.
func IsTxGasless(tx sdk.Tx, ctx sdk.Context, oracleKeeper nftoraclekeeper.Keeper, nftAdminKeeper nftadminkeeper.Keeper) (bool, sdk.AccAddress, error) {
	msgs := tx.GetMsgs()

	// Oracle transactions must contain exactly one message (no bundling allowed).
	if len(msgs) != 1 {
		return false, nil, nil
	}

	switch m := msgs[0].(type) {
	case *nftoracletypes.MsgSubmitMintResponse:
		return oracleVoteIsGasless(m, ctx, oracleKeeper, nftAdminKeeper)
	case *nftoracletypes.MsgSubmitActionResponse:
		return oracleActionResponseIsGasless(m, ctx, oracleKeeper, nftAdminKeeper)
	case *nftoracletypes.MsgSubmitVerifyCollectionOwner:
		return oracleCollectionVerifyIsGasless(m, ctx, oracleKeeper, nftAdminKeeper)
	default:
		// Non-oracle transactions are not gasless
		return false, nil, nil
	}
}

// oracleVoteIsGasless validates if an oracle mint response vote is gasless
func oracleVoteIsGasless(msg *nftoracletypes.MsgSubmitMintResponse, ctx sdk.Context, oracleKeeper nftoraclekeeper.Keeper, nftAdminKeeper nftadminkeeper.Keeper) (bool, sdk.AccAddress, error) {
	// 1. Validate oracle permission
	oracle, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return false, nil, nil
	}

	// Check if sender has oracle permission
	if !nftAdminKeeper.HasPermission(ctx, nftoracletypes.KeyPermissionOracle, oracle) {
		return false, nil, nil
	}

	// 2. Validate mint request exists and is pending
	mintRequest, found := oracleKeeper.GetMintRequest(ctx, msg.MintRequestID)
	if !found {
		return false, nil, nil
	}

	if mintRequest.Status != nftoracletypes.RequestStatus_PENDING {
		return false, nil, nil
	}

	// 3. Check for duplicate vote - spam prevention
	if hasOracleAlreadyVoted(mintRequest.Confirmers, oracle.String()) {
		return false, nil, nil
	}

	return true, oracle, nil
}

// oracleActionResponseIsGasless validates if an oracle action response is gasless
func oracleActionResponseIsGasless(msg *nftoracletypes.MsgSubmitActionResponse, ctx sdk.Context, oracleKeeper nftoraclekeeper.Keeper, nftAdminKeeper nftadminkeeper.Keeper) (bool, sdk.AccAddress, error) {
	// 1. Validate oracle permission
	oracle, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return false, nil, nil
	}

	if !nftAdminKeeper.HasPermission(ctx, nftoracletypes.KeyPermissionOracle, oracle) {
		return false, nil, nil
	}

	// 2. Validate action request exists and is pending
	actionRequest, found := oracleKeeper.GetActionRequest(ctx, msg.ActionRequestID)
	if !found {
		return false, nil, nil
	}

	if actionRequest.Status != nftoracletypes.RequestStatus_PENDING {
		return false, nil, nil
	}

	// 3. Check for duplicate vote
	if hasOracleAlreadyVoted(actionRequest.Confirmers, oracle.String()) {
		return false, nil, nil
	}

	return true, oracle, nil
}

// oracleCollectionVerifyIsGasless validates if an oracle collection verification is gasless
func oracleCollectionVerifyIsGasless(msg *nftoracletypes.MsgSubmitVerifyCollectionOwner, ctx sdk.Context, oracleKeeper nftoraclekeeper.Keeper, nftAdminKeeper nftadminkeeper.Keeper) (bool, sdk.AccAddress, error) {
	// 1. Validate oracle permission
	oracle, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return false, nil, nil
	}

	if !nftAdminKeeper.HasPermission(ctx, nftoracletypes.KeyPermissionOracle, oracle) {
		return false, nil, nil
	}

	// 2. Validate collection owner request exists and is pending
	collectionRequest, found := oracleKeeper.GetCollectionOwnerRequest(ctx, msg.VerifyRequestID)
	if !found {
		return false, nil, nil
	}

	if collectionRequest.Status != nftoracletypes.RequestStatus_PENDING {
		return false, nil, nil
	}

	// 3. Check for duplicate vote
	if hasOracleAlreadyVoted(collectionRequest.Confirmers, oracle.String()) {
		return false, nil, nil
	}

	return true, oracle, nil
}

// hasOracleAlreadyVoted checks if an oracle has already voted
func hasOracleAlreadyVoted(confirmers []string, oracleAddr string) bool {
	for _, confirmer := range confirmers {
		if confirmer == oracleAddr {
			return true
		}
	}
	return false
}

// checkAndSetSpamPreventionCounter prevents multiple oracle submissions in the same block
func checkAndSetSpamPreventionCounter(ctx sdk.Context, oracle sdk.AccAddress, oracleKeeper nftoraclekeeper.Keeper) error {
	// Get the last block height this oracle submitted a vote
	lastVoteHeight := oracleKeeper.GetOracleLastVoteHeight(ctx, oracle)
	currentHeight := ctx.BlockHeight()

	// Prevent multiple votes in the same block
	if lastVoteHeight == currentHeight {
		return errorsmod.Wrap(nftoracletypes.ErrOracleSpamPrevention, "oracle already voted in this block")
	}

	// Update the last vote height
	oracleKeeper.SetOracleLastVoteHeight(ctx, oracle, currentHeight)

	return nil
}

// VoteAloneDecorator ensures oracle votes cannot be bundled with other messages
type VoteAloneDecorator struct{}

// NewVoteAloneDecorator creates a new VoteAloneDecorator
func NewVoteAloneDecorator() VoteAloneDecorator {
	return VoteAloneDecorator{}
}

// AnteHandle implements the AnteDecorator interface
func (vad VoteAloneDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()

	if len(msgs) <= 1 {
		return next(ctx, tx, simulate)
	}

	// Check if any message is an oracle vote
	hasOracleVote := false
	for _, msg := range msgs {
		switch msg.(type) {
		case *nftoracletypes.MsgSubmitMintResponse,
			*nftoracletypes.MsgSubmitActionResponse,
			*nftoracletypes.MsgSubmitVerifyCollectionOwner:
			hasOracleVote = true
		}
		if hasOracleVote {
			break
		}
	}

	// If oracle vote is bundled with other messages, reject
	if hasOracleVote {
		return ctx, errorsmod.Wrap(
			errortypes.ErrInvalidRequest,
			"oracle votes cannot be bundled with other messages",
		)
	}

	return next(ctx, tx, simulate)
}
