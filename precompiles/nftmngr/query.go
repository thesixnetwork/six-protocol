package nftmngr

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	pcommon "github.com/thesixnetwork/six-protocol/precompiles/common"
)

func (p PrecompileExecutor) isActionExecutor(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	nftschema, err := p.stringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	executorAddress, err := p.accAddressFromArg(args[1])
	if err != nil {
		return nil, err
	}

	_, found := p.nftmngrKeeper.GetActionExecutor(ctx, nftschema, executorAddress.String())

	return method.Outputs.Pack(found)
}