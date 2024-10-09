package keeper

import (
	"context"
	"encoding/base64"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func (k msgServer) CreateVerifyCollectionOwnerRequest(goCtx context.Context, msg *types.MsgCreateVerifyCollectionOwnerRequest) (*types.MsgCreateVerifyCollectionOwnerRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if nft_schema_code exists
	schema_, found := k.nftmngrKeeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrNFTSchemaNotFound, msg.NftSchemaCode)
	}

	// check if creator is owner of the collectio
	if schema_.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(types.ErrNotCollectionOwner, msg.Creator)
	}

	// Check if nft_schema_code already verified
	if schema_.IsVerified {
		return nil, sdkerrors.Wrap(types.ErrNFTSchemaAlreadyVerified, msg.NftSchemaCode)
	}

	// TODO: Check chain origin and modify how to verify signature according to chain origin in the future

	collectionOwnerBz, err := base64.StdEncoding.DecodeString(msg.Base64VerifyRequestorSignature)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBase64, msg.Base64VerifyRequestorSignature)
	}
	data := types.CollectionOwnerSignature{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(collectionOwnerBz, &data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingCollectionOwnerSignature, err.Error())
	}
	_originContractParam, signer, err := k.ValidateCollectionOwnerSignature(data)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrVerifyingSignature, err.Error())
	}

	oracleConfig, found := k.GetOracleConfig(ctx)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrOracleConfigNotFound, "")
	}

	// Verify msg.RequiredConfirmations is less than or equal to oracleConfig.MinimumConfirmation
	if int32(msg.RequiredConfirm) < oracleConfig.MinimumConfirmation {
		return nil, sdkerrors.Wrap(types.ErrRequiredConfirmTooLess, strconv.Itoa(int(oracleConfig.MinimumConfirmation)))
	}

	createdAt := ctx.BlockTime()
	endTime := createdAt.Add(k.VerifyRequestActiveDuration(ctx))

	// validate time given format as RFC3339
	_, err = time.Parse(time.RFC3339, _originContractParam.RequestExpire.UTC().Format(time.RFC3339))
	if err != nil || len(_originContractParam.RequestExpire.String()) == 0 || _originContractParam.RequestExpire.Before(ctx.BlockHeader().Time.UTC()) {
		_originContractParam.RequestExpire = endTime
	}

	id_ := k.Keeper.AppendCollectionOwnerRequest(ctx, types.CollectionOwnerRequest{
		NftSchemaCode:   msg.NftSchemaCode,
		Signer:          *signer,
		RequiredConfirm: msg.RequiredConfirm,
		Status:          types.RequestStatus_PENDING,
		CurrentConfirm:  0,
		CreatedAt:       createdAt,
		ValidUntil:      _originContractParam.RequestExpire,
		Confirmers:      make([]string, 0),
		ContractInfo:    make([]*types.OriginContractInfo, 0),
	})

	k.Keeper.InsertActiveVerifyCollectionOwnerRequestQueue(ctx, id_, _originContractParam.RequestExpire)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeVerificationRequestCreated,
			sdk.NewAttribute(types.AttributeKeyVerifyRequestID, strconv.FormatUint(id_, 10)),
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			// sdk.NewAttribute(types.AtttibuteKeyVerifyRequestChainOrigin, schema_.OriginData.OriginChain ),
			sdk.NewAttribute(types.AttributeKeyRequiredConfirm, strconv.FormatUint(msg.RequiredConfirm, 10)),
		),
	})

	return &types.MsgCreateVerifyCollectionOwnerRequestResponse{
		Id:            id_,
		NftSchemaCode: schema_.Code,
		OwnerAddress:  *signer,
	}, nil
}

func (k msgServer) ValidateCollectionOwnerSignature(collectionOwnerSig types.CollectionOwnerSignature) (*types.OriginContractParam, *string, error) {

	sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(collectionOwnerSig.Message)), 10) + collectionOwnerSig.Message

	data := []byte(sign_msg)
	hash := crypto.Keccak256Hash(data)
	var hash_bytes = hash.Bytes()

	collectionOwnerParam := &types.OriginContractParam{}
	collectionOwnerTypeBz, err := base64.StdEncoding.DecodeString(collectionOwnerSig.Message)
	if err != nil {
		return nil, nil, err
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(collectionOwnerTypeBz, collectionOwnerParam)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(types.ErrParsingActionParam, err.Error())
	}

	//validate signature format
	decode_signature, err := hexutil.Decode(collectionOwnerSig.Signature)
	if err != nil {
		// log.Fatalf("Failed to decode signature: %v", msg.Signature)
		return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature")
	}
	signature_with_revocery_id := decode_signature

	// get pulic key from signature
	sigPublicKey, err := crypto.Ecrecover(hash_bytes, decode_signature) //recover publickey from signature and hash
	if err != nil {
		return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature or message")
	}
	// get address from public key
	pubEDCA, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "faild to unmarshal public key")
	}
	eth_address_from_pubkey := crypto.PubkeyToAddress(*pubEDCA)

	signatureNoRecoverID := signature_with_revocery_id[:len(signature_with_revocery_id)-1] // remove recovery id
	if verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID); !verified {
		return nil, nil, sdkerrors.Wrap(types.ErrVerifyingSignature, "invalid signature")
	}
	signer := eth_address_from_pubkey.Hex()
	return collectionOwnerParam, &signer, nil
}
