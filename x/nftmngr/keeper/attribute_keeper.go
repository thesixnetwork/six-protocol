package keeper

import (
	"context"
	"strconv"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	errormod "cosmossdk.io/errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) AddAttributeKeeper(ctx context.Context, creator string, nftSchemaName string, new_add_attribute types.AttributeDefinition, location types.AttributeLocation) error {
	// get existing nft schema
	schema, schemaFound := k.GetNFTSchema(ctx, nftSchemaName)
	if !schemaFound {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	// validate AttributeDefinition data
	err := k.ValidateAttributeDefinition(ctx, &new_add_attribute, &schema)
	if err != nil {
		return errormod.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	err = ValidateAttributeNames([]*types.AttributeDefinition{&new_add_attribute})
	if err != nil {
		return errormod.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// Swith location of attribute
	switch location {
	case types.AttributeLocation_NFT_ATTRIBUTE:

		_defaultMintValue, err := ConvertDefaultMintValueToSchemaAttributeValue(new_add_attribute.DefaultMintValue)
		if err != nil {
			return errormod.Wrap(types.ErrParsingMetadataMessage, err.Error())
		}

		// this case will use Msg.CreateSchemaAtribute
		k.SetSchemaAttribute(ctx, types.SchemaAttribute{
			NftSchemaCode: nftSchemaName,
			Name:          new_add_attribute.Name,
			CurrentValue:  _defaultMintValue,
			DataType:      new_add_attribute.DataType,
			Creator:       creator,
		})

		schema.OnchainData.NftAttributes = append(schema.OnchainData.NftAttributes, &new_add_attribute)

	case types.AttributeLocation_TOKEN_ATTRIBUTE:
		schema.OnchainData.TokenAttributes = append(schema.OnchainData.TokenAttributes, &new_add_attribute)

	}

	count := MergeAndCountAllAttributes(schema.OriginData.OriginAttributes, schema.OnchainData.NftAttributes, schema.OnchainData.TokenAttributes)
	// set new index to new attribute
	new_add_attribute.Index = uint64(count - 1)

	// set schema
	k.SetNFTSchema(ctx, schema)

	return nil
}

func (k Keeper) UpdateAttributeKeeper(ctx context.Context, creator, nftSchemaName string, update_attribute types.AttributeDefinition) error {
	// Check if the value exists
	valFound, isFound := k.GetSchemaAttribute(ctx, nftSchemaName, update_attribute.Name)
	if !isFound {
		return errormod.Wrap(sdkerrors.ErrKeyNotFound, "Attribute not exists in schema")
	}

	// get existing nft schema
	schema, schemaFound := k.GetNFTSchema(ctx, nftSchemaName)
	if !schemaFound {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	err := ValidateAttributeNames([]*types.AttributeDefinition{&update_attribute})
	if err != nil {
		return errormod.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// parse DefaultMintValue to SchemaAttributeValue
	schmaAttributeValue, err := ConvertDefaultMintValueToSchemaAttributeValue(update_attribute.DefaultMintValue)
	if err != nil {
		return errormod.Wrap(types.ErrParsingMetadataMessage, err.Error())
	}

	schemaAttribute := types.SchemaAttribute{
		Creator:       creator,
		NftSchemaCode: valFound.NftSchemaCode,
		Name:          valFound.Name,
		DataType:      update_attribute.DataType,
		CurrentValue:  schmaAttributeValue,
	}

	k.SetSchemaAttribute(ctx, schemaAttribute)

	return nil
}

func (k Keeper) ResyncAttibutesKeeper(ctx context.Context, creator, nftSchemaName, tokenId string) error {
	// Retrieve schema
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	// Check if creator is owner of schema
	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	// Retrieve NFT Data
	nftData, found := k.GetNftData(ctx, nftSchemaName, tokenId)
	if !found {
		return errormod.Wrap(types.ErrNftDataDoesNotExists, tokenId)
	}

	// Create map of existing attribute in nftdata
	mapExistingAttributes := make(map[string]bool)
	for _, attribute := range nftData.OnchainAttributes {
		mapExistingAttributes[attribute.Name] = true
	}

	// Loop over schema.TokenAttributes to check if exists in nftdata
	for _, attribute := range schema.OnchainData.TokenAttributes {
		if _, ok := mapExistingAttributes[attribute.Name]; !ok {
			if attribute.DefaultMintValue == nil {
				return errormod.Wrap(types.ErrNoDefaultValue, attribute.Name)
			}

			// Add attribute to nftdata with default value
			nftData.OnchainAttributes = append(nftData.OnchainAttributes, NewNFTAttributeValueFromDefaultValue(attribute.Name, attribute.DefaultMintValue))
		}
	}

	// Set nftdata
	k.SetNftData(ctx, nftData)

	return nil
}

func (k Keeper) SetAttributeOveridingKeeper(ctx context.Context, creator, nftSchemaName string, newOveridingType int32) error {
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	switch newOveridingType {
	case 0:
		schema.OriginData.AttributeOverriding = types.AttributeOverriding_ORIGIN
	case 1:
		schema.OriginData.AttributeOverriding = types.AttributeOverriding_ORIGIN
	default:
		return errormod.Wrap(types.ErrAttributeOptionDoesNotExists, strconv.Itoa(int(newOveridingType)))
	}

	k.SetNFTSchema(ctx, schema)

	return nil
}

func (k Keeper) ShowAttributeKeeper(ctx context.Context, creator, nftSchemaName string, status bool, attributesName []string) error {
	// Retrieve schema
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	// Check if creator is the owner of the schema
	if creator != schema.Owner {
		return errormod.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	type ReadAttribute struct {
		AttributeDefinition *types.AttributeDefinition
		AttributeName       string
		AttrributeIndex     int
		AttributeLocation   types.AttributeLocation
	}
	// Create map of nftattributes and tokenattributes ni ReadAttribute struct
	mapReadAttribute := make(map[string]ReadAttribute)
	for i, nftAttribute := range schema.OnchainData.NftAttributes {
		mapReadAttribute[nftAttribute.Name] = ReadAttribute{
			AttributeName:       nftAttribute.Name,
			AttributeDefinition: nftAttribute,
			AttributeLocation:   types.AttributeLocation_NFT_ATTRIBUTE,
			AttrributeIndex:     i,
		}
	}

	for i, tokenAttribute := range schema.OnchainData.TokenAttributes {
		mapReadAttribute[tokenAttribute.Name] = ReadAttribute{
			AttributeName:       tokenAttribute.Name,
			AttributeDefinition: tokenAttribute,
			AttributeLocation:   types.AttributeLocation_TOKEN_ATTRIBUTE,
			AttrributeIndex:     i,
		}
	}

	// loop over msg.AttributeNames
	for _, attributeName := range attributesName {
		// check if attribute is exist in mapReadAttribute
		if _, ok := mapReadAttribute[attributeName]; !ok {
			return errormod.Wrap(types.ErrAttributeDoesNotExists, attributeName)
		}
		readAttributeDef := mapReadAttribute[attributeName]
		if readAttributeDef.AttributeLocation == types.AttributeLocation_NFT_ATTRIBUTE {
			schema.OnchainData.NftAttributes[readAttributeDef.AttrributeIndex].HiddenToMarketplace = !status
		} else {
			schema.OnchainData.TokenAttributes[readAttributeDef.AttrributeIndex].HiddenToMarketplace = !status
		}
	}

	k.SetNFTSchema(ctx, schema)

	return nil
}

func (k Keeper) GetAttributeValue(ctx context.Context, nftSchemaName, tokenId, attributeName string) (string, error) {
	var (
		attributeValue   string
		attributeFound   = false
		isTokenAttribute = false
	)

	_, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return "", errormod.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	// Check if the data already exists
	metadata, found := k.GetNftData(ctx, nftSchemaName, tokenId)
	if !found {
		return "", errormod.Wrap(types.ErrMetadataDoesNotExists, nftSchemaName)
	}

	for _, ocAttribute := range metadata.OnchainAttributes {
		if ocAttribute.Name == attributeName {
			isTokenAttribute = true
			attributeFound = true
			// convert attribute value to string and assing to attributeValue
			attributeValue = ocAttribute.AttributeValueToString()
		}
	}

	if !isTokenAttribute {
		// find schema attribute value
		schemaAttribute, found := k.GetSchemaAttribute(ctx, nftSchemaName, attributeName)
		if found {
			attributeFound = true
			attributeValue = schemaAttribute.CurrentValueToString()
		}
	}

	if !attributeFound {
		return "", errormod.Wrap(types.ErrAttributeDoesNotExists, attributeName)
	}

	return attributeValue, nil
}
