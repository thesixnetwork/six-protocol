package nftmngr

import (
	"encoding/base64"
	"errors"
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	pcommon "github.com/thesixnetwork/six-protocol/precompiles/common"
	nftmngrtypes "github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

const (
	AddAction            = "addAction"
	AddAttribute         = "addAttribute"
	ChangeOrgOwner       = "changeOrgOwner"
	ChangeSchemaOwner    = "changeSchemaOwner"
	CreateMetadata       = "createMetadata"
	CreateSchema         = "createSchema"
	ActionByAdmin        = "actionByAdmin"
	ResyncAttribute      = "resyncAttribute"
	UpdateAttribute      = "updateSchemaAttribute"
	AttributeOveride     = "attributeOveride"
	SetBaseURI           = "setBaseURI"
	SetMetadataFormat    = "setMetadataFormat"
	SetMintAuth          = "setMintAuth"
	SetOriginChain       = "setOriginChain"
	SetOriginContract    = "setOriginContract"
	SetUriRetreival      = "setUriRetreival"
	ShowAttribute        = "showAttribute"
	ToggleAction         = "toggleAction"
	UpdateAction         = "updateAction"
	AddActionExecutor    = "addActionExecutor"
	RemoveActionExecutor = "removeActionExecutor"
	IsActionExecutor     = "isActionExecutor"
)

func (p PrecompileExecutor) addAction(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	base64NewAction, err := p.StringFromArg(args[1])
	if err != nil {
		return nil, err
	}

	// structure for new action
	var new_action nftmngrtypes.Action

	// decode base64 string to bytes
	input_action, err := base64.StdEncoding.DecodeString(base64NewAction)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingBase64, err.Error())
	}

	// unmarshal bytes to Action structure
	err = p.nftmngrKeeper.GetCodec().(*codec.ProtoCodec).UnmarshalJSON(input_action, &new_action)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingMetadataMessage, err.Error())
	}

	err = p.nftmngrKeeper.AddActionKeeper(ctx, senderCosmoAddr.String(), nftschema, new_action)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.Error{}, err.Error())
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) addActionExecutor(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftSchema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	newExecutor, err := p.AccAddressFromArg(args[1])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.AddActionExecutor(ctx, senderCosmoAddr.String(), nftSchema, newExecutor.String())
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.Error{}, err.Error())
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) removeActionExecutor(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftSchema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	newExecutor, err := p.AccAddressFromArg(args[1])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.DelActionExecutor(ctx, senderCosmoAddr.String(), nftSchema, newExecutor.String())
	if err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) addAttribute(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 3); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	locationArg, err := p.Uint32FromArg(args[1])
	if err != nil {
		return nil, err
	}

	newAttribute, err := p.StringFromArg(args[2])
	if err != nil {
		return nil, err
	}

	location := nftmngrtypes.AttributeLocation_NFT_ATTRIBUTE

	if locationArg == 1 {
		location = nftmngrtypes.AttributeLocation_TOKEN_ATTRIBUTE
	}

	var new_add_attribute nftmngrtypes.AttributeDefinition

	input_addribute, err := base64.StdEncoding.DecodeString(newAttribute)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingBase64, err.Error())
	}

	err = p.nftmngrKeeper.GetCodec().(*codec.ProtoCodec).UnmarshalJSON(input_addribute, &new_add_attribute)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingMetadataMessage, err.Error())
	}

	err = p.nftmngrKeeper.AddAttributeKeeper(ctx, senderCosmoAddr.String(), nftschema, new_add_attribute, location)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) changeOrgOwner(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	orgName, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	newOwner, err := p.AccAddressFromArg(args[1])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.ChangeOrgOwner(ctx, senderCosmoAddr.String(), newOwner.String(), orgName)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) changeSchemaOwner(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	orgName, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	newOwner, err := p.AccAddressFromArg(args[1])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.ChangeSchemaOwner(ctx, senderCosmoAddr.String(), newOwner.String(), orgName)
	if err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) createMetadata(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 3); err != nil {
		return nil, err
	}
	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
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

	base64NewMetadata, err := p.StringFromArg(args[2])
	if err != nil {
		return nil, err
	}

	newMetadata, err := base64.StdEncoding.DecodeString(base64NewMetadata)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingBase64, err.Error())
	}

	metadata := nftmngrtypes.NftData{}
	err = p.nftmngrKeeper.GetCodec().(*codec.ProtoCodec).UnmarshalJSON(newMetadata, &metadata)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingMetadataMessage, err.Error())
	}

	err = p.nftmngrKeeper.CreateNewMetadataKeeper(ctx, senderCosmoAddr.String(), nftschema, tokenId, metadata)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) createSchema(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 1); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	base64NewSchema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	jsonSchema, err := base64.StdEncoding.DecodeString(base64NewSchema)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingBase64, err.Error())
	}

	schema_input := nftmngrtypes.NFTSchemaINPUT{}
	err = p.nftmngrKeeper.GetCodec().(*codec.ProtoCodec).UnmarshalJSON(jsonSchema, &schema_input)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingSchemaMessage, err.Error())
	}

	// validate owner has using enough to pay schema fee
	schema_fee, _ := p.nftmngrKeeper.GetNFTFeeConfig(ctx)
	fee_amount, err := sdk.ParseCoinNormalized(schema_fee.SchemaFee.FeeAmount)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrInvalidFeeAmount, err.Error())
	}

	user_current_balance := p.bankKeeper.GetBalance(ctx, senderCosmoAddr, "usix")

	if !user_current_balance.Amount.GTE(fee_amount.Amount) {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrInvalidFeeAmount, "schema fee are not enough")
	}

	err = p.nftmngrKeeper.CreateNftSchemaKeeper(ctx, senderCosmoAddr.String(), schema_input)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) actionByAdmin(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}

	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 5); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
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

	actionName, err := p.StringFromArg(args[2])
	if err != nil {
		return nil, err
	}

	refId, err := p.StringFromArg(args[3])
	if err != nil {
		return nil, err
	}

	paramPointers, err := p.ParametersFromJSONArg(args[4])
	if err != nil {
		return nil, err
	}

	//  ------------------------------------
	// |                                    |
	// |          CORE NFTMODULE            |
	// |                                    |
	//  ------------------------------------

	// paramPointers := make([]*nftmngrtype.ActionParameter, 0)

	_, err = p.nftmngrKeeper.ActionByAdmin(ctx, senderCosmoAddr.String(), nftschema, tokenId, actionName, refId, paramPointers)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) resyncAttribute(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
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

	err = p.nftmngrKeeper.ResyncAttibutesKeeper(ctx, senderCosmoAddr.String(), nftschema, tokenId)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) updateAttribute(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	base64NewAttribute, err := p.StringFromArg(args[1])
	if err != nil {
		return nil, err
	}

	var new_add_attribute nftmngrtypes.AttributeDefinition

	input_addribute, err := base64.StdEncoding.DecodeString(base64NewAttribute)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingBase64, err.Error())
	}

	err = p.nftmngrKeeper.GetCodec().(*codec.ProtoCodec).UnmarshalJSON(input_addribute, &new_add_attribute)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingMetadataMessage, err.Error())
	}

	err = p.nftmngrKeeper.UpdateAttributeKeeper(ctx, senderCosmoAddr.String(), nftschema, new_add_attribute)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) attributeOveride(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	newOveride, err := p.Uint32FromArg(args[1])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.SetAttributeOveridingKeeper(ctx, senderCosmoAddr.String(), nftschema, int32(newOveride))
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) setBaseURI(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	newBaseURI, err := p.StringFromArg(args[1])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.SetBaseURIKeeper(ctx, senderCosmoAddr.String(), nftschema, newBaseURI)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) setMetadataFormat(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	newFormat, err := p.StringFromArg(args[1])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.SetMetadataFormatKeeper(ctx, senderCosmoAddr.String(), nftschema, newFormat)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) setMintAuth(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	authTo, err := p.Uint32FromArg(args[1])
	if err != nil {
		return nil, err
	}

	autorize := nftmngrtypes.AuthorizeTo_SYSTEM

	if authTo == 1 {
		autorize = nftmngrtypes.AuthorizeTo_ALL
	}

	err = p.nftmngrKeeper.SetMintAuthKeeper(ctx, senderCosmoAddr.String(), nftschema, autorize)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) setOriginChain(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	newChain, err := p.StringFromArg(args[1])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.SetOriginChainKeeper(ctx, senderCosmoAddr.String(), nftschema, newChain)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) setOriginContract(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	newContract, err := p.StringFromArg(args[1])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.SetOriginContractKeeper(ctx, senderCosmoAddr.String(), nftschema, newContract)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) setUriRetreival(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	newMethod, err := p.Uint32FromArg(args[1])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.SetURIRetrievalKeeper(ctx, senderCosmoAddr.String(), nftschema, int32(newMethod))
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) showAttribute(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 3); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	status, err := p.boolFromArg(args[1])
	if err != nil {
		return nil, err
	}

	attributeNames, err := p.ArrayOfstringFromArg(args[2])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.ShowAttributeKeeper(ctx, senderCosmoAddr.String(), nftschema, status, attributeNames)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) toggleAction(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 3); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	actionName, err := p.StringFromArg(args[1])
	if err != nil {
		return nil, err
	}

	disable, err := p.boolFromArg(args[2])
	if err != nil {
		return nil, err
	}

	err = p.nftmngrKeeper.ToggleActionKeeper(ctx, senderCosmoAddr.String(), nftschema, actionName, disable)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) updateAction(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.AccAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	nftschema, err := p.StringFromArg(args[0])
	if err != nil {
		return nil, err
	}

	base64NewAction, err := p.StringFromArg(args[1])
	if err != nil {
		return nil, err
	}

	// structure for new action
	var new_action nftmngrtypes.Action

	// decode base64 string to bytes
	input_action, err := base64.StdEncoding.DecodeString(base64NewAction)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingBase64, err.Error())
	}

	// unmarshal bytes to Action structure
	err = p.nftmngrKeeper.GetCodec().(*codec.ProtoCodec).UnmarshalJSON(input_action, &new_action)
	if err != nil {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrParsingMetadataMessage, err.Error())
	}

	err = p.nftmngrKeeper.UpdateActionKeeper(ctx, senderCosmoAddr.String(), nftschema, new_action)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
