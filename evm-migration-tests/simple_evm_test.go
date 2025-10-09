package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

// Simple EVM connectivity test without test suite
func TestSimpleEVMConnection(t *testing.T) {
	// Connect to EVM node
	client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Test chain ID
	chainID, err := client.ChainID(ctx)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(666), chainID)

	t.Logf("✅ Chain ID: %s", chainID.String())

	// Test latest block
	blockNumber, err := client.BlockNumber(ctx)
	require.NoError(t, err)
	require.True(t, blockNumber > 0)

	t.Logf("✅ Latest block: %d", blockNumber)
}

func TestSimpleAccountBalance(t *testing.T) {
	// Connect to EVM node
	client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Setup test account
	userPK := os.Getenv("PRIVATE_KEY")
	userPK = strings.TrimPrefix(userPK, "0x")
	privateKey, err := crypto.HexToECDSA(userPK)
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Check balance - this is where the error occurs
	balance, err := client.BalanceAt(ctx, address, nil)
	require.NoError(t, err)

	t.Logf("✅ Account %s balance: %s", address.Hex(), balance.String())

	// Check nonce
	nonce, err := client.PendingNonceAt(ctx, address)
	require.NoError(t, err)

	t.Logf("✅ Account nonce: %d", nonce)
}

func TestSimpleGasPrice(t *testing.T) {
	// Connect to EVM node
	client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)
	defer client.Close()

	ctx := context.Background()

	// Test gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	require.NoError(t, err)

	t.Logf("✅ Gas price: %s", gasPrice.String())
}
