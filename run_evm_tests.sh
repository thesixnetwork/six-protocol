#!/bin/bash

# EVM Migration Test Runner Script - Post v4.0.0 Upgrade
set -e

echo "🚀 Six Protocol EVM Migration Test Suite (v3.3.1 → v4.0.0)"
echo "============================================================"

# Check if EVM node is running
check_evm_node() {
    echo "🔍 Checking EVM node availability..."
    
    if curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}' \
        http://localhost:8545 > /dev/null 2>&1; then
        echo "✅ EVM node is running"
        return 0
    else
        echo "❌ EVM node is not running on localhost:8545"
        echo "Please start your EVM node first:"
        echo "  cosmovisor run start --grpc.enable true --grpc.address 0.0.0.0:9090 --json-rpc.api eth,txpool,personal,net,debug,web3 --api.enable true --home ~/.six"
        return 1
    fi
}

# Run basic connectivity tests
run_basic_tests() {
    echo "🔗 Running basic connectivity tests..."
    go test -v -timeout 60s ./evm-migration-tests -run TestEVMTestSuite/TestEVMConnection
    go test -v -timeout 60s ./evm-migration-tests -run TestEVMTestSuite/TestAccountState
    go test -v -timeout 60s ./evm-migration-tests -run TestEVMTestSuite/TestNetworkState
    echo "✅ Basic tests completed"
}

# Run transaction tests
run_transaction_tests() {
    echo "💸 Running transaction tests..."
    go test -v -timeout 120s ./evm-migration-tests -run TestEVMTestSuite/TestETHTransfer
    go test -v -timeout 120s ./evm-migration-tests -run TestEVMTestSuite/TestGasEstimation
    echo "✅ Transaction tests completed"
}

# Run contract tests
run_contract_tests() {
    echo "📄 Running smart contract tests..."
    go test -v -timeout 180s ./evm-migration-tests -run TestEVMTestSuite/TestContractDeployment
    go test -v -timeout 120s ./evm-migration-tests -run TestEVMIntegrationTestSuite/TestNFTContractIntegration
    echo "✅ Contract tests completed"
}

# Run JSON-RPC tests
run_jsonrpc_tests() {
    echo "🔌 Running JSON-RPC API tests..."
    go test -v -timeout 60s ./evm-migration-tests -run TestEVMTestSuite/TestJSONRPCMethods
    go test -v -timeout 60s ./evm-migration-tests -run TestEVMTestSuite/TestEventLogs
    echo "✅ JSON-RPC tests completed"
}

# Run performance tests
run_performance_tests() {
    echo "⚡ Running performance tests..."
    go test -v -timeout 300s ./evm-migration-tests -run TestEVMIntegrationTestSuite/TestEVMPerformance
    echo "✅ Performance tests completed"
}

# Main execution
main() {
    case "${1:-all}" in
        "basic")
            check_evm_node && run_basic_tests
            ;;
        "tx")
            check_evm_node && run_transaction_tests
            ;;
        "contract")
            check_evm_node && run_contract_tests
            ;;
        "jsonrpc")
            check_evm_node && run_jsonrpc_tests
            ;;
        "performance")
            check_evm_node && run_performance_tests
            ;;
        "all")
            if check_evm_node; then
                run_basic_tests
                run_transaction_tests
                run_contract_tests
                run_jsonrpc_tests
                run_performance_tests
                echo "🎉 All EVM migration tests completed successfully!"
            fi
            ;;
        "help")
            echo "Usage: $0 [basic|tx|contract|jsonrpc|performance|all|help]"
            echo ""
            echo "EVM Migration Test Suite - Post v4.0.0 Upgrade"
            echo "Options:"
            echo "  basic       - Run basic connectivity tests"
            echo "  tx          - Run transaction tests"
            echo "  contract    - Run smart contract tests"
            echo "  jsonrpc     - Run JSON-RPC API tests"
            echo "  performance - Run performance tests"
            echo "  all         - Run all tests (default)"
            echo "  help        - Show this help message"
            ;;
        *)
            echo "Unknown option: $1"
            echo "Run '$0 help' for usage information"
            exit 1
            ;;
    esac
}

# Execute main function with all arguments
main "$@"
