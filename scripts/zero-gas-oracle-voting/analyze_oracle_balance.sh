#!/bin/bash

# Deep Balance Analysis for Oracle Voting
# Investigate why balance doesn't change despite gas usage

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

# Detailed balance analysis
analyze_oracle_balance_behavior() {
    local oracle_name="oracle1"
    local oracle_addr=$(sixd keys show ${oracle_name} -a --keyring-backend ${KEYRING_BACKEND})
    
    log_info "🔬 Deep Balance Analysis for Oracle Voting"
    log_info "Oracle: $oracle_name ($oracle_addr)"
    log_info "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    
    # Get detailed balance before
    log_info "📊 Pre-Transaction Analysis:"
    balance_before=$(sixd query bank balances "$oracle_addr" --node ${RPC_ENDPOINT} --output json 2>/dev/null)
    usix_before=$(echo "$balance_before" | jq -r '.balances[] | select(.denom=="usix") | .amount // "0"')
    all_balances_before=$(echo "$balance_before" | jq -r '.balances[] | "\(.amount) \(.denom)"' | tr '\n' ', ')
    
    log_info "   💰 USIX Balance: $usix_before usix"
    log_info "   💎 All Balances: $all_balances_before"
    
    # Get account details
    account_info=$(sixd query auth account "$oracle_addr" --node ${RPC_ENDPOINT} --output json 2>/dev/null)
    account_number=$(echo "$account_info" | jq -r '.account_number // "unknown"')
    sequence=$(echo "$account_info" | jq -r '.sequence // "unknown"')
    
    log_info "   🆔 Account Number: $account_number"
    log_info "   🔢 Sequence: $sequence"
    
    # Check if there's a recent oracle transaction to analyze
    log_info ""
    log_info "🔍 Searching for Recent Oracle Transactions..."
    
    # Get latest mint request to test with
    latest_request=$(sixd query nftoracle list-mint-request --node ${RPC_ENDPOINT} --output json 2>/dev/null | jq -r '.MintRequest | map(select(.status == "PENDING")) | .[0] // empty')
    
    if [ -z "$latest_request" ] || [ "$latest_request" = "null" ]; then
        log_warning "No pending requests found. Creating a new one for analysis..."
        
        alice_addr=$(sixd keys show alice -a --keyring-backend ${KEYRING_BACKEND})
        token_name="balance-analysis-$(date +%s)"
        
        # Create mint request
        mint_result=$(sixd tx nftoracle create-mint-request \
            six-protocol.example \
            "${token_name}" \
            "3" \
            --from alice \
            --keyring-backend ${KEYRING_BACKEND} \
            --gas auto --gas-adjustment 1.5 --gas-prices ${GAS_PRICES} \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} -y --output json 2>&1)
        
        sleep 4
        latest_request=$(sixd query nftoracle list-mint-request --node ${RPC_ENDPOINT} --output json 2>/dev/null | jq -r '.MintRequest | map(select(.token_id == "'${token_name}'")) | .[0] // empty')
    fi
    
    request_id=$(echo "$latest_request" | jq -r '.id // "unknown"')
    
    if [ "$request_id" != "unknown" ] && [ "$request_id" != "null" ]; then
        log_success "✅ Found/Created Request ID: $request_id"
        
        # Perform oracle vote with detailed monitoring
        log_info ""
        log_info "🗳️  Performing Oracle Vote with Balance Monitoring..."
        
        # Create NFT data
        alice_addr=$(sixd keys show alice -a --keyring-backend ${KEYRING_BACKEND})
        nft_data='{"image":"https://nft.sixnetwork.io/metadata/test/balance-analysis.png","holder_address":"'${alice_addr}'","traits":[{"trait_type":"Background","value":"Balance Test"},{"trait_type":"Moon","value":"Analysis Moon"}]}'
        base64_nft_data=$(echo "$nft_data" | base64 | tr -d '\n')
        
        # Submit oracle vote with timestamps
        log_info "⏰ Timestamp: $(date '+%Y-%m-%d %H:%M:%S')"
        log_info "💰 Balance Before Vote: $usix_before usix"
        
        vote_result=$(sixd tx nftoracle submit-mint-response \
            "$request_id" \
            "$base64_nft_data" \
            --from ${oracle_name} \
            --keyring-backend ${KEYRING_BACKEND} \
            --gas auto --gas-adjustment 1.5 --gas-prices ${GAS_PRICES} \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} -y --output json 2>&1)
        
        log_info "⏰ Vote Submitted: $(date '+%Y-%m-%d %H:%M:%S')"
        
        # Extract transaction hash
        tx_hash=$(echo "$vote_result" | grep -o '"txhash":"[^"]*"' | cut -d'"' -f4 2>/dev/null)
        
        if [ -n "$tx_hash" ]; then
            log_success "🔗 Transaction Hash: $tx_hash"
            
            # Wait and check balance immediately
            sleep 2
            balance_immediate=$(sixd query bank balances "$oracle_addr" --node ${RPC_ENDPOINT} --output json 2>/dev/null)
            usix_immediate=$(echo "$balance_immediate" | jq -r '.balances[] | select(.denom=="usix") | .amount // "0"')
            log_info "💰 Balance Immediately After: $usix_immediate usix"
            
            # Wait for transaction indexing
            sleep 4
            
            # Get final balance
            balance_after=$(sixd query bank balances "$oracle_addr" --node ${RPC_ENDPOINT} --output json 2>/dev/null)
            usix_after=$(echo "$balance_after" | jq -r '.balances[] | select(.denom=="usix") | .amount // "0"')
            log_info "💰 Balance Final: $usix_after usix"
            
            # Calculate changes
            immediate_change=$((usix_before - usix_immediate))
            final_change=$((usix_before - usix_after))
            
            log_info ""
            log_info "📊 Balance Change Analysis:"
            log_info "   📉 Immediate Change: $immediate_change usix"
            log_info "   📉 Final Change: $final_change usix"
            
            # Get transaction details
            tx_details=$(sixd query tx "$tx_hash" --node ${RPC_ENDPOINT} --output json 2>/dev/null)
            
            if [ $? -eq 0 ] && [ -n "$tx_details" ]; then
                gas_used=$(echo "$tx_details" | jq -r '.gas_used // "0"')
                fee_amount=$(echo "$tx_details" | jq -r '.tx.auth_info.fee.amount[0].amount // "0"')
                
                log_info ""
                log_info "🔬 Transaction Analysis:"
                log_info "   ⛽ Gas Used: $gas_used units"
                log_info "   💵 Fee Amount: $fee_amount usix"
                log_info "   💰 Expected Deduction: $fee_amount usix"
                log_info "   💰 Actual Deduction: $final_change usix"
                
                # Check for fee events in transaction
                fee_events=$(echo "$tx_details" | jq -r '.logs[].events[]? | select(.type=="tx") | .attributes[]? | select(.key=="fee") | .value // empty' 2>/dev/null)
                if [ -n "$fee_events" ]; then
                    log_info "   🎫 Fee Events: $fee_events"
                fi
                
                # Analysis conclusion
                log_info ""
                log_info "🎯 Analysis Conclusion:"
                if [ "$final_change" = "0" ] && [ "$fee_amount" != "0" ]; then
                    log_success "   🎉 ZERO-GAS ORACLE CONFIRMED!"
                    log_success "   ✨ Oracle voting is truly gasless despite gas calculation"
                    log_info "   🏛️  This suggests protocol-level gas sponsorship for oracles"
                elif [ "$final_change" = "$fee_amount" ]; then
                    log_warning "   ⚠️  Standard gas fees applied"
                elif [ "$immediate_change" != "$final_change" ]; then
                    log_success "   🔄 Possible gas refund mechanism detected"
                else
                    log_info "   🤔 Unusual balance behavior detected"
                fi
                
            else
                log_warning "Could not retrieve transaction details"
            fi
            
        else
            log_warning "Could not extract transaction hash"
        fi
        
    else
        log_warning "Could not find or create a suitable mint request for testing"
    fi
}

main() {
    log_success "🚀 Oracle Balance Behavior Analysis"
    echo
    analyze_oracle_balance_behavior
    echo
    log_success "🏁 Analysis completed"
}

main "$@"
