package keeper

import (
	"context"
	"encoding/base64"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	errormod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateActionSigner(goCtx context.Context, msg *types.MsgCreateActionSigner) (*types.MsgCreateActionSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerActionBz, err := base64.StdEncoding.DecodeString(msg.Base64EncodedSetSignerAction)
	if err != nil {
		return nil, errormod.Wrap(types.ErrInvalidBase64, msg.Base64EncodedSetSignerAction)
	}

	data := types.SetSignerSignature{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(signerActionBz, &data)
	if err != nil {
		return nil, errormod.Wrap(types.ErrParsingBase64, err.Error())
	}
	_signerParams, signer, err := k.ValidatesetSignernature(data)
	if err != nil {
		return nil, errormod.Wrap(types.ErrVerifyingSignature, err.Error())
	}

	// Check if the value already exists
	binded, isFound := k.GetActionSigner(
		ctx,
		_signerParams.ActorAddress,
		*signer,
	)

	// switch case for found and not found
	switch isFound {
	case true:
		if binded.ExpiredAt.Before(ctx.BlockTime().UTC()) { // if expired
			// check if the expired time is same as the given time
			// prevent replay attack
			if binded.ExpiredAt.Equal(_signerParams.ExpiredAt) {
				return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "Action signer message already sent should change the expired time")
			}

			// update the action signer
			updatedAt := ctx.BlockTime()
			endTime := updatedAt.Add(k.ActionSignerActiveDuration(ctx))

			// validate time given format as RFC3339
			_, err = time.Parse(time.RFC3339, _signerParams.ExpiredAt.UTC().Format(time.RFC3339))

			if err != nil || len(_signerParams.ExpiredAt.String()) == 0 || _signerParams.ExpiredAt.Before(updatedAt) {
				_signerParams.ExpiredAt = endTime
			}

			actionSigner := types.ActionSigner{
				ActorAddress: _signerParams.ActorAddress,
				OwnerAddress: *signer,
				CreatedAt:    updatedAt,
				ExpiredAt:    _signerParams.ExpiredAt,
				Creator:      msg.Creator,
				CreationFlow: types.CreationFlow_INTERNAL_OWNER,
			}
			k.SetActionSigner(ctx, actionSigner)
			// set to bined address
			bindedList, _ := k.GetBindedSigner(ctx, *signer)

			// set to same index
			for i, bindedIndex := range bindedList.Signers {
				if bindedIndex.ActorAddress == _signerParams.ActorAddress {
					bindedList.Signers[i].ExpiredAt = _signerParams.ExpiredAt
				}
			}

			// set the binded signer
			bindedList.ActorCount = uint64(len(bindedList.Signers))
			k.SetBindedSigner(ctx, bindedList)

		} else {
			return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "Action signer already bined and not expired")
		}
	case false:
		// create the action signer
		createdAt := ctx.BlockTime()
		endTime := createdAt.Add(k.ActionSignerActiveDuration(ctx))

		// validate time given format as RFC3339
		_, err = time.Parse(time.RFC3339, _signerParams.ExpiredAt.UTC().Format(time.RFC3339))
		if err != nil || len(_signerParams.ExpiredAt.String()) == 0 || _signerParams.ExpiredAt.Before(ctx.BlockTime().UTC()) {
			_signerParams.ExpiredAt = endTime
		}

		actionSigner := types.ActionSigner{
			ActorAddress: _signerParams.ActorAddress,
			OwnerAddress: *signer,
			CreatedAt:    createdAt,
			ExpiredAt:    _signerParams.ExpiredAt,
			Creator:      msg.Creator,
			CreationFlow: types.CreationFlow_INTERNAL_OWNER,
		}

		// Get List of binded signers
		bindedList, found := k.GetBindedSigner(ctx, *signer)
		if !found {
			bindedList = types.BindedSigner{
				OwnerAddress: *signer,
				Signers:      make([]*types.BindedSignerParams, 0),
			}
		}

		// add the binded signer to the list
		bindedList.Signers = append(bindedList.Signers, &types.BindedSignerParams{
			ActorAddress: _signerParams.ActorAddress,
			ExpiredAt:    _signerParams.ExpiredAt,
		})

		// set the binded signer
		bindedList.ActorCount = uint64(len(bindedList.Signers))
		k.SetBindedSigner(ctx, bindedList)

		k.SetActionSigner(
			ctx,
			actionSigner,
		)

	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeActionSigner,
			sdk.NewAttribute(types.AttributeKeySignerOwner, _signerParams.OwnerAddress),
			sdk.NewAttribute(types.AttributeKeySignerActor, _signerParams.ActorAddress),
			sdk.NewAttribute(types.AttributeKeySginerExpireAt, _signerParams.ExpiredAt.UTC().Format(time.RFC3339)),
		),
	)

	return &types.MsgCreateActionSignerResponse{
		OwnerAddress:  _signerParams.ActorAddress,
		SignerAddress: _signerParams.OwnerAddress,
		ExpireAt:      _signerParams.ExpiredAt.UTC().Format(time.RFC3339),
	}, nil
}

func (k msgServer) UpdateActionSigner(goCtx context.Context, msg *types.MsgUpdateActionSigner) (*types.MsgUpdateActionSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerActionBz, err := base64.StdEncoding.DecodeString(msg.Base64EncodedSetSignerAction)
	if err != nil {
		return nil, errormod.Wrap(types.ErrInvalidBase64, msg.Base64EncodedSetSignerAction)
	}

	data := types.SetSignerSignature{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(signerActionBz, &data)
	if err != nil {
		return nil, errormod.Wrap(types.ErrParsingSetSignerSignature, err.Error())
	}
	_signerParams, signer, err := k.ValidatesetSignernature(data)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	isSigner, isFound := k.GetActionSigner(
		ctx,
		_signerParams.ActorAddress,
		*signer,
	)
	if !isFound {
		return nil, errormod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// prevent replay attack
	if isSigner.ExpiredAt.Equal(_signerParams.ExpiredAt) {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "Action signer message already sent should change the expired time")
	}

	// update the action signer
	updatedAt := ctx.BlockTime()
	endTime := updatedAt.Add(k.ActionSignerActiveDuration(ctx))

	// validate time given format as RFC3339
	_, err = time.Parse(time.RFC3339, _signerParams.ExpiredAt.UTC().Format(time.RFC3339))

	if err != nil || len(_signerParams.ExpiredAt.String()) == 0 || _signerParams.ExpiredAt.Before(updatedAt) {
		_signerParams.ExpiredAt = endTime
	}

	// //check if creator is "oracle"
	// if *signer != isSigner.Creator{
	// 	oracleAddress, err := sdk.AccAddressFromBech32(isSigner.Creator)
	// 	if err != nil {
	// 		return nil, errormod.Wrap(sdkerrors.ErrInvalidAddress, isSigner.Creator)
	// 	}
	// 	// get permission of the creator
	// 	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionOracle, oracleAddress)
	// 	if !granted {
	// 		return nil, errormod.Wrap(types.ErrNoOraclePermission, isSigner.Creator)
	// 	}
	// }

	if *signer != isSigner.Creator {
		return nil, errormod.Wrap(sdkerrors.ErrUnauthorized, "incorrect creator")
	}

	actionSigner := types.ActionSigner{
		ActorAddress: _signerParams.ActorAddress,
		OwnerAddress: *signer,
		CreatedAt:    updatedAt,
		ExpiredAt:    _signerParams.ExpiredAt,
		Creator:      msg.Creator,
		CreationFlow: types.CreationFlow_INTERNAL_OWNER,
	}

	k.SetActionSigner(ctx, actionSigner)

	return &types.MsgUpdateActionSignerResponse{}, nil
}

func (k msgServer) DeleteActionSigner(goCtx context.Context, msg *types.MsgDeleteActionSigner) (*types.MsgDeleteActionSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	signerActionBz, err := base64.StdEncoding.DecodeString(msg.Base64EncodedSetSignerAction)
	if err != nil {
		return nil, errormod.Wrap(types.ErrInvalidBase64, msg.Base64EncodedSetSignerAction)
	}

	data := types.SetSignerSignature{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(signerActionBz, &data)
	if err != nil {
		return nil, errormod.Wrap(types.ErrParsingSetSignerSignature, err.Error())
	}
	_signerParams, signer, err := k.ValidatesetSignernature(data)
	if err != nil {
		return nil, err
	}

	// Check if the value exists
	isSigner, isFound := k.GetActionSigner(
		ctx,
		_signerParams.ActorAddress,
		*signer,
	)

	if !isFound {
		return nil, errormod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	if msg.Creator != isSigner.Creator {
		return nil, errormod.Wrap(sdkerrors.ErrUnauthorized, "incorrect creator")
	}

	k.RemoveActionSigner(
		ctx,
		_signerParams.ActorAddress,
		*signer,
	)

	return &types.MsgDeleteActionSignerResponse{}, nil
}

func (k msgServer) ValidatesetSignernature(setSigner types.SetSignerSignature) (*types.SetSignerParams, *string, error) {
	sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(setSigner.Message)), 10) + setSigner.Message

	data := []byte(sign_msg)
	hash := crypto.Keccak256Hash(data)
	hash_bytes := hash.Bytes()

	setSignerParams := &types.SetSignerParams{}
	setSignerTypeBz, err := base64.StdEncoding.DecodeString(setSigner.Message)
	if err != nil {
		return nil, nil, err
	}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(setSignerTypeBz, setSignerParams)
	if err != nil {
		return nil, nil, errormod.Wrap(types.ErrParsingSetSignerSignature, err.Error())
	}

	// validate signature format
	decode_signature, err := hexutil.Decode(setSigner.Signature)
	if err != nil {
		// log.Fatalf("Failed to decode signature: %v", msg.Signature)
		return nil, nil, errormod.Wrap(types.ErrHexDecode, "invalid signature")
	}
	signature_with_revocery_id := decode_signature

	// get pulic key from signature
	sigPublicKey, err := crypto.Ecrecover(hash_bytes, decode_signature) // recover publickey from signature and hash
	if err != nil {
		return nil, nil, errormod.Wrap(types.ErrEcrecover, "invalid signature or message")
	}
	// get address from public key
	pubEDCA, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return nil, nil, errormod.Wrap(types.ErrUnmarshalPubkey, "faild to unmarshal public key")
	}
	eth_address_from_pubkey := crypto.PubkeyToAddress(*pubEDCA)

	signatureNoRecoverID := signature_with_revocery_id[:len(signature_with_revocery_id)-1] // remove recovery id
	if verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID); !verified {
		return nil, nil, errormod.Wrap(types.ErrVerifyingSignature, "invalid signature")
	}
	signer := eth_address_from_pubkey.Hex()
	return setSignerParams, &signer, nil
}
