#!/bin/bash

# Oracle Consensus Test with Detailed Gas Analysis
# Test oracle voting with comprehensive gas price and fee tracking

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
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }

wait_for_block() {
    echo "Waiting 3 seconds for block..."
    sleep 3
}

# Global variables to track total costs
total_gas_used=0
total_fee_paid=0
oracle_count=0
oracle_tx_hashes=()  # Array to store all transaction hashes

# Test Oracle Voting with Gas Analysis
test_oracle_voting_with_gas() {
    log_info "=== Testing Oracle Voting with Detailed Gas Analysis ==="
    
    # Verify current configuration
    oracle_config=$(sixd q nftoracle show-oracle-config --node ${RPC_ENDPOINT} 2>/dev/null)
    log_info "Current oracle configuration:"
    echo "$oracle_config"
    
    # Get alice address
    alice_addr=$(sixd keys show alice -a --keyring-backend ${KEYRING_BACKEND})
    log_info "Alice address: $alice_addr"
    
    # Create mint request
    token_name="gas-analysis-token-$(date +%s)"
    log_info "Creating mint request for token: ${token_name}"
    
    log_info "Executing mint request command..."
    request_result=$(sixd tx nftoracle create-mint-request \
        six-protocol.example \
        "${token_name}" \
        "3" \
        --from alice \
        --keyring-backend ${KEYRING_BACKEND} \
        --gas auto --gas-adjustment 1.5 --gas-prices ${GAS_PRICES} \
        --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} -y --output json 2>&1)
    
    if echo "$request_result" | grep -q '"code":0'; then
        log_success "✅ Mint request created successfully"
        
        # Extract transaction hash (handle gas estimate prefix)
        tx_hash=$(echo "$request_result" | grep -o '"txhash":"[^"]*"' | cut -d'"' -f4 2>/dev/null)
        if [ -z "$tx_hash" ]; then
            # Fallback: try to extract JSON part and parse
            json_part=$(echo "$request_result" | grep -o '{.*}' | head -1)
            tx_hash=$(echo "$json_part" | jq -r '.txhash // "unknown"' 2>/dev/null)
        fi
        [ -z "$tx_hash" ] && tx_hash="unknown"
        log_success "🔗 Mint Request Transaction Hash: $tx_hash"
        log_info "📋 Mint Request Details:"
        log_info "   🎯 Token Name: ${token_name}"
        log_info "   📝 Schema: six-protocol.example"
        log_info "   🔗 TX Hash: $tx_hash"
        
        wait_for_block
        
        # Get mint requests to find our request ID
        log_info "Fetching mint requests..."
        mint_requests=$(sixd query nftoracle list-mint-request --node ${RPC_ENDPOINT} --output json 2>/dev/null)
        
        if [ $? -eq 0 ]; then
            # Find our request
            request_id=$(echo "$mint_requests" | jq -r --arg token "$token_name" '.MintRequest | map(select(.token_id == $token)) | .[0].id // empty' 2>/dev/null)
            
            if [ -z "$request_id" ] || [ "$request_id" = "null" ]; then
                # Fallback: get the latest ID
                request_id=$(echo "$mint_requests" | jq -r '.MintRequest | map(.id | tonumber) | max // empty' 2>/dev/null)
            fi
            
            if [ -n "$request_id" ] && [ "$request_id" != "null" ]; then
                log_success "🎯 Found mint request ID: $request_id"
                
                # Show initial status
                log_info "Initial request status:"
                initial_status=$(sixd query nftoracle show-mint-request "$request_id" --node ${RPC_ENDPOINT} 2>/dev/null)
                echo "$initial_status"
                
                # Now test oracle voting with gas analysis
                log_info "🗳️  Starting Oracle Voting Process with Gas Analysis..."
                
                # Create identical NFT data for consensus (MUST be identical for all oracles)
                # Use ONLY valid trait types from the schema: Background, Moon, Plate, Tail, Whale
                nft_data='{"image":"https://nft.sixnetwork.io/metadata/test/gas-analysis.png","holder_address":"'${alice_addr}'","traits":[{"trait_type":"Background","value":"Gas Analysis Blue"},{"trait_type":"Moon","value":"Zero Gas Moon"},{"trait_type":"Plate","value":"Oracle Plate"}]}'
                
                # Create base64 data ONCE to ensure all oracles use identical data
                base64_nft_data=$(echo "$nft_data" | base64 | tr -d '\n')
                
                log_info "NFT Data: $nft_data"
                log_info "Base64 Data: $base64_nft_data"
                echo
                
                # Test with 3 oracles and detailed gas analysis
                for i in 1 2 3; do
                    oracle_name="oracle${i}"
                    oracle_addr=$(sixd keys show ${oracle_name} -a --keyring-backend ${KEYRING_BACKEND})
                    
                    log_info "🔥 Oracle${i} (${oracle_addr}) Gas Analysis (${i}/3):"
                    log_info "   ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
                    
                    # Get oracle balance before voting
                    before_balance=$(sixd query bank balances "$oracle_addr" --node ${RPC_ENDPOINT} --output json 2>/dev/null | jq -r '.balances[] | select(.denom=="usix") | .amount // "0"' 2>/dev/null)
                    log_info "   💰 Balance before voting: $before_balance usix"
                    
                    # Submit oracle vote
                    log_info "   📤 Submitting vote..."
                    oracle_result=$(sixd tx nftoracle submit-mint-response \
                        "$request_id" \
                        "$base64_nft_data" \
                        --from ${oracle_name} \
                        --keyring-backend ${KEYRING_BACKEND} \
                        --gas auto --gas-adjustment 1.5 --gas-prices ${GAS_PRICES} \
                        --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} -y --output json 2>&1)
                    
                    if echo "$oracle_result" | grep -q '"code":0'; then
                        # Extract transaction hash (handle gas estimate prefix)
                        oracle_tx_hash=$(echo "$oracle_result" | grep -o '"txhash":"[^"]*"' | cut -d'"' -f4 2>/dev/null)
                        if [ -z "$oracle_tx_hash" ]; then
                            # Fallback: try to extract JSON part and parse
                            json_part=$(echo "$oracle_result" | grep -o '{.*}' | head -1)
                            oracle_tx_hash=$(echo "$json_part" | jq -r '.txhash // "unknown"' 2>/dev/null)
                        fi
                        [ -z "$oracle_tx_hash" ] && oracle_tx_hash="unknown"
                        log_success "   ✅ Oracle${i} vote submitted successfully!"
                        log_info "   🔗 Oracle${i} Transaction Hash: $oracle_tx_hash"
                        log_info "   📋 Oracle${i} Vote Details:"
                        log_info "      🆔 Request ID: $request_id"
                        log_info "      👤 Oracle Address: $oracle_addr"
                        log_info "      🔗 TX Hash: $oracle_tx_hash"
                        
                        # Query transaction details for comprehensive gas analysis
                        if [ "$oracle_tx_hash" != "unknown" ]; then
                            log_info "   🔍 Analyzing transaction gas usage..."
                            sleep 4  # Wait for transaction to be indexed
                            
                            tx_details=$(sixd query tx "$oracle_tx_hash" --node ${RPC_ENDPOINT} --output json 2>/dev/null)
                            
                            if [ $? -eq 0 ] && [ -n "$tx_details" ]; then
                                # Extract comprehensive gas and fee information
                                gas_used=$(echo "$tx_details" | jq -r '.gas_used // "0"' 2>/dev/null)
                                gas_wanted=$(echo "$tx_details" | jq -r '.gas_wanted // "0"' 2>/dev/null)
                                fee_amount=$(echo "$tx_details" | jq -r '.tx.auth_info.fee.amount[0].amount // "0"' 2>/dev/null)
                                fee_denom=$(echo "$tx_details" | jq -r '.tx.auth_info.fee.amount[0].denom // "usix"' 2>/dev/null)
                                block_height=$(echo "$tx_details" | jq -r '.height // "unknown"' 2>/dev/null)
                                
                                # Get oracle balance after voting
                                after_balance=$(sixd query bank balances "$oracle_addr" --node ${RPC_ENDPOINT} --output json 2>/dev/null | jq -r '.balances[] | select(.denom=="usix") | .amount // "0"' 2>/dev/null)
                                
                                # Calculate balance change
                                balance_change=0
                                if [ "$before_balance" != "0" ] && [ "$after_balance" != "0" ]; then
                                    balance_change=$((before_balance - after_balance))
                                fi
                                
                                log_info "   📊 Detailed Gas Analysis:"
                                log_info "      🏗️  Block Height: $block_height"
                                log_info "      ⛽ Gas Used: $gas_used units"
                                log_info "      🎯 Gas Wanted: $gas_wanted units"
                                log_info "      💵 Fee Charged: $fee_amount $fee_denom"
                                log_info "      ⚙️  Gas Price Setting: ${GAS_PRICES}"
                                log_info "      💰 Balance Before: $before_balance usix"
                                log_info "      💰 Balance After: $after_balance usix"
                                log_info "      💸 Balance Change: $balance_change usix"
                                
                                # Calculate gas efficiency
                                if [ "$gas_used" != "0" ] && [ "$gas_wanted" != "0" ]; then
                                    efficiency=$(echo "scale=2; $gas_used * 100 / $gas_wanted" | bc 2>/dev/null || echo "N/A")
                                    log_info "      📈 Gas Efficiency: ${efficiency}% (actual/requested)"
                                fi
                                
                                # Calculate effective gas price
                                if [ "$fee_amount" != "0" ] && [ "$gas_used" != "0" ]; then
                                    effective_gas_price=$(echo "scale=6; $fee_amount / $gas_used" | bc 2>/dev/null || echo "N/A")
                                    log_info "      💎 Effective Gas Price: ${effective_gas_price} ${fee_denom}/gas"
                                fi
                                
                                # Calculate cost in SIX tokens (assuming 6 decimals)
                                if [ "$fee_amount" != "0" ]; then
                                    cost_in_six=$(echo "scale=6; $fee_amount / 1000000" | bc 2>/dev/null || echo "N/A")
                                    log_info "      🪙  Total Cost: ${cost_in_six} SIX tokens"
                                fi
                                
                                # Zero-gas analysis
                                log_info "   🔬 Zero-Gas Analysis:"
                                if [ "$fee_amount" = "0" ]; then
                                    log_success "      🎉 ZERO-GAS CONFIRMED! No fees charged for oracle voting!"
                                elif [ "$balance_change" = "0" ]; then
                                    log_success "      ✨ BALANCE UNCHANGED! Oracle voting appears to be gasless!"
                                elif [ "$fee_amount" -lt "1000" ]; then  # Less than 0.001 SIX
                                    log_success "      ⭐ MINIMAL COST: Only ${fee_amount} ${fee_denom} (< 0.001 SIX)"
                                else
                                    log_warning "      ⚠️  STANDARD GAS: ${fee_amount} ${fee_denom} charged"
                                fi
                                
                                # Add to totals
                                if [ "$gas_used" != "0" ]; then
                                    total_gas_used=$((total_gas_used + gas_used))
                                fi
                                if [ "$fee_amount" != "0" ]; then
                                    total_fee_paid=$((total_fee_paid + fee_amount))
                                fi
                                oracle_count=$((oracle_count + 1))
                                oracle_tx_hashes+=("Oracle${i}:$oracle_tx_hash")  # Store the transaction hash
                                log_info "      📊 Running Totals: Oracle#${oracle_count}, Gas: ${total_gas_used}, Fees: ${total_fee_paid} usix"
                                
                            else
                                log_warning "   ⚠️  Could not query transaction details for gas analysis"
                                oracle_count=$((oracle_count + 1))  # Still count the oracle even if gas analysis failed
                                oracle_tx_hashes+=("Oracle${i}:$oracle_tx_hash")  # Store the transaction hash even if gas analysis failed
                                log_info "      📊 Oracle counted without gas data: Oracle#${oracle_count}"
                            fi
                        else
                            # Even if we couldn't get transaction hash, the oracle still voted successfully
                            oracle_count=$((oracle_count + 1))
                            oracle_tx_hashes+=("Oracle${i}:UNKNOWN")  # Store unknown hash
                            log_warning "      ⚠️  Oracle${i} Transaction Hash: UNKNOWN (extraction failed)"
                            log_info "      📊 Oracle voted (no tx hash): Oracle#${oracle_count}"
                        fi
                        
                        wait_for_block
                        
                        # Check current consensus status
                        current_status=$(sixd query nftoracle show-mint-request "$request_id" --node ${RPC_ENDPOINT} --output json 2>/dev/null)
                        current_confirm=$(echo "$current_status" | jq -r '.MintRequest.current_confirm // "0"' 2>/dev/null)
                        status=$(echo "$current_status" | jq -r '.MintRequest.status // "UNKNOWN"' 2>/dev/null)
                        
                        log_info "   📋 Consensus Status: ${current_confirm}/3 confirmations"
                        log_info "   🎯 Request Status: $status"
                        
                        if [ "$status" = "SUCCESS_WITH_CONSENSUS" ]; then
                            log_success "   🎉 CONSENSUS ACHIEVED after Oracle${i} vote!"
                            log_info "   🏆 Final Oracle Count: $oracle_count voted successfully"
                            break
                        fi
                        
                    else
                        log_error "   ❌ Oracle${i} vote failed:"
                        echo "$oracle_result" | jq '.' 2>/dev/null || echo "$oracle_result"
                    fi
                    
                    echo  # Add spacing between oracles
                done
                
                # Final comprehensive analysis
                log_info "🏁 Final Comprehensive Gas Analysis:"
                log_info "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
                
                final_status=$(sixd query nftoracle show-mint-request "$request_id" --node ${RPC_ENDPOINT} 2>/dev/null)
                echo "$final_status"
                
                # Aggregate statistics
                log_info ""
                log_info "📈 Aggregate Oracle Voting Statistics:"
                log_info "  🔢 Total Oracles Voted: $oracle_count"
                log_info "  ⛽ Total Gas Used: $total_gas_used units"
                log_info "  💵 Total Fees Paid: $total_fee_paid usix"
                
                if [ $oracle_count -gt 0 ]; then
                    avg_gas=$(echo "scale=0; $total_gas_used / $oracle_count" | bc 2>/dev/null || echo "N/A")
                    avg_fee=$(echo "scale=0; $total_fee_paid / $oracle_count" | bc 2>/dev/null || echo "N/A")
                    log_info "  📊 Average Gas per Oracle: $avg_gas units"
                    log_info "  💰 Average Fee per Oracle: $avg_fee usix"
                fi
                
                if [ $total_fee_paid -gt 0 ]; then
                    total_cost_six=$(echo "scale=6; $total_fee_paid / 1000000" | bc 2>/dev/null || echo "N/A")
                    log_info "  🪙  Total Cost in SIX: $total_cost_six SIX tokens"
                fi
                
                # Transaction Hash Summary
                log_info ""
                log_info "🔗 Oracle Voting Transaction Hashes Summary:"
                if [ ${#oracle_tx_hashes[@]} -gt 0 ]; then
                    for hash_entry in "${oracle_tx_hashes[@]}"; do
                        oracle_name=$(echo "$hash_entry" | cut -d':' -f1)
                        tx_hash=$(echo "$hash_entry" | cut -d':' -f2-)
                        if [ "$tx_hash" = "UNKNOWN" ]; then
                            log_warning "  ⚠️  $oracle_name: Transaction hash unavailable"
                        else
                            log_success "  ✅ $oracle_name: $tx_hash"
                        fi
                    done
                else
                    log_warning "  ⚠️  No transaction hashes recorded"
                fi

                # Zero-gas summary
                log_info ""
                log_info "🔬 Zero-Gas Oracle Voting Summary:"
                if [ $total_fee_paid -eq 0 ]; then
                    log_success "  🎉 PERFECT ZERO-GAS! All oracle votes were completely free!"
                elif [ $total_fee_paid -lt 3000 ]; then  # Less than 0.003 SIX total
                    log_success "  ⭐ ULTRA-LOW COST! Total fees under 0.003 SIX"
                else
                    log_warning "  ⚠️  Standard gas fees applied. Total: $total_fee_paid usix"
                fi
                
                # Check if NFT was minted
                log_info ""
                log_info "🎨 Checking NFT minting result..."
                nft_result=$(sixd query nftmngr show-nft-data six-protocol.example "${token_name}" --node ${RPC_ENDPOINT} 2>/dev/null)
                if [ $? -eq 0 ]; then
                    log_success "🎉 NFT successfully minted with oracle consensus!"
                    echo "$nft_result" | head -15
                else
                    log_warning "NFT not found - may still be processing"
                fi
                
            else
                log_error "Could not find request ID"
            fi
            
        else
            log_error "Failed to fetch mint requests"
        fi
        
    else
        log_error "❌ Mint request creation failed:"
        echo "$request_result"
    fi
}

# Main execution
main() {
    log_info "🚀 Oracle Consensus Voting with Detailed Gas Analysis"
    log_info "=================================================="
    
    # Verify we have the accounts
    for account in alice oracle1 oracle2 oracle3; do
        if sixd keys show ${account} --keyring-backend ${KEYRING_BACKEND} >/dev/null 2>&1; then
            log_success "✅ ${account} ready"
        else
            log_error "❌ ${account} not found"
            exit 1
        fi
    done
    
    echo
    test_oracle_voting_with_gas
    
    echo
    log_success "🏁 Oracle consensus gas analysis completed!"
    log_info "Summary: $oracle_count oracles voted, $total_gas_used total gas, $total_fee_paid usix total fees"
}

main "$@"
