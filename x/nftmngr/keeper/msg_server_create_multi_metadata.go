package keeper

import (
	"context"
	"encoding/base64"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMultiMetadata(goCtx context.Context, msg *types.MsgCreateMultiMetadata) (*types.MsgCreateMultiMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//limit size token_id <= 1000
	token_size := len(msg.TokenId)
	string_token_size := string(rune(token_size))

	if token_size > 1000 {
		return nil, sdkerrors.Wrap(types.ErrLimitSizeOfInput, string_token_size)
	}

	//check if id in msg.TokenId is duplicate
	mapOfTokenId := make(map[string]bool)
	for _, tokenId := range msg.TokenId {
		if _, ok := mapOfTokenId[tokenId]; ok {
			return nil, sdkerrors.Wrap(types.ErrDuplicateInputTokenID, tokenId)
		}
		mapOfTokenId[tokenId] = true
	}

	metadata, err := base64.StdEncoding.DecodeString(msg.Base64NFTData)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}
	data := types.NftData{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(metadata, &data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}
	// Get nft schema from store
	schema, schemaFound := k.Keeper.GetNFTSchema(ctx, data.NftSchemaCode)
	// Check if the schema already exists
	if !schemaFound {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, data.NftSchemaCode)
	}

	// Validate Schema Message and return error if not valid
	valid, err := ValidateNFTData(&data, &schema)
	_ = valid
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// Append Attribute with default value to NFT Data if not exist in NFT Data yet
	mapOfTokenAttributeValues := map[string]*types.NftAttributeValue{}
	for _, attr := range data.OnchainAttributes {
		mapOfTokenAttributeValues[attr.Name] = attr
	}
	for _, attr := range schema.OnchainData.TokenAttributes {
		if attr.Required {
			if _, ok := mapOfTokenAttributeValues[attr.Name]; !ok {
				if attr.DefaultMintValue != nil {
					data.OnchainAttributes = append(data.OnchainAttributes, NewNFTAttributeValueFromDefaultValue(attr.Name, attr.DefaultMintValue))
				}
			}
		}
	}

	// TODO:: RECHCK THIS
	// // Add attributes from schema to metadata onchain attributes
	// for _, attribute := range schema.OnchainData.NftAttributesValue {
	// 	data.OnchainAttributes = append(append(data.OnchainAttributes, attribute), data.OnchainAttributes...)
	// }

	// validate flag of data.TokenID
	if data.TokenId != "MULTIMINT" {
		return nil, sdkerrors.Wrap(types.ErrInvalidFlagTokenID, data.TokenId)
	}

	// iterate through token_id list
	for _, tokenId := range msg.TokenId {

		mapOfMinters, userMintfound := k.Keeper.GetMetadataCreator(ctx, data.NftSchemaCode)
		// Check mint authorization
		switch schema.MintAuthorization {
		case types.KeyMintPermissionOnlySystem:
			// Check if creator is the schema owner
			if msg.Creator != schema.Owner {
				return nil, sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, msg.Creator)
			}
		case types.KeyMintPermissionAll:
			// Add creator to minters list
			if !userMintfound {
				mapOfMinters = types.MetadataCreator{
					NftSchemaCode:    schema.Code,
					MetadataMintedBy: make([]*types.MapTokenToMinter, 0),
				}
				mapOfMinters.MetadataMintedBy = append(mapOfMinters.MetadataMintedBy, &types.MapTokenToMinter{
					TokenId: tokenId,
					Minter:  msg.Creator,
				})
			} else {
				mapOfMinters.MetadataMintedBy = append(mapOfMinters.MetadataMintedBy, &types.MapTokenToMinter{
					TokenId: tokenId,
					Minter:  msg.Creator,
				})
			}
		}

		// replace token_id with token_id from msg
		data.TokenId = tokenId

		// Check if the data already exists
		_, dataFound := k.Keeper.GetNftData(ctx, data.NftSchemaCode, tokenId)
		if dataFound {
			return nil, sdkerrors.Wrap(types.ErrMetadataAlreadyExists, data.NftSchemaCode)
		}

		if !schema.OnchainData.GetStatusByKey(types.KeyNFTStatusFirstMintComplete) {
			schema.OnchainData.SetStatusByKey(types.KeyNFTStatusFirstMintComplete, true)
			k.Keeper.SetNFTSchema(ctx, schema)
		}

		// Add minter to minters list
		k.Keeper.SetMetadataCreator(ctx, mapOfMinters)

		// Add the data to the store
		k.Keeper.SetNftData(ctx, data)

		// Add the minted of any schema to collection
		k.Keeper.AddMetadataToCollection(ctx, &data)
	}

	// stringfy tokenId list to string token_id1,token_id2,token_id3
	tokenIdList := StringfyTokenIdList(msg.TokenId)

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateMetadata,
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyCreateMetaDataTokenID, "["+tokenIdList+"]"),
			sdk.NewAttribute(types.AttributeKeyCreateMetaDataResult, "success"),
		),
	})

	return &types.MsgCreateMultiMetadataResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}
