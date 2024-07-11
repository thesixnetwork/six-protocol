// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package utils

import (
	"fmt"
	"strings"
	"github.com/ethereum/go-ethereum/common"

	"github.com/evmos/ethermint/crypto/ethsecp256k1"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// MainnetChainID defines the Evmos EIP155 chain ID for mainnet
	MainnetChainID = "sixnet"
	// TestnetChainID defines the Evmos EIP155 chain ID for testnet
	TestnetChainID = "fivenet"
	// TestingChainID defines the Evmos EIP155 chain ID for integration test
	TestingChainID = "testnet"
	// BaseDenom defines the Evmos mainnet denomination
	BaseDenom = "asix"
)

// EthHexToCosmosAddr takes a given Hex string and derives a Cosmos SDK account address
// from it.
func EthHexToCosmosAddr(hexAddr string) sdk.AccAddress {
	return EthToCosmosAddr(common.HexToAddress(hexAddr))
}

// EthToCosmosAddr converts a given Ethereum style address to an SDK address.
func EthToCosmosAddr(addr common.Address) sdk.AccAddress {
	return sdk.AccAddress(addr.Bytes())
}

// Bech32ToHexAddr converts a given Bech32 address string and converts it to
// an Ethereum address.
func Bech32ToHexAddr(bech32Addr string) (common.Address, error) {
	accAddr, err := sdk.AccAddressFromBech32(bech32Addr)
	if err != nil {
		return common.Address{}, errortypes.Wrapf(err, "failed to convert bech32 string to address")
	}

	return CosmosToEthAddr(accAddr), nil
}

// CosmosToEthAddr converts a given SDK account address to
// an Ethereum address.
func CosmosToEthAddr(accAddr sdk.AccAddress) common.Address {
	return common.BytesToAddress(accAddr.Bytes())
}

// IsMainnet returns true if the chain-id has the Evmos mainnet EIP155 chain prefix.
func IsMainnet(chainID string) bool {
	return strings.HasPrefix(chainID, MainnetChainID)
}

// IsTestnet returns true if the chain-id has the Evmos testnet EIP155 chain prefix.
func IsTestnet(chainID string) bool {
	return strings.HasPrefix(chainID, TestnetChainID)
}

// IsTesting returns true if the chain-id has the "test" prefix.
// NOTE: for tests only
func IsTesting(chainID string) bool {
	return strings.HasPrefix(chainID, TestingChainID)
}

// IsSupportedKey returns true if the pubkey type is supported by the chain
// (i.e eth_secp256k1, amino multisig, ed25519).
// NOTE: Nested multisigs are not supported.
func IsSupportedKey(pubkey cryptotypes.PubKey) bool {
	switch pubkey := pubkey.(type) {
	case *ethsecp256k1.PubKey, *ed25519.PubKey:
		return true
	case multisig.PubKey:
		if len(pubkey.GetPubKeys()) == 0 {
			return false
		}

		for _, pk := range pubkey.GetPubKeys() {
			switch pk.(type) {
			case *ethsecp256k1.PubKey, *ed25519.PubKey:
				continue
			default:
				// Nested multisigs are unsupported
				return false
			}
		}

		return true
	default:
		return false
	}
}

// GetEvmosAddressFromBech32 returns the sdk.Account address of given address,
// while also changing bech32 human readable prefix (HRP) to the value set on
// the global sdk.Config (eg: `evmos`).
// The function fails if the provided bech32 address is invalid.
func GetEvmosAddressFromBech32(address string) (sdk.AccAddress, error) {
	bech32Prefix := strings.SplitN(address, "1", 2)[0]
	if bech32Prefix == address {
		return nil, errortypes.Wrapf(errortypes.ErrInvalidAddress, "invalid bech32 address: %s", address)
	}

	addressBz, err := sdk.GetFromBech32(address, bech32Prefix)
	if err != nil {
		return nil, errortypes.Wrapf(errortypes.ErrInvalidAddress, "invalid address %s, %s", address, err.Error())
	}

	// safety check: shouldn't happen
	if err := sdk.VerifyAddressFormat(addressBz); err != nil {
		return nil, err
	}

	return sdk.AccAddress(addressBz), nil
}

// CreateAccAddressFromBech32 creates an AccAddress from a Bech32 string.
func CreateAccAddressFromBech32(address string, bech32prefix string) (addr sdk.AccAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return sdk.AccAddress{}, fmt.Errorf("empty address string is not allowed")
	}

	bz, err := sdk.GetFromBech32(address, bech32prefix)
	if err != nil {
		return nil, err
	}

	err = sdk.VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return sdk.AccAddress(bz), nil
}