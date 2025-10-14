#!/bin/bash

# Quick Test to Debug Transaction Hash Extraction
# This will help us understand why transaction hashes are empty

# Configuration
RPC_ENDPOINT=http://localhost:26657
CHAIN_ID=testnet
GAS_PRICES=2usix
KEYRING_BACKEND=test

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}[SUCCESS]${NC} $1"; }
log_warning() { echo -e "${YELLOW}[WARNING]${NC} $1"; }

# Test transaction hash extraction
test_tx_hash_extraction() {
    log_info "🔍 Testing Transaction Hash Extraction"
    
    alice_addr=$(sixd keys show alice -a --keyring-backend ${KEYRING_BACKEND})
    token_name="debug-tx-hash-$(date +%s)"
    
    log_info "Creating mint request to test transaction hash..."
    
    # Capture the full result for debugging
    request_result=$(sixd tx nftoracle create-mint-request \
        six-protocol.example \
        "${token_name}" \
        "3" \
        --from alice \
        --keyring-backend ${KEYRING_BACKEND} \
        --gas auto --gas-adjustment 1.5 --gas-prices ${GAS_PRICES} \
        --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} -y --output json 2>&1)
    
    log_info "🔍 Full Transaction Response:"
    echo "=================================="
    echo "$request_result"
    echo "=================================="
    
    # Test different extraction methods
    log_info ""
    log_info "🧪 Testing Different Hash Extraction Methods:"
    
    # Method 1: Standard jq extraction
    tx_hash_1=$(echo "$request_result" | jq -r '.txhash // "EMPTY"' 2>/dev/null)
    log_info "Method 1 (jq .txhash): '$tx_hash_1'"
    
    # Method 2: Alternative field names
    tx_hash_2=$(echo "$request_result" | jq -r '.hash // "EMPTY"' 2>/dev/null)
    log_info "Method 2 (jq .hash): '$tx_hash_2'"
    
    # Method 3: Check if it's nested
    tx_hash_3=$(echo "$request_result" | jq -r '.tx_response.txhash // "EMPTY"' 2>/dev/null)
    log_info "Method 3 (jq .tx_response.txhash): '$tx_hash_3'"
    
    # Method 4: Raw grep
    tx_hash_4=$(echo "$request_result" | grep -o '"txhash":"[^"]*"' | cut -d'"' -f4 2>/dev/null || echo "EMPTY")
    log_info "Method 4 (grep txhash): '$tx_hash_4'"
    
    # Method 5: Check all fields
    log_info ""
    log_info "🔍 All Available Fields in Response:"
    echo "$request_result" | jq 'keys // empty' 2>/dev/null || echo "Not valid JSON or no keys found"
    
    # Check response status
    response_code=$(echo "$request_result" | jq -r '.code // "UNKNOWN"' 2>/dev/null)
    log_info "Response Code: $response_code"
    
    if echo "$request_result" | grep -q '"code":0'; then
        log_success "✅ Transaction appears successful (code 0)"
    else
        log_warning "⚠️  Transaction may have failed or response format different"
    fi
}

main() {
    log_success "🚀 Transaction Hash Extraction Debug Test"
    echo
    test_tx_hash_extraction
    echo
    log_success "🏁 Debug test completed"
}

main "$@"
