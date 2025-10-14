// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package ante

import (
	"math"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	nftadminkeeper "github.com/thesixnetwork/six-protocol/x/nftadmin/keeper"
	nftoraclekeeper "github.com/thesixnetwork/six-protocol/x/nftoracle/keeper"
	nftoracletypes "github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

const (
	// OraclePriority is the highest priority for oracle transactions
	OraclePriority = math.MaxInt64 - 100
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
	// Check if this transaction is gasless
	isGasless, err := IsTxGasless(tx, ctx, gd.nftOracleKeeper, gd.nftAdminKeeper)
	if err != nil {
		return ctx, err
	}

	// If gasless and in CheckTx, use infinite gas meter
	if isGasless && ctx.IsCheckTx() {
		// Set infinite gas meter for gasless transactions during CheckTx
		newCtx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter()).
			WithPriority(OraclePriority) // Give oracle transactions highest priority
	} else {
		newCtx = ctx
	}

	// Apply wrapped decorators (including fee deduction)
	// For gasless transactions, we skip fee deduction entirely
	if !isGasless {
		for _, decorator := range gd.wrappedDecorators {
			newCtx, err = decorator.AnteHandle(newCtx, tx, simulate, func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
				return ctx, nil
			})
			if err != nil {
				return ctx, err
			}
		}
	}

	// If gasless, ensure infinite gas meter is maintained
	if isGasless && !ctx.IsCheckTx() {
		// Reset gas meter to infinite for gasless oracle transactions in DeliverTx
		newCtx = newCtx.WithGasMeter(storetypes.NewInfiniteGasMeter())
	}

	return next(newCtx, tx, simulate)
}

// IsTxGasless determines if a transaction should be exempted from gas fees
func IsTxGasless(tx sdk.Tx, ctx sdk.Context, oracleKeeper nftoraclekeeper.Keeper, nftAdminKeeper nftadminkeeper.Keeper) (bool, error) {
	msgs := tx.GetMsgs()
	if len(msgs) == 0 {
		return false, nil
	}

	// Oracle transactions must contain only one message (no bundling allowed)
	if len(msgs) > 1 {
		return false, nil
	}

	msg := msgs[0]

	switch m := msg.(type) {
	case *nftoracletypes.MsgSubmitMintResponse:
		return oracleVoteIsGasless(m, ctx, oracleKeeper, nftAdminKeeper)
	case *nftoracletypes.MsgSubmitActionResponse:
		return oracleActionResponseIsGasless(m, ctx, oracleKeeper, nftAdminKeeper)
	case *nftoracletypes.MsgSubmitVerifyCollectionOwner:
		return oracleCollectionVerifyIsGasless(m, ctx, oracleKeeper, nftAdminKeeper)
	default:
		// Non-oracle transactions are not gasless
		return false, nil
	}
}

// oracleVoteIsGasless validates if an oracle mint response vote is gasless
func oracleVoteIsGasless(msg *nftoracletypes.MsgSubmitMintResponse, ctx sdk.Context, oracleKeeper nftoraclekeeper.Keeper, nftAdminKeeper nftadminkeeper.Keeper) (bool, error) {
	// 1. Validate oracle permission
	oracle, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return false, errorsmod.Wrap(errortypes.ErrInvalidAddress, "invalid oracle address")
	}

	// Check if sender has oracle permission
	if !nftAdminKeeper.HasPermission(ctx, nftoracletypes.KeyPermissionOracle, oracle) {
		return false, errorsmod.Wrap(nftoracletypes.ErrNoOraclePermission, msg.Creator)
	}

	// 2. Validate mint request exists and is pending
	mintRequest, found := oracleKeeper.GetMintRequest(ctx, msg.MintRequestID)
	if !found {
		return false, errorsmod.Wrap(nftoracletypes.ErrMintRequestNotFound, "mint request not found")
	}

	if mintRequest.Status != nftoracletypes.RequestStatus_PENDING {
		return false, errorsmod.Wrap(nftoracletypes.ErrMintRequestNotPending, "mint request not pending")
	}

	// 3. Check for duplicate vote - spam prevention
	if hasOracleAlreadyVoted(mintRequest.Confirmers, oracle.String()) {
		return false, errorsmod.Wrap(nftoracletypes.ErrOracleAlreadyVoted, "oracle already voted")
	}

	// 4. Check spam prevention counter
	if err := checkAndSetSpamPreventionCounter(ctx, oracle, oracleKeeper); err != nil {
		return false, err
	}

	return true, nil
}

// oracleActionResponseIsGasless validates if an oracle action response is gasless
func oracleActionResponseIsGasless(msg *nftoracletypes.MsgSubmitActionResponse, ctx sdk.Context, oracleKeeper nftoraclekeeper.Keeper, nftAdminKeeper nftadminkeeper.Keeper) (bool, error) {
	// 1. Validate oracle permission
	oracle, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return false, errorsmod.Wrap(errortypes.ErrInvalidAddress, "invalid oracle address")
	}

	if !nftAdminKeeper.HasPermission(ctx, nftoracletypes.KeyPermissionOracle, oracle) {
		return false, errorsmod.Wrap(nftoracletypes.ErrNoOraclePermission, msg.Creator)
	}

	// 2. Validate action request exists and is pending
	actionRequest, found := oracleKeeper.GetActionRequest(ctx, msg.ActionRequestID)
	if !found {
		return false, errorsmod.Wrap(nftoracletypes.ErrActionRequestNotFound, "action request not found")
	}

	if actionRequest.Status != nftoracletypes.RequestStatus_PENDING {
		return false, errorsmod.Wrap(nftoracletypes.ErrActionRequestNotPending, "action request not pending")
	}

	// 3. Check for duplicate vote
	if hasOracleAlreadyVoted(actionRequest.Confirmers, oracle.String()) {
		return false, errorsmod.Wrap(nftoracletypes.ErrOracleAlreadyVoted, "oracle already voted")
	}

	// 4. Check spam prevention counter
	if err := checkAndSetSpamPreventionCounter(ctx, oracle, oracleKeeper); err != nil {
		return false, err
	}

	return true, nil
}

// oracleCollectionVerifyIsGasless validates if an oracle collection verification is gasless
func oracleCollectionVerifyIsGasless(msg *nftoracletypes.MsgSubmitVerifyCollectionOwner, ctx sdk.Context, oracleKeeper nftoraclekeeper.Keeper, nftAdminKeeper nftadminkeeper.Keeper) (bool, error) {
	// 1. Validate oracle permission
	oracle, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return false, errorsmod.Wrap(errortypes.ErrInvalidAddress, "invalid oracle address")
	}

	if !nftAdminKeeper.HasPermission(ctx, nftoracletypes.KeyPermissionOracle, oracle) {
		return false, errorsmod.Wrap(nftoracletypes.ErrNoOraclePermission, msg.Creator)
	}

	// 2. Validate collection owner request exists and is pending
	collectionRequest, found := oracleKeeper.GetCollectionOwnerRequest(ctx, msg.VerifyRequestID)
	if !found {
		return false, errorsmod.Wrap(nftoracletypes.ErrCollectionOwnerRequestNotFound, "collection owner request not found")
	}

	if collectionRequest.Status != nftoracletypes.RequestStatus_PENDING {
		return false, errorsmod.Wrap(nftoracletypes.ErrCollectionOwnerRequestNotPending, "collection owner request not pending")
	}

	// 3. Check for duplicate vote
	if hasOracleAlreadyVoted(collectionRequest.Confirmers, oracle.String()) {
		return false, errorsmod.Wrap(nftoracletypes.ErrOracleAlreadyVoted, "oracle already voted")
	}

	// 4. Check spam prevention counter
	if err := checkAndSetSpamPreventionCounter(ctx, oracle, oracleKeeper); err != nil {
		return false, err
	}

	return true, nil
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
	// This could be implemented using transient store or in-memory tracking
	// For now, we'll use a simple block height check stored in the oracle keeper

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
