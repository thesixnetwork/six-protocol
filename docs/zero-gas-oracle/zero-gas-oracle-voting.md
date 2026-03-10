# Zero-Gas Oracle Voting Implementation for Six Protocol

## Overview

This implementation brings **Zero-Gas Oracle Voting** functionality to Six Protocol,Oracle validators can now submit votes without paying gas fees, encouraging participation and maintaining a reliable oracle network while preventing spam and abuse.

## Key Features

### 1. **Gasless Oracle Transactions**

- Oracle votes for mint requests, action requests, and collection verification are **completely free**
- No gas fees charged for legitimate oracle operations
- Maintains high priority in mempool for reliable processing

### 2. **Comprehensive Spam Prevention**

- **One vote per oracle per request** - prevents duplicate voting
- **Block-level spam prevention** - prevents multiple submissions in same block
- **Oracle permission validation** - only authorized oracles can vote
- **Transaction isolation** - oracle votes cannot be bundled with other messages

### 3. **Security & Validation**

- Validates oracle permissions through nftadmin module
- Ensures requests are in pending state before allowing votes
- Maintains all existing oracle slashing and reward mechanisms
- Robust duplicate vote detection

## Architecture Components

### Core Components

#### 1. **GaslessDecorator** (`app/ante/gasless.go`)

The main component that determines if transactions should be gas-free:

```go
type GaslessDecorator struct {
    wrappedDecorators []sdk.AnteDecorator
    nftOracleKeeper   nftoraclekeeper.Keeper
    nftAdminKeeper    nftadminkeeper.Keeper
}
```

**Key Functions:**

- `IsTxGasless()` - Determines if a transaction is eligible for zero-gas
- `AnteHandle()` - Applies infinite gas meter for gasless transactions
- Provides highest priority (`math.MaxInt64 - 100`) for oracle transactions

#### 2. **Oracle Vote Validation Functions**

- `oracleVoteIsGasless()` - Validates mint response votes
- `oracleActionResponseIsGasless()` - Validates action response votes
- `oracleCollectionVerifyIsGasless()` - Validates collection verification votes

#### 3. **VoteAloneDecorator**

Prevents oracle votes from being bundled with other transactions:

```go
// Oracle votes must be submitted alone to prevent abuse
type VoteAloneDecorator struct{}
```

#### 4. **Spam Prevention System** (`x/nftoracle/keeper/gasless_voting.go`)

- Tracks last vote height per oracle
- Prevents multiple votes in same block
- Automatic cleanup of old records

### Supported Oracle Message Types

The following message types are eligible for zero-gas voting:

1. **`MsgSubmitMintResponse`** - Oracle responses for NFT mint requests
2. **`MsgSubmitActionResponse`** - Oracle responses for action requests
3. **`MsgSubmitVerifyCollectionOwner`** - Oracle responses for collection verification

### Validation Criteria

For a transaction to be gasless, it must meet ALL criteria:

1. **Single Message** - Transaction contains only one message (no bundling)
2. **Oracle Permission** - Sender has valid oracle permission
3. **Valid Request** - Target request exists and is in PENDING state
4. **No Duplicate Vote** - Oracle hasn't already voted on this request
5. **Spam Prevention** - Oracle hasn't voted in current block

## Configuration

### Ante Handler Setup

The gasless decorator is integrated into the Cosmos ante handler chain in `app/ante/cosmos.go`:

```go
func newCosmosAnteHandler(options HandlerOptions) sdk.AnteHandler {
    return sdk.ChainAnteDecorators(
        // ... other decorators
        NewVoteAloneDecorator(),         // Prevent bundling
        // ... setup decorators
        NewGaslessDecorator(             // Zero-gas logic
            []sdk.AnteDecorator{
                ante.NewDeductFeeDecorator(...),
            },
            *options.NftOracleKeeper,
            *options.NftAdminKeeper,
        ),
        // ... remaining decorators
    )
}
```

### Handler Options

Extended `HandlerOptions` to include NFT keepers:

```go
type HandlerOptions struct {
    // ... existing fields
    NftOracleKeeper *nftoraclekeeper.Keeper
    NftAdminKeeper  *nftadminkeeper.Keeper
}
```

### App Integration

Updated `app.go` to pass required keepers:

```go
options := ante.HandlerOptions{
    // ... existing options
    NftOracleKeeper: &app.NftoracleKeeper,
    NftAdminKeeper:  &app.NftadminKeeper,
}
```

## Error Handling

New error types for gasless voting:

```go
// In x/nftoracle/types/errors.go
ErrOracleAlreadyVoted                = "Oracle already voted"
ErrOracleSpamPrevention              = "Oracle spam prevention triggered"
ErrCollectionOwnerRequestNotFound    = "Collection owner request not found"
ErrCollectionOwnerRequestNotPending  = "Collection owner request not pending"
ErrInvalidGaslessTransaction         = "Invalid gasless transaction"
```

## Testing

### Test Scripts

The implementation includes comprehensive test scripts:

1. **`test_oracle_voting_detailed.sh`** - Enhanced with zero-gas testing
   - `test_zero_gas_oracle_voting()` - Tests complete gas-free operation
   - `test_zero_gas_spam_prevention()` - Tests abuse prevention
   - Balance tracking before/after votes
   - Transaction analysis and gas usage verification

### Test Scenarios

#### Zero-Gas Voting Test

```bash
# Test oracle vote without gas fees
sixd tx nftoracle submit-mint-response "$request_id" "$base64_data" \
    --from oracle1 \
    --gas auto --gas-adjustment 1.5 \  # No --gas-prices!
    --chain-id testnet -y
```

#### Spam Prevention Test

```bash
# First vote should succeed
oracle1 vote -> ✅ SUCCESS (gasless)

# Duplicate vote should fail
oracle1 vote again -> ❌ REJECTED (duplicate)

# Different oracle should work
oracle2 vote -> ✅ SUCCESS (gasless)
```

## Benefits

### 1. **Cost Efficiency**

- **Zero operational costs** for oracle validators
- Eliminates barrier to entry for smaller validators
- Reduces oracle infrastructure costs

### 2. **Reliability**

- **Highest priority** processing ensures votes aren't delayed
- **No gas estimation issues** - votes always process
- Improved oracle uptime and participation

### 3. **Security**

- **Robust spam prevention** without compromising functionality
- **Permission-based access** maintains security model
- **Transaction isolation** prevents bundling attacks

### 4. **Scalability**

- Supports high-frequency oracle operations
- No gas market congestion for oracle votes
- Predictable oracle response times

## Comparison with Sei Protocol

| Feature               | Sei Protocol                 | Six Protocol                         |
| --------------------- | ---------------------------- | ------------------------------------ |
| **Gasless Voting**    | ✅ Exchange rate votes       | ✅ NFT oracle votes                  |
| **Spam Prevention**   | ✅ Block-level + Vote period | ✅ Block-level + Duplicate detection |
| **Priority System**   | ✅ Highest priority          | ✅ Highest priority                  |
| **Vote Isolation**    | ✅ No bundling               | ✅ No bundling                       |
| **Permission System** | ✅ Validator-based           | ✅ nftadmin permission-based         |
| **DoS Protection**    | ✅ One vote per period       | ✅ One vote per request              |

## Security Considerations

### DoS Protection

- **Oracle Permission Required** - Only authorized addresses can submit gasless votes
- **Request Validation** - Votes only accepted for valid, pending requests
- **Duplicate Prevention** - Each oracle can vote once per request
- **Block-level Rate Limiting** - Prevents spam within single block

### Economic Security

- **Slashing Still Active** - Bad oracles can still be penalized
- **Reward Mechanisms Preserved** - Incentive structures remain intact
- **Gas for Non-Oracle Transactions** - Only oracle votes are gasless

### Attack Prevention

- **Transaction Bundling Blocked** - Oracle votes must be submitted alone
- **Invalid Vote Rejection** - Comprehensive validation before gasless approval
- **Cleanup Mechanisms** - Prevents storage bloat from tracking data

## Usage Examples

### Oracle Voting (Zero Gas)

```bash
# Oracle submits mint response with zero gas
sixd tx nftoracle submit-mint-response \
    "123" \
    "base64_encoded_nft_data" \
    --from oracle1 \
    --gas auto --gas-adjustment 1.5 \
    --chain-id sixnft -y

# Result: Transaction succeeds with 0 gas fees
```

### Regular Transaction (Normal Gas)

```bash
# Regular transfer still uses gas
sixd tx bank send alice bob 100usix \
    --gas auto --gas-adjustment 1.5 --gas-prices 2usix \
    --chain-id sixnft -y

# Result: Normal gas fees apply
```

## Future Enhancements

### Potential Improvements

1. **Configurable Gasless Types** - Admin control over which message types are gasless
2. **Oracle Reputation System** - Different gas policies based on oracle performance
3. **Dynamic Rate Limiting** - Adjust spam prevention based on network conditions
4. **Analytics Dashboard** - Monitor gasless transaction volume and patterns

### Monitoring & Metrics

- Track gasless transaction volume
- Monitor oracle participation rates
- Analyze spam prevention effectiveness
- Measure oracle response times

## Conclusion

The Zero-Gas Oracle Voting implementation provides Six Protocol with a robust, scalable, and secure oracle system that eliminates operational costs while maintaining strong spam prevention and security guarantees. This enhancement enables more reliable oracle operations and encourages broader validator participation in the oracle network.

The implementation follows proven patterns from Sei Protocol while adapting to Six Protocol's unique NFT-focused oracle requirements, providing a solid foundation for high-frequency, cost-effective oracle operations.
