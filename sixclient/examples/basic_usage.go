package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/thesixnetwork/six-protocol/sixclient"
)

func main() {
	// Get mnemonic from environment variable
	mnemonic := os.Getenv("MNEMONIC")
	if mnemonic == "" {
		log.Fatal("MNEMONIC environment variable not set")
	}

	// Create client configuration
	config := sixclient.ClientConfig{
		Network:  sixclient.Testnet, // Use testnet for development
		Mnemonic: mnemonic,
		// Optional configurations (will use defaults if not set)
		GasLimit: 300000,
		GasPrice: "1.25",
		Timeout:  30 * time.Second,
	}

	// Create new SIX Protocol client
	client, err := sixclient.NewSixClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Printf("🔗 Connected to SIX Protocol %s network\n", config.Network)
	fmt.Printf("📍 Wallet address: %s\n", client.GetAddress())
	fmt.Printf("🌐 RPC endpoint: %s\n", client.GetNetwork().RPC)
	fmt.Printf("🔌 API endpoint: %s\n", client.GetNetwork().API)
	fmt.Println()

	// Example 1: Query account balances
	fmt.Println("💰 Querying account balances...")
	balances, err := client.GetBalances()
	if err != nil {
		log.Printf("Failed to get balances: %v", err)
	} else {
		if len(balances) == 0 {
			fmt.Println("   No balances found")
		} else {
			for _, balance := range balances {
				if balance.Denom == sixclient.Denom {
					// Convert usix to SIX for display
					sixAmount := sixclient.UsixToSIX(balance.Amount)
					fmt.Printf("   %s SIX (%s %s)\n", sixAmount.String(), balance.Amount.String(), balance.Denom)
				} else {
					fmt.Printf("   %s %s\n", balance.Amount.String(), balance.Denom)
				}
			}
		}
	}
	fmt.Println()

	// Example 2: Query SIX balance specifically
	fmt.Println("🪙 Querying SIX balance...")
	sixBalance, err := client.GetSIXBalance()
	if err != nil {
		log.Printf("Failed to get SIX balance: %v", err)
	} else {
		fmt.Printf("   SIX Balance: %s SIX\n", sixBalance.String())
	}
	fmt.Println()

	// Example 3: Send tokens (uncomment to test)
	/*
		fmt.Println("💸 Sending tokens...")
		recipient := "6x13g50hqdqsjk85fmgqz2h5xdxq49lsmjdwlemsp"
		amount := sixclient.SIXToUsix(1) // 1 SIX = 1,000,000 usix

		txResp, err := client.SendSIX(recipient, amount)
		if err != nil {
			log.Printf("Failed to send tokens: %v", err)
		} else {
			fmt.Printf("   ✅ Transaction successful!\n")
			fmt.Printf("   📄 Hash: %s\n", txResp.TxHash)
			fmt.Printf("   ⛽ Gas used: %d/%d\n", txResp.GasUsed, txResp.GasWanted)
			fmt.Printf("   🔗 Explorer: https://sixscan.io/fivenet/tx/%s\n", txResp.TxHash)
		}
		fmt.Println()
	*/

	// Example 4: Multi-send tokens (uncomment to test)
	/*
		fmt.Println("📤 Multi-sending tokens...")
		recipients := []sixclient.SendRequest{
			sixclient.NewSIXSendRequest("6x1recipient1address", sixclient.SIXToUsix(1)),
			sixclient.NewSIXSendRequest("6x1recipient2address", sixclient.SIXToUsix(2)),
		}

		multiTxResp, err := client.SendMultiple(recipients)
		if err != nil {
			log.Printf("Failed to multi-send: %v", err)
		} else {
			fmt.Printf("   ✅ Multi-send successful!\n")
			fmt.Printf("   📄 Hash: %s\n", multiTxResp.TxHash)
		}
		fmt.Println()
	*/

	// Example 5: Query NFT schemas
	fmt.Println("🖼️  Querying NFT schemas...")
	schemas, err := client.ListNFTSchemas()
	if err != nil {
		log.Printf("Failed to list NFT schemas: %v", err)
	} else {
		if len(schemas) == 0 {
			fmt.Println("   No NFT schemas found")
		} else {
			fmt.Printf("   Found %d NFT schemas:\n", len(schemas))
			for i, schema := range schemas {
				fmt.Printf("   %d. %s - %s\n", i+1, schema.Code, schema.Name)
				fmt.Printf("      Owner: %s\n", schema.Owner)
				fmt.Printf("      Verified: %v\n", schema.IsVerified)
			}
		}
	}
	fmt.Println()

	// Example 6: Query tokens
	fmt.Println("🪙 Querying custom tokens...")
	tokens, err := client.ListTokens()
	if err != nil {
		log.Printf("Failed to list tokens: %v", err)
	} else {
		if len(tokens) == 0 {
			fmt.Println("   No custom tokens found")
		} else {
			fmt.Printf("   Found %d tokens:\n", len(tokens))
			for i, token := range tokens {
				fmt.Printf("   %d. %s (%s)\n", i+1, token.Name, token.Base)
				if token.Creator != "system" {
					fmt.Printf("      Creator: %s\n", token.Creator)
					fmt.Printf("      Max Supply: %s\n", token.MaxSupply.String())
				} else {
					fmt.Printf("      Current Supply: %s\n", token.MaxSupply.String())
				}
			}
		}
	}
	fmt.Println()

	// Example 7: Create NFT Schema (uncomment to test)
	/*
		fmt.Println("🎨 Creating NFT schema...")
		schemaReq := sixclient.NFTSchemaRequest{
			Code:        "my-test-schema",
			Name:        "My Test NFT Collection",
			Description: "A test NFT collection created via Go SDK",
		}

		schemaTxResp, err := client.CreateNFTSchema(schemaReq)
		if err != nil {
			log.Printf("Failed to create NFT schema: %v", err)
		} else {
			fmt.Printf("   ✅ NFT Schema created successfully!\n")
			fmt.Printf("   📄 Hash: %s\n", schemaTxResp.TxHash)
		}
		fmt.Println()
	*/

	// Example 8: Create custom token (uncomment to test)
	/*
		fmt.Println("💎 Creating custom token...")
		tokenReq := sixclient.TokenRequest{
			Name:        "MyToken",
			Base:        "mytoken",
			MaxSupply:   math.NewInt(1000000),
			Mintee:      client.GetAddress(), // Allow self to mint
			Description: "My custom token created via Go SDK",
		}

		tokenTxResp, err := client.CreateToken(tokenReq)
		if err != nil {
			log.Printf("Failed to create token: %v", err)
		} else {
			fmt.Printf("   ✅ Token created successfully!\n")
			fmt.Printf("   📄 Hash: %s\n", tokenTxResp.TxHash)
		}
		fmt.Println()
	*/

	fmt.Println("✨ SixClient SDK example completed!")
	fmt.Println()
	fmt.Println("💡 Tips:")
	fmt.Println("   - Uncomment transaction examples to test sending/creating")
	fmt.Println("   - Check your balance on: https://sixscan.io/fivenet")
	fmt.Println("   - Get testnet tokens from faucet if needed")
	fmt.Printf("   - Your address: %s\n", client.GetAddress())
}
