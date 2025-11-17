# Balance Change Tracking in Precompiles

## Overview

When EVM precompiles modify bank balances through the Cosmos SDK, the EVM stateDB doesn't know about these changes and can overwrite them when committing the EVM state. This guide explains how to implement denomination-based balance change tracking to prevent this issue.

## The Problem

Consider this scenario:
1. A precompile method transfers tokens (e.g., "asix", "usix", or any other denomination)
2. The precompile modifies bank balances through the Cosmos SDK bank keeper
3. The EVM stateDB commits its state, potentially overwriting the bank keeper changes
4. The balance changes are lost or inconsistent, especially for EVM-compatible tokens

## The Solution

We track balance changes based on denomination type and apply them to the EVM stateDB before committing.

## Implementation

### 1. Basic Setup

Your precompile executor needs a reference to the precompile instance:

```go
type PrecompileExecutor struct {
    // ... other keepers
    precompile *pcommon.Precompile
    // ... other fields
}

func NewPrecompile(/* keepers */) (*pcommon.Precompile, error) {
    newAbi := GetABI()
    p := NewExecutor(/* keepers */)
    
    precompile := pcommon.NewPrecompile(newAbi, p, p.address, "your-precompile")
    p.precompile = precompile  // Set the reference
    return precompile, nil
}
```

### 2. Track Balance Changes

In your precompile methods, track balance changes when called from smart contracts:

```go
func (p *PrecompileExecutor) withdrawRewards(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
    // ... validation logic
    
    senderCosmoAddr, err := p.accAddressFromArg(caller)
    if err != nil {
        return nil, err
    }
    
    // Perform the actual operation
    coins, err := p.distributionKeeper.WithdrawDelegationRewards(ctx, senderCosmoAddr, valAddress)
    if err != nil {
        return nil, err
    }
    
    // Track balance changes if called from a smart contract
    if caller != utils.CosmosToEthAddr(senderCosmoAddr) {
        withdrawerAddr, err := p.getWithdrawerAddress(ctx, senderCosmoAddr)
        if err != nil {
            return nil, err
        }
        
        baseDenomAmount := coins.AmountOf("usix").BigInt()
        if baseDenomAmount.Sign() > 0 {
            p.precompile.SetBalanceChangeEntries(
                pcommon.NewBalanceChangeEntry(withdrawerAddr, baseDenomAmount, pcommon.Add)
            )
        }
    }
    
    return method.Outputs.Pack(true)
}
```

### 3. Using the Balance Tracker Helper

For cleaner code, use the balance tracker helper:

```go
func (p *PrecompileExecutor) withdrawRewards(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
    // ... validation logic
    
    senderCosmoAddr, err := p.accAddressFromArg(caller)
    if err != nil {
        return nil, err
    }
    
    // Perform the actual operation
    coins, err := p.distributionKeeper.WithdrawDelegationRewards(ctx, senderCosmoAddr, valAddress)
    if err != nil {
        return nil, err
    }
    
    // Track balance changes based on denomination type
    tracker := pcommon.NewBalanceTracker(p.precompile)
    
    withdrawerAddr, err := p.getWithdrawerAddress(ctx, senderCosmoAddr)
    if err != nil {
        return nil, err
    }
    
    // Track rewards for all denominations that need EVM state synchronization
    for _, coin := range coins {
        if coin.Amount.IsPositive() && pcommon.IsTrackableDenom(coin.Denom) {
            tracker.TrackRewardWithdrawal(withdrawerAddr, coin.Amount.BigInt(), coin.Denom)
        }
    }
    
    return method.Outputs.Pack(true)
}
```

## Common Patterns

### 1. Transfer Operations

```go
func (p *PrecompileExecutor) send(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
    // ... validation and parsing
    
    err := p.bankKeeper.SendCoins(ctx, senderAddr, receiverAddr, coins)
    if err != nil {
        return nil, err
    }
    
    // Track balance changes based on denomination type
    if pcommon.IsTrackableDenom(denom) {
        tracker := pcommon.NewBalanceTracker(p.precompile)
        tracker.TrackTransfer(
            utils.CosmosToEthAddr(senderAddr),
            utils.CosmosToEthAddr(receiverAddr),
            amount,
            denom,
        )
    }
    
    return method.Outputs.Pack(true)
}
```

### 2. Reward Withdrawals

```go
func (p *PrecompileExecutor) claimRewards(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
    // ... validation and claiming logic
    
    totalCoins := sdk.Coins{}
    for _, validator := range validators {
        coins, err := p.distributionKeeper.WithdrawDelegationRewards(ctx, delegatorAddr, validator.Address)
        if err != nil {
            return nil, err
        }
        totalCoins = totalCoins.Add(coins...)
    }
    
    // Track balance changes based on denomination type
    tracker := pcommon.NewBalanceTracker(p.precompile)
    
    withdrawerAddr, err := p.getWithdrawerAddress(ctx, delegatorAddr)
    if err != nil {
        return nil, err
    }
    
    // Track rewards for denominations that need EVM state synchronization
    for _, coin := range totalCoins {
        if coin.Amount.IsPositive() && pcommon.IsTrackableDenom(coin.Denom) {
            tracker.TrackRewardWithdrawal(withdrawerAddr, coin.Amount.BigInt(), coin.Denom)
        }
    }
    
    return method.Outputs.Pack(true)
}
```

### 3. Multiple Balance Changes

```go
func (p *PrecompileExecutor) complexOperation(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
    // ... complex operation affecting multiple accounts
    
    if pcommon.ShouldTrackFromContract(caller, operatorAddr) {
        tracker := pcommon.NewBalanceTracker(p.precompile)
        
        // Track multiple changes
        tracker.TrackBalanceChanges(
            pcommon.NewBalanceChangeEntry(account1, amount1, pcommon.Add),
            pcommon.NewBalanceChangeEntry(account2, amount2, pcommon.Sub),
            pcommon.NewBalanceChangeEntry(account3, amount3, pcommon.Add),
        )
    }
    
    return method.Outputs.Pack(true)
}
```

## Best Practices

### 1. Track Based on Denomination Type
Track balance changes for denominations that need EVM state synchronization:

```go
if pcommon.IsTrackableDenom(denom) {
    tracker := pcommon.NewBalanceTracker(p.precompile)
    tracker.TrackTransfer(sender, receiver, amount, denom)
}
```

### 2. Focus on EVM State Consistency
Apply balance tracking for denominations that interact with EVM state:

```go
// Track all denominations or specific ones based on your needs
if pcommon.IsTrackableDenom(denom) {
    // Apply tracking for EVM state consistency
}
```

### 3. Supported Denominations
The system can track balance changes for:
- All denominations by default (when `IsTrackableDenom` returns true for any non-empty denom)
- Specific denominations like "usix" (native gas token) and "asix" (EVM-compatible token)
- Custom filtering based on your chain's requirements

Use helper functions to check denomination eligibility:
```go
if pcommon.IsTrackableDenom(denom) {
    // Track denomination for EVM state consistency
    tracker.TrackTransfer(sender, receiver, amount, denom)
}
```

### 4. Handle Zero Amounts
Don't track zero amounts to avoid unnecessary state changes:

```go
if amount.Sign() > 0 {
    tracker.TrackBalanceChange(account, amount, pcommon.Add)
}
```

### 5. Error Handling
Always handle errors when getting withdrawer addresses or other dependencies:

```go
withdrawerAddr, err := p.getWithdrawerAddress(ctx, delegatorAddr)
if err != nil {
    return nil, err
}
```

## Testing

Test your precompiles with smart contract calls to ensure balance tracking works:

1. Deploy a test contract that calls your precompile
2. Verify that bank balances match EVM balances after operations
3. Test edge cases like zero amounts and failed operations

## Troubleshooting

### Balance Mismatches
If you see balance mismatches between bank and EVM state:
- Check that you're tracking all balance-modifying operations
- Verify you're only tracking the base denomination
- Ensure tracking only happens for smart contract calls

### Performance Issues
If tracking causes performance issues:
- Only track when necessary (smart contract context)
- Batch balance changes when possible
- Consider the gas cost of state updates

## Migration

To add balance tracking to existing precompiles:

1. Add precompile reference to your executor struct
2. Set the reference in your `NewPrecompile` function
3. Add balance tracking calls to methods that modify balances
4. Test thoroughly with both direct calls and smart contract calls