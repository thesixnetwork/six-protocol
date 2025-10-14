#!/bin/bash

# Quick Oracle Gas Analysis Demo
# Demonstrates gas price logging for oracle voting

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

# Demo function showing gas analysis implementation
demo_oracle_gas_analysis() {
    local oracle_name="oracle1"
    local oracle_addr=$(sixd keys show ${oracle_name} -a --keyring-backend ${KEYRING_BACKEND} 2>/dev/null)
    
    if [ -z "$oracle_addr" ]; then
        log_warning "Oracle1 address not found"
        return 1
    fi
    
    log_info "🔥 Oracle Gas Analysis Demo for ${oracle_name} (${oracle_addr})"
    log_info "   ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    
    # Get oracle balance
    balance=$(sixd query bank balances "$oracle_addr" --node ${RPC_ENDPOINT} --output json 2>/dev/null | jq -r '.balances[] | select(.denom=="usix") | .amount // "0"' 2>/dev/null)
    log_info "   💰 Current Balance: $balance usix"
    
    # Show gas configuration
    log_info "   ⚙️  Gas Configuration:"
    log_info "      🎯 Gas Prices: ${GAS_PRICES}"
    log_info "      🔧 Gas Mode: auto with 1.5x adjustment"
    log_info "      🏗️  RPC Endpoint: ${RPC_ENDPOINT}"
    
    # Show what the gas analysis would capture
    log_info "   📊 Gas Analysis Features:"
    log_success "      ✅ Balance Before/After Comparison"
    log_success "      ✅ Transaction Hash Tracking"
    log_success "      ✅ Gas Used vs Gas Wanted Analysis"
    log_success "      ✅ Fee Amount and Denomination"
    log_success "      ✅ Gas Efficiency Calculations"
    log_success "      ✅ Effective Gas Price Computation"
    log_success "      ✅ Cost in SIX Token Conversion"
    log_success "      ✅ Zero-Gas Detection"
    log_success "      ✅ Block Height Recording"
    
    log_info ""
    log_info "   🔬 Zero-Gas Analysis Capabilities:"
    log_info "      🎉 ZERO-GAS CONFIRMED - when fee_amount = 0"
    log_info "      ✨ BALANCE UNCHANGED - when balance_change = 0"
    log_info "      ⭐ MINIMAL COST - when fee < 1000 usix (< 0.001 SIX)"
    log_info "      ⚠️  STANDARD GAS - when normal fees apply"
    
    log_info ""
    log_info "   📈 Aggregate Statistics Tracking:"
    log_info "      🔢 Total Oracles Voted"
    log_info "      ⛽ Total Gas Used (all oracles)"
    log_info "      💵 Total Fees Paid"
    log_info "      📊 Average Gas per Oracle"
    log_info "      💰 Average Fee per Oracle"
    log_info "      🪙  Total Cost in SIX Tokens"
}

# Show example transaction analysis
demo_transaction_analysis() {
    log_info "🔍 Example Transaction Gas Analysis Output:"
    log_info "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
    
    # Show sample analysis (simulated values)
    cat << EOF
   📊 Detailed Gas Analysis:
      🏗️  Block Height: 12345
      ⛽ Gas Used: 89,432 units
      🎯 Gas Wanted: 95,000 units
      💵 Fee Charged: 178,864 usix
      ⚙️  Gas Price Setting: 2usix
      💰 Balance Before: 1,000,000 usix
      💰 Balance After: 999,821,136 usix
      💸 Balance Change: 178,864 usix
      📈 Gas Efficiency: 94.15% (actual/requested)
      💎 Effective Gas Price: 2.000000 usix/gas
      🪙  Total Cost: 0.178864 SIX tokens

   🔬 Zero-Gas Analysis:
      ⚠️  STANDARD GAS: 178,864 usix charged

🏁 Final Comprehensive Analysis:
  🔢 Total Oracles Voted: 3
  ⛽ Total Gas Used: 268,296 units  
  💵 Total Fees Paid: 536,592 usix
  📊 Average Gas per Oracle: 89,432 units
  💰 Average Fee per Oracle: 178,864 usix
  🪙  Total Cost in SIX: 0.536592 SIX tokens

🔬 Zero-Gas Oracle Voting Summary:
  ⚠️  Standard gas fees applied. Total: 536,592 usix
EOF
}

main() {
    log_success "🚀 Oracle Gas Analysis Implementation Demo"
    log_info "========================================="
    echo
    
    demo_oracle_gas_analysis
    echo
    demo_transaction_analysis
    echo
    
    log_success "🏁 Gas Analysis Implementation Complete!"
    log_info "The gas logging has been implemented with comprehensive analytics:"
    log_info "• Real-time balance tracking"  
    log_info "• Transaction hash monitoring"
    log_info "• Gas efficiency calculations"
    log_info "• Zero-gas detection"
    log_info "• Cost analysis in SIX tokens"
    log_info "• Aggregate statistics"
}

main "$@"
