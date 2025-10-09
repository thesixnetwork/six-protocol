# EVM Migration Test Suite

## Overview
Comprehensive test suite for validating EVM functionality after Six Protocol upgrade from **v3.3.1 → v4.0.0** (Cosmos SDK v0.50.10).

## Purpose
This test suite ensures that:
- EVM compatibility is maintained post-upgrade
- Smart contract functionality works correctly
- JSON-RPC APIs function as expected
- Performance benchmarks meet requirements
- No regression issues exist

## Test Categories

### 🔗 Basic Tests
- **EVM Connection**: Chain ID validation (666)
- **Account State**: Balance and nonce verification
- **Network State**: Gas price and block number checks

### 💸 Transaction Tests  
- **ETH Transfers**: Simple value transfers
- **Gas Estimation**: Transaction cost calculations
- **Receipt Validation**: Transaction success confirmation

### 📄 Smart Contract Tests
- **Contract Deployment**: Deploy and verify contracts
- **NFT Integration**: Test pre-deployed NFT contracts
- **Contract Interactions**: Call contract methods

### 🔌 JSON-RPC Tests
- **API Methods**: All standard Ethereum JSON-RPC methods
- **Event Logs**: Query and filter event logs
- **Block Data**: Block and transaction retrieval

### ⚡ Performance Tests
- **Concurrent Transactions**: Multiple simultaneous transactions
- **Load Testing**: System performance under load
- **Throughput Measurement**: TPS benchmarking

## Prerequisites

### Required Services
- Six Protocol node running with EVM enabled
- Cosmovisor managing the upgraded binary (v4.0.0)
- gRPC server accessible on port 9090
- JSON-RPC server accessible on port 8545

### Account Requirements
- Test account with sufficient SIX balance for gas fees
- Private key configured in environment

## Usage

### Quick Start
```bash
# Run all migration tests
./run_evm_tests.sh all

# Run specific test categories
./run_evm_tests.sh basic
./run_evm_tests.sh tx
./run_evm_tests.sh contract
./run_evm_tests.sh jsonrpc
./run_evm_tests.sh performance
```

### Using Makefile
```bash
# From project root
cd evm-migration-tests

# Run basic tests
make test-evm

# Run integration tests
make test-evm-integration

# Run all tests
make test-all-evm

# Generate coverage report
make test-evm-coverage

# Run specific test
make test-evm-specific TEST_NAME=TestEVMConnection
```

### Manual Testing
```bash
# Run specific test suites
go test -v -run TestEVMTestSuite
go test -v -run TestEVMIntegrationTestSuite

# Run with timeout
go test -v -timeout 300s ./...
```

## Configuration

### Environment Setup
```bash
# Start Six Protocol node with required flags
cosmovisor run start \
  --grpc.enable true \
  --grpc.address 0.0.0.0:9090 \
  --json-rpc.api eth,txpool,personal,net,debug,web3 \
  --api.enable true \
  --home ~/.six
```

### Test Account
The test suite uses a hardcoded private key. In production, use:
```bash
export PRIVATE_KEY="your_private_key_here"
export OWNER="your_address_here"
```

## Expected Results

### Success Criteria
- ✅ All basic connectivity tests pass
- ✅ Transaction execution succeeds
- ✅ Contract deployment works
- ✅ JSON-RPC APIs respond correctly
- ✅ Performance meets baseline requirements

### Key Metrics
- **Chain ID**: 666
- **Gas Price**: ~6.875 Gwei
- **Transfer Gas**: 21,000 units
- **Block Time**: ~1 second
- **Concurrent TPS**: 10+ transactions/second

## Troubleshooting

### Common Issues
1. **Connection Refused (port 9090)**: Ensure gRPC is enabled
2. **Connection Refused (port 8545)**: Ensure JSON-RPC is enabled  
3. **Insufficient Balance**: Fund test account with SIX tokens
4. **Nonce Errors**: Wait for pending transactions to complete

### Debug Commands
```bash
# Check node status
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}' \
  http://localhost:8545

# Check account balance
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x8D4Bb008181fd32E7bDac45666D6d2066781B974","latest"],"id":1}' \
  http://localhost:8545
```

## Test Files

- `evm_test.go`: Basic EVM functionality tests
- `evm_integration_test.go`: Integration and performance tests
- `simple_evm_test.go`: Simplified connectivity tests
- `test_helper.go`: Utility functions and setup
- `Makefile`: Build and test automation

## Version Compatibility

| Version | Cosmos SDK | EVM Support | Test Status |
|---------|------------|-------------|-------------|
| v3.3.1  | v0.47.x    | ✅          | Legacy      |
| v4.0.0  | v0.50.10   | ✅          | Current     |

This test suite validates the migration from v3.3.1 to v4.0.0 and ensures continued EVM compatibility.
