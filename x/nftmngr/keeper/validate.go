package keeper

import (
	"fmt"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// **** VALIDATION OF NFT METADATA ****

func CreateAttrDefMap(attrDefs []*types.AttributeDefinition) map[string]*types.AttributeDefinition {
	attrDefMap := make(map[string]*types.AttributeDefinition)
	for _, attrDef := range attrDefs {
		attrDefMap[attrDef.Name] = attrDef
	}
	return attrDefMap
}

// create ActionMap
func CreateActionMap(actions []*types.Action) map[string]*types.Action {
	actionMap := make(map[string]*types.Action)
	for _, action := range actions {
		actionMap[action.Name] = action
	}
	return actionMap
}

func CreateNftAttrValueMap(nftAttrValues []*types.NftAttributeValue) map[string]*types.NftAttributeValue {
	nftAttrValueMap := make(map[string]*types.NftAttributeValue)
	for _, nftAttrValue := range nftAttrValues {
		nftAttrValueMap[nftAttrValue.Name] = nftAttrValue
	}
	return nftAttrValueMap
}

func ValidateRequiredAttributes(schemaAttributes []*types.AttributeDefinition, mapAttributeValues map[string]*types.NftAttributeValue) (bool, string) {
	for _, schemaAttribute := range schemaAttributes {
		if schemaAttribute.Required && (schemaAttribute.DefaultMintValue == nil || schemaAttribute.DefaultMintValue.GetValue() == nil) {
			if _, ok := mapAttributeValues[schemaAttribute.Name]; !ok {
				return false, schemaAttribute.Name
			}
		}
	}
	return true, ""
}

func HasDuplicateAttributes(attributes []*types.AttributeDefinition) (bool, string) {
	mapAttributes := map[string]*types.AttributeDefinition{}
	for _, attriDef := range attributes {
		if _, ok := mapAttributes[attriDef.Name]; ok {
			return true, attriDef.Name
		}
		mapAttributes[attriDef.Name] = attriDef
	}
	return false, ""
}

func HasDuplicateNftAttributesValue(attributes []*types.NftAttributeValue) (bool, string) {
	mapAttributes := map[string]*types.NftAttributeValue{}
	for _, attriDef := range attributes {
		if _, ok := mapAttributes[attriDef.Name]; ok {
			return true, attriDef.Name
		}
		mapAttributes[attriDef.Name] = attriDef
	}
	return false, ""
}

func HasDuplicateOnchainAttributes(schemaAttributes []*types.AttributeDefinition, tokenAttributes []*types.AttributeDefinition) (bool, string) {
	mapAttributes := map[string]*types.AttributeDefinition{}
	for _, attriDef := range schemaAttributes {
		if _, ok := mapAttributes[attriDef.Name]; ok {
			return true, attriDef.Name
		}
		mapAttributes[attriDef.Name] = &types.AttributeDefinition{
			Name:                attriDef.Name,
			DataType:            attriDef.DataType,
			Required:            attriDef.Required,
			DisplayValueField:   attriDef.DisplayValueField,
			DisplayOption:       attriDef.DisplayOption,
			DefaultMintValue:    attriDef.DefaultMintValue,
			HiddenOveride:       attriDef.HiddenOveride,
			HiddenToMarketplace: attriDef.HiddenToMarketplace,
		}
	}

	for _, attriDef := range tokenAttributes {
		if _, ok := mapAttributes[attriDef.Name]; ok {
			return true, attriDef.Name
		}
		mapAttributes[attriDef.Name] = attriDef
	}

	return false, ""
}

func HasSameType(mapOriginAttributes map[string]*types.AttributeDefinition, onchainAttributes []*types.AttributeDefinition) (bool, string) {
	for _, attriVal := range onchainAttributes {
		attrDef := mapOriginAttributes[attriVal.Name]
		if attrDef == nil {
			// fmt.Println("Attribute not found: ", attriVal.Name)
			continue
		}
		if attrDef.DataType != attriVal.DataType {
			// fmt.Println("attrDef.DataType: ", attrDef.DataType)
			return false, attrDef.Name
		}
	}
	return true, ""
}

func MergeNFTDataAttributes(originAttributes []*types.AttributeDefinition, onchainAttributes []*types.AttributeDefinition) []*types.AttributeDefinition {
	mergedAttributes := make([]*types.AttributeDefinition, 0)
	for _, originAttribute := range originAttributes {
		mergedAttributes = append(append(mergedAttributes, originAttribute), onchainAttributes...)

		// ** What is is actually doing is:
		// ** mergedAttributes = append(mergedAttributes, originAttribute)
		// ** for _, attr := range onchainAttributes {
		// ** mergedAttributes = append(mergedAttributes, attr)
		// ** }
		// ** It is better performance to use append(append(mergedAttributes, originAttribute), onchainAttributes...) instead of the above
		// ** Just in case I forget, I will leave this comment here
	}
	for _, onchainAttribute := range onchainAttributes {
		mergedAttributes = append(append(mergedAttributes, onchainAttribute), originAttributes...)
	}
	return mergedAttributes
}
func HasSameTypeAsSchema(mapSchemaAttributes map[string]*types.AttributeDefinition, dataAttributes []*types.NftAttributeValue) (bool, string) {
	// If attributes have same name, then they must have same type
	for _, attriVal := range dataAttributes {
		attDef := mapSchemaAttributes[attriVal.Name]
		if attDef.DataType != GetTypeFromAttributeValue(attriVal) {
			return false, attriVal.Name
		}
	}
	return true, ""
}
func GetTypeFromAttributeValue(attribute *types.NftAttributeValue) string {
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_BooleanAttributeValue); ok {
		return "boolean"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_StringAttributeValue); ok {
		return "string"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_NumberAttributeValue); ok {
		return "number"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_FloatAttributeValue); ok {
		return "float"
	}
	return "default"
}

func DefaultMintValueHasSameType(attributes []*types.AttributeDefinition) (bool, string) {
	for _, attriDef := range attributes {
		_, attrType := HasDefaultMintValue(*attriDef)
		if attriDef.DataType != attrType {
			return false, attriDef.Name
		}
	}
	return true, ""
}

func HasDefaultMintValue(attribute types.AttributeDefinition) (bool, string) {
	// Check if onchain attribute s value exist for each attribute
	if _, ok := attribute.DefaultMintValue.GetValue().(*types.DefaultMintValue_BooleanAttributeValue); ok {
		return ok, "boolean"
	}
	if _, ok := attribute.DefaultMintValue.GetValue().(*types.DefaultMintValue_StringAttributeValue); ok {
		return ok, "string"
	}
	if _, ok := attribute.DefaultMintValue.GetValue().(*types.DefaultMintValue_NumberAttributeValue); ok {
		return ok, "number"
	}
	if _, ok := attribute.DefaultMintValue.GetValue().(*types.DefaultMintValue_FloatAttributeValue); ok {
		return ok, "float"
	}
	return false, "default"
}

func ConvertDefaultMintValueToSchemaAttributeValue(defaultMintValue *types.DefaultMintValue) (*types.SchemaAttributeValue, error) {
	schemaAttributeValue := &types.SchemaAttributeValue{}

	switch value := defaultMintValue.Value.(type) {
	case *types.DefaultMintValue_NumberAttributeValue:
		schemaAttributeValue.Value = &types.SchemaAttributeValue_NumberAttributeValue{
			NumberAttributeValue: value.NumberAttributeValue,
		}
	case *types.DefaultMintValue_StringAttributeValue:
		schemaAttributeValue.Value = &types.SchemaAttributeValue_StringAttributeValue{
			StringAttributeValue: value.StringAttributeValue,
		}
	case *types.DefaultMintValue_BooleanAttributeValue:
		schemaAttributeValue.Value = &types.SchemaAttributeValue_BooleanAttributeValue{
			BooleanAttributeValue: value.BooleanAttributeValue,
		}
	case *types.DefaultMintValue_FloatAttributeValue:
		schemaAttributeValue.Value = &types.SchemaAttributeValue_FloatAttributeValue{
			FloatAttributeValue: value.FloatAttributeValue,
		}
	default:
		return nil, fmt.Errorf("unknown value type: %T", value)
	}

	return schemaAttributeValue, nil
}

func ConvertSchemaAttributeValueToDefaultMintValue(schemaAttributeValue *types.SchemaAttributeValue) (*types.DefaultMintValue, error) {
	defaultMintValue := &types.DefaultMintValue{}

	switch value := schemaAttributeValue.Value.(type) {
	case *types.SchemaAttributeValue_NumberAttributeValue:
		defaultMintValue.Value = &types.DefaultMintValue_NumberAttributeValue{
			NumberAttributeValue: value.NumberAttributeValue,
		}
	case *types.SchemaAttributeValue_StringAttributeValue:
		defaultMintValue.Value = &types.DefaultMintValue_StringAttributeValue{
			StringAttributeValue: value.StringAttributeValue,
		}
	case *types.SchemaAttributeValue_BooleanAttributeValue:
		defaultMintValue.Value = &types.DefaultMintValue_BooleanAttributeValue{
			BooleanAttributeValue: value.BooleanAttributeValue,
		}
	case *types.SchemaAttributeValue_FloatAttributeValue:
		defaultMintValue.Value = &types.DefaultMintValue_FloatAttributeValue{
			FloatAttributeValue: value.FloatAttributeValue,
		}
	default:
		return nil, fmt.Errorf("unknown value type: %T", value)
	}

	return defaultMintValue, nil
}

func ConverSchemaAttributeToNFTAttributeValue(schemaAttributes *types.SchemaAttribute) *types.NftAttributeValue {
	nftAttributeValue := &types.NftAttributeValue{}

	switch value := schemaAttributes.CurrentValue.Value.(type) {
	case *types.SchemaAttributeValue_NumberAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_NumberAttributeValue{
			NumberAttributeValue: value.NumberAttributeValue,
		}
	case *types.SchemaAttributeValue_StringAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_StringAttributeValue{
			StringAttributeValue: value.StringAttributeValue,
		}
	case *types.SchemaAttributeValue_BooleanAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_BooleanAttributeValue{
			BooleanAttributeValue: value.BooleanAttributeValue,
		}
	case *types.SchemaAttributeValue_FloatAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_FloatAttributeValue{
			FloatAttributeValue: value.FloatAttributeValue,
		}
	default:
		return nil
	}

	return &types.NftAttributeValue{
		Name:  schemaAttributes.Name,
		Value: nftAttributeValue.Value,
	}
}

func ConverAttributeDefinitionToNFTAttributeValue(attributeDefinition *types.AttributeDefinition) *types.NftAttributeValue {
	nftAttributeValue := &types.NftAttributeValue{}

	switch value := attributeDefinition.DefaultMintValue.Value.(type) {
	case *types.DefaultMintValue_NumberAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_NumberAttributeValue{
			NumberAttributeValue: value.NumberAttributeValue,
		}
	case *types.DefaultMintValue_StringAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_StringAttributeValue{
			StringAttributeValue: value.StringAttributeValue,
		}
	case *types.DefaultMintValue_BooleanAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_BooleanAttributeValue{
			BooleanAttributeValue: value.BooleanAttributeValue,
		}
	case *types.DefaultMintValue_FloatAttributeValue:
		nftAttributeValue.Value = &types.NftAttributeValue_FloatAttributeValue{
			FloatAttributeValue: value.FloatAttributeValue,
		}
	default:
		return nil
	}

	return &types.NftAttributeValue{
		Name:                attributeDefinition.Name,
		Value:               nftAttributeValue.Value,
		HiddenToMarketplace: attributeDefinition.HiddenToMarketplace,
	}
}

// Check if NFT data attributes exists in schema
func NFTDataAttributesExistInSchema(mapAttributes map[string]*types.AttributeDefinition, dataAttributes []*types.NftAttributeValue) (bool, string) {
	// Check if dataAttributes exist in schema
	for _, attriVal := range dataAttributes {
		if _, ok := mapAttributes[attriVal.Name]; !ok {
			return false, attriVal.Name
		}
	}
	return true, ""
}
