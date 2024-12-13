package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PerformMultiTokenAction(goCtx context.Context, msg *types.MsgPerformMultiTokenAction) (*types.MsgPerformMultiTokenActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	token_size := len(msg.TokenId)

	if token_size > 500 {
		return nil, sdkerrors.Wrap(types.ErrLimitSizeOfInput, "should be less than 500")
	}

	actoion_size := len(msg.Action)

	if actoion_size > 500 {
		return nil, sdkerrors.Wrap(types.ErrLimitSizeOfInput, "should be less than 500")
	}
	// //check if id in msg.TokenId is duplicate
	// mapOfTokenId := make(map[string]bool)
	// for _, tokenId := range msg.TokenId {
	// 	if _, ok := mapOfTokenId[tokenId]; ok {
	// 		return nil, sdkerrors.Wrap(types.ErrDuplicateInputTokenID, tokenId)
	// 	}
	// 	mapOfTokenId[tokenId] = true
	// }

	// ** SCHEMA LAYER **
	// check if schema exists
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	//check if all token exists
	for _, tokenId := range msg.TokenId {
		_, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, "Schema: "+msg.NftSchemaCode+" TokenID: "+tokenId)
		}
	}

	// ** This might be different from PerformActionByAdmin but to prevent time consuming process, we will use the same code out of iteration process **
	// check if executor is authorized to perform action
	var isOwner bool
	if msg.Creator == schema.Owner {
		isOwner = true
	}

	// if not owner, check if executor is authorized to perform action
	if !isOwner {

		_, isFound := k.Keeper.GetActionExecutor(
			ctx,
			msg.NftSchemaCode,
			msg.Creator,
		)

		if !isFound {
			return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
		}
	}

	arryOfparams := strings.Split(msg.Parameters[1:len(msg.Parameters)-1], "],")
	for i, params := range arryOfparams {
		if i != len(arryOfparams)-1 {
			arryOfparams[i] = params + "]"
		} else {
			arryOfparams[i] = params
		}
	}
	// switch case for action
	// case 1 len(action) == 1 && len(tokenid) == 1
	// case 2 len(action) == 1 && len(tokenid) > 1
	// case 3 len(action) > 1 && len(tokenid) > 1
	// case 4 len(action) > 1 && len(tokenid) == 1
	// default return error
	switch {
	case len(msg.Action) == 1 && len(msg.TokenId) == 1:
		// string to json object of  []*ActionParameter
		var actionPrams_ []*types.ActionParameter
		err := json.Unmarshal([]byte(arryOfparams[0]), &actionPrams_)
		if err != nil {
			fmt.Scanln("Error in Unmarshal parameter: ", err)
		}
		msg_ := &types.MsgPerformActionByAdmin{
			Creator:       msg.Creator,
			NftSchemaCode: msg.NftSchemaCode,
			TokenId:       msg.TokenId[0],
			Action:        msg.Action[0],
			RefId:         msg.RefId,
			Parameters:    actionPrams_,
		}
		_, success := k.PerformActionByAdmin(goCtx, msg_)
		if success != nil {
			return nil, success
		}
	case len(msg.Action) == 1 && len(msg.TokenId) > 1:
		var actionPrams_ []*types.ActionParameter
		err := json.Unmarshal([]byte(arryOfparams[0]), &actionPrams_)
		if err != nil {
			fmt.Scanln("Error in Unmarshal parameter: ", err)
		}
		msg_ := &types.MsgPerformMultiTokenOneAction{
			Creator:       msg.Creator,
			NftSchemaCode: msg.NftSchemaCode,
			TokenId:       msg.TokenId,
			Action:        msg.Action[0],
			RefId:         msg.RefId,
			Parameters:    actionPrams_,
		}
		_, success := k.PerformMultiTokenOneAction(goCtx, msg_)
		if success != nil {
			return nil, success
		}
	case len(msg.Action) > 1 && len(msg.TokenId) > 1:
		msg_ := &types.MsgPerformMultiTokenMultiAction{
			Creator:       msg.Creator,
			NftSchemaCode: msg.NftSchemaCode,
			TokenId:       msg.TokenId,
			Action:        msg.Action,
			RefId:         msg.RefId,
			Parameters:    arryOfparams,
		}
		_, success := k.PerformMultiTokenMultiAction(goCtx, msg_)
		if success != nil {
			return nil, success
		}
	case len(msg.Action) > 1 && len(msg.TokenId) == 1:
		msg_ := &types.MsgPerformOneTokenMultiAction{
			Creator:       msg.Creator,
			NftSchemaCode: msg.NftSchemaCode,
			TokenId:       msg.TokenId[0],
			Action:        msg.Action,
			RefId:         msg.RefId,
			Parameters:    arryOfparams,
		}
		_, success := k.PerformOneTokenMultiAction(goCtx, msg_)
		if success != nil {
			return nil, success
		}
	default:
		return nil, sdkerrors.Wrap(types.ErrInvalidActionInput, "Cannot execute Action and TokenId length")
	}

	return &types.MsgPerformMultiTokenActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
		Action:        msg.Action,
	}, nil
}

func (k msgServer) PerformMultiTokenOneAction(goCtx context.Context, msg *types.MsgPerformMultiTokenOneAction) (*types.MsgPerformMultiTokenOneActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	//check if id in msg.TokenId is duplicate
	mapOfTokenId := make(map[string]bool)
	for _, tokenId := range msg.TokenId {
		if _, ok := mapOfTokenId[tokenId]; ok {
			return nil, sdkerrors.Wrap(types.ErrDuplicateInputTokenID, tokenId)
		}
		mapOfTokenId[tokenId] = true
	}

	// ** SCHEMA LAYER **
	// check if schema exists
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	//check if all token exists
	for _, tokenId := range msg.TokenId {
		_, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, "Schema: "+msg.NftSchemaCode+" TokenID: "+tokenId)
		}
	}

	// ** This might be different from PerformActionByAdmin but to prevent time consuming process, we will use the same code out of iteration process **
	// check if executor is authorized to perform action
	var isOwner bool
	if msg.Creator == schema.Owner {
		isOwner = true
	}

	// if not owner, check if executor is authorized to perform action
	if !isOwner {

		_, isFound := k.Keeper.GetActionExecutor(
			ctx,
			msg.NftSchemaCode,
			msg.Creator,
		)

		if !isFound {
			return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
		}
	}

	mapAction := types.Action{}
	// Check if action is disabled
	action_, found := k.Keeper.GetActionOfSchema(ctx, msg.NftSchemaCode, msg.Action)
	if found {
		action := schema.OnchainData.Actions[action_.Index]
		if action.Disable {
			return nil, sdkerrors.Wrap(types.ErrActionIsDisabled, msg.Action)
		}
		mapAction = *action
	} else {
		return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, msg.Action)
	}

	// for _, action := range schema.OnchainData.Actions {
	// 	if action.Name == msg.Action && action.Disable {
	// 		return nil, sdkerrors.Wrap(types.ErrActionIsDisabled, action.Name)
	// 	}
	// 	if action.Name == msg.Action {
	// 		mapAction = *action
	// 		break
	// 	}
	// }

	// // Check if action exists
	// if mapAction.Name == "" {
	// 	return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, msg.Action)
	// }

	// Check if AllowedAction is for system
	if mapAction.GetAllowedActioner() == types.AllowedActioner_ALLOWED_ACTIONER_USER_ONLY {
		return nil, sdkerrors.Wrap(types.ErrActionIsForUserOnly, msg.Action)
	}

	// Check if action requires parameters
	param := mapAction.GetParams()
	required_param := make([]*types.ActionParams, 0)

	for _, p := range param {
		if p.Required {
			required_param = append(required_param, p)
		}
	}

	if len(required_param) > len(msg.Parameters) {
		return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "Input parameters length is not equal to required parameters length")
	}

	for i := 0; i < len(required_param); i++ {
		if msg.Parameters[i].Name != required_param[i].Name {
			return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "input paramter name is not match to "+required_param[i].Name)
		}
		if msg.Parameters[i].Value == "" {
			msg.Parameters[i].Value = required_param[i].DefaultValue
		}
	}

	// ** TOKEN DATA LAYER **
	// iterate over token ids
	for _, tokenId := range msg.TokenId {
		tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, msg.NftSchemaCode) //! This should not happen since we already check if all token exists
		}

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

		var list_schema_attributes_ []*types.SchemaAttribute
		var map_converted_schema_attributes []*types.NftAttributeValue

		// get schema attributes and convert to NFtAttributeValue
		all_schema_attributes := k.Keeper.GetAllSchemaAttribute(ctx)
		attributeMap := make(map[string]bool)

		for _, schema_attribute := range all_schema_attributes {
			if schema_attribute.NftSchemaCode != msg.NftSchemaCode {
				continue
			}
			// Check if the attribute has already been added
			if attributeMap[schema_attribute.Name] {
				continue
			}
			// Add the attribute to the list of schema attributes
			list_schema_attributes_ = append(list_schema_attributes_, &schema_attribute)

			// Add the attribute to the map
			attributeMap[schema_attribute.Name] = true

			nftAttributeValue_ := ConverSchemaAttributeToNFTAttributeValue(&schema_attribute)
			map_converted_schema_attributes = append(map_converted_schema_attributes, nftAttributeValue_)
		}

		// ** META path ../types/meta.go **
		meta := types.NewMetadata(&schema, &tokenData, schema.OriginData.AttributeOverriding, map_converted_schema_attributes)
		meta.SetGetNFTFunction(func(tokenId string) (*types.NftData, error) {
			tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
			if !found {
				return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, msg.NftSchemaCode)
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

		err := ProcessAction(meta, &mapAction, msg.Parameters)
		if err != nil {
			return nil, err
		}
		// Check if ChangeList is empty, error if empty
		if len(meta.ChangeList) == 0 {
			return nil, sdkerrors.Wrap(types.ErrEmptyChangeList, msg.Action)
		}

		// Update back to nftdata
		k.Keeper.SetNftData(ctx, tokenData)

		// Udpate to target
		// loop over meta.OtherUpdatedTokenDatas
		for _, otherTokenData := range meta.OtherUpdatedTokenDatas {
			k.Keeper.SetNftData(ctx, *otherTokenData)
		}

		for _, change := range meta.ChangeList {
			val, found := k.Keeper.GetSchemaAttribute(ctx, msg.NftSchemaCode, change.Key)
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

				k.Keeper.SetSchemaAttribute(ctx, val)
			}
		}

		// Emit events on metadata change
		// Check action with reference exists
		refId := msg.RefId + "_token-id_" + tokenId
		if msg.RefId != "" {

			_, found := k.Keeper.GetActionByRefId(ctx, refId)
			if found {
				return nil, sdkerrors.Wrap(types.ErrRefIdAlreadyExists, refId)
			}

			k.Keeper.SetActionByRefId(ctx, types.ActionByRefId{
				RefId:         refId,
				Creator:       msg.Creator,
				NftSchemaCode: msg.NftSchemaCode,
				TokenId:       tokenId,
				Action:        mapAction.Name,
			})
		}

		// Emit events on metadata change
		// appply change list to nftdata
		changeList, _ := json.Marshal(meta.ChangeList)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
				sdk.NewAttribute(types.AttributeKeyRunActionResult, "success"),
				sdk.NewAttribute(types.AttributeKeyTokenId, tokenId),
				sdk.NewAttribute(types.AttributeKeyRunActionRefId, refId),
				// Emit change list attributes
				sdk.NewAttribute(types.AttributeKeyRunActionChangeList, string(changeList)),
			),
		)

	}

	return &types.MsgPerformMultiTokenOneActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}

func (k msgServer) PerformMultiTokenMultiAction(goCtx context.Context, msg *types.MsgPerformMultiTokenMultiAction) (*types.MsgPerformMultiTokenMultiActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// check action len and parameters len are suitable
	if len(msg.Action) != len(msg.Parameters) {
		return nil, sdkerrors.Wrap(types.ErrActionAndParametersNotMatch, "Action: "+string(int32(len(msg.Action)))+" Parameters: "+string(int32(len(msg.Parameters))))
	}

	// ** SCHEMA LAYER **
	// check if schema exists
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	//check if all token exists
	for _, tokenId := range msg.TokenId {
		_, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, "Schema: "+msg.NftSchemaCode+" TokenID: "+tokenId)
		}
	}

	// ** This might be different from PerformActionByAdmin but to prevent time consuming process, we will use the same code out of iteration process **
	// check if executor is authorized to perform action
	var isOwner bool
	if msg.Creator == schema.Owner {
		isOwner = true
	}

	// if not owner, check if executor is authorized to perform action
	if !isOwner {

		_, isFound := k.Keeper.GetActionExecutor(
			ctx,
			msg.NftSchemaCode,
			msg.Creator,
		)

		if !isFound {
			return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
		}
	}

	mapAction := []types.Action{}
	for index, actionIter := range msg.Action {
		// Check if action is disabled
		action_, found := k.Keeper.GetActionOfSchema(ctx, msg.NftSchemaCode, actionIter)
		if found {
			action := schema.OnchainData.Actions[action_.Index]
			if action.Disable {
				return nil, sdkerrors.Wrap(types.ErrActionIsDisabled, actionIter)
			}
			mapAction = append(mapAction, *action)
		} else {
			return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, actionIter)
		}

		// for _, _action := range schema.OnchainData.Actions {
		// 	if _action.Name == actionIter && _action.Disable {
		// 		return nil, sdkerrors.Wrap(types.ErrActionIsDisabled, _action.Name)
		// 	}
		// 	if _action.Name == actionIter {
		// 		mapAction = append(mapAction, *_action)
		// 		break
		// 	}
		// }

		// // Check if action exists
		// if mapAction[index].Name == "" {
		// 	return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, actionIter)
		// }

		// Check if AllowedAction is for system
		if mapAction[index].GetAllowedActioner() == types.AllowedActioner_ALLOWED_ACTIONER_USER_ONLY {
			return nil, sdkerrors.Wrap(types.ErrActionIsForUserOnly, actionIter)
		}
	}

	for index, params_ := range msg.Parameters {
		// string to json object of  []*ActionParameter
		var actionPrams_ []*types.ActionParameter
		err := json.Unmarshal([]byte(params_), &actionPrams_)
		if err != nil {
			sdkerrors.Wrap(types.ErrInvalidParameter, "Error in Unmarshal required parameters ")
		}
		// Check if action requires parameters
		param := mapAction[index].GetParams()
		required_param := make([]*types.ActionParams, 0)

		for _, p := range param {
			if p.Required {
				required_param = append(required_param, p)
			}
		}

		if len(required_param) > len(msg.Parameters[index]) {
			return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "Input parameters length is not equal to required parameters length")
		}

		for i := 0; i < len(required_param); i++ {
			if actionPrams_[i].Name != required_param[i].Name {
				return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "input paramter name is not match to "+required_param[i].Name)
			}
			if actionPrams_[i].Value == "" {
				actionPrams_[i].Value = required_param[i].DefaultValue
			}
		}

	}

	// ** TOKEN DATA LAYER **
	// iterate over token ids
	for index, tokenId := range msg.TokenId {
		tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, msg.NftSchemaCode) //! This should not happen since we already check if all token exists
		}

		// unmarshal parameters
		var actionPrams_ []*types.ActionParameter
		err := json.Unmarshal([]byte(msg.Parameters[index]), &actionPrams_)
		if err != nil {
			sdkerrors.Wrap(types.ErrInvalidParameter, "Error in Unmarshal: TOKEN DATA LAYER cannot unmarshal parameters")
		}

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

		var list_schema_attributes_ []*types.SchemaAttribute
		var map_converted_schema_attributes []*types.NftAttributeValue

		// get schema attributes and convert to NFtAttributeValue
		all_schema_attributes := k.Keeper.GetAllSchemaAttribute(ctx)
		attributeMap := make(map[string]bool)

		for _, schema_attribute := range all_schema_attributes {
			if schema_attribute.NftSchemaCode != msg.NftSchemaCode {
				continue
			}
			// Check if the attribute has already been added
			if attributeMap[schema_attribute.Name] {
				continue
			}
			// Add the attribute to the list of schema attributes
			list_schema_attributes_ = append(list_schema_attributes_, &schema_attribute)

			// Add the attribute to the map
			attributeMap[schema_attribute.Name] = true

			nftAttributeValue_ := ConverSchemaAttributeToNFTAttributeValue(&schema_attribute)
			map_converted_schema_attributes = append(map_converted_schema_attributes, nftAttributeValue_)
		}

		// ** META path ../types/meta.go **
		meta := types.NewMetadata(&schema, &tokenData, schema.OriginData.AttributeOverriding, map_converted_schema_attributes)
		meta.SetGetNFTFunction(func(tokenId string) (*types.NftData, error) {
			tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
			if !found {
				return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, msg.NftSchemaCode)
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

		err = ProcessAction(meta, &mapAction[index], actionPrams_)
		if err != nil {
			return nil, err
		}
		// Check if ChangeList is empty, error if empty
		if len(meta.ChangeList) == 0 {
			return nil, sdkerrors.Wrap(types.ErrEmptyChangeList, msg.Action[index])
		}

		// Update back to nftdata
		k.Keeper.SetNftData(ctx, tokenData)

		// Udpate to target
		// loop over meta.OtherUpdatedTokenDatas
		for _, otherTokenData := range meta.OtherUpdatedTokenDatas {
			k.Keeper.SetNftData(ctx, *otherTokenData)
		}

		for _, change := range meta.ChangeList {
			val, found := k.Keeper.GetSchemaAttribute(ctx, msg.NftSchemaCode, change.Key)
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

				k.Keeper.SetSchemaAttribute(ctx, val)
			}
		}

		// Emit events on metadata change
		// Check action with reference exists
		refId := msg.RefId + "_token-id_" + tokenId + "_action-id_" + string(rune(index))
		if msg.RefId != "" {

			_, found := k.Keeper.GetActionByRefId(ctx, refId)
			if found {
				return nil, sdkerrors.Wrap(types.ErrRefIdAlreadyExists, refId)
			}

			k.Keeper.SetActionByRefId(ctx, types.ActionByRefId{
				RefId:         refId,
				Creator:       msg.Creator,
				NftSchemaCode: msg.NftSchemaCode,
				TokenId:       tokenId,
				Action:        mapAction[index].Name,
			})
		}

		// Emit events on metadata change
		// appply change list to nftdata
		changeList, _ := json.Marshal(meta.ChangeList)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
				sdk.NewAttribute(types.AttributeKeyRunActionResult, "success"),
				sdk.NewAttribute(types.AttributeKeyTokenId, tokenId),
				sdk.NewAttribute(types.AttributeKeyRunActionRefId, refId),
				// Emit change list attributes
				sdk.NewAttribute(types.AttributeKeyRunActionChangeList, string(changeList)),
			),
		)

	}

	return &types.MsgPerformMultiTokenMultiActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}

func (k msgServer) PerformOneTokenMultiAction(goCtx context.Context, msg *types.MsgPerformOneTokenMultiAction) (*types.MsgPerformOneTokenMultiActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// check action len and parameters len are suitable
	if len(msg.Action) != len(msg.Parameters) {
		return nil, sdkerrors.Wrap(types.ErrActionAndParametersNotMatch, "Action: "+string(int32(len(msg.Action)))+" Parameters: "+string(int32(len(msg.Parameters))))
	}

	// ** SCHEMA LAYER **
	// check if schema exists
	schema, found := k.Keeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrSchemaDoesNotExists, msg.NftSchemaCode)
	}

	//check if all token exists
	tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, msg.TokenId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, "Schema: "+msg.NftSchemaCode+" TokenID: "+msg.TokenId)
	}

	// ** This might be different from PerformActionByAdmin but to prevent time consuming process, we will use the same code out of iteration process **
	// check if executor is authorized to perform action
	var isOwner bool
	if msg.Creator == schema.Owner {
		isOwner = true
	}

	// if not owner, check if executor is authorized to perform action
	if !isOwner {

		_, isFound := k.Keeper.GetActionExecutor(
			ctx,
			msg.NftSchemaCode,
			msg.Creator,
		)

		if !isFound {
			return nil, sdkerrors.Wrap(types.ErrUnauthorized, msg.Creator)
		}
	}

	mapAction := []types.Action{}
	for index, actionIter := range msg.Action {
		// Check if action is disabled
		action_, found := k.Keeper.GetActionOfSchema(ctx, msg.NftSchemaCode, actionIter)
		if found {
			action := schema.OnchainData.Actions[action_.Index]
			if action.Disable {
				return nil, sdkerrors.Wrap(types.ErrActionIsDisabled, actionIter)
			}
			mapAction = append(mapAction, *action)
		} else {
			return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, actionIter)
		}

		// for _, _action := range schema.OnchainData.Actions {
		// 	if _action.Name == actionIter && _action.Disable {
		// 		return nil, sdkerrors.Wrap(types.ErrActionIsDisabled, _action.Name)
		// 	}
		// 	if _action.Name == actionIter {
		// 		mapAction = append(mapAction, *_action)
		// 		break
		// 	}
		// }

		// // Check if action exists
		// if mapAction[index].Name == "" {
		// 	return nil, sdkerrors.Wrap(types.ErrActionDoesNotExists, actionIter)
		// }

		// Check if AllowedAction is for system
		if mapAction[index].GetAllowedActioner() == types.AllowedActioner_ALLOWED_ACTIONER_USER_ONLY {
			return nil, sdkerrors.Wrap(types.ErrActionIsForUserOnly, actionIter)
		}
	}

	for index, params_ := range msg.Parameters {
		// string to json object of  []*ActionParameter
		var actionPrams_ []*types.ActionParameter
		err := json.Unmarshal([]byte(params_), &actionPrams_)
		if err != nil {
			sdkerrors.Wrap(types.ErrInvalidParameter, "Error in Unmarshal parameters index")
		}
		// Check if action requires parameters
		param := mapAction[index].GetParams()
		required_param := make([]*types.ActionParams, 0)

		for _, p := range param {
			if p.Required {
				required_param = append(required_param, p)
			}
		}

		if len(required_param) > len(msg.Parameters[index]) {
			return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "Input parameters length is not equal to required parameters length")
		}

		for i := 0; i < len(required_param); i++ {
			if actionPrams_[i].Name != required_param[i].Name {
				return nil, sdkerrors.Wrap(types.ErrInvalidParameter, "input paramter name is not match to "+required_param[i].Name)
			}
			if actionPrams_[i].Value == "" {
				actionPrams_[i].Value = required_param[i].DefaultValue
			}
		}

	}

	// ** TOKEN DATA LAYER **
	// iterate over token ids
	count := 0
	for index, msg_ := range msg.Action {
		_ = msg_ // unused
		// unmarshal parameters
		var actionPrams_ []*types.ActionParameter
		err := json.Unmarshal([]byte(msg.Parameters[index]), &actionPrams_)
		if err != nil {
			sdkerrors.Wrap(types.ErrInvalidParameter, "Error in Unmarshal: TOKEN DATA LAYER  index msg.Action")
		}

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

		var list_schema_attributes_ []*types.SchemaAttribute
		var map_converted_schema_attributes []*types.NftAttributeValue

		// get schema attributes and convert to NFtAttributeValue
		all_schema_attributes := k.Keeper.GetAllSchemaAttribute(ctx)
		attributeMap := make(map[string]bool)

		for _, schema_attribute := range all_schema_attributes {
			if schema_attribute.NftSchemaCode != msg.NftSchemaCode {
				continue
			}
			// Check if the attribute has already been added
			if attributeMap[schema_attribute.Name] {
				continue
			}
			// Add the attribute to the list of schema attributes
			list_schema_attributes_ = append(list_schema_attributes_, &schema_attribute)

			// Add the attribute to the map
			attributeMap[schema_attribute.Name] = true

			nftAttributeValue_ := ConverSchemaAttributeToNFTAttributeValue(&schema_attribute)
			map_converted_schema_attributes = append(map_converted_schema_attributes, nftAttributeValue_)
		}

		// ** META path ../types/meta.go **
		meta := types.NewMetadata(&schema, &tokenData, schema.OriginData.AttributeOverriding, map_converted_schema_attributes)
		meta.SetGetNFTFunction(func(tokenId string) (*types.NftData, error) {
			tokenData, found := k.Keeper.GetNftData(ctx, msg.NftSchemaCode, tokenId)
			if !found {
				return nil, sdkerrors.Wrap(types.ErrMetadataDoesNotExists, msg.NftSchemaCode)
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

		err = ProcessAction(meta, &mapAction[index], actionPrams_)
		if err != nil {
			return nil, err
		}
		// Check if ChangeList is empty, error if empty
		if len(meta.ChangeList) == 0 {
			return nil, sdkerrors.Wrap(types.ErrEmptyChangeList, msg.Action[index])
		}

		// Update back to nftdata
		k.Keeper.SetNftData(ctx, tokenData)

		// Udpate to target
		// loop over meta.OtherUpdatedTokenDatas
		for _, otherTokenData := range meta.OtherUpdatedTokenDatas {
			k.Keeper.SetNftData(ctx, *otherTokenData)
		}

		for _, change := range meta.ChangeList {
			val, found := k.Keeper.GetSchemaAttribute(ctx, msg.NftSchemaCode, change.Key)
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

				k.Keeper.SetSchemaAttribute(ctx, val)
			}
		}

		// Emit events on metadata change
		// Check action with reference exists
		refId := msg.RefId + "_action-id_" + string(int32(count))
		if msg.RefId != "" {

			_, found := k.Keeper.GetActionByRefId(ctx, refId)
			if found {
				return nil, sdkerrors.Wrap(types.ErrRefIdAlreadyExists, refId)
			}

			k.Keeper.SetActionByRefId(ctx, types.ActionByRefId{
				RefId:         refId,
				Creator:       msg.Creator,
				NftSchemaCode: msg.NftSchemaCode,
				TokenId:       msg.TokenId,
				Action:        mapAction[index].Name,
			})
		}

		// Emit events on metadata change
		// appply change list to nftdata
		changeList, _ := json.Marshal(meta.ChangeList)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
				sdk.NewAttribute(types.AttributeKeyRunActionResult, "success"),
				sdk.NewAttribute(types.AttributeKeyTokenId, msg.TokenId),
				sdk.NewAttribute(types.AttributeKeyRunActionRefId, refId),
				// Emit change list attributes
				sdk.NewAttribute(types.AttributeKeyRunActionChangeList, string(changeList)),
			),
		)
		count++
	}

	return &types.MsgPerformOneTokenMultiActionResponse{
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
		Action:        msg.Action,
	}, nil
}
