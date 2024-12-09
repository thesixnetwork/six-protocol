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

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	executorAddress, err := p.AccAddressFromArg(args[1])
	if err != nil {
		return nil, err
	}

	_, found := p.nftmngrKeeper.GetActionExecutor(ctx, nftschema, executorAddress.String())

	return method.Outputs.Pack(found)
}

func (p PrecompileExecutor) isSchemaOwner(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	ownerAddress, err := p.AccAddressFromArg(args[1])
	if err != nil {
		return nil, err
	}

	found, err := p.nftmngrKeeper.IsSchemaOwner(ctx, nftschema, ownerAddress.String())
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(found)
}

func (p PrecompileExecutor) getAttributeValue(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 3); err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	tokenId, err := p.StringFromArg(args[1])
	if err != nil {
		return nil, err
	}

	attributeName, err := p.StringFromArg(args[2])
	if err != nil {
		return nil, err
	}

	attributeValue, err := p.nftmngrKeeper.GetAttributeValue(ctx, nftschema, tokenId, attributeName)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(attributeValue)
}
