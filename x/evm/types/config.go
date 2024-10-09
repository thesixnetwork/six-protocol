// Copyright (C) 2024 SIX Network
// This file is part of the modified EVM module from Ethermint, 
// and is licensed under the terms of the GNU Lesser General Public License v3
package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

// EVMConfig encapsulates common parameters needed to create an EVM to execute a message
// It's mainly to reduce the number of method parameters
type EVMConfig struct {
	Params      Params
	ChainConfig *params.ChainConfig
	CoinBase    common.Address
	BaseFee     *big.Int
}
