package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

// TestMain runs before all tests
func TestMain(m *testing.M) {
	// Setup test environment
	if err := setupTestEnvironment(); err != nil {
		log.Fatalf("Failed to setup test environment: %v", err)
	}

	// Run tests
	code := m.Run()

	// Cleanup
	cleanupTestEnvironment()

	os.Exit(code)
}

func setupTestEnvironment() error {
	// Wait for EVM node to be ready
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Printf("EVM node not ready, waiting...")
		return waitForEVMNode()
	}
	defer client.Close()

	// Check if node is responsive
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.ChainID(ctx)
	if err != nil {
		log.Printf("EVM node not responsive, waiting...")
		return waitForEVMNode()
	}

	log.Println("EVM node is ready for testing")
	return nil
}

func waitForEVMNode() error {
	maxRetries := 30
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		client, err := ethclient.Dial("http://localhost:8545")
		if err == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			_, err = client.ChainID(ctx)
			client.Close()
			cancel()

			if err == nil {
				log.Println("EVM node is now ready")
				return nil
			}
		}

		log.Printf("Waiting for EVM node... (%d/%d)", i+1, maxRetries)
		time.Sleep(retryDelay)
	}

	return fmt.Errorf("EVM node failed to become ready after %d retries", maxRetries)
}

func cleanupTestEnvironment() {
	log.Println("Cleaning up test environment")
	// Add cleanup logic if needed
}

// Test helper functions
func skipIfNoEVMNode(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Skip("EVM node not available, skipping test")
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = client.ChainID(ctx)
	if err != nil {
		t.Skip("EVM node not responsive, skipping test")
	}
}
