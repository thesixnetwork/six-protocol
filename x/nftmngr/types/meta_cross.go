package types

import (
	"regexp"
	"strconv"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type CrossSchemaMetadata struct {
	nftDatas     map[string]*NftData
	nftSchemas   map[string]*NFTSchema
	changeList   map[string]ChangeList
	mapSchemaKey map[string]MapAllKey
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
		mapSchemaKey: make(map[string]MapAllKey),
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

func (c *CrossSchemaMetadata) GetNftData(schemaCode string) *NftData {
	if err := c.validateSchemaName(schemaCode); err != nil {
		panic(err)
	}
	return c.nftDatas[schemaCode]
}

func (c *CrossSchemaMetadata) GetTokenURI(schemaCode string) string {
	return c.nftDatas[schemaCode].TokenUri
}

func (c *CrossSchemaMetadata) SetTokenURI(schemaCode, uri string) {
	if err := c.validateSchemaName(schemaCode); err != nil {
		panic(err)
	}
	c.changeList[schemaCode] = append(c.changeList[schemaCode], &MetadataChange{
		Key:           "tokenURI",
		PreviousValue: c.nftDatas[schemaCode].TokenUri,
		NewValue:      uri,
	})
	c.nftDatas[schemaCode].TokenUri = uri
}

func (c *CrossSchemaMetadata) GetImage(schemaCode string) string {
	if c.nftDatas[schemaCode].OnchainImage != "" {
		return c.nftDatas[schemaCode].OnchainImage
	}

	return c.nftDatas[schemaCode].OriginImage
}

func (c *CrossSchemaMetadata) SetImage(schemaCode, imagePath string) {
	currentImage := c.nftDatas[schemaCode].OnchainImage
	if currentImage == "" {
		currentImage = c.nftDatas[schemaCode].OriginImage
	}
	c.changeList[schemaCode] = append(c.changeList[schemaCode], &MetadataChange{
		Key:           "image",
		PreviousValue: currentImage,
		NewValue:      imagePath,
	})
	c.nftDatas[schemaCode].OnchainImage = imagePath
}

func (c *CrossSchemaMetadata) GetTokenId(schemaCode string) string {
	if err := c.validateSchemaName(schemaCode); err != nil {
		panic(err)
	}
	return c.nftDatas[schemaCode].TokenId
}

func (c *CrossSchemaMetadata) GetNftSchema(schemaCode string) *NFTSchema {
	if err := c.validateSchemaName(schemaCode); err != nil {
		panic(err)
	}
	return c.nftSchemas[schemaCode]
}

func (c *CrossSchemaMetadata) GetChangeList(schemaCode string) ChangeList {
	return c.changeList[schemaCode]
}

func (c *CrossSchemaMetadata) validateSchemaName(schemaCode string) error {
	if _, exists := c.nftSchemas[schemaCode]; !exists {
		return sdkerrors.Wrap(ErrSchemaNotFound, schemaCode)
	}

	return nil
}

func (c *CrossSchemaMetadata) getAttribute(schemaCode, key string) (*MetadataAttribute, error) {
	// Validate schema exists
	if err := c.validateSchemaName(schemaCode); err != nil {
		return nil, sdkerrors.Wrapf(err, "schema validation failed for %s", schemaCode)
	}

	schemaKey, exists := c.mapSchemaKey[schemaCode]
	if !exists {
		return nil, sdkerrors.Wrapf(ErrSchemaNotFound, "schema %s not found in mapSchemaKey", schemaCode)
	}

	attri, exists := schemaKey[key]
	if !exists {
		return nil, sdkerrors.Wrapf(ErrAttributeNotFoundForAction, "attribute %s not found in schema %s", key, schemaCode)
	}

	if attri == nil {
		return nil, sdkerrors.Wrapf(ErrInvalidOperation, "attribute %s is nil in schema %s", key, schemaCode)
	}

	return attri, nil
}

func (c *CrossSchemaMetadata) GetNumber(schemaCode string, key string) int64 {
	v, err := c.MustGetNumber(schemaCode, key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *CrossSchemaMetadata) MustGetNumber(schemaCode, key string) (int64, error) {
	attri, err := c.getAttribute(schemaCode, key)
	if err != nil {
		return 0, err
	}
	if attri == nil {
		return 0, sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue); ok {
		return int64(attri.AttributeValue.GetNumberAttributeValue().Value), nil
	}
	return 0, sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
}

func (c *CrossSchemaMetadata) GetString(schemaCode, key string) string {
	v, err := c.MustGetString(schemaCode, key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *CrossSchemaMetadata) GetSubString(schemaCode, key string, start int64, end int64) string {
	v, err := c.MustGetString(schemaCode, key)
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

func (c *CrossSchemaMetadata) SetString(schemaCode, key, value string) error {
	attri, err := c.getAttribute(schemaCode, key)
	if err != nil {
		return err
	}
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
			c.changeList[schemaCode] = append(c.changeList[schemaCode], &MetadataChange{
				Key:           key,
				PreviousValue: attri.AttributeValue.GetStringAttributeValue().Value,
				NewValue:      value,
			})
			c.mapSchemaKey[schemaCode][key].AttributeValue = newAttributeValue
			c.nftDatas[schemaCode].OnchainAttributes[attri.Index] = newAttributeValue
		case "schema":
			c.changeList[schemaCode] = append(c.changeList[schemaCode], &MetadataChange{
				Key:           key,
				PreviousValue: attri.AttributeValue.GetStringAttributeValue().Value,
				NewValue:      value,
			})
			c.mapSchemaKey[schemaCode][key].AttributeValue = newAttributeValue
		default:
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}
	return nil
}

func (c *CrossSchemaMetadata) SetNumber(schemaCode, key string, value int64) error {
	attri, err := c.getAttribute(schemaCode, key)
	if err != nil {
		return err
	}
	if attri == nil {
		return sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
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
			c.changeList[schemaCode] = append(c.changeList[schemaCode], &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatUint(attri.AttributeValue.GetNumberAttributeValue().Value, 10),
				NewValue:      strconv.FormatUint(uint64(value), 10),
			})
			c.mapSchemaKey[schemaCode][key].AttributeValue = newAttributeValue
			c.nftDatas[schemaCode].OnchainAttributes[attri.Index] = newAttributeValue
		case "schema":
			c.changeList[schemaCode] = append(c.changeList[schemaCode], &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatUint(attri.AttributeValue.GetNumberAttributeValue().Value, 10),
				NewValue:      strconv.FormatUint(uint64(value), 10),
			})
			c.mapSchemaKey[schemaCode][key].AttributeValue = newAttributeValue
		default:
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}
	return nil
}

func (c *CrossSchemaMetadata) SetFloat(schemaCode, key string, value float64) error {
	attri, err := c.getAttribute(schemaCode, key)
	if err != nil {
		return err
	}
	if attri == nil {
		return sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
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
			c.changeList[schemaCode] = append(c.changeList[schemaCode], &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatFloat(attri.AttributeValue.GetFloatAttributeValue().Value, 'f', -1, 64),
				NewValue:      strconv.FormatFloat(value, 'f', -1, 64),
			})
			c.mapSchemaKey[schemaCode][key].AttributeValue = newAttributeValue
			c.nftDatas[schemaCode].OnchainAttributes[attri.Index] = newAttributeValue
		case "schema":
			c.changeList[schemaCode] = append(c.changeList[schemaCode], &MetadataChange{
				Key:           key,
				PreviousValue: strconv.FormatFloat(attri.AttributeValue.GetFloatAttributeValue().Value, 'f', -1, 64),
				NewValue:      strconv.FormatFloat(value, 'f', -1, 64),
			})
			c.mapSchemaKey[schemaCode][key].AttributeValue = newAttributeValue
		default:
			return sdkerrors.Wrap(ErrAttributeOverriding, "can not override the origin attribute")
		}
	} else {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}
	return nil
}

func (c *CrossSchemaMetadata) ToLowercase(schemaCode, key string) string {
	v, err := c.MustGetString(schemaCode, key)
	if err != nil {
		panic(err)
	}
	return strings.ToLower(v)
}

func (c *CrossSchemaMetadata) ToUppercase(schemaCode, key string) string {
	v, err := c.MustGetString(schemaCode, key)
	if err != nil {
		panic(err)
	}
	return strings.ToUpper(v)
}

func (c *CrossSchemaMetadata) MustGetString(schemaCode, key string) (string, error) {
	attri, err := c.getAttribute(schemaCode, key)
	if err != nil {
		return "", err
	}
	if attri == nil {
		return "", sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_StringAttributeValue); ok {
		return attri.AttributeValue.GetStringAttributeValue().Value, nil
	}
	return "", sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
}

func (c *CrossSchemaMetadata) GetFloat(schemaCode, key string) float64 {
	v, err := c.MustGetFloat(schemaCode, key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *CrossSchemaMetadata) MustGetFloat(schemaCode, key string) (float64, error) {
	attri, err := c.getAttribute(schemaCode, key)
	if err != nil {
		return 0, err
	}
	if attri == nil {
		return 0, sdkerrors.Wrap(ErrAttributeNotFoundForAction, key)
	}
	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_FloatAttributeValue); ok {
		return attri.AttributeValue.GetFloatAttributeValue().Value, nil
	}
	return 0, sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
}

func (c *CrossSchemaMetadata) GetBoolean(schemaCode, key string) bool {
	v, err := c.MustGetBool(schemaCode, key)
	if err != nil {
		panic(err)
	}
	return v
}

func (c *CrossSchemaMetadata) MustGetBool(schemaCode, key string) (bool, error) {
	attri, err := c.getAttribute(schemaCode, key)
	if err != nil {
		return false, err
	}
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
