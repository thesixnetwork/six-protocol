package nftmngr

import (
	"bytes"
	"embed"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	pcommon "github.com/thesixnetwork/six-protocol/v4/precompiles/common"
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
	accountKeeper pcommon.AccountKeeper
	bankKeeper    pcommon.BankKeeper
	address       common.Address

	/*
	   #################
	   #### GETTER #####
	   #################
	*/
	GetActionExecutorID []byte
	IsActionExecutorID  []byte
	IsSchemaOwnerID     []byte
	GetAttributeValueID []byte

	/*
	   #################
	   #### SETTER #####
	   #################
	*/
	AddActionID             []byte
	AddAttributeID          []byte
	ChangeOrgOwnerID        []byte
	ChangeSchemaOwnerID     []byte
	CreateMetadataID        []byte
	CreateSchemaID          []byte
	ResyncAttributeID       []byte
	UpdateAttributeID       []byte
	AttributeOverideID      []byte
	SetBaseURIID            []byte
	SetMetadataFormatID     []byte
	SetMintAuthID           []byte
	SetOriginChainID        []byte
	SetOriginContractID     []byte
	SetUriRetreivalID       []byte
	ShowAttributeID         []byte
	ToggleActionID          []byte
	UpateActionID           []byte
	AddActionExecutorID     []byte
	RemoveActionExecutorID  []byte
	ActionByAdminID         []byte
	VirtualSchemaProposalId []byte
	VoteVirtualId           []byte
	PerformVirtualActionId  []byte
}

func NewExecutor(nftmngrKeeper pcommon.NftmngrKeeper, accountKeeper pcommon.AccountKeeper, bankKeeper pcommon.BankKeeper) *PrecompileExecutor {
	return &PrecompileExecutor{
		nftmngrKeeper: nftmngrKeeper,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		address:       common.HexToAddress(NftmngrAddress),
	}
}

func NewPrecompile(nftmngrKeeper pcommon.NftmngrKeeper, accountKeeper pcommon.AccountKeeper, bankKeeper pcommon.BankKeeper) (*pcommon.Precompile, error) {
	newAbi := GetABI()
	p := NewExecutor(nftmngrKeeper, accountKeeper, bankKeeper)

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
		case IsActionExecutor:
			p.IsActionExecutorID = m.ID
		case IsSchemaOwner:
			p.IsSchemaOwnerID = m.ID
		case GetAttributeValue:
			p.GetAttributeValueID = m.ID
		case VirtualSchemaProposal:
			p.VirtualSchemaProposalId = m.ID
		case VoteVirtualSchema:
			p.VoteVirtualId = m.ID
		case PerformVirtualAction:
			p.PerformVirtualActionId = m.ID
		}
	}

	return pcommon.NewPrecompile(newAbi, p, p.address, "nftmngr"), nil
}

// Address implements common.PrecompileExecutor.
func (p *PrecompileExecutor) Address() common.Address {
	return p.address
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	//if method.Name == "voteVirtualSchema" {
	//  return 5000
	//}
	return pcommon.DefaultGasCost(input, p.IsTransaction(method.Name))
}

func (PrecompileExecutor) IsTransaction(method string) bool {
	_, ok := transactionMethods[method]
	return ok
}

func (p PrecompileExecutor) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("precompile", "nftmngr")
}
