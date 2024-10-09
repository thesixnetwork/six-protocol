package keeper_test

import (
	"fmt"
	"testing"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/gogo/protobuf/jsonpb"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
)

func TestCreateSchema(t *testing.T) {
	fmt.Println("Start Test")
	// base64Data := "ewogICAgIm5mdF9zY2hlbWFfY29kZSI6ICJidWFrYXcxIiwKICAgICJ0b2tlbl9pZCI6ICIxIiwKICAgICJ0b2tlbl9vd25lciI6ICI2bmZ0MTlwNXl2d3h0YzJxbmdud2RxbGhmemRjZDBzc25jd3d1bXJ4MmphIiwKICAgICJvcmlnaW5faW1hZ2UiOiAiaHR0cHM6Ly9iazFuZnQuc2l4bmV0d29yay5pby9pcGZzL1FtYUJVZHFUY3RWdGVhanBvTVZFcG81NVhFWGUyWVZ3V0RnZU5pMVVUNDdVS1UvMS5wbmciLAogICAgIm9yaWdpbl9hdHRyaWJ1dGVzIjogWwogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAiYmFja2dyb3VuZF9sIiwKICAgICAgICAgICAgInZhbHVlIjogIkwgQkcgUmluZ3NpZGUgR29sZGVuIgogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJiYWNrZ3JvdW5kX3IiLAogICAgICAgICAgICAidmFsdWUiOiAiUiBUaWdlciIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAicGxhdGVfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIEJhbmNoYW1layBCbGFjayIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAicGxhdGVfciIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJSIEJhbmNoYW1layBHb2xkZW4iCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImJvZHlfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIEJvZHkgTm9ybWFsIEdlbnRsZW1hbiIKICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICAgIm5hbWUiOiAiaGVhZF9sIiwKICAgICAgICAgICAgInZhbHVlIjogIkwgTk9OIEhFQUQiCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImNsb3RoZXNfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIFJvYmUgR29sZGVuIgogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJleHRyYV9sIiwKICAgICAgICAgICAgInZhbHVlIjogIkwgTk8gRVhUUkEiCiAgICAgICAgfSwKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImhhbmRfbCIsCiAgICAgICAgICAgICJ2YWx1ZSI6ICJMIEdsb3ZlIEJsYWNrIEJWIgogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJpbmZsdWVuY2VyIiwKICAgICAgICAgICAgInZhbHVlIjogIlRlZXJhd2F0IFlpb3lpbSIKICAgICAgICB9CiAgICBdLAogICAgIm9uY2hhaW5fYXR0cmlidXRlcyI6IFsKICAgICAgICB7CiAgICAgICAgICAgICJuYW1lIjogImV4Y2x1c2l2ZV9wYXJ0eV9hY2Nlc3MiLAogICAgICAgICAgICAidmFsdWUiOiBmYWxzZQogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgICAibmFtZSI6ICJmaXJzdF9kaXNjb3VudF91c2VkIiwKICAgICAgICAgICAgInZhbHVlIjogZmFsc2UKICAgICAgICB9CiAgICBdCn0K"
	// k, _ := keepertest.NftmngrKeeper(t)
	// intput_nft_schema, err := base64.StdEncoding.DecodeString(base64Data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	intput_nft_schema := `
	{
		"code": "sixnetwork.nftexpo4",
		"name": "NFTexpo",
		"owner": "0xNFTOWNER",
		"origin_data": {
			"origin_base_uri": "https://ipfs.io/ipfs/QmcVhbj3Vt8W3YcuyyyDmbbRSJjFSsprR4EHwsgeJnVHfw/",
			"uri_retrieval_method": "BASE",
			"origin_chain": "ethereum",
			"origin_contract_address": "0xaA83FA374645E875Ea58Bb94596d4adB467A06Ff",
			"attribute_overriding": "CHAIN",
			"metadata_format": "opensea",
			"origin_attributes": [
				{
					"name": "background",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Background"
						}
					}
				},
				{
					"name": "foreground",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Foreground"
						}
					}
				},
				{
					"name": "head",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Head"
						}
					}
				},
				{
					"name": "moon",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Moon"
						}
					}
				},
				{
					"name": "tail",
					"data_type": "string",
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Tail"
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
					"name": "points",
					"default_mint_value": {
						"number_attribute_value": {
							"value": 0
						}
					},
					"data_type": "number",
					"required": true,
					"display_value_field": "value",
					"hidden_to_marketplace": true,
					"display_option": {
						"opensea": {
							"trait_type": "Points"
						}
					}
				},
				{
					"name": "check_point_1",
					"default_mint_value": {
						"boolean_attribute_value": {
							"value": false
						}
					},
					"data_type": "boolean",
					"required": true,
					"display_value_field": "value",
					"hidden_to_marketplace": false,
					"display_option": {
						"bool_true_value": "Checked",
						"bool_false_value": "No",
						"opensea": {
							"trait_type": "Check Point 1"
						}
					}
				},
				{
					"name": "check_point_2",
					"default_mint_value": {
						"boolean_attribute_value": {
							"value": false
						}
					},
					"data_type": "boolean",
					"required": true,
					"display_value_field": "value",
					"hidden_to_marketplace": false,
					"display_option": {
						"bool_true_value": "Checked",
						"bool_false_value": "No",
						"opensea": {
							"trait_type": "Check Point 2"
						}
					}
				},
				{
					"name": "mission_milestones",
					"default_mint_value": {
						"number_attribute_value": {
							"value": 0
						}
					},
					"data_type": "number",
					"required": true,
					"display_value_field": "value",
					"display_option": {
						"opensea": {
							"trait_type": "Milestones",
							"max_value": 3,
							"display_type": "number"
						}
					}
				},
				{
					"name": "bonus_milestone",
					"default_mint_value": {
						"boolean_attribute_value": {
							"value": false
						}
					},
					"data_type": "boolean",
					"required": true,
					"display_value_field": "value",
					"display_option": {
						"bool_true_value": "Yes",
						"bool_false_value": "No",
						"opensea": {
							"trait_type": "Bonus Milestone"
						}
					}
				},
				{
					"name": "ais_customer",
					"default_mint_value": {
						"boolean_attribute_value": {
							"value": false
						}
					},
					"data_type": "boolean",
					"required": true,
					"display_value_field": "value",
					"display_option": {
						"bool_true_value": "Yes",
						"bool_false_value": "No",
						"opensea": {
							"trait_type": "AIS Customer"
						}
					}
				},
				{
					"name": "ais_gift_redeemed",
					"default_mint_value": {
						"boolean_attribute_value": {
							"value": false
						}
					},
					"data_type": "boolean",
					"required": true,
					"display_value_field": "value",
					"display_option": {
						"bool_true_value": "Yes",
						"bool_false_value": "No",
						"opensea": {
							"trait_type": "AIS Gift Redeemed"
						}
					}
				},
				{
					"name": "transformed",
					"default_mint_value": {
						"boolean_attribute_value": {
							"value": false
						}
					},
					"data_type": "boolean",
					"required": true,
					"display_value_field": "value",
					"display_option": {
						"bool_true_value": "Yes",
						"bool_false_value": "No",
						"opensea": {
							"trait_type": "Transformed"
						}
					}
				}
			],
			"actions": [
				{
					"name": "use_checkpoint_1",
					"desc": "Attend the event",
					"when": "meta.GetBoolean('check_point_1') == false",
					"then": [
						"meta.SetBoolean('check_point_1', true)",
						"meta.SetNumber('mission_milestones', meta.GetNumber('mission_milestones') + 1)"
					]
				},
				{
					"name": "use_checkpoint_2",
					"desc": "Go to stage 1",
					"when": "meta.GetBoolean('check_point_2') == false",
					"then": [
						"meta.SetBoolean('check_point_2', true)",
						"meta.SetNumber('mission_milestones', meta.GetNumber('mission_milestones') + 1)"
					]
				},
				{
					"name": "claim_bonus",
					"desc": "Claim bonus points",
					"when": "meta.GetBoolean('bonus_milestone') == false",
					"then": [
						"meta.SetNumber('points', meta.GetNumber('points') + 100)",
						"meta.SetBoolean('bonus_milestone', true)",
						"meta.SetNumber('mission_milestones', meta.GetNumber('mission_milestones') + 1)"
					]
				},
				{
					"name": "redeem_merchandise_100",
					"desc": "Redeem merch for 100 points",
					"when": "meta.GetBoolean('points') > 100",
					"then": [
						"meta.SetNumber('points', meta.GetNumber('points') - 100)"
					]
				},
				{
					"name": "redeem_merchandise_200",
					"desc": "Redeem merch for 200 points",
					"when": "meta.GetBoolean('points') > 200",
					"then": [
						"meta.SetNumber('points', meta.GetNumber('points') - 200)"
					]
				},
				{
					"name": "redeem_merchandise_300",
					"desc": "Redeem merch for 300 points",
					"when": "meta.GetBoolean('points') > 300",
					"then": [
						"meta.SetNumber('points', meta.GetNumber('points') - 300)"
					]
				},
				{
					"name": "redeem_merchandise_400",
					"desc": "Redeem merch for 400 points",
					"when": "meta.GetBoolean('points') > 400",
					"then": [
						"meta.SetNumber('points', meta.GetNumber('points') - 400)"
					]
				},
				{
					"name": "redeem_merchandise_500",
					"desc": "Redeem merch for 500 points",
					"when": "meta.GetBoolean('points') > 500",
					"then": [
						"meta.SetNumber('points', meta.GetNumber('points') - 500)"
					]
				},
				{
					"name": "redeem_merchandise_ais",
					"desc": "Redeem AIS gift",
					"when": "meta.GetBoolean('points') > 500 && meta.GetBoolean('ais_customer') == true && meta.GetBoolean('ais_gift_redeemed') == false",
					"then": [
						"meta.SetNumber('points', meta.GetNumber('points') - 500)",
						"meta.SetBoolean('ais_gift_redeemed', true)"
					]
				},
				{
					"name": "transform",
					"desc": "Transform",
					"when": "meta.GetBoolean('transformed') == false && meta.GetNumber('mission_milestones') == 3",
					"then": [
						"meta.SetBoolean('transformed', true)",
						"meta.SetImage(meta.ReplaceAllString(meta.GetImage(),'.jpeg','-t.jpeg'))"
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
						"value": 1726694811
					}
				}
			]
		}
	}
	
	`
	schema_input := types.NFTSchemaINPUT{}
	// err := k.GetCodec().(*codec.ProtoCodec).UnmarshalJSON([]byte(intput_nft_schema), &data)
	err := jsonpb.UnmarshalString(intput_nft_schema, &schema_input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	valid, err := keeper.ValidateNFTSchema(&schema_input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// merge all attributes and label with index
	keeper.MergeAllAttributesAndAlterOrderIndex(schema_input.OriginData.OriginAttributes, schema_input.OnchainData.NftAttributes, schema_input.OnchainData.TokenAttributes)

	fmt.Println("Valid: ", valid)
	// print data output to console as json and formatted
	marshaler := jsonpb.Marshaler{Indent: "  "}
	json, err := marshaler.MarshalToString(&schema_input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(json)

}

func CreateAttrDefMap(attrDefs []*types.AttributeDefinition) map[string]*types.AttributeDefinition {
	attrDefMap := make(map[string]*types.AttributeDefinition)
	for n, attrDef := range attrDefs {
		attrDefMap[attrDef.Name] = attrDef
		attrDefMap[attrDef.Name].Index = uint64(n)
	}
	fmt.Println(attrDefMap)
	return attrDefMap
}

func CreateNftAttrValueMap(nftAttrValues []*types.NftAttributeValue) map[string]*types.NftAttributeValue {
	nftAttrValueMap := make(map[string]*types.NftAttributeValue)
	for _, nftAttrValue := range nftAttrValues {
		nftAttrValueMap[nftAttrValue.Name] = nftAttrValue
	}
	return nftAttrValueMap
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

func HasSameType(mapOriginAttributes map[string]*types.AttributeDefinition, onchainAttributes []*types.AttributeDefinition) (bool, string) {
	for _, attriVal := range onchainAttributes {
		attrDef := mapOriginAttributes[attriVal.Name]
		if attrDef == nil {
			continue
		}
		if attrDef.DataType != attriVal.DataType {
			return false, attrDef.Name
		}
	}
	return true, ""
}

func DefaultMintValueHasSameTypeAs(attributes []*types.AttributeDefinition) (bool, string) {
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

func SchemaAttributeHasDefaultMintValue(attribute types.AttributeDefinition) (bool, string) {
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
