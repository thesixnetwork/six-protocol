package precompiles

import (
	"bytes"
	"math/big"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v20/x/evm/core/vm"

	pcommon "github.com/thesixnetwork/six-protocol/v4/precompiles/common"
	"github.com/thesixnetwork/six-protocol/v4/utils"
)

// ExamplePrecompileExecutor demonstrates how to implement balance change tracking
// in your precompiles to prevent EVM stateDB from overwriting bank keeper changes
type ExamplePrecompileExecutor struct {
	bankKeeper pcommon.BankKeeper
	address    common.Address
	precompile *pcommon.Precompile
}

// NewExamplePrecompile creates a new example precompile with balance tracking
func NewExamplePrecompile(bankKeeper pcommon.BankKeeper) (*pcommon.Precompile, error) {
	executor := &ExamplePrecompileExecutor{
		bankKeeper: bankKeeper,
		address:    common.HexToAddress("0x0000000000000000000000000000000000001999"), // Example address
	}

	// Create a minimal ABI for demonstration
	abiJSON := `[
		{
			"name": "transferTokens",
			"type": "function",
			"stateMutability": "nonpayable",
			"inputs": [
				{"name": "to", "type": "address"},
				{"name": "amount", "type": "uint256"}
			],
			"outputs": [
				{"name": "success", "type": "bool"}
			]
		}
	]`

	newAbi, err := abi.JSON(bytes.NewReader([]byte(abiJSON)))
	if err != nil {
		return nil, err
	}

	precompile := pcommon.NewPrecompile(newAbi, executor, executor.address, "example")
	executor.precompile = precompile

	return precompile, nil
}

func (p *ExamplePrecompileExecutor) Address() common.Address {
	return p.address
}

func (p *ExamplePrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	return pcommon.DefaultGasCost(input, true)
}

func (p *ExamplePrecompileExecutor) Execute(
	ctx sdk.Context,
	method *abi.Method,
	caller common.Address,
	callingContract common.Address,
	args []interface{},
	value *big.Int,
	readOnly bool,
	evm *vm.EVM,
) ([]byte, error) {
	switch method.Name {
	case "transferTokens":
		return p.transferTokens(ctx, caller, method, args, value, readOnly)
	}
	return nil, nil
}

// transferTokens demonstrates balance tracking for a token transfer operation
func (p *ExamplePrecompileExecutor) transferTokens(
	ctx sdk.Context,
	caller common.Address,
	method *abi.Method,
	args []interface{},
	value *big.Int,
	readOnly bool,
) ([]byte, error) {
	if readOnly {
		return nil, vm.ErrExecutionReverted
	}

	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	// Parse arguments
	toAddr := args[0].(common.Address)
	amount := args[1].(*big.Int)

	// Convert caller to cosmos address
	senderCosmoAddr := utils.EthToCosmosAddr(caller)
	receiverCosmoAddr := utils.EthToCosmosAddr(toAddr)

	// Perform the bank operation
	coins := sdk.NewCoins(sdk.NewCoin(pcommon.BaseDenom, sdkmath.NewIntFromBigInt(amount)))
	err := p.bankKeeper.SendCoins(ctx, senderCosmoAddr, receiverCosmoAddr, coins)
	if err != nil {
		return nil, err
	}

	// CRITICAL: Track balance changes when called from smart contract
	// This prevents the EVM stateDB from overwriting the bank keeper changes
	if pcommon.ShouldTrackFromContract(caller, senderCosmoAddr) {
		tracker := pcommon.NewBalanceTracker(p.precompile)

		// Track the transfer operation
		tracker.TrackTransfer(caller, toAddr, amount, pcommon.BaseDenom)

		// Alternative approach - track individual balance changes:
		// tracker.TrackBalanceChanges(
		//     pcommon.NewBalanceChangeEntry(caller, amount, pcommon.Sub),
		//     pcommon.NewBalanceChangeEntry(toAddr, amount, pcommon.Add),
		// )
	}

	return method.Outputs.Pack(true)
}

/*
USAGE EXAMPLE:

1. Deploy a smart contract that calls this precompile:

```solidity
contract ExampleCaller {
    address constant EXAMPLE_PRECOMPILE = 0x0000000000000000000000000000000000001999;

    function callTransferTokens(address to, uint256 amount) external {
        (bool success, ) = EXAMPLE_PRECOMPILE.call(
            abi.encodeWithSignature("transferTokens(address,uint256)", to, amount)
        );
        require(success, "Transfer failed");
    }
}
```

2. Without balance tracking:
   - Smart contract calls precompile
   - Precompile modifies bank balances via Cosmos SDK
   - EVM stateDB commits and overwrites the bank changes
   - Balances become inconsistent

3. With balance tracking:
   - Smart contract calls precompile
   - Precompile modifies bank balances via Cosmos SDK
   - Precompile records balance changes using SetBalanceChangeEntries
   - EVM applies the recorded changes to stateDB before committing
   - Balances remain consistent between bank keeper and EVM state

KEY POINTS:
- Only track balance changes when called from smart contracts (caller != original sender)
- Only track the base denomination ("usix") to prevent conflicts
- Use the helper functions for cleaner code
- Always handle errors when getting withdrawal addresses or other dependencies
*/
