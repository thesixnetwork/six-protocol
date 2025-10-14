#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
RPC_ENDPOINT=http://localhost:26657
CHAIN_ID=testnet
GAS_PRICES=2usix

# Function to print colored messages
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# Function to wait for block confirmation
wait_for_block() {
    log_info "Waiting for block confirmation..."
    sleep 3
}

# Function to grant oracle permission
grantOracle() {
    local oracle_address=$1
    local oracle_name=$2
    
    log_info "Granting 'oracle' permission to $oracle_name ($oracle_address)"
    
    local result=$(sixd tx nftadmin grant-permission oracle $oracle_address \
        --from super-admin \
        --gas auto --gas-adjustment 1.5 --gas-prices $GAS_PRICES \
        -y --node $RPC_ENDPOINT --chain-id $CHAIN_ID \
        --output json 2>&1)
    
    if echo "$result" | grep -q '"code":0'; then
        local txhash=$(echo "$result" | jq -r '.txhash // empty')
        log_success "$oracle_name oracle permission granted! TxHash: $txhash"
    else
        log_error "Failed to grant oracle permission to $oracle_name"
        echo "$result"
        return 1
    fi
    
    wait_for_block
}

log_info "Starting Zero-Gas Oracle Voting Setup..."

# Check if nft-schema.json exists
if [ ! -f "nft-schema.json" ]; then
    log_error "nft-schema.json file not found!"
    exit 1
fi

BASE64_SCHEMA=$(cat nft-schema.json | base64 | tr -d '\n')

log_info "Step 1: Granting oracle_admin permission to alice..."
result=$(sixd tx nftadmin grant-permission oracle_admin $(sixd keys show alice -a) \
    --from super-admin \
    --gas-prices $GAS_PRICES \
    -y --node $RPC_ENDPOINT --chain-id $CHAIN_ID \
    --output json 2>&1)

if echo "$result" | grep -q '"code":0'; then
    txhash=$(echo "$result" | jq -r '.txhash // empty')
    log_success "Oracle admin permission granted to alice! TxHash: $txhash"
else
    log_error "Failed to grant oracle_admin permission to alice"
    echo "$result"
    exit 1
fi

wait_for_block

log_info "Step 2: Setting minimum confirmation to 1..."
result=$(sixd tx nftoracle set-minimum-confirmation 1 \
    --from super-admin \
    --gas-prices $GAS_PRICES \
    -y --node $RPC_ENDPOINT --chain-id $CHAIN_ID \
    --output json 2>&1)

if echo "$result" | grep -q '"code":0'; then
    txhash=$(echo "$result" | jq -r '.txhash // empty')
    log_success "Minimum confirmation set to 1! TxHash: $txhash"
else
    log_error "Failed to set minimum confirmation"
    echo "$result"
    exit 1
fi

wait_for_block

log_info "Step 3: Creating NFT schema..."
result=$(sixd tx nftmngr create-nft-schema \
    --from alice \
    --gas auto --gas-adjustment 1.5 --gas-prices $GAS_PRICES \
    -y --chain-id $CHAIN_ID --node $RPC_ENDPOINT \
    --output json \
    $BASE64_SCHEMA 2>&1)

if echo "$result" | grep -q '"code":0'; then
    txhash=$(echo "$result" | jq -r '.txhash // empty')
    log_success "NFT schema created! TxHash: $txhash"
else
    log_error "Failed to create NFT schema"
    echo "$result"
    exit 1
fi

wait_for_block

log_info "Step 4: Granting oracle permissions to oracle accounts..."

# Grant oracle permissions
grantOracle $(sixd keys show oracle1 -a) "oracle1" || exit 1
grantOracle $(sixd keys show oracle2 -a) "oracle2" || exit 1
grantOracle $(sixd keys show oracle3 -a) "oracle3" || exit 1
grantOracle $(sixd keys show oracle4 -a) "oracle4" || exit 1

log_info "Step 5: Verifying setup..."
log_info "Checking oracle permissions..."

sixd q nftadmin show-authorization --node $RPC_ENDPOINT

log_success "🎉 Zero-gas oracle voting setup completed successfully!"
log_info "Available oracles: oracle1, oracle2, oracle3, oracle4"
log_info "Oracle admin: alice"
log_info "Minimum confirmation: 1"
log_info "NFT schema: six-protocol.example"
