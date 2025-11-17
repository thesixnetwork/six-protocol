# SixClient - Go SDK for SIX Protocol

A native Go SDK for interacting with the SIX Protocol blockchain without requiring keyring files. This SDK supports mnemonic-based signing and uses HTTP REST API calls for maximum compatibility with load balancers and standard web infrastructure.

## Features

🔑 **Mnemonic-Based Authentication** - No keyring files required  
🌐 **HTTP REST API** - Compatible with load balancers and firewalls  
⚡ **Native Go Integration** - Type-safe and performant  
🏦 **Banking Operations** - Send tokens, query balances, multi-send  
🖼️ **NFT Management** - Create schemas, manage NFTs, query collections  
🪙 **Custom Tokens** - Create, mint, burn, and manage custom tokens  
🔧 **Transaction Broadcasting** - Sign and broadcast transactions seamlessly  
📊 **Account Management** - Query account info, balances, and transaction history  

## Installation

```bash
go get github.com/thesixnetwork/six-protocol/sixclient
```

## Quick Start

### 1. Basic Setup

```go
package main

import (
    "fmt"
    "log"
    "os"
    "time"
    
    "github.com/thesixnetwork/six-protocol/sixclient"
)

func main() {
    // Get mnemonic from environment
    mnemonic := os.Getenv("MNEMONIC")
    if mnemonic == "" {
        log.Fatal("MNEMONIC environment variable not set")
    }
    
    // Create client
    client, err := sixclient.NewSixClient(sixclient.ClientConfig{
        Network:  sixclient.Testnet,
        Mnemonic: mnemonic,
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Connected! Address: %s\n", client.GetAddress())
}
```

### 2. Environment Setup

Create a `.env` file or set environment variables:

```bash
export MNEMONIC="your twelve word mnemonic phrase here"
```

### 3. Network Configuration

The SDK supports three networks:

```go
// Testnet (recommended for development)
client, _ := sixclient.NewSixClient(sixclient.ClientConfig{
    Network:  sixclient.Testnet,
    Mnemonic: mnemonic,
})

// Mainnet (for production)
client, _ := sixclient.NewSixClient(sixclient.ClientConfig{
    Network:  sixclient.Mainnet,
    Mnemonic: mnemonic,
})

// Local (for local development)
client, _ := sixclient.NewSixClient(sixclient.ClientConfig{
    Network:  sixclient.Local,
    Mnemonic: mnemonic,
})
```

## Banking Operations

### Query Balances

```go
// Get all balances
balances, err := client.GetBalances()
if err != nil {
    log.Fatal(err)
}

for _, balance := range balances {
    fmt.Printf("%s: %s\n", balance.Denom, balance.Amount.String())
}

// Get SIX balance specifically (in SIX units, not usix)
sixBalance, err := client.GetSIXBalance()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("SIX Balance: %s SIX\n", sixBalance.String())
```

### Send Tokens

```go
import "cosmossdk.io/math"

// Send SIX tokens (convenience method)
recipient := "6x13g50hqdqsjk85fmgqz2h5xdxq49lsmjdwlemsp"
amount := sixclient.SIXToUsix(10) // 10 SIX tokens

txResp, err := client.SendSIX(recipient, amount)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Transaction Hash: %s\n", txResp.TxHash)
fmt.Printf("Gas Used: %d/%d\n", txResp.GasUsed, txResp.GasWanted)

// Send any denomination
txResp, err = client.SendTokens(recipient, math.NewInt(1000000), "usix")
```

### Multi-Send

```go
// Send to multiple recipients in one transaction
recipients := []sixclient.SendRequest{
    sixclient.NewSIXSendRequest("6x1address1...", sixclient.SIXToUsix(5)),
    sixclient.NewSIXSendRequest("6x1address2...", sixclient.SIXToUsix(3)),
    sixclient.NewSendRequest("6x1address3...", math.NewInt(2000000), "usix"),
}

txResp, err := client.SendMultiple(recipients)
if err != nil {
    log.Fatal(err)
}
```

## NFT Operations

### Create NFT Schema

```go
schemaReq := sixclient.NFTSchemaRequest{
    Code:        "my-nft-collection",
    Name:        "My NFT Collection",
    Description: "A collection of unique digital assets",
}

txResp, err := client.CreateNFTSchema(schemaReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Schema created: %s\n", txResp.TxHash)
```

### Query NFT Schema

```go
schema, err := client.QueryNFTSchema("my-nft-collection")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Schema: %s (%s)\n", schema.Name, schema.Code)
fmt.Printf("Owner: %s\n", schema.Owner)
fmt.Printf("Verified: %v\n", schema.IsVerified)
```

### List All NFT Schemas

```go
schemas, err := client.ListNFTSchemas()
if err != nil {
    log.Fatal(err)
}

for _, schema := range schemas {
    fmt.Printf("%s - %s (Owner: %s)\n", schema.Code, schema.Name, schema.Owner)
}
```

### Transfer Schema Ownership

```go
newOwner := "6x1newowneraddress..."
txResp, err := client.TransferSchemaOwnership("my-nft-collection", newOwner)
if err != nil {
    log.Fatal(err)
}
```

## Custom Token Operations

### Create Custom Token

```go
import "cosmossdk.io/math"

tokenReq := sixclient.TokenRequest{
    Name:        "MyToken",
    Base:        "mytoken",
    MaxSupply:   math.NewInt(1000000),
    Mintee:      client.GetAddress(), // Allow self to mint
    Description: "My custom utility token",
}

txResp, err := client.CreateToken(tokenReq)
if err != nil {
    log.Fatal(err)
}
```

### Mint Tokens

```go
mintReq := sixclient.TokenMintRequest{
    Name:   "MyToken",
    Amount: math.NewInt(10000),
    Mintee: "6x1recipientaddress...",
}

txResp, err := client.MintToken(mintReq)
if err != nil {
    log.Fatal(err)
}
```

### Burn Tokens

```go
burnReq := sixclient.TokenBurnRequest{
    Name:   "MyToken", 
    Amount: math.NewInt(5000),
}

txResp, err := client.BurnToken(burnReq)
if err != nil {
    log.Fatal(err)
}
```

### Query Token Information

```go
token, err := client.QueryToken("MyToken")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Token: %s (%s)\n", token.Name, token.Base)
fmt.Printf("Creator: %s\n", token.Creator)
fmt.Printf("Max Supply: %s\n", token.MaxSupply.String())
```

## Configuration Options

### Client Configuration

```go
config := sixclient.ClientConfig{
    Network:    sixclient.Testnet,        // Required: Network to connect to
    Mnemonic:   "your mnemonic here...",  // Required: Wallet mnemonic
    HTTPClient: &http.Client{},           // Optional: Custom HTTP client
    GasLimit:   300000,                   // Optional: Default gas limit
    GasPrice:   "1.25",                   // Optional: Gas price in usix
    Timeout:    30 * time.Second,         // Optional: Request timeout
}

client, err := sixclient.NewSixClient(config)
```

### Network Endpoints

| Network | Chain ID | RPC Endpoint | API Endpoint |
|---------|----------|--------------|--------------|
| Mainnet | sixnet | https://sixnet-rpc.sixprotocol.net:443 | https://sixnet-api.sixprotocol.net |
| Testnet | fivenet | https://rpc1.fivenet.sixprotocol.net:443 | https://api1.fivenet.sixprotocol.net |
| Local | testnet | http://localhost:26657 | http://localhost:1317 |

## Account Management

### Get Account Information

```go
// Get wallet address
address := client.GetAddress()
fmt.Printf("Address: %s\n", address)

// Get public key
pubKey := client.GetPublicKey()
fmt.Printf("Public Key: %s\n", pubKey.String())

// Get detailed account info
account, err := client.GetAccount()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Account Number: %d\n", account.AccountNumber)
fmt.Printf("Sequence: %d\n", account.Sequence)
```

### Query Account Balance

```go
// Get specific denomination balance
balance, err := client.GetBalance("usix")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("USIX Balance: %s\n", balance.Amount.String())
```

## Utility Functions

### Amount Conversion

```go
import "cosmossdk.io/math"

// Convert SIX to usix (multiply by 1,000,000)
usixAmount := sixclient.SIXToUsix(10) // 10 SIX = 10,000,000 usix

// Convert usix to SIX (divide by 1,000,000)  
sixAmount := sixclient.UsixToSIX(math.NewInt(10000000)) // 10,000,000 usix = 10 SIX

// Create send requests
sendReq := sixclient.NewSIXSendRequest("6x1address...", sixclient.SIXToUsix(5))
```

## Error Handling

The SDK uses standard Go error handling patterns:

```go
// Always check for errors
txResp, err := client.SendSIX(recipient, amount)
if err != nil {
    // Handle different error types
    switch {
    case strings.Contains(err.Error(), "insufficient funds"):
        fmt.Println("Not enough balance for this transaction")
    case strings.Contains(err.Error(), "invalid address"):
        fmt.Println("Recipient address is invalid")
    default:
        fmt.Printf("Transaction failed: %v\n", err)
    }
    return
}

// Check transaction result
if txResp.Code != 0 {
    fmt.Printf("Transaction failed with code %d: %s\n", txResp.Code, txResp.RawLog)
    return
}

fmt.Printf("Transaction successful: %s\n", txResp.TxHash)
```

## Examples

Check the `examples/` directory for complete working examples:

- [`basic_usage.go`](examples/basic_usage.go) - Basic SDK usage examples
- [`token_management.go`](examples/token_management.go) - Custom token operations
- [`nft_operations.go`](examples/nft_operations.go) - NFT schema and data management

### Running Examples

```bash
# Set your mnemonic
export MNEMONIC="your twelve word mnemonic phrase here"

# Run basic usage example
go run examples/basic_usage.go

# Run specific examples
go run examples/token_management.go
go run examples/nft_operations.go
```

## Security Best Practices

### 1. Mnemonic Security

```go
// ✅ Good: Use environment variables
mnemonic := os.Getenv("MNEMONIC")

// ❌ Bad: Never hardcode mnemonics
// mnemonic := "word1 word2 word3..." // DON'T DO THIS
```

### 2. Network Selection

```go
// ✅ Good: Use testnet for development
config := sixclient.ClientConfig{
    Network: sixclient.Testnet, // Safe for testing
    // ...
}

// ⚠️ Caution: Only use mainnet for production
config := sixclient.ClientConfig{
    Network: sixclient.Mainnet, // Real money!
    // ...
}
```

### 3. Amount Validation

```go
// ✅ Good: Validate amounts before sending
balance, err := client.GetSIXBalance()
if err != nil {
    log.Fatal(err)
}

sendAmount := sixclient.SIXToUsix(10)
if balance.LT(sendAmount) {
    fmt.Println("Insufficient balance")
    return
}

txResp, err := client.SendSIX(recipient, sendAmount)
```

## Testing

### Testnet Setup

1. **Get testnet SIX tokens** from the faucet
2. **Set your mnemonic** in environment variables
3. **Use testnet network** in your client configuration

```go
client, err := sixclient.NewSixClient(sixclient.ClientConfig{
    Network:  sixclient.Testnet, // Always use testnet for testing
    Mnemonic: os.Getenv("MNEMONIC"),
})
```

### Unit Tests

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestSendTokens ./...
```

## Troubleshooting

### Common Issues

#### 1. Invalid Mnemonic

```
Error: invalid mnemonic phrase
```

**Solution:** Ensure your mnemonic is exactly 12 or 24 words separated by spaces.

#### 2. Insufficient Balance

```
Error: transaction failed (code 5): insufficient funds
```

**Solution:** Check your balance and ensure you have enough tokens + fees.

#### 3. Network Connection Issues

```
Error: failed to connect to API
```

**Solution:** Check network connectivity and ensure the API endpoint is accessible.

#### 4. Gas Estimation Issues

```
Error: out of gas
```

**Solution:** Increase gas limit in client configuration:

```go
config := sixclient.ClientConfig{
    // ...
    GasLimit: 500000, // Increase from default 300000
}
```

### Debug Mode

Enable debug logging for troubleshooting:

```go
import "log"

// Enable debug logging
log.SetFlags(log.LstdFlags | log.Lshortfile)

// Your client operations...
```

## API Reference

### Client Methods

| Method | Description | Returns |
|--------|-------------|---------|
| `NewSixClient(config)` | Creates new client | `*SixClient, error` |
| `GetAddress()` | Get wallet address | `string` |
| `GetPublicKey()` | Get public key | `cryptotypes.PubKey` |
| `GetBalances()` | Get all balances | `[]Balance, error` |
| `GetBalance(denom)` | Get balance for denom | `*Balance, error` |
| `SendTokens(to, amount, denom)` | Send tokens | `*TxResponse, error` |
| `SendSIX(to, amount)` | Send SIX tokens | `*TxResponse, error` |
| `BroadcastTx(msgs, memo)` | Broadcast transaction | `*TxResponse, error` |

### NFT Methods

| Method | Description | Returns |
|--------|-------------|---------|
| `CreateNFTSchema(req)` | Create NFT schema | `*TxResponse, error` |
| `QueryNFTSchema(code)` | Query NFT schema | `*NFTSchema, error` |
| `ListNFTSchemas()` | List all schemas | `[]NFTSchema, error` |
| `TransferSchemaOwnership(code, newOwner)` | Transfer ownership | `*TxResponse, error` |

### Token Methods

| Method | Description | Returns |
|--------|-------------|---------|
| `CreateToken(req)` | Create custom token | `*TxResponse, error` |
| `MintToken(req)` | Mint tokens | `*TxResponse, error` |
| `BurnToken(req)` | Burn tokens | `*TxResponse, error` |
| `QueryToken(name)` | Query token info | `*Token, error` |
| `ListTokens()` | List all tokens | `[]Token, error` |

## Contributing

We welcome contributions to the SixClient SDK! Please read our contributing guidelines and submit pull requests to our repository.

### Development Setup

```bash
# Clone repository
git clone https://github.com/thesixnetwork/six-protocol.git
cd six-protocol/sixclient

# Install dependencies
go mod tidy

# Run tests
go test ./...

# Run examples
export MNEMONIC="your test mnemonic"
go run examples/basic_usage.go
```

## License

This project is licensed under the MIT License. See the [LICENSE](../LICENSE) file for details.

## Support

- 📚 **Documentation:** [SIX Protocol Docs](https://docs.sixprotocol.net)
- 💬 **Discord:** [SIX Protocol Community](https://discord.gg/sixprotocol)
- 🐛 **Issues:** [GitHub Issues](https://github.com/thesixnetwork/six-protocol/issues)
- 🌐 **Website:** [sixprotocol.net](https://sixprotocol.net)

## Related Projects

- [SIX Protocol Chain](../) - The main SIX Protocol blockchain implementation
- [SixChain JS SDK](../../sixchainJS) - JavaScript/TypeScript SDK for web applications
- [SIX Protocol Explorer](https://sixscan.io) - Blockchain explorer for SIX Protocol

---

Built with ❤️ by the SIX Protocol team