package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	NftmngrKeeper "github.com/thesixnetwork/sixnft/x/nftmngr/keeper"
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
			schemaAttibuteConverted, _ := NftmngrKeeper.ConvertDefaultMintValueToSchemaAttributeValue(nftAttribute.DefaultMintValue)
			app.NftmngrKeeper.SetSchemaAttribute(ctx, nftmngrtypes.SchemaAttribute{
				NftSchemaCode:       nftSchemaV1.Code,
				Name:                nftAttribute.Name,
				DataType:            nftAttribute.DataType,
				Required:            nftAttribute.Required,
				DisplayValueField:   nftAttribute.DisplayValueField,
				DisplayOption:       nftAttribute.DisplayOption,
				CurrentValue:        schemaAttibuteConverted,
				HiddenOveride:       nftAttribute.HiddenOveride,
				HiddenToMarketplace: nftAttribute.HiddenToMarketplace,
				Creator:             nftSchemaV1.Owner,
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
