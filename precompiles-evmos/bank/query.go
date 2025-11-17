package bank

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Balances returns all the native token balances (contract address, amount) for a given account.
func (e BankExecutor) Balances(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if len(args) < 1 {
		return nil, ErrInvalidArgs
	}
	account, err := ParseBalancesArgs(args)
	if err != nil {
		return nil, err
	}
	i := 0
	balances := make([]Balance, 0)

	e.bankKeeper.IterateAccountBalances(ctx, account, func(coin sdk.Coin) bool {
		defer func() { i++ }()
		if i > 0 {
			ctx.GasMeter().ConsumeGas(GasBalanceOf, "ERC-20 extension balances method")
		}
		contractAddress, err := e.erc20Keeper.GetCoinAddress(ctx, coin.Denom)
		if err != nil {
			return false
		}
		balances = append(balances, Balance{
			ContractAddress: contractAddress,
			Amount:          coin.Amount.BigInt(),
		})
		return false
	})
	return method.Outputs.Pack(balances)
}

// TotalSupply returns the total supply of all the native tokens.
func (e *BankExecutor) TotalSupply(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	i := 0
	var totalSupply []Balance

	e.bankKeeper.IterateTotalSupply(ctx, func(coin sdk.Coin) bool {
		defer func() { i++ }()
		if i > 0 {
			ctx.GasMeter().ConsumeGas(GasTotalSupply, "ERC-20 extension totalSupply method")
		}
		contractAddress, err := e.erc20Keeper.GetCoinAddress(ctx, coin.Denom)
		if err != nil {
			return false
		}
		totalSupply = append(totalSupply, Balance{
			ContractAddress: contractAddress,
			Amount:          coin.Amount.BigInt(),
		})
		return false
	})
	return method.Outputs.Pack(totalSupply)
}

// SupplyOf returns the total native supply of a given registered erc20 token.
func (e *BankExecutor) SupplyOf(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if len(args) < 1 {
		return nil, ErrInvalidArgs
	}
	erc20ContractAddress, err := ParseSupplyOfArgs(args)
	if err != nil {
		return nil, err
	}

	tokenPairID := e.erc20Keeper.GetERC20Map(ctx, erc20ContractAddress)
	tokenPair, found := e.erc20Keeper.GetTokenPair(ctx, tokenPairID)
	if !found {
		return method.Outputs.Pack(big.NewInt(0))
	}

	supply := e.bankKeeper.GetSupply(ctx, tokenPair.Denom)
	return method.Outputs.Pack(supply.Amount.BigInt())
}

// You'd define this error in a common error location:
var ErrInvalidArgs = fmt.Errorf("bank precompile: invalid arguments")
