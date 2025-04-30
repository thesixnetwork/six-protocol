package staking_test

import (
	// "math/big"
	// "time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/suite"
	"github.com/thesixnetwork/six-protocol/precompiles/staking"
	testkeyring "github.com/thesixnetwork/six-protocol/testutil/integation/keyring"
	"github.com/thesixnetwork/six-protocol/testutil/network"
)

type PrecompileTestSuite struct {
	suite.Suite

	precompile *staking.PrecompileExecutor
	keyring    testkeyring.Keyring
	network    *network.UnitTestNetwork
}

func (p *PrecompileTestSuite) TestIsTransaction() {
	testCases := []struct {
		name   string
		method abi.Method
		isTx   bool
	}{
		{
			staking.DelegateMethod,
			p.precompile.Methods[staking.DelegateMethod],
			true,
		},
		{
			staking.RedelegateMethod,
			p.precompile.Methods[staking.RedelegateMethod],
			true,
		},
		{
			staking.UndelegateMethod,
			p.precompile.Methods[staking.UndelegateMethod],
			true,
		},
		{
			staking.DelegationMethod,
			p.precompile.Methods[staking.DelegationMethod],
			false,
		},
	}

	for _, tc := range testCases {
		p.Run(tc.name, func() {
			p.Require().Equal(p.precompile.IsTransaction(tc.method.Name), tc.isTx)
		})
	}

}

// func (p *PrecompileTestSuite) TestRequiredGas() {
// 	testcases := []struct {
// 		name     string
// 		malleate func() []byte
// 		expGas   uint64
// 		method   abi.Method
// 	}{
// 		{
// 			"success - delegate transaction with correct gas estimation",
// 			func() []byte {
// 				input, err := p.precompile.Pack(
// 					staking.DelegateMethod,
// 					p.keyring.GetAddr(0),
// 					p.network.GetValidators()[0].GetOperator(),
// 					big.NewInt(10000000000),
// 				)
// 				p.Require().NoError(err)
// 				return input
// 			},
// 			7760,
// 			p.precompile.Methods[staking.DelegateMethod],
// 		},
// 		{
// 			"success - undelegate transaction with correct gas estimation",
// 			func() []byte {
// 				input, err := p.precompile.Pack(
// 					staking.UndelegateMethod,
// 					p.keyring.GetAddr(0),
// 					p.network.GetValidators()[0].GetOperator(),
// 					big.NewInt(1),
// 				)
// 				p.Require().NoError(err)
// 				return input
// 			},
// 			7760,
// 			p.precompile.Methods[staking.UndelegateMethod],
// 		},
// 	}

// 	for _, tc := range testcases {
// 		p.Run(tc.name, func() {

// 			// malleate contract input
// 			input := tc.malleate()
// 			gas := p.precompile.RequiredGas(input, &tc.method)

// 			p.Require().Equal(gas, tc.expGas)
// 		})
// 	}

// }




