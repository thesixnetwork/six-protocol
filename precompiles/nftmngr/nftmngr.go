package nftmngr

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/utils"
	"github.com/tendermint/tendermint/libs/log"
	pcommon "github.com/thesixnetwork/six-protocol/precompiles/common"
	nftmngrtype "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

const (
	NftmngrAddress = "0x0000000000000000000000000000000000001055"
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

func GetABI() abi.ABI {
	abiBz, err := f.ReadFile("abi.json")
	if err != nil {
		panic(err)
	}

	newAbi, err := abi.JSON(bytes.NewReader(abiBz))
	if err != nil {
		panic(err)
	}
	return newAbi
}

type ActionParameter struct {
	Name  string
	Value string
}

type PrecompileExecutor struct {
	nftmngrKeeper pcommon.NftmngrKeeper
	bankKeeper    pcommon.BankKeeper
	/*
	   #################
	   #### GETTER #####
	   #################
	*/
	GetActionExecutorId []byte
	/*
	   #################
	   #### SETTER #####
	   #################
	*/
	AddActionID            []byte
	AddAttributeID         []byte
	ChangeOrgOwnerID       []byte
	ChangeSchemaOwnerID    []byte
	CreateMetadataID       []byte
	CreateSchemaID         []byte
	ResyncAttributeID      []byte
	UpdateAttributeID      []byte
	AttributeOverideID     []byte
	SetBaseURIID           []byte
	SetMetadataFormatID    []byte
	SetMintAuthID          []byte
	SetOriginChainID       []byte
	SetOriginContractID    []byte
	SetUriRetreivalID      []byte
	ShowAttributeID        []byte
	ToggleActionID         []byte
	UpateActionID          []byte
	AddActionExecutorID    []byte
	RemoveActionExecutorID []byte
	ActionByAdminID        []byte
	address                common.Address
}

func NewExecutor(nftmngrKeeper pcommon.NftmngrKeeper, bankKeeper pcommon.BankKeeper) (*PrecompileExecutor, error) {
	p := &PrecompileExecutor{
		nftmngrKeeper: nftmngrKeeper,
		bankKeeper:    bankKeeper,
		address:       common.HexToAddress(NftmngrAddress),
	}

	return p, nil
}

func NewPrecompile(nftmngrKeeper pcommon.NftmngrKeeper, bankKeeper pcommon.BankKeeper) (*pcommon.Precompile, error) {
	newAbi := GetABI()
	p := &PrecompileExecutor{
		nftmngrKeeper: nftmngrKeeper,
		bankKeeper:    bankKeeper,
		address:       common.HexToAddress(NftmngrAddress),
	}

	for name, m := range newAbi.Methods {
		switch name {
		case ActionByAdmin:
			p.ActionByAdminID = m.ID
		case AddAction:
			p.AddActionID = m.ID
		case AddAttribute:
			p.AddAttributeID = m.ID
		case ChangeOrgOwner:
			p.ChangeOrgOwnerID = m.ID
		case ChangeSchemaOwner:
			p.ChangeSchemaOwnerID = m.ID
		case CreateMetadata:
			p.CreateMetadataID = m.ID
		case CreateSchema:
			p.CreateSchemaID = m.ID
		case ResyncAttribute:
			p.ResyncAttributeID = m.ID
		case UpdateAttribute:
			p.UpdateAttributeID = m.ID
		case AttributeOveride:
			p.AttributeOverideID = m.ID
		case SetBaseURI:
			p.SetBaseURIID = m.ID
		case SetMetadataFormat:
			p.SetMetadataFormatID = m.ID
		case SetMintAuth:
			p.SetMintAuthID = m.ID
		case SetOriginChain:
			p.SetOriginChainID = m.ID
		case SetOriginContract:
			p.SetOriginContractID = m.ID
		case SetUriRetreival:
			p.SetUriRetreivalID = m.ID
		case ShowAttribute:
			p.ShowAttributeID = m.ID
		case ToggleAction:
			p.ToggleActionID = m.ID
		case UpdateAction:
			p.UpateActionID = m.ID
		case AddActionExecutor:
			p.AddActionExecutorID = m.ID
		case RemoveActionExecutor:
			p.RemoveActionExecutorID = m.ID
		}
	}

	return pcommon.NewPrecompile(newAbi, p, p.address, "nftmngr"), nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	return pcommon.DefaultGasCost(input, p.IsTransaction(method.Name))
}

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

func (p PrecompileExecutor) AccAddressFromBech32(arg interface{}) (bec32Addr sdk.AccAddress, err error) {
	addr := arg.(string)
	bec32Addr, err = sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid bech32 address")
	}
	return bec32Addr, nil
}

func (p PrecompileExecutor) AccAddressFromArg(arg interface{}) (sdk.AccAddress, error) {
	addr := arg.(common.Address)
	if addr == (common.Address{}) {
		return nil, errors.New("invalid addr")
	}
	bec32Addr := utils.EthToCosmosAddr(addr)
	return bec32Addr, nil
}

func (p PrecompileExecutor) StringFromArg(arg interface{}) (string, error) {
	stringArg, ok := arg.(string)
	if !ok {
		return "", errors.New("invalid argument type string")
	}
	return stringArg, nil
}

func (p PrecompileExecutor) ArrayOfstringFromArg(arg interface{}) ([]string, error) {
	arrayStringArg, ok := arg.([]string)
	if !ok {
		return nil, errors.New("invalid argument type string")
	}
	return arrayStringArg, nil
}

func (p PrecompileExecutor) boolFromArg(arg interface{}) (bool, error) {
	boolArg, ok := arg.(bool)
	if !ok {
		return false, errors.New("invalid argument type string")
	}

	return boolArg, nil
}

func (p PrecompileExecutor) Uint64FromArg(arg interface{}) (uint64, error) {
	uint64Arg, ok := arg.(uint64)
	if !ok {
		return 0, errors.New("invalid argument type string")
	}

	return uint64Arg, nil
}

func (p PrecompileExecutor) Uint32FromArg(arg interface{}) (uint32, error) {
	uint32Arg, ok := arg.(uint32)
	if !ok {
		return 0, errors.New("invalid argument type string")
	}

	return uint32Arg, nil
}

func (p PrecompileExecutor) ParametersFromJSONArg(arg interface{}) ([]*nftmngrtype.ActionParameter, error) {
	jsonStr, ok := arg.(string)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid argument type, expected string")
	}

	var params []nftmngrtype.ActionParameter
	if err := json.Unmarshal([]byte(jsonStr), &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid JSON format")
	}

	// Convert to slice of pointers to ActionParameter
	paramPointers := make([]*nftmngrtype.ActionParameter, len(params))
	for i := range params {
		paramPointers[i] = &params[i]
	}

	return paramPointers, nil
}

type TransactionMetadata struct {
	// RequiresAuth  bool
	// ModifiesState bool
	// MinGas        uint64
	Description string
}

var transactionMethods = map[string]TransactionMetadata{
	ActionByAdmin: {
		Description: "Perform action",
	},
	AddAction: {
		Description: "Add new action to schema",
	},
	AddAttribute: {
		Description: "Add new attribute to schema",
	},
	ChangeOrgOwner: {
		Description: "Change organization owner",
	},
	ChangeSchemaOwner: {
		Description: "Change schema owner",
	},
	CreateMetadata: {
		Description: "Create new metadata",
	},
	CreateSchema: {
		Description: "Create new NFT schema",
	},
	ResyncAttribute: {
		Description: "Resynchronize attribute",
	},
	UpdateAttribute: {
		Description: "Update existing attribute",
	},
	AttributeOveride: {
		Description: "Override attribute properties",
	},
	SetBaseURI: {
		Description: "Set base URI for NFTs",
	},
	SetMetadataFormat: {
		Description: "Set metadata format",
	},
	SetMintAuth: {
		Description: "Set minting authorization",
	},
	SetOriginChain: {
		Description: "Set origin chain",
	},
	SetOriginContract: {
		Description: "Set origin contract",
	},
	SetUriRetreival: {
		Description: "Set URI retrieval method",
	},
	ShowAttribute: {
		Description: "Show attribute details",
	},
	ToggleAction: {
		Description: "Toggle action state",
	},
	UpdateAction: {
		Description: "Update existing action",
	},
	AddActionExecutor: {
		Description: "Add new action executor",
	},
	RemoveActionExecutor: {
		Description: "Remove action executor",
	},
}

func (PrecompileExecutor) IsTransaction(method string) bool {
	_, ok := transactionMethods[method]
	return ok
}

func (p PrecompileExecutor) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("precompile", "nftmngr")
}
