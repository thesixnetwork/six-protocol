// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
package evm_test

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	testutil "github.com/evmos/evmos/v20/testutil"
	"github.com/evmos/evmos/v20/testutil/integration/evmos/grpc"
	testkeyring "github.com/evmos/evmos/v20/testutil/integration/evmos/keyring"
	"github.com/evmos/evmos/v20/testutil/integration/evmos/network"
	evmtypes "github.com/evmos/evmos/v20/x/evm/types"

	"github.com/thesixnetwork/six-protocol/v4/app/ante/evm"
)

func (suite *EvmAnteTestSuite) TestIncrementSequence() {
	keyring := testkeyring.New(1)
	unitNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	grpcHandler := grpc.NewIntegrationHandler(unitNetwork)
	accAddr := keyring.GetAccAddr(0)

	testCases := []struct {
		name          string
		expectedError error
		malleate      func(acct sdk.AccountI) uint64
	}{
		{
			name:          "fail: invalid sequence",
			expectedError: errortypes.ErrInvalidSequence,
			malleate: func(acct sdk.AccountI) uint64 {
				return acct.GetSequence() + 1
			},
		},
		{
			name:          "success: increments sequence",
			expectedError: nil,
			malleate: func(acct sdk.AccountI) uint64 {
				return acct.GetSequence()
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			account, err := grpcHandler.GetAccount(accAddr.String())
			suite.Require().NoError(err)
			preSequence := account.GetSequence()

			nonce := tc.malleate(account)

			// Create mock Ethereum transaction with the nonce
			evmTxArgs := evmtypes.EvmTxArgs{
				ChainID:  big.NewInt(1),
				Nonce:    nonce,
				To:       &common.Address{},
				Amount:   big.NewInt(0),
				GasLimit: 100000,
				GasPrice: big.NewInt(1),
				Input:    []byte{},
			}

			// Create mock MsgEthereumTx
			msgEthereumTx := evmtypes.NewTx(&evmTxArgs)
			msgEthereumTx.From = accAddr.String()

			// Create mock Tx containing the MsgEthereumTx
			mockTx := &testutil.MockTx{
				Msgs: []sdk.Msg{msgEthereumTx},
			}

			// Function under test
			err = evm.IncrementNonce(
				unitNetwork.GetContext(),
				unitNetwork.App.AccountKeeper,
				account,
				mockTx,
				tc.expectedError == nil, // If no error expected, use true (safe ordering), otherwise false
			)

			if tc.expectedError != nil {
				suite.Require().Error(err)
				suite.Contains(err.Error(), tc.expectedError.Error())
			} else {
				suite.Require().NoError(err)

				suite.Require().Equal(preSequence+1, account.GetSequence())
				updatedAccount, err := grpcHandler.GetAccount(accAddr.String())
				suite.Require().NoError(err)
				suite.Require().Equal(preSequence+1, updatedAccount.GetSequence())
			}
		})
	}
}

// TestIncrementSequenceUnorderedCheckTx is a regression test for the foundry
// "fast" deploy failure ("nonce too low; got 1, expected >= 5"). In unordered
// mode during mempool admission (CheckTx), several txs from one sender arrive
// out of order. They must all be admitted without error, and the sequence must
// NOT advance during admission — otherwise a lower nonce arriving after a batch
// of higher ones would be wrongly rejected as stale. The nonce mempool re-orders
// them and the strict DeliverTx check enforces ordering at execution time.
func (suite *EvmAnteTestSuite) TestIncrementSequenceUnorderedCheckTx() {
	keyring := testkeyring.New(1)
	unitNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	grpcHandler := grpc.NewIntegrationHandler(unitNetwork)
	accAddr := keyring.GetAccAddr(0)

	account, err := grpcHandler.GetAccount(accAddr.String())
	suite.Require().NoError(err)
	baseSeq := account.GetSequence()

	// CheckTx context: mempool admission path.
	checkCtx := unitNetwork.GetContext().WithIsCheckTx(true)

	newTx := func(nonce uint64) sdk.Tx {
		evmTxArgs := evmtypes.EvmTxArgs{
			ChainID:  big.NewInt(1),
			Nonce:    nonce,
			To:       &common.Address{},
			Amount:   big.NewInt(0),
			GasLimit: 100000,
			GasPrice: big.NewInt(1),
			Input:    []byte{},
		}
		msgEthereumTx := evmtypes.NewTx(&evmTxArgs)
		msgEthereumTx.From = accAddr.String()
		return &testutil.MockTx{Msgs: []sdk.Msg{msgEthereumTx}}
	}

	call := func(nonce uint64) error {
		return evm.IncrementNonce(
			checkCtx,
			unitNetwork.App.AccountKeeper,
			account,
			newTx(nonce),
			true, // unsafeUnorderedTx (unordered mode)
		)
	}

	// Out-of-order arrival of nonces baseSeq+{0,2,3,4,5} then the late baseSeq+1.
	// Every current-or-future nonce is admitted, and the late lower nonce (+1)
	// is NOT rejected because admission does not advance the sequence.
	for _, off := range []uint64{0, 2, 3, 4, 5, 1} {
		suite.Require().NoError(call(baseSeq+off), "nonce %d should be admitted in unordered CheckTx", baseSeq+off)
	}
	suite.Require().Equal(baseSeq, account.GetSequence(), "sequence must not advance during CheckTx admission")

	// A nonce below the committed sequence is a replay/stale tx and must be
	// rejected even during admission, keeping it out of the mempool.
	if baseSeq > 0 {
		suite.Require().Error(call(baseSeq-1), "stale nonce below committed sequence must be rejected")
	}
}
