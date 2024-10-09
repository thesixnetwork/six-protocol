// Copyright (C) 2024 SIX Network
// This file is part of the modified EVM module from Ethermint (https://github.com/evmos/ethermint),
// and is licensed under the terms of the GNU Lesser General Public License v3
package keeper_test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
	evmtypes "github.com/thesixnetwork/six-protocol/x/evm/types"
)

var templateAccessListTx = &ethtypes.AccessListTx{
	GasPrice: big.NewInt(1),
	Gas:      21000,
	To:       &common.Address{},
	Value:    big.NewInt(0),
	Data:     []byte{},
}

var templateLegacyTx = &ethtypes.LegacyTx{
	GasPrice: big.NewInt(1),
	Gas:      21000,
	To:       &common.Address{},
	Value:    big.NewInt(0),
	Data:     []byte{},
}

var templateDynamicFeeTx = &ethtypes.DynamicFeeTx{
	GasFeeCap: big.NewInt(10),
	GasTipCap: big.NewInt(2),
	Gas:       21000,
	To:        &common.Address{},
	Value:     big.NewInt(0),
	Data:      []byte{},
}

func newSignedEthTx(
	txData ethtypes.TxData,
	nonce uint64,
	addr sdk.Address,
	krSigner keyring.Signer,
	ethSigner ethtypes.Signer,
) (*ethtypes.Transaction, error) {
	var ethTx *ethtypes.Transaction
	switch txData := txData.(type) {
	case *ethtypes.AccessListTx:
		txData.Nonce = nonce
		ethTx = ethtypes.NewTx(txData)
	case *ethtypes.LegacyTx:
		txData.Nonce = nonce
		ethTx = ethtypes.NewTx(txData)
	case *ethtypes.DynamicFeeTx:
		txData.Nonce = nonce
		ethTx = ethtypes.NewTx(txData)
	default:
		return nil, errors.New("unknown transaction type!")
	}

	sig, _, err := krSigner.SignByAddress(addr, ethTx.Hash().Bytes())
	if err != nil {
		return nil, err
	}

	ethTx, err = ethTx.WithSignature(ethSigner, sig)
	if err != nil {
		return nil, err
	}

	return ethTx, nil
}

func newNativeMessage(
	nonce uint64,
	blockHeight int64,
	blockTime uint64,
	address common.Address,
	cfg *params.ChainConfig,
	krSigner keyring.Signer,
	ethSigner ethtypes.Signer,
	txType byte,
	data []byte,
	accessList ethtypes.AccessList,
) (core.Message, error) {
	msgSigner := ethtypes.MakeSigner(cfg, big.NewInt(blockHeight), blockTime)

	var (
		ethTx   *ethtypes.Transaction
		baseFee *big.Int
	)

	switch txType {
	case ethtypes.LegacyTxType:
		templateLegacyTx.Nonce = nonce
		if data != nil {
			templateLegacyTx.Data = data
		}
		ethTx = ethtypes.NewTx(templateLegacyTx)
	case ethtypes.AccessListTxType:
		templateAccessListTx.Nonce = nonce
		if data != nil {
			templateAccessListTx.Data = data
		} else {
			templateAccessListTx.Data = []byte{}
		}

		templateAccessListTx.AccessList = accessList
		ethTx = ethtypes.NewTx(templateAccessListTx)
	case ethtypes.DynamicFeeTxType:
		templateDynamicFeeTx.Nonce = nonce

		if data != nil {
			templateAccessListTx.Data = data
		} else {
			templateAccessListTx.Data = []byte{}
		}
		templateAccessListTx.AccessList = accessList
		ethTx = ethtypes.NewTx(templateDynamicFeeTx)
		baseFee = big.NewInt(3)
	default:
		return core.Message{}, errors.New("unsupport tx type")
	}

	msg := &evmtypes.MsgEthereumTx{}
	msg.FromEthereumTx(ethTx)
	msg.From = address.Hex()

	if err := msg.Sign(ethSigner, krSigner); err != nil {
		return core.Message{}, err
	}

	m, err := msg.AsMessage(ethTx, msgSigner, baseFee)
	if err != nil {
		return core.Message{}, err
	}

	return m, nil
}

func BenchmarkApplyTransaction(b *testing.B) {
	suite := KeeperTestSuite{enableLondonHF: true}
	suite.SetupTest()

	ethSigner := ethtypes.LatestSignerForChainID(suite.app.EVMKeeper.ChainID())

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tx, err := newSignedEthTx(templateAccessListTx,
			suite.app.EVMKeeper.GetNonce(suite.ctx, suite.address),
			sdk.AccAddress(suite.address.Bytes()),
			suite.signer,
			ethSigner,
		)
		require.NoError(b, err)

		b.StartTimer()
		resp, err := suite.app.EVMKeeper.ApplyTransaction(suite.ctx, tx)
		b.StopTimer()

		require.NoError(b, err)
		require.False(b, resp.Failed())
	}
}

func BenchmarkApplyTransactionWithLegacyTx(b *testing.B) {
	suite := KeeperTestSuite{enableLondonHF: true}
	suite.SetupTest()

	ethSigner := ethtypes.LatestSignerForChainID(suite.app.EVMKeeper.ChainID())

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tx, err := newSignedEthTx(templateLegacyTx,
			suite.app.EVMKeeper.GetNonce(suite.ctx, suite.address),
			sdk.AccAddress(suite.address.Bytes()),
			suite.signer,
			ethSigner,
		)
		require.NoError(b, err)

		b.StartTimer()
		resp, err := suite.app.EVMKeeper.ApplyTransaction(suite.ctx, tx)
		b.StopTimer()

		require.NoError(b, err)
		require.False(b, resp.Failed())
	}
}

func BenchmarkApplyTransactionWithDynamicFeeTx(b *testing.B) {
	suite := KeeperTestSuite{enableFeemarket: true, enableLondonHF: true}
	suite.SetupTest()

	ethSigner := ethtypes.LatestSignerForChainID(suite.app.EVMKeeper.ChainID())

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		tx, err := newSignedEthTx(templateDynamicFeeTx,
			suite.app.EVMKeeper.GetNonce(suite.ctx, suite.address),
			sdk.AccAddress(suite.address.Bytes()),
			suite.signer,
			ethSigner,
		)
		require.NoError(b, err)

		b.StartTimer()
		resp, err := suite.app.EVMKeeper.ApplyTransaction(suite.ctx, tx)
		b.StopTimer()

		require.NoError(b, err)
		require.False(b, resp.Failed())
	}
}

func BenchmarkApplyMessage(b *testing.B) {
	suite := KeeperTestSuite{enableLondonHF: true}
	suite.SetupTest()

	params := suite.app.EVMKeeper.GetParams(suite.ctx)
	ethCfg := params.ChainConfig.EthereumConfig(suite.app.EVMKeeper.ChainID())
	signer := ethtypes.LatestSignerForChainID(suite.app.EVMKeeper.ChainID())

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		m, err := newNativeMessage(
			suite.app.EVMKeeper.GetNonce(suite.ctx, suite.address),
			suite.ctx.BlockHeight(),
			uint64(suite.ctx.BlockTime().Unix()),
			suite.address,
			ethCfg,
			suite.signer,
			signer,
			ethtypes.AccessListTxType,
			nil,
			nil,
		)
		require.NoError(b, err)

		b.StartTimer()
		resp, err := suite.app.EVMKeeper.ApplyMessage(suite.ctx, m, nil, true)
		b.StopTimer()

		require.NoError(b, err)
		require.False(b, resp.Failed())
	}
}

func BenchmarkApplyMessageWithLegacyTx(b *testing.B) {
	suite := KeeperTestSuite{enableLondonHF: true}
	suite.SetupTest()

	params := suite.app.EVMKeeper.GetParams(suite.ctx)
	ethCfg := params.ChainConfig.EthereumConfig(suite.app.EVMKeeper.ChainID())
	signer := ethtypes.LatestSignerForChainID(suite.app.EVMKeeper.ChainID())

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		m, err := newNativeMessage(
			suite.app.EVMKeeper.GetNonce(suite.ctx, suite.address),
			suite.ctx.BlockHeight(),
			uint64(suite.ctx.BlockTime().Unix()),
			suite.address,
			ethCfg,
			suite.signer,
			signer,
			ethtypes.LegacyTxType,
			nil,
			nil,
		)
		require.NoError(b, err)

		b.StartTimer()
		resp, err := suite.app.EVMKeeper.ApplyMessage(suite.ctx, m, nil, true)
		b.StopTimer()

		require.NoError(b, err)
		require.False(b, resp.Failed())
	}
}

func BenchmarkApplyMessageWithDynamicFeeTx(b *testing.B) {
	suite := KeeperTestSuite{enableFeemarket: true, enableLondonHF: true}
	suite.SetupTest()

	params := suite.app.EVMKeeper.GetParams(suite.ctx)
	ethCfg := params.ChainConfig.EthereumConfig(suite.app.EVMKeeper.ChainID())
	signer := ethtypes.LatestSignerForChainID(suite.app.EVMKeeper.ChainID())

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		m, err := newNativeMessage(
			suite.app.EVMKeeper.GetNonce(suite.ctx, suite.address),
			suite.ctx.BlockHeight(),
			uint64(suite.ctx.BlockTime().Unix()),
			suite.address,
			ethCfg,
			suite.signer,
			signer,
			ethtypes.DynamicFeeTxType,
			nil,
			nil,
		)
		require.NoError(b, err)

		b.StartTimer()
		resp, err := suite.app.EVMKeeper.ApplyMessage(suite.ctx, m, nil, true)
		b.StopTimer()

		require.NoError(b, err)
		require.False(b, resp.Failed())
	}
}
