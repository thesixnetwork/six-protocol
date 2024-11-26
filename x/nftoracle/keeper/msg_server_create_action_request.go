package keeper

import (
	"context"
	"encoding/base64"
	"strconv"
	"time"

	nftmngrtypes "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k msgServer) CreateActionRequest(goCtx context.Context, msg *types.MsgCreateActionRequest) (*types.MsgCreateActionRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	actionSigBz, err := base64.StdEncoding.DecodeString(msg.Base64ActionSignature)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBase64, msg.Base64ActionSignature)
	}
	data := types.ActionSignature{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(actionSigBz, &data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingActionSignature, err.Error())
	}

	actionOralceParam, signer, err := k.ValidateActionSignature(data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrValidateSignature, err.Error())
	}

	switch {
	case actionOralceParam.OnBehalfOf == "":
		break
	case actionOralceParam.OnBehalfOf != "" && actionOralceParam.OnBehalfOf == *signer:
		break
	case actionOralceParam.OnBehalfOf != "":
		_actionSigner, isActionSigner := k.GetActionSigner(ctx, *signer, actionOralceParam.OnBehalfOf)
		if isActionSigner && _actionSigner.ExpiredAt.After(ctx.BlockTime().UTC()) && actionOralceParam.OnBehalfOf == _actionSigner.OwnerAddress {
			*signer = _actionSigner.OwnerAddress
		} else {
			return nil, sdkerrors.Wrap(types.ErrInvalidSigningOnBehalfOf, "invalid onBehalfOf or ActionSigner is expired")
		}
	}

	// get schema from action request
	schema, found := k.nftmngrKeeper.GetNFTSchema(ctx, actionOralceParam.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNFTSchemaNotFound, actionOralceParam.NftSchemaCode)
	}

	mapAction := nftmngrtypes.Action{}
	for _, action := range schema.OnchainData.Actions {
		if action.Name == actionOralceParam.Action && action.Disable {
			return nil, sdkerrors.Wrap(nftmngrtypes.ErrActionIsDisabled, action.Name)
		}
		if action.Name == actionOralceParam.Action {
			mapAction = *action
			break
		}
	}

	// Check if action requires parameters
	param := mapAction.GetParams()
	required_param := make([]*nftmngrtypes.ActionParams, 0)

	for _, p := range param {
		if p.Required {
			required_param = append(required_param, p)
		}
	}

	if len(required_param) > len(actionOralceParam.Params) {
		return nil, sdkerrors.Wrap(nftmngrtypes.ErrInvalidParameter, "Input parameters length is not equal to required parameters length")
	}

	for i := 0; i < len(required_param); i++ {
		if actionOralceParam.Params[i].Name != required_param[i].Name {
			return nil, sdkerrors.Wrap(nftmngrtypes.ErrInvalidParameter, "input paramter name is not match to "+required_param[i].Name)
		}
		if actionOralceParam.Params[i].Value == "" {
			actionOralceParam.Params[i].Value = required_param[i].DefaultValue
		}
	}

	// var actionRequest types.ActionRequest

	// Check if the token is already Actioned
	_, found = k.nftmngrKeeper.GetNftData(ctx, actionOralceParam.NftSchemaCode, actionOralceParam.TokenId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrMetadataNotExists, actionOralceParam.NftSchemaCode)
	}

	// Check action with reference exists
	if actionOralceParam.RefId != "" {

		_, found := k.nftmngrKeeper.GetActionByRefId(ctx, actionOralceParam.RefId)
		if found {
			return nil, sdkerrors.Wrap(types.ErrRefIdAlreadyExists, actionOralceParam.RefId)
		}
	}

	oracleConfig, found := k.GetOracleConfig(ctx)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOracleConfigNotFound, "")
	}

	// Verify msg.RequiredConfirmations is less than or equal to oracleConfig.MinimumConfirmation
	if int32(msg.RequiredConfirm) < oracleConfig.MinimumConfirmation {
		return nil, sdkerrors.Wrap(types.ErrRequiredConfirmTooLess, strconv.Itoa(int(oracleConfig.MinimumConfirmation)))
	}

	createdAt := ctx.BlockTime() //now time of block
	endTime := createdAt.Add(k.ActionRequestActiveDuration(ctx))

	// validate time given format as RFC3339
	_, err = time.Parse(time.RFC3339, actionOralceParam.ExpiredAt.UTC().Format(time.RFC3339))
	if err != nil || len(actionOralceParam.ExpiredAt.String()) == 0 || actionOralceParam.ExpiredAt.Before(ctx.BlockHeader().Time.UTC()) {
		actionOralceParam.ExpiredAt = endTime
	}

	id_ := k.Keeper.AppendActionRequest(ctx, types.ActionOracleRequest{
		NftSchemaCode:   actionOralceParam.NftSchemaCode,
		TokenId:         actionOralceParam.TokenId,
		RequiredConfirm: msg.RequiredConfirm,
		Caller:          *signer,
		Action:          actionOralceParam.Action,
		Params:          actionOralceParam.Params,
		RefId:           actionOralceParam.RefId,
		Status:          types.RequestStatus_PENDING,
		CurrentConfirm:  0,
		CreatedAt:       createdAt,
		ValidUntil:      actionOralceParam.ExpiredAt,
		Confirmers:      make([]string, 0),
		DataHashes:      make([]*types.DataHash, 0),
	})

	k.Keeper.InsertActiveActionRequestQueue(ctx, id_, actionOralceParam.ExpiredAt)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeActionRequestCreated,
			sdk.NewAttribute(types.AttributeKeyActionRequestID, strconv.FormatUint(id_, 10)),
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, actionOralceParam.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyTokenID, actionOralceParam.TokenId),
			sdk.NewAttribute(types.AttributeKeyRequiredConfirm, strconv.FormatUint(msg.RequiredConfirm, 10)),
		),
	})

	return &types.MsgCreateActionRequestResponse{
		Id: id_,
	}, nil
}

func (k msgServer) ValidateActionSignature(actionSig types.ActionSignature) (*types.ActionOracleParam, *string, error) {

	sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(actionSig.Message)), 10) + actionSig.Message

	data := []byte(sign_msg)
	hash := crypto.Keccak256Hash(data)
	var hash_bytes = hash.Bytes()

	actionOralceParam := &types.ActionOracleParam{}
	actionParamBz, err := base64.StdEncoding.DecodeString(actionSig.Message)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(types.ErrParsingActionParam, "Error to DecodeString")
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(actionParamBz, actionOralceParam)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(types.ErrParsingActionParam, "Error to UnmarshalJSON")
	}

	//validate signature format
	decode_signature, err := hexutil.Decode(actionSig.Signature)
	if err != nil {
		// log.Fatalf("Failed to decode signature: %v", msg.Signature)
		return nil, nil, sdkerrors.Wrap(types.ErrHexDecode, "Failed to decode signature")
	}
	signature_with_revocery_id := decode_signature
	// remove last byte coz is is recovery id

	// get pulic key from signature
	sigPublicKey, err := crypto.Ecrecover(hash_bytes, decode_signature) //recover publickey from signature and hash
	if err != nil {
		return nil, nil, sdkerrors.Wrap(types.ErrEcrecover, "invalid signature or message")
	}

	// fmt.Println("sigPublicKey:", hexutil.Encode(sigPublicKey))
	// get address from public key
	pubEDCA, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(types.ErrUnmarshalPubkey, "faild to unmarshal public key")
	}
	eth_address_from_pubkey := crypto.PubkeyToAddress(*pubEDCA)

	signatureNoRecoverID := signature_with_revocery_id[:len(signature_with_revocery_id)-1] // remove recovery id
	if verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID); !verified {
		return nil, nil, sdkerrors.Wrap(types.ErrVerifyingSignature, "invalid signature")
	}
	signer := eth_address_from_pubkey.Hex()
	return actionOralceParam, &signer, nil
}
