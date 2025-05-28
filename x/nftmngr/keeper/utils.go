package keeper

import (
	"bytes"
	"encoding/json"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	sdkerrors "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// validate Action data
func ValidateAction(action *types.Action, schema *types.NFTSchema) error {
	// Onchain Data Nft Actions Map
	mapNftActions := CreateActionMap(schema.OnchainData.Actions)
	// check if action name is unique
	if _, found := mapNftActions[action.Name]; found {
		return sdkerrors.Wrap(types.ErrActionAlreadyExists, action.Name)
	}

	// validate action struct
	if action.Name == "" || action.Name == " " {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action name is empty")
	}
	if action.Desc == "" || action.Desc == " " {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action description is empty")
	}
	if action.When == "" || action.When == " " {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action type is empty")
	}
	// validate array of action.Then is not empty
	if len(action.Then) == 0 {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action.Then is empty")
	}
	return nil
}

func ValidateVirutualAction(action *types.Action) error {
	// validate action struct
	if action.Name == "" || action.Name == " " {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action name is empty")
	}
	if action.Desc == "" || action.Desc == " " {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action description is empty")
	}
	if action.When == "" || action.When == " " {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action type is empty")
	}
	// validate array of action.Then is not empty
	if len(action.Then) == 0 {
		return sdkerrors.Wrap(types.ErrInvalidActionAttribute, "action.Then is empty")
	}
	return nil
}

func NewNFTAttributeValueFromDefaultValue(name string, defaultValue *types.DefaultMintValue) *types.NftAttributeValue {
	nftAttributeValue := &types.NftAttributeValue{
		Name: name,
	}
	switch defaultValue.Value.(type) {
	case *types.DefaultMintValue_NumberAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_NumberAttributeValue{NumberAttributeValue: defaultValue.GetNumberAttributeValue()}
	case *types.DefaultMintValue_StringAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_StringAttributeValue{StringAttributeValue: defaultValue.GetStringAttributeValue()}
	case *types.DefaultMintValue_BooleanAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_BooleanAttributeValue{BooleanAttributeValue: defaultValue.GetBooleanAttributeValue()}
	case *types.DefaultMintValue_FloatAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_FloatAttributeValue{FloatAttributeValue: defaultValue.GetFloatAttributeValue()}
	default:
		return nil
	}
	return nftAttributeValue
}

// validate AttributeDefinition data
func (k Keeper) ValidateAttributeDefinition(ctx sdk.Context, attribute *types.AttributeDefinition, schema *types.NFTSchema) error {
	valFound, found := k.GetSchemaAttribute(ctx, schema.Code, attribute.Name)
	if found {
		return sdkerrors.Wrap(types.ErrAttributeAlreadyExists, valFound.Name)
	}
	// Onchain Data Nft Attributes Map
	mapOriginAttributes := CreateAttrDefMap(schema.OriginData.OriginAttributes)
	mapTokenAttributes := CreateAttrDefMap(schema.OnchainData.TokenAttributes)
	// check if attribute name is unique
	if _, found := mapOriginAttributes[attribute.Name]; found {
		return sdkerrors.Wrap(types.ErrAttributeAlreadyExists, attribute.Name)
	}
	if _, found := mapTokenAttributes[attribute.Name]; found {
		return sdkerrors.Wrap(types.ErrAttributeAlreadyExists, attribute.Name)
	}

	return nil
}

// merge all attributes and count the index
func MergeAndCountAllAttributes(originAttributes []*types.AttributeDefinition, onchainNFTAttributes []*types.AttributeDefinition, onchainTokenAttribute []*types.AttributeDefinition) int {
	// length or originAttributes
	length_originAttributes := len(originAttributes)
	// length or onchainNFTAttributes
	length_onchainNFTAttributes := len(onchainNFTAttributes)
	// length or onchainTokenAttribute
	length_onchainTokenAttribute := len(onchainTokenAttribute)

	// length of all attributes
	length_allAttributes := length_originAttributes + length_onchainNFTAttributes + length_onchainTokenAttribute
	return length_allAttributes
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// stringfy tokenId list to string token_id1,token_id2,token_id3
func StringfyTokenIdList(list []string) string {
	var result string
	for i, item := range list {
		if i == 0 {
			result = item
		} else {
			result = result + "," + item
		}
	}
	return result
}
