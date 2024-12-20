package keeper

import (
	"encoding/json"
	// "fmt"
	"strconv"
	"time"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SetupSchemaAndMetadata(ctx sdk.Context, schemaName, tokenId string) (*types.NFTSchema, *types.NftData, []*types.NftAttributeValue) {
	var (
		schema                    = types.NFTSchema{}
		tokenData                 = types.NftData{}
		convertedSchemaAttributes = []*types.NftAttributeValue{}
		mapExistingAttributes     = make(map[string]bool)
		attributeMap              = make(map[string]bool)
	)

	schema, found := k.GetNFTSchema(ctx, schemaName)
	if !found {
		return nil, nil, nil
	}

	tokenData, found = k.GetNftData(ctx, schemaName, tokenId)
	if !found {
		return nil, nil, nil
	}

	// Add missing attributes with default values
	for _, attribute := range tokenData.OnchainAttributes {
		mapExistingAttributes[attribute.Name] = true
	}

	// Loop over schema.TokenAttributes to check if exists in nftdata
	for _, attribute := range schema.OnchainData.TokenAttributes {
		if _, ok := mapExistingAttributes[attribute.Name]; !ok {
			if attribute.DefaultMintValue == nil {
				return nil, nil, nil
			}
			// Add attribute to nftdata with default value
			tokenData.OnchainAttributes = append(tokenData.OnchainAttributes,
				NewNFTAttributeValueFromDefaultValue(attribute.Name, attribute.DefaultMintValue))
		}
	}

	global_attributes := schema.OnchainData.NftAttributes

	for _, schema_attribute := range global_attributes {
		// Check if the attribute has already been added
		if attributeMap[schema_attribute.Name] {
			continue
		}

		nftAttributeValue, found := k.GetSchemaAttribute(ctx, schema.Code, schema_attribute.Name)

		if !found {
			return nil, nil, nil
		}

		// Add the attribute to the map
		attributeMap[schema_attribute.Name] = true

		nftAttributeValue_ := ConverSchemaAttributeToNFTAttributeValue(&nftAttributeValue)
		convertedSchemaAttributes = append(convertedSchemaAttributes, nftAttributeValue_)
	}

	return &schema, &tokenData, convertedSchemaAttributes
}

func (k Keeper) ActionByAdmin(ctx sdk.Context, creator, nftSchemaName, tokenId, actionName, refId string, parameters []*types.ActionParameter) (changelist []byte, err error) {
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	var isOwner bool

	if creator == schema.Owner {
		isOwner = true
	}

	// if not owner, check if executor is authorized to perform action
	if !isOwner {

		_, isFound := k.GetActionExecutor(
			ctx,
			nftSchemaName,
			creator,
		)

		if !isFound {
			return nil, sdkerrors.Wrap(types.ErrUnauthorized, creator)
		}
	}

	mapAction := types.Action{}

	// Check if action is disabled
	action_, found := k.GetActionOfSchema(ctx, nftSchemaName, actionName)
	if found {
		action := schema.OnchainData.Actions[action_.Index]
		if action.Disable {
			return nil, sdkerrors.Wrap(types.ErrActionIsDisabled, actionName)
		}
		mapAction = *action
	} else {
		return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, actionName)
	}

	if mapAction.GetAllowedActioner() == types.AllowedActioner_ALLOWED_ACTIONER_USER_ONLY {
		return nil, sdkerrors.Wrap(types.ErrActionIsForUserOnly, actionName)
	}

	// Check if action requires parameters
	param := mapAction.GetParams()
	required_param := make([]*types.ActionParams, 0)

	for _, p := range param {
		if p.Required {
			required_param = append(required_param, p)
		}
	}

	if len(required_param) > len(parameters) {
		return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "Input parameters length is not equal to required parameters length")
	}

	for i := 0; i < len(required_param); i++ {
		if parameters[i].Name != required_param[i].Name {
			return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "input paramter name is not match to "+required_param[i].Name)
		}
		if parameters[i].Value == "" {
			parameters[i].Value = required_param[i].DefaultValue
		}
	}

	tokenData, found := k.GetNftData(ctx, nftSchemaName, tokenId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, "Schema: "+nftSchemaName+" TokenID: "+tokenId)
	}

	// ** TOKEN DATA LAYER **
	// Create map of existing attribute in nftdata
	mapExistingAttributes := make(map[string]bool)
	for _, attribute := range tokenData.OnchainAttributes {
		mapExistingAttributes[attribute.Name] = true
	}

	// Loop over schema.TokenAttributes to check if exists in nftdata
	for _, attribute := range schema.OnchainData.TokenAttributes {
		if _, ok := mapExistingAttributes[attribute.Name]; !ok {
			if attribute.DefaultMintValue == nil {
				return nil, sdkerrors.Wrap(types.ErrNoDefaultValue, attribute.Name)
			}
			// Add attribute to nftdata with default value
			tokenData.OnchainAttributes = append(tokenData.OnchainAttributes,
				NewNFTAttributeValueFromDefaultValue(attribute.Name, attribute.DefaultMintValue))
		}
	}

	var map_converted_schema_attributes []*types.NftAttributeValue

	global_attributes := schema.OnchainData.NftAttributes

	attributeMap := make(map[string]bool)

	for _, schema_attribute := range global_attributes {
		// Check if the attribute has already been added
		if attributeMap[schema_attribute.Name] {
			continue
		}

		nftAttributeValue, found := k.GetSchemaAttribute(ctx, schema.Code, schema_attribute.Name)

		if !found {
			return nil, sdkerrors.Wrap(types.ErrNoDefaultValue, schema_attribute.Name+" NOT FOUND")
		}

		// Add the attribute to the map
		attributeMap[schema_attribute.Name] = true

		nftAttributeValue_ := ConverSchemaAttributeToNFTAttributeValue(&nftAttributeValue)
		map_converted_schema_attributes = append(map_converted_schema_attributes, nftAttributeValue_)
	}

	meta := types.NewMetadata(&schema, &tokenData, schema.OriginData.AttributeOverriding, map_converted_schema_attributes)
	meta.SetGetNFTFunction(func(tokenId string) (*types.NftData, error) {
		tokenData, found := k.GetNftData(ctx, nftSchemaName, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, nftSchemaName)
		}
		return &tokenData, nil
	})

	// utils function
	meta.SetGetBlockTimeFunction(func() time.Time {
		return ctx.BlockTime()
	})

	// utils function
	meta.SetGetBlockHeightFunction(func() int64 {
		return ctx.BlockHeight()
	})

	err = ProcessAction(meta, &mapAction, parameters)
	if err != nil {
		return nil, err
	}

	// Check if ChangeList is empty, error if empty
	if len(meta.ChangeList) == 0 {
		return nil, sdkerrors.Wrap(types.ErrEmptyChangeList, actionName)
	}

	// Update back to nftdata
	k.SetNftData(ctx, tokenData)

	// Udpate to target
	// loop over meta.OtherUpdatedTokenDatas
	for _, otherTokenData := range meta.OtherUpdatedTokenDatas {
		k.SetNftData(ctx, *otherTokenData)
	}

	for _, change := range meta.ChangeList {
		val, found := k.GetSchemaAttribute(ctx, nftSchemaName, change.Key)
		if found {
			switch val.DataType {
			case "string":
				val.CurrentValue.Value = &types.SchemaAttributeValue_StringAttributeValue{
					StringAttributeValue: &types.StringAttributeValue{
						Value: change.NewValue,
					},
				}
			case "boolean":
				boolValue, err := strconv.ParseBool(change.NewValue)
				if err != nil {
					return nil, err
				}
				val.CurrentValue.Value = &types.SchemaAttributeValue_BooleanAttributeValue{
					BooleanAttributeValue: &types.BooleanAttributeValue{
						Value: boolValue,
					},
				}
			case "number":
				uintValue, err := strconv.ParseUint(change.NewValue, 10, 64)
				if err != nil {
					return nil, err
				}
				val.CurrentValue.Value = &types.SchemaAttributeValue_NumberAttributeValue{
					NumberAttributeValue: &types.NumberAttributeValue{
						Value: uintValue,
					},
				}
			case "float":
				floatValue, err := strconv.ParseFloat(change.NewValue, 64)
				if err != nil {
					return nil, err
				}
				val.CurrentValue.Value = &types.SchemaAttributeValue_FloatAttributeValue{
					FloatAttributeValue: &types.FloatAttributeValue{
						Value: floatValue,
					},
				}
			default:
				return nil, sdkerrors.Wrap(types.ErrParsingAttributeValue, val.DataType)
			}

			k.SetSchemaAttribute(ctx, val)
		}
	}

	// Check action with reference exists
	if refId != "" {

		_, found := k.GetActionByRefId(ctx, refId)
		if found {
			return nil, sdkerrors.Wrap(types.ErrRefIdAlreadyExists, refId)
		}

		k.SetActionByRefId(ctx, types.ActionByRefId{
			RefId:         refId,
			Creator:       creator,
			NftSchemaCode: nftSchemaName,
			TokenId:       tokenId,
			Action:        mapAction.Name,
		})
	}

	changeList, _ := json.Marshal(meta.ChangeList)

	return changeList, nil
}

func (k Keeper) AddActionKeeper(ctx sdk.Context, creator string, nftSchemaName string, newAction types.Action) error {
	// get existing action in schema
	schema, schemaFound := k.GetNFTSchema(ctx, nftSchemaName)
	if !schemaFound {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	if creator != schema.Owner {
		return sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	// validate Action data
	err := ValidateAction(&newAction, &schema)
	if err != nil {
		return sdkerrors.Wrap(types.ErrValidatingMetadata, err.Error())
	}

	// append new action
	schema.OnchainData.Actions = append(schema.OnchainData.Actions, &newAction)

	// save index of action
	k.SetActionOfSchema(ctx, types.ActionOfSchema{
		Name:          newAction.Name,
		NftSchemaCode: schema.Code,
		Index:         uint64(len(schema.OnchainData.Actions) - 1),
	})

	// save schema
	k.SetNFTSchema(ctx, schema)

	return nil
}

func (k Keeper) UpdateActionKeeper(ctx sdk.Context, creator, nftSchemaName string, updateAction types.Action) error {
	// get existing action
	actionOfSchema, found := k.GetActionOfSchema(ctx, nftSchemaName, updateAction.Name)
	if !found {
		return sdkerrors.Wrap(types.ErrActionDoesNotExists, updateAction.Name)
	}

	// get existing nft schema
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	// updator is valid
	if creator != schema.Owner {
		return sdkerrors.Wrap(types.ErrUnauthorized, creator)
	}

	// update action by its index
	schema.OnchainData.Actions[actionOfSchema.Index] = &updateAction

	// update schema
	k.SetNFTSchema(ctx, schema)
	return nil
}

func (k Keeper) ToggleActionKeeper(ctx sdk.Context, creator, nftSchemaName, actionName string, status bool) error {
	schema, found := k.GetNFTSchema(ctx, nftSchemaName)
	if !found {
		return sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}
	// Check if creator is owner of schema
	if creator != schema.Owner {
		return sdkerrors.Wrap(types.ErrCreatorDoesNotMatch, creator)
	}

	// Update is_active in schema
	for i, action := range schema.OnchainData.Actions {
		if action.Name == actionName {
			schema.OnchainData.Actions[i].Disable = status
		}
	}

	k.SetNFTSchema(ctx, schema)

	return nil
}

func (k Keeper) PerfromVirtualAction(ctx sdk.Context, creator, nftSchemaName, tokenId, actionName, refId string, parameters []*types.ActionParameter) (changeList [][]byte, err error) {
	var (
		change_list            = [][]byte{}
		schemaList             = []*types.NFTSchema{}
		tokenDataList          = []*types.NftData{}
		crossSchemaOveride     = types.CrossSchemaAttributeOverriding{}
		schemaGlobalAttributes = types.CrossSchemaGlobalAttributes{}
		shareAttributeName     = types.CrossSchemaSharedAttributeName{}
	)

	// get virtual schema
	virtualSchema, found := k.GetVirtualSchema(ctx, nftSchemaName)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, nftSchemaName)
	}

	// get virtual action
	vitualAction, found := k.GetVirtualAction(ctx, nftSchemaName, actionName)
	if found {
		if vitualAction.Disable {
			return nil, sdkerrors.Wrap(types.ErrActionIsDisabled, actionName)
		}
	} else {
		return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, actionName)
	}

	// Check if action requires parameters
	param := vitualAction.GetParams()
	required_param := make([]*types.ActionParams, 0)
	for _, p := range param {
		if p.Required {
			required_param = append(required_param, p)
		}
	}

	if len(required_param) > len(parameters) {
		return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "Input parameters length is not equal to required parameters length")
	}

	for i := 0; i < len(required_param); i++ {
		if parameters[i].Name != required_param[i].Name {
			return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "input paramter name is not match to "+required_param[i].Name)
		}
		if parameters[i].Value == "" {
			parameters[i].Value = required_param[i].DefaultValue
		}
	}

	// get schema component
	for _, schemaRegistry := range virtualSchema.Registry {
		schema, tokenData, convertedSchemaAttributes := k.SetupSchemaAndMetadata(ctx, schemaRegistry.NftSchemaCode, tokenId)
		if (schema == nil) || (tokenData == nil) || (convertedSchemaAttributes == nil) {
			return changeList, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, nftSchemaName)
		}
		schemaList = append(schemaList, schema)
		tokenDataList = append(tokenDataList, tokenData)
		crossSchemaOveride[schema.Code] = schema.OriginData.AttributeOverriding
		schemaGlobalAttributes[schema.Code] = convertedSchemaAttributes
		shareAttributeName[schema.Code] = schemaRegistry.SharedAttributes
	}

	crossMetadata := types.NewCrossSchemaMetadata(schemaList, tokenDataList, crossSchemaOveride, schemaGlobalAttributes, shareAttributeName)

	err = ProcessCrossSchemaAction(crossMetadata, vitualAction.ToAction(), parameters)
	if err != nil {
		return nil, err
	}

	someValueChange := false
	// Check if ChangeList is empty, error if empty
	for _, schemaRegistry := range virtualSchema.Registry {
		if len(crossMetadata.GetChangeList(schemaRegistry.NftSchemaCode)) > 0 {
			someValueChange = true
		}
	}

	if !someValueChange {
		return nil, sdkerrors.Wrap(types.ErrEmptyChangeList, actionName)
	}

	for _, schemaRegistry := range virtualSchema.Registry {
		k.SetNftData(ctx, *crossMetadata.GetNftData(schemaRegistry.NftSchemaCode))

		for _, change := range crossMetadata.GetChangeList(schemaRegistry.NftSchemaCode) {
			val, found := k.GetSchemaAttribute(ctx, nftSchemaName, change.Key)
			if found {
				switch val.DataType {
				case "string":
					val.CurrentValue.Value = &types.SchemaAttributeValue_StringAttributeValue{
						StringAttributeValue: &types.StringAttributeValue{
							Value: change.NewValue,
						},
					}
				case "boolean":
					boolValue, err := strconv.ParseBool(change.NewValue)
					if err != nil {
						return nil, err
					}
					val.CurrentValue.Value = &types.SchemaAttributeValue_BooleanAttributeValue{
						BooleanAttributeValue: &types.BooleanAttributeValue{
							Value: boolValue,
						},
					}
				case "number":
					uintValue, err := strconv.ParseUint(change.NewValue, 10, 64)
					if err != nil {
						return nil, err
					}
					val.CurrentValue.Value = &types.SchemaAttributeValue_NumberAttributeValue{
						NumberAttributeValue: &types.NumberAttributeValue{
							Value: uintValue,
						},
					}
				case "float":
					floatValue, err := strconv.ParseFloat(change.NewValue, 64)
					if err != nil {
						return nil, err
					}
					val.CurrentValue.Value = &types.SchemaAttributeValue_FloatAttributeValue{
						FloatAttributeValue: &types.FloatAttributeValue{
							Value: floatValue,
						},
					}
				default:
					return nil, sdkerrors.Wrap(types.ErrParsingAttributeValue, val.DataType)
				}

				k.SetSchemaAttribute(ctx, val)
			}
		}

		individualChangeList, _ := json.Marshal(crossMetadata.GetChangeList(schemaRegistry.NftSchemaCode))
		change_list = append(change_list, individualChangeList)
	}

	// Check action with reference exists
	if refId != "" {

		_, found := k.GetActionByRefId(ctx, refId)
		if found {
			return nil, sdkerrors.Wrap(types.ErrRefIdAlreadyExists, refId)
		}

		k.SetActionByRefId(ctx, types.ActionByRefId{
			RefId:         refId,
			Creator:       creator,
			NftSchemaCode: nftSchemaName,
			TokenId:       tokenId,
			Action:        vitualAction.Name,
		})
	}

	return change_list, nil
}
