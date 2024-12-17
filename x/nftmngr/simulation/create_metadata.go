package simulation

import (
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func GenNFTSchemaFromInput(schemaInput types.NFTSchemaINPUT) types.NFTSchema {
	_schema := types.NFTSchema{
		Code:        schemaInput.Code,
		Name:        schemaInput.Name,
		Owner:       schemaInput.Owner,
		Description: schemaInput.Description,
		OriginData:  schemaInput.OriginData,
		OnchainData: &types.OnChainData{
			TokenAttributes: schemaInput.OnchainData.TokenAttributes,
			Actions:         schemaInput.OnchainData.Actions,
			Status:          schemaInput.OnchainData.Status,
		},
		IsVerified:        schemaInput.IsVerified,
		MintAuthorization: schemaInput.MintAuthorization,
	}
	return _schema
}

func SimulateCreateSchemaMetadata(schemaInput types.NFTSchemaINPUT, metaInput types.NftData) (schema_ types.NFTSchema, metadata_ types.NftData, schemaAttributesValue []*types.NftAttributeValue) {
	_schema := types.NFTSchema{
		Code:        schemaInput.Code,
		Name:        schemaInput.Name,
		Owner:       schemaInput.Owner,
		Description: schemaInput.Description,
		OriginData:  schemaInput.OriginData,
		OnchainData: &types.OnChainData{
			TokenAttributes: schemaInput.OnchainData.TokenAttributes,
			Actions:         schemaInput.OnchainData.Actions,
			Status:          schemaInput.OnchainData.Status,
		},
		IsVerified:        schemaInput.IsVerified,
		MintAuthorization: schemaInput.MintAuthorization,
	}
	var listOfGlobalAttributeValue []*types.NftAttributeValue
	for _, schemaAttribute := range schemaInput.OnchainData.NftAttributes {
		schemaAttribute_ := keeper.ConverAttributeDefinitionToNFTAttributeValue(schemaAttribute)
		listOfGlobalAttributeValue = append(listOfGlobalAttributeValue, schemaAttribute_)
	}

	_, err := keeper.ValidateNFTData(&metaInput, &_schema)
	if err != nil {
		panic(err)
	}

	// Append Attribute with default value to NFT Data if not exist in NFT Data yet
	mapOfTokenAttributeValues := map[string]*types.NftAttributeValue{}
	for _, attr := range metaInput.OnchainAttributes {
		mapOfTokenAttributeValues[attr.Name] = attr
	}
	for _, attr := range _schema.OnchainData.TokenAttributes {
		if attr.Required {
			if _, ok := mapOfTokenAttributeValues[attr.Name]; !ok {
				if attr.DefaultMintValue != nil {
					metaInput.OnchainAttributes = append(metaInput.OnchainAttributes, keeper.NewNFTAttributeValueFromDefaultValue(attr.Name, attr.DefaultMintValue))
				}
			}
		}
	}

	globalAttributeValues := []*types.NftAttributeValue{}
	for _, attr := range schemaInput.OnchainData.NftAttributes {
		if attr.Required {
			if _, ok := mapOfTokenAttributeValues[attr.Name]; !ok {
				if attr.DefaultMintValue != nil {
					globalAttributeValues = append(globalAttributeValues, keeper.NewNFTAttributeValueFromDefaultValue(attr.Name, attr.DefaultMintValue))
				}
			}
		}
	}

	return _schema, metaInput, globalAttributeValues
}
