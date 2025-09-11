# ChangeDelegatorAddress Function Tests

This document describes the tests created for the `ChangeDelegatorAddress` function and how to run them.

## Overview

The `ChangeDelegatorAddress` function allows changing the delegator address for all staking-related records in the Cosmos SDK. This is a powerful administrative function that transfers:

1. **Delegations** - Active staking positions with validators
2. **Unbonding Delegations** - Tokens in the process of being unstaked
3. **Redelegations** - Tokens being moved between validators

## Test Files Created

### 1. `change_delegator_address_test.go`
Basic unit tests focusing on:
- Address validation
- Address conversion utilities
- Input validation edge cases

### 2. `change_delegator_address_comprehensive_test.go` 
More comprehensive tests including:
- Data structure testing
- Staking type understanding
- Integration test templates

### 3. Mock utilities in `testutil/mock_keepers.go`
Mock implementations for:
- `MockStakingKeeper` - For testing staking operations
- `MockAccountKeeper` - For testing account operations
- `MockAddressCodec` - For testing address conversions

## How to Run Tests

Due to dependency complexity in the Cosmos SDK environment, here are the recommended approaches:

### Option 1: Simple Unit Tests (Recommended)
```bash
cd /path/to/six-protocol
go test ./x/tokenmngr/keeper -run TestAddressConversion -v
```

### Option 2: Full Integration Tests (Requires Setup)
For full integration testing, you would need to:
1. Set up a simulated blockchain environment
2. Create test validators 
3. Set up initial delegations
4. Test the complete workflow

Example integration test structure:
```go
func TestChangeDelegatorAddress_Integration(t *testing.T) {
    // Setup simapp environment
    app := simapp.Setup(false)
    ctx := app.BaseApp.NewContext(false, tmproto.Header{})
    
    // Create test accounts with funds
    oldAddr := simapp.AddTestAddrs(app, ctx, 1, sdk.NewInt(10000000))[0]
    newAddr := simapp.AddTestAddrs(app, ctx, 1, sdk.NewInt(10000000))[0]
    
    // Create validators and delegate tokens
    // ... setup delegations, unbonding delegations, redelegations ...
    
    // Execute the function
    err := app.TokenmgrKeeper.ChangeDelegatorAddress(ctx, oldAddr, newAddr)
    require.NoError(t, err)
    
    // Verify all records were transferred correctly
    // ... verification code ...
}
```

## Function Usage Example

```go
package main

import (
    "context"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func ExampleUsage(keeper *Keeper, ctx context.Context) {
    // Convert string addresses to sdk.AccAddress
    oldAddrStr := "cosmos1oldaddress..."
    newAddrStr := "cosmos1newaddress..."
    
    oldAddr, err := sdk.AccAddressFromBech32(oldAddrStr)
    if err != nil {
        panic(err)
    }
    
    newAddr, err := sdk.AccAddressFromBech32(newAddrStr)
    if err != nil {
        panic(err)
    }
    
    // Get staking info before change (optional)
    infoBefore, err := keeper.GetDelegatorStakingInfo(ctx, oldAddr)
    if err != nil {
        panic(err)
    }
    
    // Execute the address change
    err = keeper.ChangeDelegatorAddress(ctx, oldAddr, newAddr)
    if err != nil {
        panic(err)
    }
    
    // Verify staking info after change (optional)
    infoAfter, err := keeper.GetDelegatorStakingInfo(ctx, newAddr)
    if err != nil {
        panic(err)
    }
    
    // Compare totals to ensure nothing was lost
    if !infoBefore.TotalBonded.Equal(infoAfter.TotalBonded) {
        panic("Total bonded amount mismatch!")
    }
}
```

## Test Scenarios Covered

### Input Validation
- ✅ Empty old address
- ✅ Empty new address  
- ✅ Same old and new address
- ✅ Invalid address formats

### Address Conversion
- ✅ Valid bech32 conversion
- ✅ Invalid bech32 strings
- ✅ Address equality checks

### Data Structures
- ✅ DelegatorStakingInfo struct creation
- ✅ Delegation record structure
- ✅ Unbonding delegation structure  
- ✅ Redelegation record structure

### Security Considerations
- The function includes validation to prevent:
  - Transferring to an address that already has delegations
  - Operating with empty/invalid addresses
  - Self-transfers (same old and new address)

## Production Considerations

⚠️ **WARNING**: This is a powerful administrative function. Before using in production:

1. **Test thoroughly** in a development environment
2. **Backup state** before running in production
3. **Validate addresses** before calling the function
4. **Monitor gas usage** for large numbers of delegations
5. **Consider rewards distribution** - may need separate handling
6. **Add proper logging/events** to track changes

## Troubleshooting

If tests fail to compile:
1. Run `go mod tidy` to update dependencies
2. Ensure Cosmos SDK version compatibility
3. Check that all imports are available
4. Verify the keeper interfaces are properly defined

For integration testing, you may need to:
1. Set up the full simapp environment
2. Initialize validators with proper voting power
3. Fund test accounts with sufficient tokens
4. Handle genesis state properly
