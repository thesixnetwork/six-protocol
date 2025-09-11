// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
package tx

import (


	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"



	"github.com/evmos/ethermint/crypto/ethsecp256k1"
)

// NewAddrKey generates an Ethereum address and its corresponding private key.
func NewAddrKey() (common.Address, *ethsecp256k1.PrivKey) {
	privkey, _ := ethsecp256k1.GenerateKey()
	key, err := privkey.ToECDSA()
	if err != nil {
		return common.Address{}, nil
	}

	addr := crypto.PubkeyToAddress(key.PublicKey)

	return addr, privkey
}