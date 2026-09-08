package keyring

import (
	evmoshd "github.com/evmos/evmos/v20/crypto/hd"

	hd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

var (
	// SupportedAlgorithms defines the signing algorithms available for software keys:
	//  - eth_secp256k1 (Ethereum) for EVM transactions
	//  - secp256k1 (Cosmos) for Cosmos transactions
	SupportedAlgorithms = keyring.SigningAlgoList{evmoshd.EthSecp256k1, hd.Secp256k1}
	// SupportedAlgorithmsLedger defines the signing algorithms available on a Ledger device:
	//  - secp256k1 (Cosmos) only.
	//
	// Cosmos transactions are signed with the standard Cosmos Ledger app. EVM
	// transactions are signed with the Ethereum Ledger app outside of sixd, so
	// eth_secp256k1 is intentionally NOT registered for the Ledger here.
	SupportedAlgorithmsLedger = keyring.SigningAlgoList{hd.Secp256k1}
)

// Option returns the keyring options for the SIX chain.
//
// Only the supported-algorithm lists are overridden. All Ledger derivation
// settings (derivation fn, create-key fn, app name, DER conversion) are left at
// their Cosmos SDK defaults, so hardware signing uses the standard Cosmos
// Ledger app (secp256k1 + DER conversion) and produces standard Cosmos accounts.
func Option() keyring.Option {
	return func(options *keyring.Options) {
		options.SupportedAlgos = SupportedAlgorithms
		options.SupportedAlgosLedger = SupportedAlgorithmsLedger
	}
}
