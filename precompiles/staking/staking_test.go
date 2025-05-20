package staking_test

import (
	"fmt"

	stakingmodule "github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/thesixnetwork/six-protocol/precompiles/staking"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (p *StakingPrecompileTestSuite) TestIsTransaction() {

	testCases := []struct {
		name string
		// method abi.Method
		isTx bool
	}{
		{
			staking.DelegateMethod,
			// p.precompile.Methods[staking.DelegateMethod],
			true,
		},
		{
			staking.RedelegateMethod,
			// p.precompile.Methods[staking.RedelegateMethod],
			true,
		},
		{
			staking.UndelegateMethod,
			// p.precompile.Methods[staking.UndelegateMethod],
			true,
		},
		{
			staking.DelegationMethod,
			// p.precompile.Methods[staking.DelegationMethod],
			false,
		},
	}

	for _, tc := range testCases {
		p.Run(tc.name, func() {
			p.Require().Equal(p.stakingPrecompile.IsTransaction(tc.name), tc.isTx)
		})
	}
}

func (p *StakingPrecompileTestSuite) TestExecute() {
	// Prepare test data for different methods
	caller := common.BytesToAddress(p.user.Address.Bytes()) // Convert Cosmos address to Ethereum address
	fmt.Println("caller:", caller)
	callingContract := common.Address{} // Example calling contract address

	// Convert string amount to *big.Int
	delegateAmount := new(big.Int)
	delegateAmount.SetString("10000000000000000", 10) // 0.01 SIX (in wei)
	fmt.Println("delegateAmount:", delegateAmount)

	expected := common.LeftPadBytes([]byte{0x01}, 32)

	// Create arguments for each method type
	delegateArgs := []interface{}{
		p.validator.Address.String(), // Validator's Bech32 address
		delegateAmount,
	}
	// redelegateArgs := []interface{}{
	// 	p.srcValidator.Address.String(), // Source Validator
	// 	p.dstValidator.Address.String(), // Destination Validator
	// 	delegateAmount,
	// }
	undelegateArgs := []interface{}{
		p.validator.Address.String(), // Validator's Bech32 address
		delegateAmount,
	}

	// Method setup
	methods := map[string]abi.Method{
		staking.DelegateMethod:   p.stakingPrecompile.Methods[staking.DelegateMethod],
		staking.RedelegateMethod: p.stakingPrecompile.Methods[staking.RedelegateMethod],
		staking.UndelegateMethod: p.stakingPrecompile.Methods[staking.UndelegateMethod],
	}

	// Test cases for all methods
	testCases := []struct {
		methodName     string
		args           []interface{}
		expectedResult []byte
	}{
		{
			methodName:     staking.DelegateMethod,
			args:           delegateArgs,
			expectedResult: expected,
		},
		// {
		// 	methodName:     staking.RedelegateMethod,
		// 	args:           redelegateArgs,
		// 	expectedResult: expected,
		// },
		{
			methodName:     staking.UndelegateMethod,
			args:           undelegateArgs,
			expectedResult: expected,
		},
	}

	// Loop over each test case
	for _, tc := range testCases {
		p.Run(tc.methodName, func() {
			method := methods[tc.methodName]

			// Pre-delegate to srcValidator using precompile if testing Redelegate
			if tc.methodName == staking.RedelegateMethod {

				fmt.Println("srcValidator:", p.srcValidator.Address.String())
				fmt.Println("dstValidator:", p.dstValidator.Address.String())

				preDelegateMethod := methods[staking.DelegateMethod]

				// preDelegatedstArgs := []interface{}{
				// 	p.dstValidator.Address.String(), // Delegate to srcValidator
				// 	delegateAmount,
				// }

				// preDelegatedstResult, err := p.stakingPrecompile.Execute(
				// 	p.ctx,
				// 	&preDelegateMethod,
				// 	caller,
				// 	callingContract,
				// 	preDelegatedstArgs,
				// 	big.NewInt(0),
				// 	false,
				// 	nil,
				// )
				// p.Require().NoError(err, "pre-delegate to dstValidator via precompile failed")
				// p.Require().Equal(expected, preDelegatedstResult, "unexpected result from pre-delegate")
				// delegationdst, found := p.stakingKeeper.GetDelegation(p.ctx, caller.Bytes(), p.dstValidator.Address)
				// fmt.Println("Has delegation dst:", found, "amount:", delegationdst.GetShares())

				preDelegateArgs := []interface{}{
					p.srcValidator.Address.String(), // Delegate to srcValidator
					delegateAmount,
				}

				preDelegateResult, err := p.stakingPrecompile.Execute(
					p.ctx,
					&preDelegateMethod,
					caller,
					callingContract,
					preDelegateArgs,
					big.NewInt(0),
					false,
					nil,
				)
				p.Require().NoError(err, "pre-delegate to srcValidator via precompile failed")
				// 💡 Finalize validator state
				stakingmodule.EndBlocker(p.ctx, *p.stakingKeeper)

				bondedVals := p.stakingKeeper.GetBondedValidatorsByPower(p.ctx)
				for i, v := range bondedVals {
					fmt.Printf("✅ Active validator #%d: %s (Power: %v)\n", i, v.OperatorAddress, v.Tokens)
				}

				val, _ := p.stakingKeeper.GetValidator(p.ctx, p.srcValidator.Address)
				fmt.Println("srcValidator status after EndBlocker:", val.Status.String())

				vals, _ := p.stakingKeeper.GetValidator(p.ctx, p.srcValidator.Address)
				fmt.Println("src tokens:", vals.Tokens.String(), "delegator shares:", vals.DelegatorShares.String())

				vald, _ := p.stakingKeeper.GetValidator(p.ctx, p.dstValidator.Address)
				fmt.Println("dst tokens:", vald.Tokens.String(), "delegator shares:", vald.DelegatorShares.String())

				p.Require().Equal(expected, preDelegateResult, "unexpected result from pre-delegate")
				delegation, found := p.stakingKeeper.GetDelegation(p.ctx, caller.Bytes(), p.srcValidator.Address)
				fmt.Println("Has delegation:", found, "amount:", delegation.GetShares())

				fmt.Println("src is bonded:", vals.IsBonded())
				fmt.Println("dst is bonded:", vald.IsBonded())

			}

			// Execute the actual test case
			result, err := p.stakingPrecompile.Execute(
				p.ctx,
				&method,
				caller,
				callingContract,
				tc.args,
				big.NewInt(0),
				false,
				nil,
			)

			if err != nil {
				fmt.Printf("Execute failed for method %s: %v\n", tc.methodName, err)
			}

			fmt.Println("tc.expectedResult:", tc.expectedResult, "result:", result)
			p.Require().NotNil(result)
			p.Require().Equal(tc.expectedResult, result)
		})
	}

}
