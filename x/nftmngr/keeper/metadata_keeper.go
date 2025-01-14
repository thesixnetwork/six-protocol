package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k Keeper) CreateNewMetadataKeeper(ctx sdk.Context, creator, nftSchemaName, tokenId string, metadata types.NftData) error {
	schema, schemaFound := k.GetNFTSchema(ctx, nftSchemaName)
	if !schemaFound {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	mapOfMinters, userMintfound := k.GetMetadataCreator(ctx, nftSchemaName)

	// Check mint authorization
	switch schema.MintAuthorization {
	case types.KeyMintPermissionOnlySystem:
		// Check if creator is the schema owner
		if creator != schema.Owner {
			return sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, creator)
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
				Minter:  creator,
			})
		} else {
			mapOfMinters.MetadataMintedBy = append(mapOfMinters.MetadataMintedBy, &types.MapTokenToMinter{
				TokenId: tokenId,
				Minter:  creator,
			})
		}
	}

	// Validate Schema Message and return error if not valid
	valid, err := ValidateNFTData(&metadata, &schema)
	_ = valid
	if err != nil {
		return sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// Append Attribute with default value to NFT Data if not exist in NFT Data yet
	mapOfTokenAttributeValues := map[string]*types.NftAttributeValue{}
	for _, attr := range metadata.OnchainAttributes {
		mapOfTokenAttributeValues[attr.Name] = attr
	}
	for _, attr := range schema.OnchainData.TokenAttributes {
		if attr.Required {
			if _, ok := mapOfTokenAttributeValues[attr.Name]; !ok {
				if attr.DefaultMintValue != nil {
					metadata.OnchainAttributes = append(metadata.OnchainAttributes, NewNFTAttributeValueFromDefaultValue(attr.Name, attr.DefaultMintValue))
				}
			}
		}
	}

	// Check if the data already exists
	_, dataFound := k.GetNftData(ctx, metadata.NftSchemaCode, metadata.TokenId)
	if dataFound {
		return sdkerrors.Wrap(types.ErrMetadataAlreadyExists, metadata.NftSchemaCode)
	}

	if !schema.OnchainData.GetStatusByKey(types.KeyNFTStatusFirstMintComplete) {
		schema.OnchainData.SetStatusByKey(types.KeyNFTStatusFirstMintComplete, true)
		k.SetNFTSchema(ctx, schema)
	}

	// Add minter to minters list
	k.SetMetadataCreator(ctx, mapOfMinters)

	// Add the data to the store
	k.SetNftData(ctx, metadata)

	// Add the minted of any schema to collection
	k.AddMetadataToCollection(ctx, &metadata)

	return nil
}

// Validate NFT Data
func ValidateNFTData(data *types.NftData, schema *types.NFTSchema) (bool, error) {
	mapAttributeDefinition := CreateAttrDefMap(schema.OriginData.OriginAttributes)

	// NOTE: No need to merge schema attributes because we will retrieve the attributes value from its directly
	// WE WILL NOT CHANGE SCHEMA VALUE IN THIS FUNCTION
	mergedAttributes := MergeNFTDataAttributes(schema.OriginData.OriginAttributes, schema.OnchainData.TokenAttributes)
	mergedMap := CreateAttrDefMap(mergedAttributes)

	// Check if attributes exist in schema
	attributesExistsInSchema, err := NFTDataAttributesExistInSchema(mergedMap, data.OnchainAttributes)
	if !attributesExistsInSchema {
		return false, sdkerrors.Wrap(types.ErrOnchainAttributesNotExistsInSchema, fmt.Sprintf("Attribute does not exist in schema: %s", err))
	}

	// Check if origin attributes exist in schema
	attributesOriginExistsInSchema, err := NFTDataAttributesExistInSchema(mapAttributeDefinition, data.OriginAttributes)
	if !attributesOriginExistsInSchema {
		return false, sdkerrors.Wrap(types.ErrOnchainAttributesNotExistsInSchema, fmt.Sprintf("Attribute does not exist in schema: %s", err))
	}

	// Validate required attributes
	validated, requiredAttributeName := ValidateRequiredAttributes(schema.OnchainData.TokenAttributes, CreateNftAttrValueMap(data.OnchainAttributes))
	if !validated {
		return false, sdkerrors.Wrap(types.ErrRequiredAttributeMissing, requiredAttributeName)
	}

	// Validate Onchain Attributes Value
	duplicated, err := HasDuplicateNftAttributesValue(data.OnchainAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOnchainAttributesValue, fmt.Sprintf("Duplicate attribute name: %s", err))
	}

	// Validate Origin Attributes Value
	duplicated, err = HasDuplicateNftAttributesValue(data.OriginAttributes)
	if duplicated {
		return false, sdkerrors.Wrap(types.ErrDuplicateOriginAttributesValue, fmt.Sprintf("Duplicate attribute name: %s", err))
	}

	// Validate Origin Attributes Exist in Schema
	hasSameType, err := HasSameTypeAsSchema(mapAttributeDefinition, data.OriginAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrOriginAttributesNotSameTypeAsSchema, fmt.Sprintf("Does not have same type as schema: %s", err))
	}

	// Validate Onchain Attributes Exist in Schema
	hasSameType, err = HasSameTypeAsSchema(mergedMap, data.OnchainAttributes)
	if !hasSameType {
		return false, sdkerrors.Wrap(types.ErrOnchainTokenAttributesNotSameTypeAsSchema, fmt.Sprintf("Does not have same type as schema: %s", err))
	}
	return true, nil
}
