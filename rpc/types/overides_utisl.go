package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

func ToProtoOverrideAccount(oa *OverrideAccount) *evmtypes.OverrideAccount {
	// Transform Nonce
	nonce := uint64(0)
	if oa.Nonce != nil {
		nonce = uint64(*oa.Nonce)
	}

	// Transform Code
	code := []byte{}
	if oa.Code != nil {
		code = *oa.Code
	}

	// Transform Balance
	balance := ""
	if oa.Balance != nil && *oa.Balance != nil {
		balance = (*oa.Balance).String()
	}

	// Transform State
	state := make(map[string]string)
	if oa.State != nil {
		for k, v := range *oa.State {
			state[k.Hex()] = v.Hex()
		}
	}

	// Transform StateDiff
	stateDiff := make(map[string]string)
	if oa.StateDiff != nil {
		for k, v := range *oa.StateDiff {
			stateDiff[k.Hex()] = v.Hex()
		}
	}

	return &evmtypes.OverrideAccount{
		Nonce:     nonce,
		Code:      code,
		Balance:   balance,
		State:     state,
		StateDiff: stateDiff,
	}
}

// Detransform from the protobuf-generated OverrideAccount to the original OverrideAccount
func FromProtoOverrideAccount(protoOA evmtypes.OverrideAccount) OverrideAccount {
	// Detransform Nonce
	var nonce *hexutil.Uint64
	if protoOA.Nonce != 0 {
		convertedNonce := hexutil.Uint64(protoOA.Nonce)
		nonce = &convertedNonce
	}

	// Detransform Code
	var code *hexutil.Bytes
	if len(protoOA.Code) > 0 {
		convertedCode := hexutil.Bytes(protoOA.Code)
		code = &convertedCode
	}

	// Detransform Balance
    var balance **hexutil.Big
    if protoOA.Balance != "" {
        bigIntBalance := new(big.Int)
        bigIntBalance.SetString(protoOA.Balance, 10)
        convertedBalance := (*hexutil.Big)(bigIntBalance)
        balance = &convertedBalance
    }

	// Detransform State
	var state *map[common.Hash]common.Hash
	if len(protoOA.State) > 0 {
		transformedState := make(map[common.Hash]common.Hash)
		for k, v := range protoOA.State {
			transformedState[common.HexToHash(k)] = common.HexToHash(v)
		}
		state = &transformedState
	}

	// Detransform StateDiff
	var stateDiff *map[common.Hash]common.Hash
	if len(protoOA.StateDiff) > 0 {
		transformedStateDiff := make(map[common.Hash]common.Hash)
		for k, v := range protoOA.StateDiff {
			transformedStateDiff[common.HexToHash(k)] = common.HexToHash(v)
		}
		stateDiff = &transformedStateDiff
	}

	return OverrideAccount{
		Nonce:     nonce,
		Code:      code,
		Balance:   balance,
		State:     state,
		StateDiff: stateDiff,
	}
}


// ToProtoStateOverride transforms a StateOverride to its ProtoStateOverride equivalent.
func ToProtoStateOverride(so *StateOverride) *evmtypes.StateOverride {
    protoSO := &evmtypes.StateOverride{
        Accounts: make(map[string]*evmtypes.OverrideAccount),
    }

    for address, overrideAccount := range *so {
        protoSO.Accounts[address.Hex()] = ToProtoOverrideAccount(&overrideAccount)
    }

    return protoSO
}

// FromProtoStateOverride detransforms a ProtoStateOverride back to a StateOverride.
func FromProtoStateOverride(protoSO *evmtypes.StateOverride) StateOverride {
    so := make(StateOverride)

    for address, protoOverrideAccount := range protoSO.Accounts {
        so[common.HexToAddress(address)] = FromProtoOverrideAccount(*protoOverrideAccount)
    }

    return so
}


func ToProtoBlockOverride(bo *BlockOverrides) *evmtypes.BlockOverrides {
	protoBO := &evmtypes.BlockOverrides{
		Number:      bo.Number.String(),
		Defficulty:  bo.Difficulty.String(),
		Time:        uint64(*bo.Time),
		GasLimit:    uint64(*bo.GasLimit),
		Coinbase:    bo.Coinbase.String(),
		Random:      bo.Random.String(),
		BaseFee:     bo.BaseFee.String(),
		BlobBaseFee: bo.BlobBaseFee.String(),
	}
	return protoBO
}

func FromProtoBlockOverride(protoBO *evmtypes.BlockOverrides) BlockOverrides{
	var number *hexutil.Big
	if protoBO.Number != "" {
		bigIntBNumber := new(big.Int)
		bigIntBNumber.SetString(protoBO.Number, 10)
		number = (*hexutil.Big)(bigIntBNumber)
	}

	var difficulty *hexutil.Big
	if protoBO.Defficulty != "" {
		bigIntDiff := new(big.Int)
		bigIntDiff.SetString(protoBO.Defficulty, 10)
		difficulty = (*hexutil.Big)(bigIntDiff)
	}

	var coinBase *common.Address
	if protoBO.Coinbase != "" {
		coinBaseAddree := common.HexToAddress(protoBO.Coinbase)
		coinBase = &coinBaseAddree
	}

	var time *hexutil.Uint64
	if protoBO.Time != 0 {
		convertedTime := hexutil.Uint64(protoBO.Time)
		time = &convertedTime
	}

	var gasLimit *hexutil.Uint64
	if protoBO.Time != 0 {
		convertedGasLimit := hexutil.Uint64(protoBO.GasLimit)
		gasLimit = &convertedGasLimit
	}

	var random *common.Hash
	if protoBO.Random != "" {
		randomHash := common.HexToHash(protoBO.Random)
		random = &randomHash
	}

	var baseFee *hexutil.Big
	if protoBO.BaseFee != "" {
		bigIntBaseFee := new(big.Int)
		bigIntBaseFee.SetString(protoBO.BaseFee, 10)
		baseFee = (*hexutil.Big)(bigIntBaseFee)
	}

	var blobFee *hexutil.Big
	if protoBO.BlobBaseFee != "" {
		bigIntBlobFee := new(big.Int)
		bigIntBlobFee.SetString(protoBO.BlobBaseFee, 10)
		blobFee = (*hexutil.Big)(bigIntBlobFee)
	}

	bo := BlockOverrides{
		Number:      number,
		Difficulty:  difficulty,
		Time:        time,
		GasLimit:    gasLimit,
		Coinbase:    coinBase,
		Random:      random,
		BaseFee:     baseFee,
		BlobBaseFee: blobFee,
	}

	return bo
}