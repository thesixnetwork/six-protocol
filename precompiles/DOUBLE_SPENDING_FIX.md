# Double Spending Issue and Solution in EVM Precompiles

## Problem Statement

When EVM precompiles modify bank balances through Cosmos SDK operations, there's a potential for **double spending** where recipients receive tokens twice:

1. **Bank Module**: Updates Cosmos SDK bank state (recipient gets tokens)
2. **EVM State Tracking**: Updates EVM stateDB (recipient gets tokens AGAIN)

### Example Scenario

User transfers 1000 ASIX via bank precompile:
```
Before: User A has 5000 ASIX, User B has 0 ASIX

Step 1 - Bank Transfer:
  Bank Module: A = 4000 ASIX, B = 1000 ASIX ✅

Step 2 - EVM State Sync:
  EVM stateDB.AddBalance: B += 1000 ASIX
  Result: B = 2000 ASIX in EVM state ❌ DOUBLE SPENDING!
```

## Root Cause Analysis

The issue stems from misunderstanding the relationship between bank balances and EVM state:

### Current Problematic Flow
```go
// 1. Bank operation (modifies Cosmos state)
bankKeeper.SendCoins(ctx, sender, receiver, coins)

// 2. EVM tracking (modifies EVM state ADDITIVELY)
if IsEvmDenom(denom) {
    tracker.TrackTransfer(sender, receiver, amount, denom) 
    // This calls stateDB.AddBalance() which ADDS to existing EVM balance
}
```

### The Core Problem
- **Bank state** and **EVM state** should be **synchronized**, not **additive**
- `stateDB.AddBalance()` adds to existing EVM balance
- This creates double crediting when bank operations already updated balances

## Solution Approaches

### Approach 1: EVM-Only for EVM Tokens (Recommended)

For EVM-compatible tokens (like "asix"), handle transfers entirely in EVM state:

```go
func (p PrecompileExecutor) send(ctx sdk.Context, caller common.Address, ...) ([]byte, error) {
    if pcommon.IsEvmDenom(denom) {
        // EVM tokens: Handle via EVM state only
        tracker := pcommon.NewBalanceTracker(p.precompile)
        tracker.TrackTransfer(senderEth, receiverEth, amount, denom)
        // NO bank operation for EVM tokens
    } else {
        // Non-EVM tokens: Handle via bank keeper only
        bankKeeper.SendCoins(ctx, sender, receiver, coins)
        // NO EVM tracking for non-EVM tokens
    }
}
```

### Approach 2: Bank-First with Proper Sync

Use bank as source of truth and sync EVM state correctly:

```go
func (p PrecompileExecutor) send(ctx sdk.Context, caller common.Address, ...) ([]byte, error) {
    // Always use bank keeper for actual transfer
    bankKeeper.SendCoins(ctx, sender, receiver, coins)
    
    if pcommon.IsEvmDenom(denom) {
        // For EVM tokens, ensure EVM state matches bank state
        // This should SYNC, not ADD to EVM balances
        p.syncEvmStateWithBank(ctx, sender, receiver, amount, denom)
    }
}

func (p PrecompileExecutor) syncEvmStateWithBank(ctx sdk.Context, sender, receiver sdk.AccAddress, amount *big.Int, denom string) {
    // Get actual bank balances after transfer
    senderBankBalance := p.bankKeeper.GetBalance(ctx, sender, denom)
    receiverBankBalance := p.bankKeeper.GetBalance(ctx, receiver, denom)
    
    // Set EVM state to match bank state (not additive)
    stateDB.SetBalance(senderEth, senderBankBalance.Amount.BigInt())
    stateDB.SetBalance(receiverEth, receiverBankBalance.Amount.BigInt())
}
```

### Approach 3: Balance Verification and Correction

Track what the EVM balance should be and correct any discrepancies:

```go
func (p PrecompileExecutor) send(ctx sdk.Context, caller common.Address, ...) ([]byte, error) {
    // Record EVM balances before bank operation
    senderEvmBefore := stateDB.GetBalance(senderEth)
    receiverEvmBefore := stateDB.GetBalance(receiverEth)
    
    // Execute bank transfer
    bankKeeper.SendCoins(ctx, sender, receiver, coins)
    
    if pcommon.IsEvmDenom(denom) {
        // Calculate what EVM balances should be after bank operation
        expectedSenderEvm := new(big.Int).Sub(senderEvmBefore, amount)
        expectedReceiverEvm := new(big.Int).Add(receiverEvmBefore, amount)
        
        // Set EVM state to expected values
        stateDB.SetBalance(senderEth, expectedSenderEvm)
        stateDB.SetBalance(receiverEth, expectedReceiverEvm)
    }
}
```

## Current Implementation Issues

### Bank Precompile
```go
// PROBLEMATIC: This can cause double spending
bankKeeper.SendCoins(ctx, sender, receiver, coins) // Bank: B += 1000
if IsEvmDenom(denom) {
    tracker.TrackTransfer(sender, receiver, amount, denom) // EVM: B += 1000 AGAIN!
}
```

### Staking Precompile
```go
// PROBLEMATIC: ASIX consumption not properly tracked
tokenmngrKeeper.AttoCoinConverter(ctx, user, user, amount) // Converts asix->usix
// Missing: EVM state update for asix consumption
```

### Token Factory Precompile
```go
// PROBLEMATIC: ASIX conversion not properly tracked  
tokenmngrKeeper.AttoCoinConverter(ctx, sender, receiver, amount) // Converts asix
// Missing: Proper EVM state synchronization
```

## Recommended Fix

### 1. Modify SetBalanceChangeEntries Logic

Instead of additive operations, use synchronization:

```go
// Current (PROBLEMATIC)
func (p *Precompile) applyBalanceChanges(stateDB *statedb.StateDB) {
    for _, entry := range p.balanceChanges {
        switch entry.Op {
        case Add:
            stateDB.AddBalance(entry.Account, entry.Amount) // ADDS to existing
        case Sub:
            stateDB.SubBalance(entry.Account, entry.Amount) // SUBS from existing
        }
    }
}

// Fixed (SYNCHRONIZED)
func (p *Precompile) applyBalanceChanges(stateDB *statedb.StateDB) {
    for _, entry := range p.balanceChanges {
        switch entry.Op {
        case Sync:
            // Set EVM balance to match bank balance exactly
            stateDB.SetBalance(entry.Account, entry.Amount)
        case Add:
            // Only use for operations that don't involve bank
            stateDB.AddBalance(entry.Account, entry.Amount)
        case Sub:
            // Only use for operations that don't involve bank
            stateDB.SubBalance(entry.Account, entry.Amount)
        }
    }
}
```

### 2. Add Sync Operation Type

```go
type Operation int8

const (
    Sub Operation = iota
    Add
    Sync // NEW: Set EVM balance to exact value
)

// Use for bank operations
tracker.SyncBalance(account, exactBankBalance)

// Use for EVM-only operations  
tracker.TrackBalanceChange(account, amount, Add)
```

### 3. Updated Bank Precompile

```go
func (p PrecompileExecutor) send(ctx sdk.Context, caller common.Address, ...) ([]byte, error) {
    // Execute bank transfer (source of truth)
    bankKeeper.SendCoins(ctx, sender, receiver, coins)
    
    if pcommon.IsEvmDenom(denom) {
        // Get post-transfer bank balances
        senderBalance := p.bankKeeper.GetBalance(ctx, sender, denom)
        receiverBalance := p.bankKeeper.GetBalance(ctx, receiver, denom)
        
        // Sync EVM state to match bank state exactly
        tracker := pcommon.NewBalanceTracker(p.precompile)
        tracker.SyncBalance(senderEth, senderBalance.Amount.BigInt())
        tracker.SyncBalance(receiverEth, receiverBalance.Amount.BigInt())
    }
}
```

## Testing Strategy

### Test Cases Required

1. **Double Spending Detection**
   ```go
   // Transfer 1000 ASIX
   // Verify recipient gets exactly 1000, not 2000
   ```

2. **Cross-Module Consistency**
   ```go
   // Check bank balance == EVM balance after operations
   ```

3. **Multiple Operations**
   ```go
   // Multiple transfers in sequence should be cumulative, not multiplicative
   ```

4. **Failed Operations**
   ```go
   // Failed bank operations should not update EVM state
   ```

## Migration Plan

1. **Phase 1**: Add Sync operation type
2. **Phase 2**: Update bank precompile to use sync for EVM tokens
3. **Phase 3**: Update staking precompile for asix consumption
4. **Phase 4**: Update token factory precompile for conversions
5. **Phase 5**: Comprehensive testing and validation

## Verification

After implementing the fix:

```go
// Before any operation
bankBalanceBefore := bankKeeper.GetBalance(account, "asix")
evmBalanceBefore := stateDB.GetBalance(account)

// Execute operation
precompile.Send(...)

// After operation  
bankBalanceAfter := bankKeeper.GetBalance(account, "asix")
evmBalanceAfter := stateDB.GetBalance(account)

// Verify consistency
assert.Equal(bankBalanceAfter.BigInt(), evmBalanceAfter)
```

This ensures bank and EVM states remain synchronized without double spending.