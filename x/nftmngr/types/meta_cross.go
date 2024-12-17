package types

import (
	"regexp"
	"strconv"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type CrossSchemaMetadata struct {
	nftDatas                     map[string]*NftData
	nftSchemas                   map[string]*NFTSchema
	changeList                   map[string]ChangeList
	mapSchemaKey                 map[string]MapAllKey
	OtherSchemaUpdatedTokenDatas map[string]*OtherUpdatedTokenDatas
	CrossSchemaNFTFunction       func(schemaName string, tokenId string) (*NFTSchema, error)
}

type (
	CrossSchemaAttributeOverriding map[string]AttributeOverriding
	CrossSchemaGlobalAttributes    map[string][]*NftAttributeValue
)

func NewCrossSchemaMetadata(schemaList []*NFTSchema, tokenList []*NftData, attributesOverriding CrossSchemaAttributeOverriding, schemaGlobalAttriubutes CrossSchemaGlobalAttributes) *CrossSchemaMetadata {
	nftSchemas := make(map[string]*NFTSchema)
	nftDatas := make(map[string]*NftData)
	crossSchemaMetadata := &CrossSchemaMetadata{
		nftDatas:     nftDatas,
		nftSchemas:   nftSchemas,
		changeList:   make(map[string]ChangeList),
		mapSchemaKey: make(map[string]MapAllKey), // Ensure this is initialized
	}

	if len(schemaList) != len(tokenList) {
		return nil
	}

	for i, schema := range schemaList {
		if schema == nil || tokenList[i] == nil {
			return nil
		}

		nftSchemas[schema.Code] = schema
		nftDatas[schema.Code] = tokenList[i]

		// Ensure the inner map is initialized for this schema code
		crossSchemaMetadata.mapSchemaKey[schema.Code] = make(MapAllKey)

		for j, attri := range tokenList[i].OriginAttributes {
			crossSchemaMetadata.mapSchemaKey[schema.Code][attri.Name] = &MetadataAttribute{
				Index:          j,
				AttributeValue: attri,
				From:           "origin",
			}
		}

		for j, attri := range tokenList[i].OnchainAttributes {
			if _, ok := crossSchemaMetadata.mapSchemaKey[schema.Code][attri.Name]; ok {
				if attributesOverriding[schema.Code] == AttributeOverriding_CHAIN {
					crossSchemaMetadata.mapSchemaKey[schema.Code][attri.Name] = &MetadataAttribute{
						AttributeValue: attri,
						From:           "chain",
						Index:          j,
					}
				}
			} else {
				crossSchemaMetadata.mapSchemaKey[schema.Code][attri.Name] = &MetadataAttribute{
					AttributeValue: attri,
					From:           "chain",
					Index:          j,
				}
			}
		}

		for j, attri := range schemaGlobalAttriubutes[schema.Code] {
			crossSchemaMetadata.mapSchemaKey[schema.Code][attri.Name] = &MetadataAttribute{
				AttributeValue: attri,
				From:           "schema",
				Index:          j,
			}
		}
	}

	return crossSchemaMetadata
}

func (c *CrossSchemaMetadata) SetCrossSchemaNFTFunction(f func(schemaName string, tokenId string) (*NFTSchema, error)) {
	c.CrossSchemaNFTFunction = f
}

func (c *CrossSchemaMetadata) GetNftData(schemaName string) *NftData {
	return c.nftDatas[schemaName]
}

func (c *CrossSchemaMetadata) GetNftSchema(schemaName string) *NFTSchema {
	return c.nftSchemas[schemaName]
}

func (c *CrossSchemaMetadata) GetChangeList(schemaName string) ChangeList {
	return c.changeList[schemaName]
}

func (c *CrossSchemaMetadata) GetNumber(schemaName string, key string) int64 {
	v, err := c.MustGetNumber(schemaName, key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *CrossSchemaMetadata) MustGetNumber(schemaName, key string) (int64, error) {
	schemaKey := c.mapSchemaKey[schemaName]
	attri := schemaKey[key]

	if attri == nil {
		return 0, sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue); ok {
		return int64(attri.AttributeValue.GetNumberAttributeValue().Value), nil
	}
	return 0, sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
}

func (c *CrossSchemaMetadata) GetString(schemaName, key string) string {
	v, err := c.MustGetString(schemaName, key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *CrossSchemaMetadata) GetSubString(schemaName, key string, start int64, end int64) string {
	v, err := c.MustGetString(schemaName, key)
	if end > int64(len(v)) {
		panic(sdkerrors.Wrap(ErrInvalidActionInput, "end can not be greater than string length"))
	}
	if start == end {
		return ""
	}
	if start < 0 {
		start = int64(len(v)) + (start + 1)
	}
	if end < 0 {
		end = int64(len(v)) + (end + 1)
	}
	if start > end {
		panic(sdkerrors.Wrap(ErrInvalidActionInput, "start can not be greater than end"))
	}
	if err != nil {
		panic(err)
	}
	return v[start:end]
}

func (c *CrossSchemaMetadata) SetString(schemaName, key, value string) error {
	schemaKey := c.mapSchemaKey[schemaName]
	attri := schemaKey[key]
	if attri == nil {
		panic(sdkerrors.Wrap(ErrAttributeNotFoundForAction, key))
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_StringAttributeValue); ok {
		// String
		newAttributeValue := &NftAttributeValue{
			Name: attri.AttributeValue.Name,
			Value: &NftAttributeValue_StringAttributeValue{
				StringAttributeValue: &StringAttributeValue{
					Value: value,
				},
			},
		}
		switch attri.From {
		case "chain":
			c.changeList[schemaName] = append(c.changeList[schemaName], &MetadataChange{
				Key:           key,
				PreviousValue: attri.AttributeValue.GetStringAttributeValue().Value,
				NewValue:      value,
			})
			c.mapSchemaKey[schemaName][key].AttributeValue = newAttributeValue
			c.nftDatas[schemaName].OnchainAttributes[attri.Index] = newAttributeValue
		case "schema":
			c.changeList[schemaName] = append(c.changeList[schemaName], &MetadataChange{
				Key:           key,
				PreviousValue: attri.AttributeValue.GetStringAttributeValue().Value,
				NewValue:      value,
			})
			c.mapSchemaKey[schemaName][key].AttributeValue = newAttributeValue
		default:
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}
	return nil
}

func (c *CrossSchemaMetadata) SetNumber(schemaName, key string, value int64) error {
	schemaKey := c.mapSchemaKey[schemaName]
	attri := schemaKey[key]
	if attri == nil {
		panic(sdkerrors.Wrap(ErrAttributeNotFoundForAction, key))
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue); ok {
		// Number
		newAttributeValue := &NftAttributeValue{
			Name: attri.AttributeValue.Name,
			Value: &NftAttributeValue_NumberAttributeValue{
				NumberAttributeValue: &NumberAttributeValue{
					Value: uint64(value),
				},
			},
		}
		switch attri.From {
		case "chain":
			c.changeList[schemaName] = append(c.changeList[schemaName], &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatUint(attri.AttributeValue.GetNumberAttributeValue().Value, 10),
				NewValue:      strconv.FormatUint(uint64(value), 10),
			})
			c.mapSchemaKey[schemaName][key].AttributeValue = newAttributeValue
			c.nftDatas[schemaName].OnchainAttributes[attri.Index] = newAttributeValue
		case "schema":
			c.changeList[schemaName] = append(c.changeList[schemaName], &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatUint(attri.AttributeValue.GetNumberAttributeValue().Value, 10),
				NewValue:      strconv.FormatUint(uint64(value), 10),
			})
			c.mapSchemaKey[schemaName][key].AttributeValue = newAttributeValue
		default:
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}
	return nil
}

func (c *CrossSchemaMetadata) SetFloat(schemaName, key string, value float64) error {
	schemaKey := c.mapSchemaKey[schemaName]
	attri := schemaKey[key]
	if attri == nil {
		panic(sdkerrors.Wrap(ErrAttributeNotFoundForAction, key))
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_FloatAttributeValue); ok {
		// Float
		newAttributeValue := &NftAttributeValue{
			Name: attri.AttributeValue.Name,
			Value: &NftAttributeValue_FloatAttributeValue{
				FloatAttributeValue: &FloatAttributeValue{
					Value: value,
				},
			},
		}
		switch attri.From {
		case "chain":
			c.changeList[schemaName] = append(c.changeList[schemaName], &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatFloat(attri.AttributeValue.GetFloatAttributeValue().Value, 'f', -1, 64),
				NewValue:      strconv.FormatFloat(value, 'f', -1, 64),
			})
			c.mapSchemaKey[schemaName][key].AttributeValue = newAttributeValue
			c.nftDatas[schemaName].OnchainAttributes[attri.Index] = newAttributeValue
		case "schema":
			c.changeList[schemaName] = append(c.changeList[schemaName], &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatFloat(attri.AttributeValue.GetFloatAttributeValue().Value, 'f', -1, 64),
				NewValue:      strconv.FormatFloat(value, 'f', -1, 64),
			})
			c.mapSchemaKey[schemaName][key].AttributeValue = newAttributeValue
		default:
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}
	return nil
}

func (c *CrossSchemaMetadata) ToLowercase(schemaName, key string) string {
	v, err := c.MustGetString(schemaName, key)
	if err != nil {
		panic(err)
	}
	return strings.ToLower(v)
}

func (c *CrossSchemaMetadata) ToUppercase(schemaName, key string) string {
	v, err := c.MustGetString(schemaName, key)
	if err != nil {
		panic(err)
	}
	return strings.ToUpper(v)
}

func (c *CrossSchemaMetadata) MustGetString(schemaName, key string) (string, error) {
	schemaKey := c.mapSchemaKey[schemaName]
	attri := schemaKey[key]
	if attri == nil {
		return "", sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_StringAttributeValue); ok {
		return attri.AttributeValue.GetStringAttributeValue().Value, nil
	}
	return "", sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
}

func (c *CrossSchemaMetadata) GetFloat(schemaName, key string) float64 {
	v, err := c.MustGetFloat(schemaName, key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *CrossSchemaMetadata) MustGetFloat(schemaName, key string) (float64, error) {
	schemaKey := c.mapSchemaKey[schemaName]
	attri := schemaKey[key]
	if attri == nil {
		return 0, sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_FloatAttributeValue); ok {
		return attri.AttributeValue.GetFloatAttributeValue().Value, nil
	}
	return 0, sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
}

func (c *CrossSchemaMetadata) GetBoolean(schemaName, key string) bool {
	v, err := c.MustGetBool(schemaName, key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *CrossSchemaMetadata) MustGetBool(schemaName, key string) (bool, error) {
	schemaKey := c.mapSchemaKey[schemaName]
	attri := schemaKey[key]
	if attri == nil {
		return false, sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_BooleanAttributeValue); ok {
		return attri.AttributeValue.GetBooleanAttributeValue().Value, nil
	}
	return false, sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
}

func (c *CrossSchemaMetadata) ReplaceAllString(input, regexpStr, replaceStr string) string {
	reg := regexp.MustCompile(regexpStr)
	return reg.ReplaceAllString(input, replaceStr)
}