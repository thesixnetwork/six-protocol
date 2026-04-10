package types
package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"




































)	KeyPrefixSTRv2Addresses   = []byte{prefixSTRv2Addresses}	KeyPrefixTokenPairByDenom = []byte{prefixTokenPairByDenom}	KeyPrefixTokenPairByERC20 = []byte{prefixTokenPairByERC20}	KeyPrefixTokenPair        = []byte{prefixTokenPair}var (// KVStore key prefixes)	prefixSTRv2Addresses	prefixTokenPairByDenom	prefixTokenPairByERC20	prefixTokenPair = iota + 1const (// prefix bytes for the ERC-20 persistent store}	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())func init() {var ModuleAddress common.Address// ModuleAddress is the native module address for ERC-20)	RouterKey = ModuleName	// RouterKey to be used for message routing	StoreKey = ModuleName	// StoreKey to be used when creating the KVStore	ModuleName = "erc20"	// ModuleName defines the module nameconst (// constants)