package keeper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"strconv"

	nftmngrkeeper "github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	nftmngrtypes "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	errormod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitMintResponse(goCtx context.Context, msg *types.MsgSubmitMintResponse) (*types.MsgSubmitMintResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oracle, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionOracle, oracle)
	if !granted {
		return nil, errormod.Wrap(types.ErrNoOraclePermission, msg.Creator)
	}

	mintRequest, found := k.GetMintRequest(ctx, msg.MintRequestID)
	if !found {
		return nil, errormod.Wrap(types.ErrMintRequestNotFound, strconv.FormatUint(msg.MintRequestID, 10))
	}

	// Check if mint request is still pending
	if mintRequest.Status != types.RequestStatus_PENDING {
		return nil, errormod.Wrap(types.ErrMintRequestNotPending, strconv.FormatUint(msg.MintRequestID, 10))
	}

	// Check if current confirmation count is less than required confirmation count
	if mintRequest.CurrentConfirm >= mintRequest.RequiredConfirm {
		return nil, errormod.Wrap(types.ErrMintRequestConfirmedAlreadyComplete, strconv.FormatUint(msg.MintRequestID, 10))
	}

	// Convert msg.Base64NftMetadata to bytes
	nftMetadata, err := base64.StdEncoding.DecodeString(msg.Base64NftData)
	if err != nil {
		return nil, errormod.Wrap(types.ErrParsingBase64, err.Error())
	}

	nftOriginData := types.NftOriginData{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(nftMetadata, &nftOriginData)
	if err != nil {
		return nil, errormod.Wrap(types.ErrParsingOriginData, err.Error())
	}

	if nftOriginData.HolderAddress == "" || nftOriginData.Image == "" || nftOriginData.Traits == nil || len(nftOriginData.Traits) == 0 {
		return nil, errormod.Wrap(types.ErrParsingOriginData, "missing required fields")
	}

	// ! :: Check Deterministic Hash from concurrence response
	if mintRequest.CurrentConfirm == 0 {
		// Create sha256 hash of nftMetadata
		dataHash := sha256.Sum256(nftMetadata)
		mintRequest.DataHashes = append(mintRequest.DataHashes, &types.DataHash{
			OriginData: &nftOriginData,
			Hash:       dataHash[:],
			Confirmers: []string{msg.Creator},
		})
	} else {
		// Check if creator has already confirmed this mint request
		// fetch confirmed data
		for _, confirmer := range mintRequest.Confirmers {
			if confirmer == msg.Creator {
				return nil, errormod.Wrap(types.ErrOracleConfirmedAlready, strconv.FormatUint(msg.MintRequestID, 10))
			}
		}
		// if _, ok := mintRequest.Confirmers[msg.Creator]; ok {
		// 	return nil, errormod.Wrap(types.ErrOracleConfirmedAlready, strconv.FormatUint(msg.MintRequestID, 10)+", "+msg.Creator)
		// }
		// Compare data hash with previous data hash
		dataHash := sha256.Sum256(nftMetadata)
		dataHashMatch := false
		for _, hash := range mintRequest.DataHashes {
			if res := bytes.Compare(dataHash[:], hash.Hash); res == 0 {
				dataHashMatch = true
				hash.Confirmers = append(hash.Confirmers, msg.Creator)
				break
			}
		}
		if !dataHashMatch {
			mintRequest.DataHashes = append(mintRequest.DataHashes, &types.DataHash{
				OriginData: &nftOriginData,
				Hash:       dataHash[:],
				Confirmers: []string{msg.Creator},
			})
		}
	}

	if mintRequest.Confirmers == nil {
		mintRequest.Confirmers = make([]string, 0)
	}
	// Mark creator as confirmed
	// mintRequest.Confirmers[msg.Creator] = true
	mintRequest.Confirmers = append(mintRequest.Confirmers, msg.Creator)

	// increase mintRequest.CurrentConfirm
	mintRequest.CurrentConfirm++

	if mintRequest.CurrentConfirm == mintRequest.RequiredConfirm {
		// Mint NFT
		// Check if there is only one data hash
		if len(mintRequest.DataHashes) > 1 {
			// Update mintRequest.Status to be FAILED
			mintRequest.Status = types.RequestStatus_FAILED_WITHOUT_CONSENSUS
		} else {
			// Update mintRequest.Status to be SUCCESS
			mintRequest.Status = types.RequestStatus_SUCCESS_WITH_CONSENSUS
		}

		// Emit event a consensus result
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeMintRequestConfirmed,
				sdk.NewAttribute(types.AttributeKeyMintRequestID, strconv.FormatUint(mintRequest.Id, 10)),
				sdk.NewAttribute(types.AttributeKeyMintRequestStatus, mintRequest.Status.String()),
			),
		)
		if mintRequest.Status == types.RequestStatus_SUCCESS_WITH_CONSENSUS {
			nftData, err := k.CreateMetaDataFromOriginData(ctx, mintRequest, mintRequest.DataHashes[0].OriginData)
			if err != nil {
				return nil, err
			}
			k.nftmngrKeeper.SetNftData(ctx, *nftData)
		}
	}

	// Update Mint Request Back to storage
	k.SetMintRequest(ctx, mintRequest)

	return &types.MsgSubmitMintResponseResponse{}, nil
}

func (k msgServer) CreateMetaDataFromOriginData(ctx sdk.Context, mintRequest types.MintRequest, originData *types.NftOriginData) (*nftmngrtypes.NftData, error) {
	nftData := &nftmngrtypes.NftData{
		NftSchemaCode:    mintRequest.NftSchemaCode,
		TokenId:          mintRequest.TokenId,
		TokenOwner:       originData.HolderAddress,
		OwnerAddressType: nftmngrtypes.OwnerAddressType_ORIGIN_ADDRESS,
		OriginImage:      originData.Image,
	}

	schema, found := k.nftmngrKeeper.GetNFTSchema(ctx, mintRequest.NftSchemaCode)
	if !found {
		return nil, errormod.Wrap(types.ErrNFTSchemaNotFound, mintRequest.NftSchemaCode)
	}

	originAttributes, err := k.FromOriginDataToNftOriginAttribute(ctx, &schema, originData)
	if err != nil {
		return nil, err
	}

	nftData.OriginAttributes = originAttributes
	nftData.OnchainAttributes = make([]*nftmngrtypes.NftAttributeValue, 0)

	// Poppulate Onchain Data
	for _, attr := range schema.OnchainData.TokenAttributes {
		if attr.DefaultMintValue != nil {
			nftData.OnchainAttributes = append(nftData.OnchainAttributes, nftmngrkeeper.NewNFTAttributeValueFromDefaultValue(attr.Name, attr.DefaultMintValue))
		}
	}

	return nftData, nil
}

func (k Keeper) FromOriginDataToNftOriginAttribute(ctx sdk.Context, schema *nftmngrtypes.NFTSchema, originData *types.NftOriginData) ([]*nftmngrtypes.NftAttributeValue, error) {
	attributes := make([]*nftmngrtypes.NftAttributeValue, 0)

	if schema.OriginData == nil {
		return nil, errormod.Wrap(types.ErrParsingOriginData, "origin data is nil")
	}

	if schema.OriginData.MetadataFormat != "opensea" {
		return nil, errormod.Wrap(types.ErrUnsupportedMetadataFormat, schema.OriginData.MetadataFormat)
	}
	// Create map of attribute definition by trait type
	attributeByTrait := make(map[string]*nftmngrtypes.AttributeDefinition)
	for _, attriDef := range schema.OriginData.OriginAttributes {
		attributeByTrait[attriDef.DisplayOption.Opensea.TraitType] = attriDef
	}

	for _, trait := range originData.Traits {
		if attributeByTrait[trait.TraitType] == nil {
			return nil, errormod.Wrap(types.ErrNFTSchemaAttributeNotFound, trait.TraitType)
		}
		attributeValue, err := TranslateToAttributeValue(attributeByTrait[trait.TraitType], trait.Value)
		if err != nil {
			return nil, err
		}

		// Check if attribute already exists and update it, otherwise append
		found := false
		for i, existingAttr := range attributes {
			if existingAttr.Name == attributeValue.Name {
				attributes[i] = attributeValue // Update existing attribute
				found = true
				break
			}
		}
		if !found {
			attributes = append(attributes, attributeValue) // Append new attribute
		}

		delete(attributeByTrait, trait.TraitType)
	}
	// Looking for _HIDDEN_ boolean attribute
	for _, attriDef := range attributeByTrait {
		if attriDef.DataType == "boolean" && attriDef.DisplayOption != nil {
			var attributeValue *nftmngrtypes.NftAttributeValue
			if attriDef.DisplayOption.BoolTrueValue == "_HIDDEN_" {
				attributeValue = &nftmngrtypes.NftAttributeValue{
					Name: attriDef.Name,
					Value: &nftmngrtypes.NftAttributeValue_BooleanAttributeValue{
						BooleanAttributeValue: &nftmngrtypes.BooleanAttributeValue{
							Value: true,
						},
					},
				}
				attributes = append(attributes, attributeValue)
			} else if attriDef.DisplayOption.BoolFalseValue == "_HIDDEN_" {
				attributeValue = &nftmngrtypes.NftAttributeValue{
					Name: attriDef.Name,
					Value: &nftmngrtypes.NftAttributeValue_BooleanAttributeValue{
						BooleanAttributeValue: &nftmngrtypes.BooleanAttributeValue{
							Value: false,
						},
					},
				}
				attributes = append(attributes, attributeValue)
			}
		}
	}

	return attributes, nil
}

func TranslateToAttributeValue(attriDef *nftmngrtypes.AttributeDefinition, value string) (*nftmngrtypes.NftAttributeValue, error) {
	attributeValue := &nftmngrtypes.NftAttributeValue{}
	attributeValue.Name = attriDef.Name
	switch attriDef.DataType {
	case "string":
		attributeValue.Value = &nftmngrtypes.NftAttributeValue_StringAttributeValue{
			StringAttributeValue: &nftmngrtypes.StringAttributeValue{
				Value: value,
			},
		}
	case "number":
		numberValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nil, err
		}
		attributeValue.Value = &nftmngrtypes.NftAttributeValue_NumberAttributeValue{
			NumberAttributeValue: &nftmngrtypes.NumberAttributeValue{
				Value: numberValue,
			},
		}
	case "boolean":
		booleanValue := false
		if attriDef.DisplayOption != nil && attriDef.DisplayOption.BoolTrueValue == value {
			booleanValue = true
		}
		attributeValue.Value = &nftmngrtypes.NftAttributeValue_BooleanAttributeValue{
			BooleanAttributeValue: &nftmngrtypes.BooleanAttributeValue{
				Value: booleanValue,
			},
		}
	case "float":
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		attributeValue.Value = &nftmngrtypes.NftAttributeValue_FloatAttributeValue{
			FloatAttributeValue: &nftmngrtypes.FloatAttributeValue{
				Value: floatValue,
			},
		}
	default:
		return nil, errormod.Wrap(types.ErrParsingOriginData, "unsupported data type: "+attriDef.DataType)
	}
	return attributeValue, nil
}
