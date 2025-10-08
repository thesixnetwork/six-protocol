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

	"github.com/thesixnetwork/six-protocol/app/ante/evm"
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
