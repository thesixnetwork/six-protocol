package nftmngr

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM) (bz []byte, err error) {
	switch method.Name {
	case ActionByAdmin:
		return p.actionByAdmin(ctx, caller, method, args, value, readOnly)
	case AddAction:
		return p.addAction(ctx, caller, method, args, value, readOnly)
	case AddAttribute:
		return p.addAttribute(ctx, caller, method, args, value, readOnly)
	case ChangeOrgOwner:
		return p.changeOrgOwner(ctx, caller, method, args, value, readOnly)
	case ChangeSchemaOwner:
		return p.changeSchemaOwner(ctx, caller, method, args, value, readOnly)
	case CreateMetadata:
		return p.createMetadata(ctx, caller, method, args, value, readOnly)
	case CreateSchema:
		return p.createSchema(ctx, caller, method, args, value, readOnly)
	case ResyncAttribute:
		return p.resyncAttribute(ctx, caller, method, args, value, readOnly)
	case UpdateAttribute:
		return p.updateAttribute(ctx, caller, method, args, value, readOnly)
	case AttributeOveride:
		return p.attributeOveride(ctx, caller, method, args, value, readOnly)
	case SetBaseURI:
		return p.setBaseURI(ctx, caller, method, args, value, readOnly)
	case SetMetadataFormat:
		return p.setMetadataFormat(ctx, caller, method, args, value, readOnly)
	case SetMintAuth:
		return p.setMintAuth(ctx, caller, method, args, value, readOnly)
	case SetOriginChain:
		return p.setOriginChain(ctx, caller, method, args, value, readOnly)
	case SetOriginContract:
		return p.setOriginContract(ctx, caller, method, args, value, readOnly)
	case SetUriRetreival:
		return p.setUriRetreival(ctx, caller, method, args, value, readOnly)
	case ShowAttribute:
		return p.showAttribute(ctx, caller, method, args, value, readOnly)
	case ToggleAction:
		return p.toggleAction(ctx, caller, method, args, value, readOnly)
	case UpdateAction:
		return p.updateAction(ctx, caller, method, args, value, readOnly)
	case AddActionExecutor:
		return p.addActionExecutor(ctx, caller, method, args, value, readOnly)
	case RemoveActionExecutor:
		return p.removeActionExecutor(ctx, caller, method, args, value, readOnly)
	case IsActionExecutor:
		return p.isActionExecutor(ctx, method, args, value)
	}
	return
}
