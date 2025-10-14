package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/suite"
)

type EVMTestSuite struct {
	suite.Suite
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainID    *big.Int
}

func TestEVMTestSuite(t *testing.T) {
	suite.Run(t, new(EVMTestSuite))
}

func (suite *EVMTestSuite) SetupSuite() {
	// Connect to local EVM node
	client, err := ethclient.Dial("http://localhost:8545")
	suite.Require().NoError(err)
	suite.client = client

	// Setup private key (from .env)
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	suite.Require().NotEmpty(privateKeyHex, "PRIVATE_KEY environment variable must be set")

	// Remove 0x prefix if present
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	suite.Require().NoError(err)
	suite.privateKey = privateKey

	// Get address from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	suite.Require().True(ok)
	suite.address = crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	suite.Require().NoError(err)
	suite.chainID = chainID
}

// Test 1: Basic EVM Connection
func (suite *EVMTestSuite) TestEVMConnection() {
	// Test if we can connect to EVM and get chain ID
	chainID, err := suite.client.ChainID(context.Background())
	suite.Require().NoError(err)
	suite.Equal(big.NewInt(666), chainID, "Chain ID should be 666")

	// Test network ID
	networkID, err := suite.client.NetworkID(context.Background())
	suite.Require().NoError(err)
	suite.NotNil(networkID)
}

// Test 2: Account Balance and Nonce
func (suite *EVMTestSuite) TestAccountState() {
	ctx := context.Background()

	// Check balance
	balance, err := suite.client.BalanceAt(ctx, suite.address, nil)
	suite.Require().NoError(err)
	suite.True(balance.Cmp(big.NewInt(0)) >= 0, "Balance should be >= 0")

	// Check nonce
	nonce, err := suite.client.PendingNonceAt(ctx, suite.address)
	suite.Require().NoError(err)
	suite.T().Logf("Account %s has nonce: %d, balance: %s", suite.address.Hex(), nonce, balance.String())
}

// Test 3: Gas Price and Block Number
func (suite *EVMTestSuite) TestNetworkState() {
	ctx := context.Background()

	// Test gas price
	gasPrice, err := suite.client.SuggestGasPrice(ctx)
	suite.Require().NoError(err)
	suite.True(gasPrice.Cmp(big.NewInt(0)) > 0, "Gas price should be > 0")

	// Test latest block
	latestBlock, err := suite.client.BlockNumber(ctx)
	suite.Require().NoError(err)
	suite.True(latestBlock > 0, "Latest block should be > 0")

	suite.T().Logf("Gas price: %s, Latest block: %d", gasPrice.String(), latestBlock)
}

// Test 4: Simple ETH Transfer
func (suite *EVMTestSuite) TestETHTransfer() {
	ctx := context.Background()

	// Create recipient address
	recipientKey, err := crypto.GenerateKey()
	suite.Require().NoError(err)
	recipient := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Get initial balances
	senderBalance, err := suite.client.BalanceAt(ctx, suite.address, nil)
	suite.Require().NoError(err)

	recipientBalance, err := suite.client.BalanceAt(ctx, recipient, nil)
	suite.Require().NoError(err)

	// Skip if sender has no balance
	if senderBalance.Cmp(big.NewInt(0)) == 0 {
		suite.T().Skip("Sender has no balance for transfer test")
		return
	}

	// Create transaction
	nonce, err := suite.client.PendingNonceAt(ctx, suite.address)
	suite.Require().NoError(err)

	gasLimit := uint64(21000)
	gasPrice, err := suite.client.SuggestGasPrice(ctx)
	suite.Require().NoError(err)

	transferAmount := big.NewInt(1000000000000000) // 0.001 ETH
	tx := types.NewTransaction(nonce, recipient, transferAmount, gasLimit, gasPrice, nil)

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(suite.chainID), suite.privateKey)
	suite.Require().NoError(err)

	// Send transaction
	err = suite.client.SendTransaction(ctx, signedTx)
	suite.Require().NoError(err)

	suite.T().Logf("Transfer transaction sent: %s", signedTx.Hash().Hex())

	// Wait for confirmation
	receipt, err := suite.waitForTxReceipt(ctx, signedTx.Hash())
	suite.Require().NoError(err)
	suite.Equal(types.ReceiptStatusSuccessful, receipt.Status)

	// Verify balances changed
	newRecipientBalance, err := suite.client.BalanceAt(ctx, recipient, nil)
	suite.Require().NoError(err)

	expectedBalance := new(big.Int).Add(recipientBalance, transferAmount)
	suite.Equal(expectedBalance, newRecipientBalance, "Recipient balance should increase by transfer amount")
}

// Test 5: Smart Contract Deployment
func (suite *EVMTestSuite) TestContractDeployment() {
	ctx := context.Background()

	// Simple contract bytecode (empty contract)
	contractBytecode := "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220c7b3a89b8f442a9f8d6d32342b8e5ed4c8653e1a9c36ac1a84a0266c8b0a290c64736f6c634300081a0033"

	// Get nonce and gas price
	nonce, err := suite.client.PendingNonceAt(ctx, suite.address)
	suite.Require().NoError(err)

	gasPrice, err := suite.client.SuggestGasPrice(ctx)
	suite.Require().NoError(err)

	// Create deployment transaction
	gasLimit := uint64(500000)
	data := common.FromHex(contractBytecode)
	tx := types.NewContractCreation(nonce, big.NewInt(0), gasLimit, gasPrice, data)

	// Sign and send
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(suite.chainID), suite.privateKey)
	suite.Require().NoError(err)

	err = suite.client.SendTransaction(ctx, signedTx)
	suite.Require().NoError(err)

	suite.T().Logf("Contract deployment transaction: %s", signedTx.Hash().Hex())

	// Wait for receipt
	receipt, err := suite.waitForTxReceipt(ctx, signedTx.Hash())
	suite.Require().NoError(err)
	suite.Equal(types.ReceiptStatusSuccessful, receipt.Status)
	suite.NotEqual(common.Address{}, receipt.ContractAddress)

	suite.T().Logf("Contract deployed at: %s", receipt.ContractAddress.Hex())
}

// Test 6: JSON-RPC Methods
func (suite *EVMTestSuite) TestJSONRPCMethods() {
	ctx := context.Background()

	// Test eth_blockNumber
	blockNumber, err := suite.client.BlockNumber(ctx)
	suite.Require().NoError(err)
	suite.True(blockNumber > 0)

	// Test eth_getBalance
	balance, err := suite.client.BalanceAt(ctx, suite.address, nil)
	suite.Require().NoError(err)
	suite.NotNil(balance)

	// Test eth_gasPrice
	gasPrice, err := suite.client.SuggestGasPrice(ctx)
	suite.Require().NoError(err)
	suite.True(gasPrice.Cmp(big.NewInt(0)) > 0)

	// Test eth_chainId
	chainID, err := suite.client.ChainID(ctx)
	suite.Require().NoError(err)
	suite.Equal(big.NewInt(666), chainID)

	suite.T().Log("All JSON-RPC methods working correctly")
}

// Test 7: Event Logs
func (suite *EVMTestSuite) TestEventLogs() {
	ctx := context.Background()

	// Get latest block
	latestBlock, err := suite.client.BlockNumber(ctx)
	suite.Require().NoError(err)

	// Query logs from recent blocks
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(latestBlock - 10)),
		ToBlock:   big.NewInt(int64(latestBlock)),
	}

	logs, err := suite.client.FilterLogs(ctx, query)
	suite.Require().NoError(err)

	suite.T().Logf("Found %d logs in last 10 blocks", len(logs))
}

// Test 8: Gas Estimation
func (suite *EVMTestSuite) TestGasEstimation() {
	ctx := context.Background()

	// Create a simple transfer message
	recipientKey, err := crypto.GenerateKey()
	suite.Require().NoError(err)
	recipient := crypto.PubkeyToAddress(recipientKey.PublicKey)

	msg := ethereum.CallMsg{
		From:     suite.address,
		To:       &recipient,
		Value:    big.NewInt(1000),
		Gas:      0,
		GasPrice: big.NewInt(1),
		Data:     nil,
	}

	// Estimate gas
	gasEstimate, err := suite.client.EstimateGas(ctx, msg)
	suite.Require().NoError(err)
	suite.Equal(uint64(21000), gasEstimate, "Simple transfer should use 21000 gas")

	suite.T().Logf("Gas estimate for transfer: %d", gasEstimate)
}

// Helper function to wait for transaction receipt
func (suite *EVMTestSuite) waitForTxReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return nil, context.DeadlineExceeded
		case <-ticker.C:
			receipt, err := suite.client.TransactionReceipt(ctx, txHash)
			if err == nil {
				return receipt, nil
			}
		}
	}
}
