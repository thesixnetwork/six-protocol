package app

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	nftmngrKeeper "github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
	nftmngrtypes "github.com/thesixnetwork/sixnft/x/nftmngr/types"
)

func (app *App) MigrationFromV1ToV2Handlers(ctx sdk.Context) {

	// get all NFTSchema
	nftSchemasV1 := app.NftmngrKeeper.GetAllNFTSchemaV1(ctx)

	for _, nftSchemaV1 := range nftSchemasV1 {
		// migrate list of system actioners to new map
		for _, systemActioner := range nftSchemaV1.SystemActioners {
			// set system actioner to new map
			app.NftmngrKeeper.SetActionExecutor(ctx, nftmngrtypes.ActionExecutor{
				NftSchemaCode:   nftSchemaV1.Code,
				ExecutorAddress: systemActioner,
				Creator:         nftSchemaV1.Owner,
			})

			val, found := app.NftmngrKeeper.GetExecutorOfSchema(ctx, nftSchemaV1.Code)
			if !found {
				val = nftmngrtypes.ExecutorOfSchema{
					NftSchemaCode:   nftSchemaV1.Code,
					ExecutorAddress: []string{},
				}
			}

			// set executorOfSchema
			val.ExecutorAddress = append(val.ExecutorAddress, systemActioner)

			app.NftmngrKeeper.SetExecutorOfSchema(ctx, nftmngrtypes.ExecutorOfSchema{
				NftSchemaCode:   nftSchemaV1.Code,
				ExecutorAddress: val.ExecutorAddress,
			})
		}

		// migrate schema to new schema
		app.NftmngrKeeper.SetNFTSchema(ctx, nftmngrtypes.NFTSchema{
			Code:        nftSchemaV1.Code,
			Name:        nftSchemaV1.Name,
			Owner:       nftSchemaV1.Owner,
			Description: nftSchemaV1.Name,
			OriginData:  nftSchemaV1.OriginData,
			OnchainData: &nftmngrtypes.OnChainData{
				TokenAttributes: nftSchemaV1.OnchainData.TokenAttributes,
				Actions:         nftSchemaV1.OnchainData.Actions,
				Status:          nftSchemaV1.OnchainData.Status,
			},
			IsVerified:        nftSchemaV1.IsVerified,
			MintAuthorization: nftSchemaV1.MintAuthorization,
		})

		// migrate NFT attributes to new schema attributes
		for _, nftAttribute := range nftSchemaV1.OnchainData.NftAttributes {
			schemaAttibuteConverted, _ := nftmngrKeeper.ConvertDefaultMintValueToSchemaAttributeValue(nftAttribute.DefaultMintValue)
			app.NftmngrKeeper.SetSchemaAttribute(ctx, nftmngrtypes.SchemaAttribute{
				NftSchemaCode: nftSchemaV1.Code,
				Name:          nftAttribute.Name,
				DataType:      nftAttribute.DataType,
				CurrentValue:  schemaAttibuteConverted,
				Creator:       nftSchemaV1.Owner,
			})

		}

		// migrate NFT actions to new schema actions
		for i, nftAction := range nftSchemaV1.OnchainData.Actions {
			app.NftmngrKeeper.SetActionOfSchema(ctx, nftmngrtypes.ActionOfSchema{
				NftSchemaCode: nftSchemaV1.Code,
				Name:          nftAction.Name,
				Index:         uint64(i),
			})
		}
	}

}

func (app *App) RollbackFromV2toV1(ctx sdk.Context) {

	// get all NFTSchema
	nftSchemasV2 := app.NftmngrKeeper.GetAllNFTSchemaV2(ctx)

	for _, nftSchemaV2 := range nftSchemasV2 {

		nft_attributes_from_schema_attribute := []*nftmngrtypes.AttributeDefinition{}

		attribute_of_schema, found := app.NftmngrKeeper.GetAttributeOfSchema(ctx, nftSchemaV2.Code)
		if !found {
			continue
		}

		for i, attribute := range attribute_of_schema.SchemaAttributes {
			schemaAttribute, _ := ConvertSchemaAttributeToNftAttributeDefinition(attribute, i)
			nft_attributes_from_schema_attribute = append(nft_attributes_from_schema_attribute, schemaAttribute)
		}

		originOutput, nftOutput, tokenOutput := MergeAllAttributesAndAlterOrderIndex(nftSchemaV2.OriginData.OriginAttributes, nft_attributes_from_schema_attribute, nftSchemaV2.OnchainData.TokenAttributes)

		// parse schema_input to NFTSchema
		schema := nftmngrtypes.NFTSchema{
			Code:        nftSchemaV2.Code,
			Name:        nftSchemaV2.Name,
			Owner:       nftSchemaV2.Owner,
			Description: nftSchemaV2.Description,
			OriginData: &nftmngrtypes.OriginData{
				OriginChain:           nftSchemaV2.OriginData.OriginChain,
				OriginContractAddress: nftSchemaV2.OriginData.OriginContractAddress,
				OriginBaseUri:         nftSchemaV2.OriginData.OriginBaseUri,
				AttributeOverriding:   nftSchemaV2.OriginData.AttributeOverriding,
				MetadataFormat:        nftSchemaV2.OriginData.MetadataFormat,
				OriginAttributes:      originOutput,
				UriRetrievalMethod:    nftSchemaV2.OriginData.UriRetrievalMethod,
			},
			OnchainData: &nftmngrtypes.OnChainData{
				TokenAttributes: tokenOutput,
				NftAttributes:   nftOutput,
				Actions:         nftSchemaV2.OnchainData.Actions,
				Status:          nftSchemaV2.OnchainData.Status,
			},
			IsVerified:        nftSchemaV2.IsVerified,
			MintAuthorization: nftSchemaV2.MintAuthorization,
		}

		for _, schemaDefaultMintAttribute := range schema.OnchainData.NftAttributes {
			schmaAttributeValue, _ := nftmngrKeeper.ConvertDefaultMintValueToSchemaAttributeValue(schemaDefaultMintAttribute.DefaultMintValue)

			app.NftmngrKeeper.SetSchemaAttribute(ctx, nftmngrtypes.SchemaAttribute{
				NftSchemaCode: schema.Code,
				Name:          schemaDefaultMintAttribute.Name,
				DataType:      schemaDefaultMintAttribute.DataType,
				CurrentValue:  schmaAttributeValue,
				Creator:       schema.Owner,
			})
		}

		app.NftmngrKeeper.SetNFTSchema(ctx, schema)

	}

}

// function to convert SchemaAttribute to NftAttributeDefinition
func ConvertSchemaAttributeToNftAttributeDefinition(schemaAttributes *nftmngrtypes.SchemaAttributeV1, index int) (*nftmngrtypes.AttributeDefinition, error) {
	attributeDef := &nftmngrtypes.AttributeDefinition{}

	switch value := schemaAttributes.CurrentValue.Value.(type) {
	case *nftmngrtypes.SchemaAttributeValue_NumberAttributeValue:
		attributeDef.DefaultMintValue = &nftmngrtypes.DefaultMintValue{
			Value: &nftmngrtypes.DefaultMintValue_NumberAttributeValue{
				NumberAttributeValue: value.NumberAttributeValue,
			},
		}
	case *nftmngrtypes.SchemaAttributeValue_StringAttributeValue:
		attributeDef.DefaultMintValue = &nftmngrtypes.DefaultMintValue{
			Value: &nftmngrtypes.DefaultMintValue_StringAttributeValue{
				StringAttributeValue: value.StringAttributeValue,
			},
		}
	case *nftmngrtypes.SchemaAttributeValue_BooleanAttributeValue:
		attributeDef.DefaultMintValue = &nftmngrtypes.DefaultMintValue{
			Value: &nftmngrtypes.DefaultMintValue_BooleanAttributeValue{
				BooleanAttributeValue: value.BooleanAttributeValue,
			},
		}
	case *nftmngrtypes.SchemaAttributeValue_FloatAttributeValue:
		attributeDef.DefaultMintValue = &nftmngrtypes.DefaultMintValue{
			Value: &nftmngrtypes.DefaultMintValue_FloatAttributeValue{
				FloatAttributeValue: value.FloatAttributeValue,
			},
		}
	default:
		return nil, fmt.Errorf("unknown value type: %T", value)
	}

	return &nftmngrtypes.AttributeDefinition{
		Name:                schemaAttributes.Name,
		DataType:            schemaAttributes.DataType,
		Required:            schemaAttributes.Required,
		DisplayValueField:   schemaAttributes.DisplayValueField,
		DisplayOption:       schemaAttributes.DisplayOption,
		DefaultMintValue:    attributeDef.DefaultMintValue,
		HiddenOveride:       schemaAttributes.HiddenOveride,
		HiddenToMarketplace: schemaAttributes.HiddenToMarketplace,
		Index:               uint64(index),
	}, nil
}

func MergeAllAttributesAndAlterOrderIndex(originAttributes []*nftmngrtypes.AttributeDefinition, nftAttribute []*nftmngrtypes.AttributeDefinition, tokenAttribute []*nftmngrtypes.AttributeDefinition) (originAttributeWithIndex []*nftmngrtypes.AttributeDefinition, nftAttributeWithIndex []*nftmngrtypes.AttributeDefinition, tokenAttributeWithIndex []*nftmngrtypes.AttributeDefinition) {
	orignOutput := make([]*nftmngrtypes.AttributeDefinition, 0)
	nftOutput := make([]*nftmngrtypes.AttributeDefinition, 0)
	tokenOutput := make([]*nftmngrtypes.AttributeDefinition, 0)
	// var index uint64 = 0
	// for _, attribute := range append(originAttributes, onchainTokenAttribute...) {
	// 	attribute.Index = index
	// 	mergedAttributes = append(mergedAttributes, attribute)
	// 	index++
	// }

	var index uint64 = 0
	for _, attribute := range originAttributes {
		attribute.Index = index
		orignOutput = append(orignOutput, attribute)
		index++
	}
	for _, attribute := range nftAttribute {
		attribute.Index = index
		nftOutput = append(nftOutput, attribute)
		index++
	}

	for _, attribute := range tokenAttribute {
		attribute.Index = index
		tokenOutput = append(tokenOutput, attribute)
		index++
	}

	return orignOutput, nftOutput, tokenOutput
}
