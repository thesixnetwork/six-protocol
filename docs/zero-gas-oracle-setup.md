# Zero-Gas Oracle Voting - Quick Setup Guide

## Prerequisites

1. **Running Six Protocol Node**
   ```bash
   # Start your local testnet
   sixd start --minimum-gas-prices="0usix"
   ```

2. **Required Accounts**
   ```bash
   # Create/import these accounts in your keyring
   sixd keys add super-admin    # Admin account
   sixd keys add alice         # Oracle admin
   sixd keys add oracle1       # Oracle validator 1  
   sixd keys add oracle2       # Oracle validator 2
   sixd keys add oracle3       # Oracle validator 3
   ```

3. **Fund Accounts**
   ```bash
   # Fund accounts with initial balance (for non-gasless transactions)
   sixd tx bank send <faucet> $(sixd keys show alice -a) 1000000usix
   sixd tx bank send <faucet> $(sixd keys show oracle1 -a) 1000000usix
   # ... repeat for other accounts
   ```

## Oracle Setup

### 1. Grant Oracle Permissions
```bash
# Grant oracle admin permission to alice
sixd tx nftadmin grant-permission oracle_admin $(sixd keys show alice -a) \
    --from super-admin --chain-id testnet -y

# Grant oracle permission to oracle accounts
sixd tx nftadmin grant-permission oracle $(sixd keys show oracle1 -a) \
    --from alice --chain-id testnet -y

sixd tx nftadmin grant-permission oracle $(sixd keys show oracle2 -a) \
    --from alice --chain-id testnet -y

sixd tx nftadmin grant-permission oracle $(sixd keys show oracle3 -a) \
    --from alice --chain-id testnet -y
```

### 2. Configure Oracle Settings
```bash
# Set minimum confirmation (number of oracle votes needed)
sixd tx nftoracle set-minimum-confirmation 2 \
    --from alice --chain-id testnet -y

# Create NFT schema (required for mint requests)
BASE64_SCHEMA=$(cat nft-schema.json | base64 | tr -d '\n')
sixd tx nftmngr create-nft-schema \
    --from alice --gas auto --gas-adjustment 1.5 \
    --chain-id testnet ${BASE64_SCHEMA} -y
```

## Testing Zero-Gas Oracle Voting

### 1. Create Mint Request
```bash
# Create a mint request (this uses normal gas)
sixd tx nftoracle create-mint-request \
    "six-protocol.example" \
    "test-token-$(date +%s)" \
    "2" \
    --from alice \
    --gas auto --gas-adjustment 1.5 --gas-prices 2usix \
    --chain-id testnet -y
```

### 2. Check Oracle Balance (Before)
```bash
# Record oracle balance before voting
ORACLE1_ADDR=$(sixd keys show oracle1 -a)
BEFORE_BALANCE=$(sixd query bank balances "$ORACLE1_ADDR" --output json | \
    jq -r '.balances[] | select(.denom=="usix") | .amount')

echo "Oracle1 balance before: $BEFORE_BALANCE usix"
```

### 3. Submit Zero-Gas Oracle Vote
```bash
# Get the mint request ID
REQUEST_ID=$(sixd query nftoracle list-mint-request --output json | \
    jq -r '.MintRequest[-1].id')

# Create NFT metadata
NFT_DATA='{
    "image": "https://example.com/nft.png", 
    "holder_address": "'$(sixd keys show alice -a)'",
    "traits": [{"trait_type": "Test", "value": "Zero-Gas"}]
}'

BASE64_NFT_DATA=$(echo "$NFT_DATA" | base64 | tr -d '\n')

# Submit oracle response WITHOUT gas prices (zero-gas!)
sixd tx nftoracle submit-mint-response \
    "$REQUEST_ID" \
    "$BASE64_NFT_DATA" \
    --from oracle1 \
    --gas auto --gas-adjustment 1.5 \
    --chain-id testnet -y
```

### 4. Verify Zero-Gas Operation
```bash
# Check oracle balance after voting
AFTER_BALANCE=$(sixd query bank balances "$ORACLE1_ADDR" --output json | \
    jq -r '.balances[] | select(.denom=="usix") | .amount')

echo "Oracle1 balance after: $AFTER_BALANCE usix"
echo "Gas used: $((BEFORE_BALANCE - AFTER_BALANCE)) usix"

# Should show 0 gas used!
if [ "$BEFORE_BALANCE" = "$AFTER_BALANCE" ]; then
    echo "✅ SUCCESS: Zero-gas oracle voting working!"
else
    echo "❌ ISSUE: Gas was charged ($((BEFORE_BALANCE - AFTER_BALANCE)) usix)"
fi
```

### 5. Complete Oracle Consensus
```bash
# Submit second oracle vote to reach consensus
sixd tx nftoracle submit-mint-response \
    "$REQUEST_ID" \
    "$BASE64_NFT_DATA" \
    --from oracle2 \
    --gas auto --gas-adjustment 1.5 \
    --chain-id testnet -y

# Check if mint request is completed
sixd query nftoracle show-mint-request "$REQUEST_ID"
```

## Test Script

Use the enhanced test script:
```bash
# Make executable
chmod +x scripts/tests/test_oracle_voting_detailed.sh

# Run zero-gas tests
./scripts/tests/test_oracle_voting_detailed.sh
```

## Troubleshooting

### Common Issues

1. **"No oracle permission" Error**
   ```bash
   # Verify oracle permission
   sixd query nftadmin show-authorization
   # Re-grant if needed
   ```

2. **"Oracle already voted" Error**
   ```bash
   # This is expected - each oracle can only vote once per request
   # Use a different oracle account
   ```

3. **"Request not found" Error**
   ```bash
   # Check if request exists
   sixd query nftoracle list-mint-request
   # Create new request if needed
   ```

4. **Gas Still Being Charged**
   ```bash
   # Check if you're using --gas-prices flag (remove it for oracle votes)
   # Verify transaction is oracle vote message type
   # Check ante handler configuration
   ```

### Verification Commands

```bash
# Check oracle permissions
sixd query nftadmin show-authorization

# List all mint requests  
sixd query nftoracle list-mint-request

# Check oracle configuration
sixd query nftoracle show-oracle-config

# View transaction details
sixd query tx <TX_HASH>
```

## Expected Results

When zero-gas oracle voting is working correctly:

1. **Oracle votes cost 0 gas** - Balance unchanged after voting
2. **High priority processing** - Votes confirmed quickly  
3. **Spam prevention works** - Duplicate votes rejected
4. **Regular transactions still use gas** - Only oracle votes are free
5. **Transaction isolation** - Oracle votes cannot be bundled

## Next Steps

- Test different oracle message types (action responses, collection verification)
- Monitor oracle participation rates
- Test spam prevention mechanisms
- Implement monitoring and alerting

For detailed implementation information, see `docs/zero-gas-oracle-voting.md`.
