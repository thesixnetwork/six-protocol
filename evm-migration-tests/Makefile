# EVM Migration Test Makefile - Post v4.0.0 Upgrade

.PHONY: test-evm test-evm-integration test-all-evm clean-test

# Variables
GO_MODULE := github.com/thesixnetwork/six-protocol
TEST_TIMEOUT := 300s
TEST_VERBOSITY := -v

# Run basic EVM migration tests
test-evm:
	@echo "Running EVM migration tests..."
	@go test -timeout $(TEST_TIMEOUT) $(TEST_VERBOSITY) -run TestEVMTestSuite

# Run integration tests
test-evm-integration:
	@echo "Running EVM integration migration tests..."
	@go test -timeout $(TEST_TIMEOUT) $(TEST_VERBOSITY) -run TestEVMIntegrationTestSuite

# Run all EVM migration tests
test-all-evm:
	@echo "Running all EVM migration tests..."
	@go test -timeout $(TEST_TIMEOUT) $(TEST_VERBOSITY) ./...

# Run tests with coverage
test-evm-coverage:
	@echo "Running EVM migration tests with coverage..."
	@go test -timeout $(TEST_TIMEOUT) -coverprofile=coverage.out -run TestEVMTestSuite
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: evm-migration-tests/coverage.html"

# Run specific test
test-evm-specific:
	@echo "Running specific EVM migration test: $(TEST_NAME)"
	@go test -timeout $(TEST_TIMEOUT) $(TEST_VERBOSITY) -run $(TEST_NAME)

# Clean test artifacts
clean-test:
	@echo "Cleaning test artifacts..."
	@rm -f coverage.out coverage.html
	@rm -f *.log

# Setup test dependencies
setup-test-deps:
	@echo "Installing test dependencies..."
	@go mod tidy
	@go get github.com/stretchr/testify/suite
	@go get github.com/ethereum/go-ethereum

# Run tests in CI mode
test-ci:
	@echo "Running migration tests in CI mode..."
	@go test -timeout $(TEST_TIMEOUT) -json ./... | tee test_results.json

# Examples of running tests
help:
	@echo "Available EVM migration test commands:"
	@echo "  make test-evm              - Run basic EVM migration tests"
	@echo "  make test-evm-integration  - Run integration migration tests"
	@echo "  make test-all-evm          - Run all EVM migration tests"
	@echo "  make test-evm-coverage     - Run tests with coverage"
	@echo "  make test-evm-specific TEST_NAME=TestEVMConnection - Run specific test"
	@echo "  make clean-test            - Clean test artifacts"
	@echo "  make setup-test-deps       - Install test dependencies"
