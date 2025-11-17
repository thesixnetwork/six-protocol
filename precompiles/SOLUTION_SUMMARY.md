# EVM StateDB Balance Tracking Solution

## Problem Statement

When EVM precompiles modify bank balances through the Cosmos SDK, the EVM stateDB doesn't know about these changes and can overwrite them when committing the EVM state. This leads to inconsistent balances between the bank keeper and EVM state.

## Root Cause

The issue occurs in this sequence:
1. Smart contract calls precompile
2. Precompile modifies bank balances via Cosmos SDK (e.g., `distributionKeeper.WithdrawDelegationRewards`)
3. EVM stateDB commits its state without knowing about the bank changes
4. The stateDB overwrites the bank keeper changes
5. Balance inconsistencies arise

## Solution Overview

We implemented a balance change tracking mechanism that:
1. Records balance changes when precompiles modify bank balances
2. Applies these changes to the EVM stateDB before committing
3. Only activates when called from smart contracts (not direct transactions)
4. Only tracks the base denomination to prevent conflicts

## Implementation Details

### 1. Core Types (`precompiles/common/precompiles.go`)

Added balance tracking types and functionality:
```go
type Operation int8
const (
    Sub Operation = iota
    Add
)

type BalanceChangeEntry struct {
    Account common.Address
    Amount  *big.Int
    Op      Operation
}
```

### 2. Precompile Base Class Updates

Modified the `Precompile` struct to include:
- `balanceChanges []BalanceChangeEntry` field
- `SetBalanceChangeEntries()` method to record changes
- `applyBalanceChanges()` method to apply changes to stateDB
- Integration in the `Run()` method

### 3. Helper Utilities (`precompiles/common/balance_tracking.go`)

Created helper functions for clean integration:
```go
// Check if tracking should be applied
func ShouldTrackFromContract(caller common.Address, delegatorAddr sdk.AccAddress) bool

// Extract base denomination amount
func GetBaseDenomAmount(coins sdk.Coins) *big.Int

// Check denomination types
func IsBaseDenom(denom string) bool      // "usix"
func IsEvmDenom(denom string) bool       // "asix"  
func IsTrackableDenom(denom string) bool // both "usix" and "asix"

// Track various types of operations for ANY denomination
func (bt *BalanceTracker) TrackTransfer(sender, receiver common.Address, amount *big.Int, denom string)
func (bt *BalanceTracker) TrackRewardWithdrawal(withdrawerAddr common.Address, amount *big.Int, denom string)
```

### 4. Integration Pattern

For any precompile method that modifies bank balances:

```go
func (p *PrecompileExecutor) someMethod(ctx sdk.Context, caller common.Address, /* other params */) ([]byte, error) {
    // 1. Perform the actual operation
    coins, err := p.keeper.SomeOperation(ctx, /* params */)
    if err != nil {
        return nil, err
    }
    
    // 2. Track balance changes if called from smart contract
    if pcommon.ShouldTrackFromContract(caller, accountAddr) {
        tracker := pcommon.NewBalanceTracker(p.precompile)
        
        // Get the withdrawal/recipient address
        recipientAddr, err := p.getRecipientAddress(ctx, accountAddr)
        if err != nil {
            return nil, err
        }
        
        // Track the balance change
        baseDenomAmount := pcommon.GetBaseDenomAmount(coins)
        tracker.TrackRewardWithdrawal(recipientAddr, baseDenomAmount, pcommon.BaseDenom)
    }
    
    return method.Outputs.Pack(true)
}
```

## Updated Precompiles

### 1. Distribution Precompile
- Added balance tracking to `withdrawRewards` method
- Tracks reward withdrawals to the appropriate withdrawal address
- Uses the helper utilities for clean code

### 2. Bank Precompile
- Added balance tracking to `send` method
- Tracks both sender (subtract) and receiver (add) balance changes
- Only tracks base denomination transfers

## Key Features

### 1. Smart Contract Context Detection
Only applies tracking when `caller != originalSender`, meaning the precompile was called from a smart contract rather than directly.

### 2. All Denominations Tracked
Tracks changes for ALL denominations when called from smart contracts to prevent EVM state conflicts for any token transfers.

### 3. Error Handling
Graceful fallback when withdrawal addresses can't be retrieved - defaults to the delegator address.

### 4. Performance Optimization
- Minimal overhead for direct transactions (no tracking)
- Efficient balance change application
- Clear separation of concerns

## Usage Examples

### For New Precompiles
```go
// In your precompile constructor
func NewPrecompile(keeper SomeKeeper) (*pcommon.Precompile, error) {
    executor := &PrecompileExecutor{keeper: keeper}
    precompile := pcommon.NewPrecompile(abi, executor, address, "name")
    executor.precompile = precompile  // Important: set the reference
    return precompile, nil
}

// In methods that modify balances
if pcommon.ShouldTrackFromContract(caller, accountAddr) {
    tracker := pcommon.NewBalanceTracker(p.precompile)
    tracker.TrackBalanceChange(address, amount, pcommon.Add)
}
```

### For Existing Precompiles
1. Add `precompile *pcommon.Precompile` field to executor
2. Set the reference in constructor: `executor.precompile = precompile`
3. Add tracking calls to balance-modifying methods
4. Use helper utilities for cleaner code

## Testing Considerations

### Test Scenarios
1. Direct precompile calls (should not track)
2. Smart contract calls (should track)
3. Zero amount operations (should not track)
4. Multiple balance changes in one operation
5. Error conditions during tracking

### Validation
- Verify bank balances match EVM balances after operations
- Test with both successful and failed transactions
- Ensure no tracking overhead for direct calls

## Best Practices

1. **Only Track When Necessary**: Use `ShouldTrackFromContract()` check
2. **Track All Denominations**: Track any denomination for smart contract calls to prevent state conflicts
3. **Handle Errors**: Always handle withdrawal address retrieval errors
4. **Zero Amount Check**: Don't track zero amounts
5. **Clean Code**: Use helper utilities from `balance_tracking.go`

## Migration Guide

To add this solution to existing precompiles:

1. **Update Precompile Structure**:
   ```go
   type PrecompileExecutor struct {
       // existing fields...
       precompile *pcommon.Precompile
   }
   ```

2. **Set Precompile Reference**:
   ```go
   func NewPrecompile(/*args*/) (*pcommon.Precompile, error) {
       executor := NewExecutor(/*args*/)
       precompile := pcommon.NewPrecompile(abi, executor, address, name)
       executor.precompile = precompile  // Add this line
       return precompile, nil
   }
   ```

3. **Add Balance Tracking**:
   ```go
   // In methods that modify bank balances
   if pcommon.ShouldTrackFromContract(caller, accountAddr) {
       tracker := pcommon.NewBalanceTracker(p.precompile)
       // Track ANY denomination for smart contract calls
       tracker.TrackBalanceChange(address, amount, operation)
   }
   ```

## Files Modified

- `precompiles/common/precompiles.go` - Core balance tracking functionality
- `precompiles/common/balance_tracking.go` - Helper utilities (new file)
- `precompiles/common/expected_keepers.go` - Added GetDelegatorWithdrawAddr interface
- `precompiles/distribution/distribution.go` - Added tracking to withdrawRewards
- `precompiles/bank/bank.go` - Added tracking to send method
- `precompiles/example_integration.go` - Complete example (new file)

## Result

With this solution, precompiles can safely modify bank balances through the Cosmos SDK without worrying about EVM stateDB overwrites. The system maintains consistency between bank keeper state and EVM state for ALL denominations, ensuring reliable operation when called from smart contracts regardless of the token type being transferred.