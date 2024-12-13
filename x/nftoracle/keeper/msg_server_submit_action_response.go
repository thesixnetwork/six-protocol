package keeper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	nftmngrkeeper "github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	nftmngrtypes "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	utils "github.com/thesixnetwork/six-protocol/utils"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

func (k msgServer) SubmitActionResponse(goCtx context.Context, msg *types.MsgSubmitActionResponse) (*types.MsgSubmitActionResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oracle, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionOracle, oracle)
	if !granted {
		return nil, sdkerrors.Wrap(types.ErrNoOraclePermission, msg.Creator)
	}

	actionRequest, found := k.GetActionRequest(ctx, msg.ActionRequestID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrMintRequestNotFound, strconv.FormatUint(msg.ActionRequestID, 10))
	}

	// Check if mint request is still pending
	if actionRequest.Status != types.RequestStatus_PENDING {
		return nil, sdkerrors.Wrap(types.ErrActionRequestNotPending, strconv.FormatUint(msg.ActionRequestID, 10))
	}

	// Check if currernt confirmation count is less than required confirmation count
	if actionRequest.CurrentConfirm >= actionRequest.RequiredConfirm {
		return nil, sdkerrors.Wrap(types.ErrActionRequestConfirmedAlreadyComplete, strconv.FormatUint(msg.ActionRequestID, 10))
	}

	// Convert msg.Base64NftMetadata to bytes
	nftMetadata, err := base64.StdEncoding.DecodeString(msg.Base64NftData)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	nftOriginData := types.NftOriginData{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(nftMetadata, &nftOriginData)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingOriginData, err.Error())
	}

	if nftOriginData.HolderAddress == "" || nftOriginData.Image == "" || nftOriginData.Traits == nil || len(nftOriginData.Traits) == 0 {
		return nil, sdkerrors.Wrap(types.ErrParsingOriginData, "missing required fields")
	}

	if actionRequest.CurrentConfirm == 0 {
		// Create sha512 hash of nftMetadata
		dataHash := sha256.Sum256(nftMetadata)
		actionRequest.DataHashes = append(actionRequest.DataHashes, &types.DataHash{
			OriginData: &nftOriginData,
			Hash:       dataHash[:],
			Confirmers: []string{msg.Creator},
		})
	} else {
		// Check if creator has already confirmed this mint request
		// fetch confirmed data
		for _, confirmer := range actionRequest.Confirmers {
			if confirmer == msg.Creator {
				return nil, sdkerrors.Wrap(types.ErrOracleConfirmedAlready, strconv.FormatUint(msg.ActionRequestID, 10))
			}

		}
		// if _, ok := actionRequest.Confirmers[msg.Creator]; ok {
		// 	return nil, sdkerrors.Wrap(types.ErrOracleConfirmedAlready, strconv.FormatUint(msg.ActionRequestID, 10)+", "+msg.Creator)
		// }
		// Compare data hash with previous data hash
		dataHash := sha256.Sum256(nftMetadata)
		dataHashMatch := false
		for _, hash := range actionRequest.DataHashes {
			if res := bytes.Compare(dataHash[:], hash.Hash); res == 0 {
				dataHashMatch = true
				hash.Confirmers = append(hash.Confirmers, msg.Creator)
				break
			}
		}
		if !dataHashMatch {
			actionRequest.DataHashes = append(actionRequest.DataHashes, &types.DataHash{
				OriginData: &nftOriginData,
				Hash:       dataHash[:],
				Confirmers: []string{msg.Creator},
			})
		}
	}
	if actionRequest.Confirmers == nil {
		actionRequest.Confirmers = make([]string, 0)
	}
	// Mark creator as confirmed
	// actionRequest.Confirmers[msg.Creator] = true
	actionRequest.Confirmers = append(actionRequest.Confirmers, msg.Creator)

	// increase actionRequest.CurrentConfirm
	actionRequest.CurrentConfirm++

	if actionRequest.CurrentConfirm == actionRequest.RequiredConfirm {
		// Mint NFT
		// Check if there is only one data hash
		if len(actionRequest.DataHashes) > 1 {
			// Update mintRequest.Status to be FAILED
			actionRequest.Status = types.RequestStatus_FAILED_WITHOUT_CONCENSUS
		} else {
			// Update mintRequest.Status to be SUCCESS
			actionRequest.Status = types.RequestStatus_SUCCESS_WITH_CONSENSUS
		}

		// Emit event a consensus result
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeMintRequestConfirmed,
				sdk.NewAttribute(types.AttributeKeyActionRequestID, strconv.FormatUint(actionRequest.Id, 10)),
				sdk.NewAttribute(types.AttributeKeyActionRequestStatus, actionRequest.Status.String()),
			),
		)
		if actionRequest.Status == types.RequestStatus_SUCCESS_WITH_CONSENSUS {
			err = nil
			// Udpate NFT Data
			nftData, found := k.nftmngrKeeper.GetNftData(ctx, actionRequest.NftSchemaCode, actionRequest.TokenId)
			if !found {
				return nil, sdkerrors.Wrap(types.ErrMetaDataNotFound, actionRequest.NftSchemaCode+":"+actionRequest.TokenId)
			}
			//ctx sdk.Context, nftData nftmngrtypes.NftData, originData *types.NftOriginData
			err = k.UpdateMetaDataFromOriginData(ctx, &nftData, actionRequest.DataHashes[0].OriginData)
			if err != nil {
				actionRequest.Status = types.RequestStatus_FAILED_ON_EXECUTION
				actionRequest.ExecutionErrorMessage = err.Error()
			}
			// Perform Action
			if err == nil {
				err = k.PerformAction(ctx, &actionRequest, &nftData)
				if err != nil {
					actionRequest.Status = types.RequestStatus_FAILED_ON_EXECUTION
					actionRequest.ExecutionErrorMessage = err.Error()
				}
			}
		}
	}

	// Update Mint Request Back to storage
	k.SetActionRequest(ctx, actionRequest)

	return &types.MsgSubmitActionResponseResponse{}, nil
}

func (k msgServer) PerformAction(ctx sdk.Context, actionRequest *types.ActionOracleRequest, tokenData *nftmngrtypes.NftData) error {
	schema, found := k.nftmngrKeeper.GetNFTSchema(ctx, actionRequest.NftSchemaCode)
	if !found {
		return sdkerrors.Wrap(types.ErrNFTSchemaNotFound, actionRequest.NftSchemaCode)
	}

	if actionRequest.Caller != actionRequest.DataHashes[0].OriginData.HolderAddress {
		return sdkerrors.Wrap(types.ErrUnauthorizedCaller, actionRequest.Caller)
	}

	mapAction := nftmngrtypes.Action{}
	// Check if action is disabled
	action_, found := k.nftmngrKeeper.GetActionOfSchema(ctx, actionRequest.NftSchemaCode, actionRequest.Action)
	if found {
		action := schema.OnchainData.Actions[action_.Index]
		if action.Disable {
			return sdkerrors.Wrap(nftmngrtypes.ErrActionIsDisabled, actionRequest.Action)
		}
		mapAction = *action
	} else {
		return sdkerrors.Wrap(nftmngrtypes.ErrActionDoesNotExists, actionRequest.Action)
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

	// Check if AllowedAction is for user
	if mapAction.GetAllowedActioner() == nftmngrtypes.AllowedActioner_ALLOWED_ACTIONER_SYSTEM_ONLY {
		return sdkerrors.Wrap(nftmngrtypes.ErrActionIsForSystemOnly, mapAction.Name)
	}

	// Check if action requires parameters
	param := mapAction.GetParams()
	required_param := make([]*nftmngrtypes.ActionParams, 0)

	for _, p := range param {
		if p.Required {
			required_param = append(required_param, p)
		}
	}

	if len(required_param) > len(actionRequest.Params) {
		return sdkerrors.Wrap(nftmngrtypes.ErrInvalidParameter, "Input parameters length is not equal to required parameters length")
	}

	for i := 0; i < len(required_param); i++ {
		if actionRequest.Params[i].Name != required_param[i].Name {
			return sdkerrors.Wrap(nftmngrtypes.ErrInvalidParameter, "input paramter name is not match to "+required_param[i].Name)
		}
		if actionRequest.Params[i].Value == "" {
			actionRequest.Params[i].Value = required_param[i].DefaultValue
		}
	}

	input_param := make([]*nftmngrtypes.ActionParameter, 0)
	for _, p := range actionRequest.Params {
		input_param = append(input_param, &nftmngrtypes.ActionParameter{
			Name:  p.Name,
			Value: p.Value,
		})
	}

	var list_schema_attributes_ []*nftmngrtypes.SchemaAttribute
	var map_converted_schema_attributes []*nftmngrtypes.NftAttributeValue

	// get schema attributes and convert to NFtAttributeValue
	all_schema_attributes := k.nftmngrKeeper.GetAllSchemaAttribute(ctx)

	attributeMap := make(map[string]bool)

	for _, schema_attribute := range all_schema_attributes {
		if schema_attribute.NftSchemaCode != actionRequest.NftSchemaCode {
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

		nftAttributeValue_ := nftmngrkeeper.ConverSchemaAttributeToNFTAttributeValue(&schema_attribute)
		map_converted_schema_attributes = append(map_converted_schema_attributes, nftAttributeValue_)
	}

	meta := nftmngrtypes.NewMetadata(&schema, tokenData, schema.OriginData.AttributeOverriding, map_converted_schema_attributes)
	meta.SetGetNFTFunction(func(tokenId string) (*nftmngrtypes.NftData, error) {
		tokenData, found := k.nftmngrKeeper.GetNftData(ctx, schema.Code, tokenId)
		if !found {
			return nil, sdkerrors.Wrap(nftmngrtypes.ErrMetadataDoesNotExists, schema.Code)
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

	err := ProcessAction(meta, &mapAction, input_param)
	if err != nil {
		return err
	}
	// Check if ChangeList is empty, error if empty
	if len(meta.ChangeList) == 0 {
		return sdkerrors.Wrap(types.ErrEmptyChangeList, actionRequest.Action)
	}

	k.nftmngrKeeper.SetNftData(ctx, *tokenData)

	// Udpate to target
	// loop over meta.OtherUpdatedTokenDatas
	for _, otherTokenData := range meta.OtherUpdatedTokenDatas {
		k.nftmngrKeeper.SetNftData(ctx, *otherTokenData)
	}

	for _, change := range meta.ChangeList {
		val, found := k.nftmngrKeeper.GetSchemaAttribute(ctx, actionRequest.NftSchemaCode, change.Key)
		if found {
			switch val.DataType {
			case "string":
				val.CurrentValue.Value = &nftmngrtypes.SchemaAttributeValue_StringAttributeValue{
					StringAttributeValue: &nftmngrtypes.StringAttributeValue{
						Value: change.NewValue,
					},
				}
			case "boolean":
				boolValue, err := strconv.ParseBool(change.NewValue)
				if err != nil {
					return err
				}
				val.CurrentValue.Value = &nftmngrtypes.SchemaAttributeValue_BooleanAttributeValue{
					BooleanAttributeValue: &nftmngrtypes.BooleanAttributeValue{
						Value: boolValue,
					},
				}
			case "number":
				uintValue, err := strconv.ParseUint(change.NewValue, 10, 64)
				if err != nil {
					return err
				}
				val.CurrentValue.Value = &nftmngrtypes.SchemaAttributeValue_NumberAttributeValue{
					NumberAttributeValue: &nftmngrtypes.NumberAttributeValue{
						Value: uintValue,
					},
				}
			case "float":
				floatValue, err := strconv.ParseFloat(change.NewValue, 64)
				if err != nil {
					return err
				}
				val.CurrentValue.Value = &nftmngrtypes.SchemaAttributeValue_FloatAttributeValue{
					FloatAttributeValue: &nftmngrtypes.FloatAttributeValue{
						Value: floatValue,
					},
				}
			default:
				return sdkerrors.Wrap(nftmngrtypes.ErrParsingAttributeValue, val.DataType)
			}

			k.nftmngrKeeper.SetSchemaAttribute(ctx, val)
		}
	}

	// Check action with reference exists
	if actionRequest.RefId != "" {

		_, found := k.nftmngrKeeper.GetActionByRefId(ctx, actionRequest.RefId)
		if found {
			return sdkerrors.Wrap(types.ErrRefIdAlreadyExists, actionRequest.RefId)
		}

		k.nftmngrKeeper.SetActionByRefId(ctx, nftmngrtypes.ActionByRefId{
			RefId:         actionRequest.RefId,
			Creator:       actionRequest.Caller,
			NftSchemaCode: actionRequest.NftSchemaCode,
			TokenId:       actionRequest.TokenId,
			Action:        mapAction.Name,
		})
	}
	// Emit events on metadata change
	changeList, _ := json.Marshal(meta.ChangeList)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.EventMessage, types.EventTypeRunAction),
			sdk.NewAttribute(types.AttributeKeyRunActionResult, "success"),
			// Emit change list attributes
			sdk.NewAttribute(types.AttributeKeyRunActionChangeList, string(changeList)),
		),
	)

	return nil
}

func ProcessAction(meta *nftmngrtypes.Metadata, action *nftmngrtypes.Action, params []*nftmngrtypes.ActionParameter) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()
	// Create params map from types.ActionParameter
	paramsMap := make(map[string]*nftmngrtypes.ActionParameter)
	for _, param := range params {
		paramsMap[param.Name] = param
	}

	dataContext := ast.NewDataContext()
	err = dataContext.Add("meta", meta)
	if err != nil {
		return err
	}
	err = dataContext.Add("params", paramsMap)
	if err != nil {
		return err
	}

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	ruleAction := &nftmngrkeeper.RuleAction{
		Name:     action.Name,
		Desc:     action.Desc,
		Salience: 10,
		When:     action.When,
		Then:     action.Then,
	}

	ruleAction.Then = append(ruleAction.Then, "Retract('"+action.Name+"');")
	ruleBytes, _ := utils.JSONMarshal(ruleAction)

	ruleResouce := pkg.NewBytesResource(ruleBytes)
	resource := pkg.NewJSONResourceFromResource(ruleResouce)

	err = ruleBuilder.BuildRuleFromResource(action.Name, "0.1.1", resource)
	if err != nil {
		return err
	}
	kb := lib.NewKnowledgeBaseInstance(action.Name, "0.1.1")
	eng1 := &engine.GruleEngine{MaxCycle: 100}
	err = eng1.Execute(dataContext, kb)
	if err != nil {
		return err
	}
	return nil
}

func (k msgServer) UpdateMetaDataFromOriginData(ctx sdk.Context, nftData *nftmngrtypes.NftData, originData *types.NftOriginData) error {

	schema, found := k.nftmngrKeeper.GetNFTSchema(ctx, nftData.NftSchemaCode)
	if !found {
		return sdkerrors.Wrap(types.ErrNFTSchemaNotFound, nftData.NftSchemaCode)
	}

	originAttributes, err := k.FromOriginDataToNftOriginAttribute(ctx, &schema, originData)
	if err != nil {
		return err
	}

	nftData.OriginAttributes = originAttributes

	mapOfTokenAttributeValues := make(map[string]*nftmngrtypes.NftAttributeValue)
	for _, attr := range nftData.OnchainAttributes {
		mapOfTokenAttributeValues[attr.Name] = attr
	}
	for _, attr := range schema.OnchainData.TokenAttributes {
		if attr.Required {
			if _, ok := mapOfTokenAttributeValues[attr.Name]; !ok {
				if attr.DefaultMintValue != nil {
					nftData.OnchainAttributes = append(nftData.OnchainAttributes, nftmngrkeeper.NewNFTAttributeValueFromDefaultValue(attr.Name, attr.DefaultMintValue))
				}
			}
		}
	}

	return nil
}

