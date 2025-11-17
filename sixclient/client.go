// Package sixclient provides a native Go SDK for interacting with SIX Protocol blockchain
// without requiring keyring files. It supports mnemonic-based signing and HTTP REST API calls.
package sixclient

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/go-bip39"

	// SixProtocol modules
	nftmngrmoduletypes "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	tokenmngrmoduletypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// Network represents different SIX Protocol networks
type Network string

const (
	// Network constants
	Mainnet Network = "mainnet"
	Testnet Network = "testnet"
	Local   Network = "local"

	// Chain configurations
	MainnetRPC     = "https://sixnet-rpc.sixprotocol.net:443"
	MainnetAPI     = "https://sixnet-api.sixprotocol.net"
	MainnetChainID = "sixnet"

	TestnetRPC     = "https://rpc1.fivenet.sixprotocol.net:443"
	TestnetAPI     = "https://api1.fivenet.sixprotocol.net"
	TestnetChainID = "fivenet"

	LocalRPC     = "http://localhost:26657"
	LocalAPI     = "http://localhost:1317"
	LocalChainID = "testnet"

	// SIX Protocol specifics
	AddressPrefix = "6x"
	Denom         = "usix"
	DefaultGas    = uint64(300000)
	DefaultFee    = "1.25"
)

// NetworkConfig holds network-specific configuration
type NetworkConfig struct {
	RPC     string
	API     string
	ChainID string
}

// ClientConfig holds client configuration options
type ClientConfig struct {
	Network    Network
	Mnemonic   string
	HTTPClient *http.Client
	GasLimit   uint64
	GasPrice   string
	Timeout    time.Duration
}

// SixClient is the main client for interacting with SIX Protocol
type SixClient struct {
	config     NetworkConfig
	mnemonic   string
	privateKey cryptotypes.PrivKey
	address    sdk.AccAddress
	httpClient *http.Client
	codec      codec.ProtoCodecMarshaler
	txConfig   authtx.TxConfig
	gasLimit   uint64
	gasPrice   string
	timeout    time.Duration
}

// Account holds account information
type Account struct {
	Address       string `json:"address"`
	AccountNumber uint64 `json:"account_number"`
	Sequence      uint64 `json:"sequence"`
}

// Balance represents a coin balance
type Balance struct {
	Denom  string   `json:"denom"`
	Amount math.Int `json:"amount"`
}

// TxResponse represents a transaction response
type TxResponse struct {
	TxHash    string `json:"txhash"`
	Code      uint32 `json:"code"`
	RawLog    string `json:"raw_log"`
	GasUsed   uint64 `json:"gas_used"`
	GasWanted uint64 `json:"gas_wanted"`
	Height    int64  `json:"height"`
}

// NFTSchema represents an NFT schema
type NFTSchema struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	IsVerified  bool   `json:"is_verified"`
}

// Token represents a custom token
type Token struct {
	Name      string   `json:"name"`
	Base      string   `json:"base"`
	Creator   string   `json:"creator"`
	MaxSupply math.Int `json:"max_supply"`
}

func init() {
	// Set SIX Protocol address prefixes
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AddressPrefix, AddressPrefix+"pub")
	config.SetBech32PrefixForValidator(AddressPrefix+"valoper", AddressPrefix+"valoperpub")
	config.SetBech32PrefixForConsensusNode(AddressPrefix+"valcons", AddressPrefix+"valconspub")
	config.Seal()
}

// NewSixClient creates a new SIX Protocol client
func NewSixClient(config ClientConfig) (*SixClient, error) {
	// Get network configuration
	netConfig := getNetworkConfig(config.Network)
	if netConfig == nil {
		return nil, fmt.Errorf("unsupported network: %s", config.Network)
	}

	// Validate mnemonic
	if !bip39.IsMnemonicValid(config.Mnemonic) {
		return nil, fmt.Errorf("invalid mnemonic phrase")
	}

	// Derive private key from mnemonic
	seed, err := bip39.NewSeedWithErrorChecking(config.Mnemonic, "")
	if err != nil {
		return nil, fmt.Errorf("failed to generate seed: %w", err)
	}

	hdPath := hd.CreateHDPath(118, 0, 0)
	masterPriv, ch := hd.ComputeMastersFromSeed(seed)
	derivedPriv, err := hd.DerivePrivateKeyForPath(masterPriv, ch, hdPath.String())
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key: %w", err)
	}

	privKey := &secp256k1.PrivKey{Key: derivedPriv}
	address := sdk.AccAddress(privKey.PubKey().Address())

	// Set default values
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	}
	if config.GasLimit == 0 {
		config.GasLimit = DefaultGas
	}
	if config.GasPrice == "" {
		config.GasPrice = DefaultFee
	}
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second
	}

	// Create codec
	registry := types.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(registry)
	banktypes.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry)
	nftmngrmoduletypes.RegisterInterfaces(registry)
	tokenmngrmoduletypes.RegisterInterfaces(registry)

	cdc := codec.NewProtoCodec(registry)
	txConfig := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)

	return &SixClient{
		config:     *netConfig,
		mnemonic:   config.Mnemonic,
		privateKey: privKey,
		address:    address,
		httpClient: config.HTTPClient,
		codec:      cdc,
		txConfig:   txConfig,
		gasLimit:   config.GasLimit,
		gasPrice:   config.GasPrice,
		timeout:    config.Timeout,
	}, nil
}

// GetAddress returns the client's wallet address
func (c *SixClient) GetAddress() string {
	return c.address.String()
}

// GetPublicKey returns the client's public key
func (c *SixClient) GetPublicKey() cryptotypes.PubKey {
	return c.privateKey.PubKey()
}

// GetNetwork returns the current network configuration
func (c *SixClient) GetNetwork() NetworkConfig {
	return c.config
}

// httpGet performs an HTTP GET request to the API
func (c *SixClient) httpGet(path string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.config.API, path)

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	return io.ReadAll(resp.Body)
}

// httpPost performs an HTTP POST request to the API
func (c *SixClient) httpPost(path string, data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	url := fmt.Sprintf("%s%s", c.config.API, path)

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP POST failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// GetAccount retrieves account information
func (c *SixClient) GetAccount() (*Account, error) {
	path := fmt.Sprintf("/cosmos/auth/v1beta1/accounts/%s", c.address.String())
	body, err := c.httpGet(path)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Account struct {
			Type          string `json:"@type"`
			Address       string `json:"address"`
			AccountNumber string `json:"account_number"`
			Sequence      string `json:"sequence"`
		} `json:"account"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account response: %w", err)
	}

	accountNumber, _ := strconv.ParseUint(resp.Account.AccountNumber, 10, 64)
	sequence, _ := strconv.ParseUint(resp.Account.Sequence, 10, 64)

	return &Account{
		Address:       resp.Account.Address,
		AccountNumber: accountNumber,
		Sequence:      sequence,
	}, nil
}

// GetBalances retrieves account balances
func (c *SixClient) GetBalances() ([]Balance, error) {
	path := fmt.Sprintf("/cosmos/bank/v1beta1/balances/%s", c.address.String())
	body, err := c.httpGet(path)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Balances []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"balances"`
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal balance response: %w", err)
	}

	balances := make([]Balance, len(resp.Balances))
	for i, bal := range resp.Balances {
		amount, ok := math.NewIntFromString(bal.Amount)
		if !ok {
			amount = math.ZeroInt()
		}
		balances[i] = Balance{
			Denom:  bal.Denom,
			Amount: amount,
		}
	}

	return balances, nil
}

// GetBalance retrieves balance for a specific denomination
func (c *SixClient) GetBalance(denom string) (*Balance, error) {
	balances, err := c.GetBalances()
	if err != nil {
		return nil, err
	}

	for _, balance := range balances {
		if balance.Denom == denom {
			return &balance, nil
		}
	}

	return &Balance{Denom: denom, Amount: math.ZeroInt()}, nil
}

// BroadcastTx signs and broadcasts a transaction
func (c *SixClient) BroadcastTx(msgs []sdk.Msg, memo string) (*TxResponse, error) {
	// Get account info for sequence number
	account, err := c.GetAccount()
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	// Build transaction
	txBuilder := c.txConfig.NewTxBuilder()
	if err := txBuilder.SetMsgs(msgs...); err != nil {
		return nil, fmt.Errorf("failed to set messages: %w", err)
	}

	// Set gas and fees
	txBuilder.SetGasLimit(c.gasLimit)

	gasPrice, err := sdk.NewDecFromStr(c.gasPrice)
	if err != nil {
		return nil, fmt.Errorf("invalid gas price: %w", err)
	}
	feeAmount := gasPrice.MulInt64(int64(c.gasLimit)).TruncateInt()
	fee := sdk.NewCoins(sdk.NewCoin(Denom, feeAmount))
	txBuilder.SetFeeAmount(fee)
	txBuilder.SetMemo(memo)

	// Sign transaction
	signerData := authsigning.SignerData{
		ChainID:       c.config.ChainID,
		AccountNumber: account.AccountNumber,
		Sequence:      account.Sequence,
	}

	// Create signature placeholder
	sigV2 := signing.SignatureV2{
		PubKey: c.privateKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
			Signature: nil,
		},
		Sequence: account.Sequence,
	}

	if err := txBuilder.SetSignatures(sigV2); err != nil {
		return nil, fmt.Errorf("failed to set signatures: %w", err)
	}

	// Generate sign bytes
	signBytes, err := authsigning.GetSignBytesAdapter(
		context.Background(),
		c.txConfig.SignModeHandler(),
		signing.SignMode_SIGN_MODE_DIRECT,
		signerData,
		txBuilder.GetTx(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get sign bytes: %w", err)
	}

	// Sign the bytes
	signature, err := c.privateKey.Sign(signBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Set final signature
	sigV2 = signing.SignatureV2{
		PubKey: c.privateKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
			Signature: signature,
		},
		Sequence: account.Sequence,
	}

	if err := txBuilder.SetSignatures(sigV2); err != nil {
		return nil, fmt.Errorf("failed to set final signature: %w", err)
	}

	// Encode transaction
	txBytes, err := c.txConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, fmt.Errorf("failed to encode transaction: %w", err)
	}

	// Broadcast transaction
	broadcastReq := struct {
		TxBytes string `json:"tx_bytes"`
		Mode    string `json:"mode"`
	}{
		TxBytes: hex.EncodeToString(txBytes),
		Mode:    "BROADCAST_MODE_SYNC",
	}

	respBody, err := c.httpPost("/cosmos/tx/v1beta1/txs", broadcastReq)
	if err != nil {
		return nil, fmt.Errorf("failed to broadcast transaction: %w", err)
	}

	var broadcastResp struct {
		TxResponse struct {
			TxHash    string `json:"txhash"`
			Code      uint32 `json:"code"`
			RawLog    string `json:"raw_log"`
			GasUsed   string `json:"gas_used"`
			GasWanted string `json:"gas_wanted"`
			Height    string `json:"height"`
		} `json:"tx_response"`
	}

	if err := json.Unmarshal(respBody, &broadcastResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal broadcast response: %w", err)
	}

	resp := broadcastResp.TxResponse
	if resp.Code != 0 {
		return nil, fmt.Errorf("transaction failed (code %d): %s", resp.Code, resp.RawLog)
	}

	gasUsed, _ := strconv.ParseUint(resp.GasUsed, 10, 64)
	gasWanted, _ := strconv.ParseUint(resp.GasWanted, 10, 64)
	height, _ := strconv.ParseInt(resp.Height, 10, 64)

	return &TxResponse{
		TxHash:    resp.TxHash,
		Code:      resp.Code,
		RawLog:    resp.RawLog,
		GasUsed:   gasUsed,
		GasWanted: gasWanted,
		Height:    height,
	}, nil
}

// getNetworkConfig returns network configuration for the given network
func getNetworkConfig(network Network) *NetworkConfig {
	switch network {
	case Mainnet:
		return &NetworkConfig{
			RPC:     MainnetRPC,
			API:     MainnetAPI,
			ChainID: MainnetChainID,
		}
	case Testnet:
		return &NetworkConfig{
			RPC:     TestnetRPC,
			API:     TestnetAPI,
			ChainID: TestnetChainID,
		}
	case Local:
		return &NetworkConfig{
			RPC:     LocalRPC,
			API:     LocalAPI,
			ChainID: LocalChainID,
		}
	default:
		return nil
	}
}
