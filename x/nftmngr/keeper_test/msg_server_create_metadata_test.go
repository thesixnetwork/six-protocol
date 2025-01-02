package keeper_test

import (
	"fmt"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func TestCreateData(t *testing.T) {
	fmt.Println("Start Test")
	// base64Data := "ewogICAgIm5mdF9zY2hlbWFfY29kZSI6ICJidWFrYXcxIiwKICAgICJ0b2tlbl9pZCI6ICIxIiwKICAgICJ0b2tlbl9vd25lciI6ICI2bmZ0MTlwNXl2d3h0YzJxbmdud2RxbGhmemRjZDBzc25jd3d1bXJ4MmphIiwKICAgICJvcmlnaW5faW1hZ2UiOiAiaHR0cHM6Ly9iazFuZnQuc2l4bmV0d29yay5pby9pcGZzL1FtYUJVZHFUY3RWdGVhanBvTVZFcG81NVhFWGUyWVZ3V0RnZU5pMVVUNDdVS1UvMS5wbmciLAogICAgIm9yaWdpbl9hdHRyaWJ1dGVzIjogWwogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAiYmFja2dyb3VuZF9sIiwKICAgICAgICAgICAgInZhbHVlIjogIkwgQkcgUmluZ3NpZGUgR29sZGVuIgogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJiYWNrZ3JvdW5kX3IiLAogICAgICAgICAgICAidmFsdWUiOiAiUiBUaWdlciIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAicGxhdGVfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIEJhbmNoYW1layBCbGFjayIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAicGxhdGVfciIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJSIEJhbmNoYW1layBHb2xkZW4iCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImJvZHlfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIEJvZHkgTm9ybWFsIEdlbnRsZW1hbiIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAiaGVhZF9sIiwKICAgICAgICAgICAgInZhbHVlIjogIkwgTk9OIEhFQUQiCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImNsb3RoZXNfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIFJvYmUgR29sZGVuIgogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJleHRyYV9sIiwKICAgICAgICAgICAgInZhbHVlIjogIkwgTk8gRVhUUkEiCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImhhbmRfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIEdsb3ZlIEJsYWNrIEJWIgogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJpbmZsdWVuY2VyIiwKICAgICAgICAgICAgInZhbHVlIjogIlRlZXJhd2F0IFlpb3lpbSIKICAgICAgICB9CiAgICBdLAogICAgIm9uY2hhaW5fYXR0cmlidXRlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImV4Y2x1c2l2ZV9wYXJ0eV9hY2Nlc3MiLAogICAgICAgICAgICAidmFsdWUiOiBmYWxzZQogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJmaXJzdF9kaXNjb3VudF91c2VkIiwKICAgICAgICAgICAgInZhbHVlIjogZmFsc2UKICAgICAgICB9CiAgICBdCn0K"
	// k, _ := keepertest.NftmngrKeeper(t)
	// metadata, err := base64.StdEncoding.DecodeString(base64Data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	schema := `
	{
		"code": "buakaw1",
		"name": "Buakaw1",
		"owner": "0xNFTOWNER",
		"origin_data": {
			"origin_base_uri": "https://bk1nft.sixnetwork.io/ipfs/QmcovEcZxiM1M9gf3535jUCpbqFZSjc2hHuYnZEbWNmt5a",
			"uri_retrieval_method": "BASE",
			"origin_chain": "ethereum",
			"origin_contract_address": "0x9F1CC70b11f4129d042d0037c2066d12E16d9a52",
			"attribute_overriding": "CHAIN",
			"metadata_format": "opensea",
			"origin_attributes": [
				{
					"name": "background_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Background L"
						}
					}
				},
				{
					"name": "background_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Background R"
						}
					}
				},
				{
					"name": "plate_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Plate L"
						}
					}
				},
				{
					"name": "plate_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Plate R"
						}
					}
				},
				{
					"name": "body_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Body L"
						}
					}
				},
				{
					"name": "body_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Body R"
						}
					}
				},
				{
					"name": "head_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Head L"
						}
					}
				},
				{
					"name": "head_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Head R"
						}
					}
				},
				{
					"name": "clothes_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Clothes L"
						}
					}
				},
				{
					"name": "clothes_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Clothes R"
						}
					}
				},
				{
					"name": "extra_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Extra L"
						}
					}
				},
				{
					"name": "extra_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Extra R"
						}
					}
				},
				{
					"name": "hand_l",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Hand L"
						}
					}
				},
				{
					"name": "hand_r",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Hand R"
						}
					}
				},
				{
					"name": "influencer",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Influencer"
						}
					}
				}
			]
		},
		"onchain_data": {
			"reveal_required": true,
			"reveal_secret": "",
			"nft_attributes": [
				{
					"name": "expire_date",
					"data_type": "number",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"display_type": "date",
							"trait_type": "Expire Date"
						}
					}
				}
			],
			"token_attributes": [
				{
					"name": "exclusive_party_access",
					"data_type": "boolean",
					"required": true,
					"hidden_to_marketplace": true
				},
				{
					"name": "first_discount_used",
					"data_type": "boolean",
					"required": true,
					"hidden_to_marketplace": true,
					"display_option": {
						"bool_true_value": "Yes",
						"bool_false_value": "_HIDDEN_",
						"opensea": {
							"display_type": "boolean",
							"trait_type": "First Discount Used"
						}
					}
				},
				{
					"name": "percent_discount_10",
					"data_type": "number",
					"default_mint_value": {
						"number_attribute_value": {
							"value": 10
						}
					},
					"required": true,
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Available 10 Percent Discount",
							"max_value": 10,
							"display_type": "number"
						}
					}
				},
				{
					"name": "discount_percentage",
					"data_type": "number",
					"required": true,
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Discount Percentage",
							"display_type": "boost_percentage"
						}
					}
				},
				{
					"name": "discount_amount",
					"data_type": "number",
					"required": true,
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Discount",
							"display_type": "boost_number"
						}
					}
				}
			],
			"actions": [
				{
					"name": "use10percentdiscount",
					"desc": "Use 10% discount on purchase",
					"when": "meta.GetNumber('percent_discount_10') > 0",
					"then": [
						"meta.SetNumber('percent_discount_10',meta.GetNumber('percent_discount_10') - 1)",
						"meta.SetBoolean('first_discount_used', true)"
					]
				}
			],
			"status": [{
				"status_name": "first_mint_complete",
				"status_value": false
			}],
			"nft_attributes_value": [
				{
					"name": "expire_date",
					"number_attribute_value": {
						"value": 1665009678
					}
				}
			]
		}
	}`
	metadata := `
	{
		"nft_schema_code": "buakaw1",
		"token_id": "1",
		"token_owner": "0xTOKENOWNER",
		"owner_address_type": "ORIGIN_ADDRESS",
		"origin_image": "https://bk1nft.sixnetwork.io/ipfs/QmaBUdqTctVteajpoMVEpo55XEXe2YVwWDgeNi1UT47UKU/1.png",
		"origin_attributes": [
			{
				"name": "background_l",
				"string_attribute_value": {
					"value": "L BG Ringside Golden"
				}
			},
			{
				"name": "background_r",
				"string_attribute_value": {
					"value": "R Tiger"
				}
			},
			{
				"name": "plate_l",
				"string_attribute_value": {
					"value": "L Banchamek Black"
				}
			},
			{
				"name": "plate_r",
				"string_attribute_value": {
					"value": "R Banchamek Golden"
				}
			},
			{
				"name": "body_l",
				"string_attribute_value": {
					"value": "L Body Normal Gentleman"
				}
			},
			{
				"name": "head_l",
				"string_attribute_value": {
					"value": "L NON HEAD"
				}
			},
			{
				"name": "clothes_l",
				"string_attribute_value": {
					"value": "L Robe Golden"
				}
			},
			{
				"name": "extra_l",
				"string_attribute_value": {
					"value": "L NO EXTRA"
				}
			},
			{
				"name": "hand_l",
				"string_attribute_value": {
					"value": "L Glove Black BV"
				}
			},
			{
				"name": "influencer",
				"string_attribute_value": {
					"value": "Teerawat Yioyim"
				}
			}
		],
		"onchain_attributes": [
			{
				"name": "exclusive_party_access",
				"boolean_attribute_value": {
					"value": true
				}
			},
			{
				"name": "first_discount_used",
				"boolean_attribute_value": {
					"value": false
				}
			},
			{
				"name": "influencer",
				"string_attribute_value": {
					"value": "Teerawat Yioyim"
				}
			}
		]
	}`
	data := types.NftData{}
	// err := k.GetCodec().(*codec.ProtoCodec).UnmarshalJSON([]byte(metadata), &data)
	err := jsonpb.UnmarshalString(metadata, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_schema := types.NFTSchema{}
	// err := k.GetCodec().(*codec.ProtoCodec).UnmarshalJSON([]byte(metadata), &data)
	err = jsonpb.UnmarshalString(schema, &_schema)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Append Attribute with default value to NFT Data if not exist in NFT Data yet
	mapOfTokenAttributeValues := map[string]*types.NftAttributeValue{}
	for _, attr := range data.OnchainAttributes {
		mapOfTokenAttributeValues[attr.Name] = attr
	}
	for _, attr := range _schema.OnchainData.TokenAttributes {
		if attr.Required {
			if _, ok := mapOfTokenAttributeValues[attr.Name]; !ok {
				if attr.DefaultMintValue != nil {
					data.OnchainAttributes = append(data.OnchainAttributes, NewNFTAttributeValueFromDefaultValue(attr.Name, attr.DefaultMintValue))
				}
			}
		}
	}
	fmt.Println("data", data)

	// Set default mint value
	// for _, attr := range data.OnchainAttributes {
	// 	hasValue, attrValueType := HasNftAttrubuteValue(*attr)
	// 	fmt.Println("name: ", attr.Name, " | hasValue: ", hasValue, " | attrValueType: ", attrValueType)
	// 	if !hasValue {
	// 		_data, valid := SetDefaultNftAttributeValue(data, *attr, _schema)
	// 		fmt.Println("Data: ", _data)
	// 		fmt.Println("Valid: ", valid)
	// 	}
	// }

	valid, err := ValidateNFTData(&data, &_schema)
	fmt.Println("Valid: ", valid)
	fmt.Println("Valid Err: ", err)

	// fmt.Println("DataOutput", data)

	// data := types.NftData{}
	// attribute := types.NftAttributeValue{
	// 	Name: "exclusive_party_access",
	// }
	// attribute.Value = &types.NftAttributeValue_BooleanAttributeValue_{BooleanAttributeValue: &types.NftAttributeValue_BooleanAttributeValue{Value: true}}
	// data.OnchainAttributes = append(data.OnchainAttributes, &attribute)

	// m := jsonpb.Marshaler{}
	// str, _ := m.MarshalToString(&attribute)
	// fmt.Println(str)
}

func DataCreateAttrDefMap(attrDefs []*types.AttributeDefinition) map[string]*types.AttributeDefinition {
	attrDefMap := make(map[string]*types.AttributeDefinition)
	for _, attrDef := range attrDefs {
		attrDefMap[attrDef.Name] = attrDef
	}
	return attrDefMap
}

// Validate NFT Data
func ValidateNFTData(data *types.NftData, schema *types.NFTSchema) (bool, error) {
	// Origin Data Origin Attributes Map
	mapAttributeDefinition := DataCreateAttrDefMap(schema.OriginData.OriginAttributes)

	// Merge Origin Attributes and Onchain Attributes together
	mergedAttributes := MergeNFTDataAttributes(schema.OriginData.OriginAttributes, schema.OnchainData.TokenAttributes)
	mergedMap := DataCreateAttrDefMap(mergedAttributes)

	// Check if attributes exist in schema
	attributesExistsInSchema, err := NFTDataAttributesExistInSchema(mergedMap, data.OnchainAttributes)
	if !attributesExistsInSchema {
		return false, sdkerrors.Wrap(types.ErrOnchainAttributesNotExistsInSchema, fmt.Sprintf("Attribute does not exist in schema: %s", err))
	}

	fmt.Println("1 NFTDataHasDuplicateNftAttributesValue")
	// Validate Onchain Attributes Value
	duplicated, err := NFTDataHasDuplicateNftAttributesValue(data.OnchainAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOnchainAttributesValue, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	fmt.Println("2 NFTDataHasDuplicateNftAttributesValue")
	// Validate Origin Attributes Value
	duplicated, err = NFTDataHasDuplicateNftAttributesValue(data.OriginAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOriginAttributesValue, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	fmt.Println("3 HasSameTypeAsSchema")
	// Validate Origin Attributes Exist in Schema
	hasSameType, err := HasSameTypeAsSchema(mapAttributeDefinition, data.OriginAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrOriginAttributesNotSameTypeAsSchema, fmt.Sprintf("Duplicate attribute name: %s", err))
	}
	fmt.Println("4 HasSameTypeAsSchema")

	// Validate Onchain Attributes Exist in Schema
	hasSameType, err = HasSameTypeAsSchema(mergedMap, data.OnchainAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrOnchainTokenAttributesNotSameTypeAsSchema, fmt.Sprintf("Duplicate attribute name: %s", err))
	}

	return true, nil
}

func MergeNFTDataAttributes(originAttributes []*types.AttributeDefinition, onchainAttributes []*types.AttributeDefinition) []*types.AttributeDefinition {
	mergedAttributes := make([]*types.AttributeDefinition, 0)
	for _, originAttribute := range originAttributes {
		mergedAttributes = append(append(mergedAttributes, originAttribute), onchainAttributes...)
	}
	for _, onchainAttribute := range onchainAttributes {
		mergedAttributes = append(append(mergedAttributes, onchainAttribute), originAttributes...)
	}
	return mergedAttributes
}

func NFTDataHasDuplicateNftAttributesValue(attributes []*types.NftAttributeValue) (bool, string) {
	mapAttributes := map[string]*types.NftAttributeValue{}
	for _, attriDef := range attributes {
		if _, ok := mapAttributes[attriDef.Name]; ok {
			return true, attriDef.Name
		}
		mapAttributes[attriDef.Name] = attriDef
	}
	return false, ""
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

// If onchain attributes value does not exist use default_mint_value from schema instead
func HasNftAttrubuteValue(attribute types.NftAttributeValue) (bool, string) {
	// Check if onchain attribute s value exist for each attribute
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_BooleanAttributeValue); ok {
		return ok, "boolean"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_StringAttributeValue); ok {
		return ok, "string"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_NumberAttributeValue); ok {
		return ok, "number"
	}
	if _, ok := attribute.GetValue().(*types.NftAttributeValue_FloatAttributeValue); ok {
		return ok, "float"
	}
	return false, "default"
}

func SetDefaultNftAttributeValue(data types.NftData, attribute types.NftAttributeValue, schema types.NFTSchema) (types.NftData, bool) {
	mapAttributeDefinition := DataCreateAttrDefMap(schema.OnchainData.TokenAttributes)

	for _, attriDef := range data.OnchainAttributes {
		if attriDef.Name != attribute.Name {
			continue
		}
		if mapAttributeDefinition[attriDef.Name].DataType == "number" {
			newValue := mapAttributeDefinition[attriDef.Name].DefaultMintValue.GetNumberAttributeValue()
			attriDef.Value = &types.NftAttributeValue_NumberAttributeValue{NumberAttributeValue: newValue}
		}
		return data, true
	}
	return data, false
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

// Check if NFT data attributes exists in schema
func NFTDataAttributesExistInSchema(mapAttributes map[string]*types.AttributeDefinition, dataAttributes []*types.NftAttributeValue) (bool, string) {
	// Check if dataAttributes exist in schema
	for _, attriVal := range dataAttributes {
		fmt.Println("NFTDataAttributesExistInSchema: ", attriVal.Name)
		if _, ok := mapAttributes[attriVal.Name]; !ok {
			return false, attriVal.Name
		}
	}
	return true, ""
}
