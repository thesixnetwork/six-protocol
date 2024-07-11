// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	balancetypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	// BalancesMethod defines the ABI method name for the bank Balances
	// query.
	BalancesMethod = "balances"
	// TotalSupplyMethod defines the ABI method name for the bank TotalSupply
	// query.
	TotalSupplyMethod = "totalSupply"
	// SupplyOfMethod defines the ABI method name for the bank SupplyOf
	// query.
	SupplyOfMethod = "supplyOf"
)

// Balances returns all the native token balances (address, amount) for a given
// account. This method charges the account the corresponding value of an ERC-20
// balanceOf call for each token returned.
func (p Precompile) Balances(
	ctx sdk.Context,
	_ *vm.Contract,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	account, err := ParseBalancesArgs(args)
	if err != nil {
		return nil, err
	}

	balance, err := p.bankKeeper.AllBalances(ctx.Context(), &balancetypes.QueryAllBalancesRequest{
		Address: account.String(),
	})

	return method.Outputs.Pack(balance)
}
