# ASIX Transfer Balance Tracking Guide

## Overview

This guide specifically covers balance tracking for ASIX (EVM denomination) transfers when using the bank precompile from smart contracts, particularly for scenarios like your Solidity script.

## Your Use Case

### Script Analysis
Your Solidity script calls:
```solidity
contractAddress.call(
    abi.encodeWithSignature(
        "send(address,address,string,uint256)",
        ownerAddress,
        0xd907f36f7D83344057a619b6D83A45B3288c3c21,
        "asix",
        990 * 1e18
    )
);
```

This is **exactly** the scenario where balance tracking is critical!

## Why ASIX Transfers Need Balance Tracking

### The Problem
1. Your script calls the bank precompile from a smart contract
2. Bank precompile transfers 990 ASIX via Cosmos SDK bank keeper
3. EVM stateDB doesn't know about this balance change
4. When EVM commits state, it overwrites the bank keeper changes
5. **Result**: Inconsistent balances between bank keeper and EVM state

### The Solution
Balance tracking ensures EVM stateDB is updated with the bank changes:
1. Script calls bank precompile ✅
2. Bank keeper executes transfer ✅
3. **Balance tracking detects contract call** ✅
4. **Balance tracking records ASIX balance changes** ✅
5. **EVM stateDB is updated before commit** ✅
6. Balances remain consistent ✅

## How It Works

### Detection Logic
```go
// 1. Detect if called from smart contract
if pcommon.ShouldTrackFromContract(caller, senderCosmoAddr) {
    // 2. Check if denomination needs tracking
    if pcommon.IsTrackableDenom("asix") { // Returns true for "asix"
        // 3. Track balance changes
        tracker.TrackTransfer(senderEth, receiverEth, amount, "asix")
    }
}
```

### What Gets Tracked
For your 990 ASIX transfer:
- **Sender**: -990 ASIX (subtraction from your owner address)
- **Receiver**: +990 ASIX (addition to 0xd907f36f7D83344057a619b6D83A45B3288c3c21)

## Implementation Status

### ✅ Current Implementation
The balance tracking is already implemented and will work for your script:

1. **Bank Precompile** (`precompiles/bank/bank.go`):
   - Detects smart contract calls
   - Tracks ASIX denomination transfers
   - Applies balance changes to EVM stateDB

2. **Balance Tracker** (`precompiles/common/balance_tracking.go`):
   - Recognizes "asix" as trackable denomination
   - Creates balance entries for sender/receiver
   - Integrates with EVM state management

### Debug Logging
Your script will generate these logs:
```
Bank precompile: tracking balance changes for smart contract call
Bank precompile: denomination is trackable, applying balance tracking
[BalanceTracker] TrackTransfer called: sender=0x..., receiver=0xd907..., amount=990000000000000000000, denom=asix
[BalanceTracker] Denomination asix is trackable, creating balance entries
Bank precompile: balance tracking applied successfully
```

## Supported Denominations

| Denomination | Tracked | Use Case | Your Script |
|--------------|---------|----------|-------------|
| `"asix"` | ✅ **YES** | EVM transactions | **✅ Used** |
| `"usix"` | ✅ YES | Native gas/staking | ❌ Not used |
| Other tokens | ❌ No | Custom tokens | ❌ Not applicable |

## Testing Your Script

### Expected Behavior
When you run your script:

1. **Before Transfer**:
   - Owner has X ASIX in both bank keeper and EVM state
   - Receiver has Y ASIX in both bank keeper and EVM state

2. **During Transfer**:
   - Bank keeper: Owner -990 ASIX, Receiver +990 ASIX
   - Balance tracking: Records same changes for EVM stateDB

3. **After Transfer**:
   - Owner has (X-990) ASIX in both bank keeper and EVM state ✅
   - Receiver has (Y+990) ASIX in both bank keeper and EVM state ✅
   - **No inconsistencies!** ✅

### Verification Steps
1. Check balances before running script
2. Run your script
3. Verify balances are updated correctly
4. Check logs for balance tracking messages

## Troubleshooting

### Script Fails with Balance Issues
If you see inconsistent balances:
1. Check that balance tracking logs appear
2. Verify denomination is exactly "asix" (case sensitive)
3. Ensure sufficient gas for state updates

### No Balance Tracking Logs
If tracking isn't triggered:
1. Verify calling from smart contract (not direct transaction)
2. Check denomination spelling: must be "asix"
3. Ensure precompile address is correct

### Transaction Succeeds but Balances Wrong
If transfer works but balances are inconsistent:
1. Check EVM state commitment process
2. Verify balance entries are applied to stateDB
3. Review gas limits for state updates

## Configuration

### Required Constants
```go
// In precompiles/common/balance_tracking.go
const (
    BaseDenom = "usix"  // Native denomination
    EvmDenom  = "asix"  // EVM denomination - YOUR USE CASE
)
```

### Precompile Address
Your script uses the correct bank precompile address:
```solidity
address contractAddress = BANK_PRECOMPILE_ADDRESS;
```

Make sure `BANK_PRECOMPILE_ADDRESS` equals `0x0000000000000000000000000000000000001001`

## Best Practices

### For ASIX Transfers
1. **Always use smart contracts**: Direct transfers don't need tracking
2. **Use exact denomination**: "asix" (lowercase, exact spelling)
3. **Sufficient gas**: Include gas for balance tracking operations
4. **Error handling**: Check for balance tracking success

### Script Optimization
```solidity
// Your current approach is correct:
(bool success, bytes memory result) = contractAddress.call(
    abi.encodeWithSignature(
        "send(address,address,string,uint256)",
        ownerAddress,
        receiverAddress,
        "asix",  // ✅ Correct EVM denomination
        amount   // ✅ Proper amount
    )
);
require(success, "Transaction failed"); // ✅ Good error handling
```

## Summary

✅ **Your script is perfectly set up for balance tracking!**

The current implementation will:
- Detect your smart contract call
- Recognize "asix" as a trackable denomination
- Create balance entries for sender and receiver
- Apply changes to EVM stateDB
- Maintain consistency between bank keeper and EVM state

Your 990 ASIX transfer will work correctly with full balance tracking enabled.