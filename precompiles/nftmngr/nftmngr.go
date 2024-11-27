package nftmngr

import (
	"bytes"
	"embed"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tendermint/tendermint/libs/log"
	pcommon "github.com/thesixnetwork/six-protocol/precompiles/common"
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

func (PrecompileExecutor) IsTransaction(method string) bool {
	_, ok := transactionMethods[method]
	return ok
}

func (p PrecompileExecutor) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("precompile", "nftmngr")
}
