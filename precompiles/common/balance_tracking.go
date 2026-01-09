package common

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/thesixnetwork/six-protocol/v4/utils"
)

const (
	// BaseDenom is the base denomination used in the chain
	BaseDenom = "usix"
	// EvmDenom is the EVM denomination used in the chain
	EvmDenom = "asix"
)

// BalanceTracker provides helper functions for tracking balance changes in precompiles
type BalanceTracker struct {
	precompile *Precompile
}

// NewBalanceTracker creates a new balance tracker
func NewBalanceTracker(precompile *Precompile) *BalanceTracker {
	return &BalanceTracker{
		precompile: precompile,
	}
}

// TrackBalanceChange records a balance change for a single account
// This should be called when a precompile modifies bank balances to prevent
// the EVM stateDB from overwriting the changes when committing
func (bt *BalanceTracker) TrackBalanceChange(account common.Address, amount *big.Int, op Operation) {
	bt.precompile.SetBalanceChangeEntries(NewBalanceChangeEntry(account, amount, op))
}

// TrackBalanceChanges records multiple balance changes
// This is useful when a single operation affects multiple accounts (e.g., transfers)
func (bt *BalanceTracker) TrackBalanceChanges(entries ...BalanceChangeEntry) {
	bt.precompile.SetBalanceChangeEntries(entries...)
}

// TrackTransfer records balance changes for a transfer operation
// This is a convenience function for the common case of moving tokens between accounts
func (bt *BalanceTracker) TrackTransfer(sender, receiver common.Address, amount *big.Int, denom string) {
	if amount.Sign() == 0 {
		return
	}

	if denom != "" && !IsTrackableDenom(denom) {
		return
	}

	if denom != EvmDenom {
		return
	}

	bt.TrackBalanceChanges(
		NewBalanceChangeEntry(sender, amount, Sub),
		NewBalanceChangeEntry(receiver, amount, Add),
	)
}

// TrackRewardWithdrawal records balance changes for reward withdrawal operations
// This handles the case where rewards are withdrawn to a different address than the delegator
func (bt *BalanceTracker) TrackRewardWithdrawal(withdrawerAddr common.Address, amount *big.Int, denom string) {
	// Track balance changes for any denomination when amount is positive
	if amount.Sign() > 0 {
		bt.TrackBalanceChange(withdrawerAddr, amount, Add)
	}
}

// ShouldTrackFromContract checks if balance tracking should be applied
// This returns true when the precompile is called from a smart contract
// (caller != original transaction sender)
// NOTE: This is kept for compatibility but denomination-based tracking is now preferred
func ShouldTrackFromContract(caller common.Address, delegatorAddr sdk.AccAddress) bool {
	return caller != utils.CosmosToEthAddr(delegatorAddr)
}

// GetBaseDenomAmount extracts the base denomination amount from a coin collection
// Returns the amount as *big.Int, or big.NewInt(0) if the base denom is not found
func GetBaseDenomAmount(coins sdk.Coins) *big.Int {
	return coins.AmountOf(BaseDenom).BigInt()
}

// IsBaseDenom checks if the given denomination is the base denomination
func IsBaseDenom(denom string) bool {
	return denom == BaseDenom
}

// IsEvmDenom checks if the given denomination is the EVM denomination
func IsEvmDenom(denom string) bool {
	return denom == EvmDenom
}

// IsTrackableDenom checks if the given denomination should be tracked for EVM state changes
// Returns true for all valid denominations to ensure EVM stateDB consistency
func IsTrackableDenom(denom string) bool {
	// Track all denominations to ensure EVM stateDB consistency for all token types
	return denom != ""
}
