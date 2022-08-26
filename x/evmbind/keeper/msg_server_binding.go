package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"

	// "crypto/ecdsa"
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k msgServer) CreateBinding(goCtx context.Context, msg *types.MsgCreateBinding) (*types.MsgCreateBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetBinding(
		ctx,
		msg.EthAddress,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	//chaeck creator is valid
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	// hash the message
	sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(msg.SignMessage)), 10) + msg.SignMessage
	data := []byte(sign_msg)
	hash := crypto.Keccak256Hash(data)
	var hash_bytes = hash.Bytes()

	//validate signature format
	decode_signature, err := hexutil.Decode(msg.EthSignature)
	if err != nil {
		// log.Fatalf("Failed to decode signature: %v", msg.Signature)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature")
	}
	signature_with_revocery_id := decode_signature
	// remove last byte coz is is recovery id
	decode_signature[64] -= 27 // this on should be checked whether it can be a weak point later // remove recovery id

	// get pulic key from signature
	sigPublicKey, err := crypto.Ecrecover(hash_bytes, decode_signature) //recover publickey from signature and hash
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature or message")
	}
	// pubKey_encode := hexutil.Encode(sigPublicKey)

	// get address from public key
	pubEDCA, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "faild to unmarshal public key")
	}
	eth_address_from_pubkey := crypto.PubkeyToAddress(*pubEDCA)

	eth_address := common.HexToAddress(msg.EthAddress)
	if matches := bytes.Equal([]byte(eth_address_from_pubkey.Hex()), []byte(eth_address.Hex())); !matches {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fail to validate eth address, check signature and message are correct")
	}

	signatureNoRecoverID := signature_with_revocery_id[:len(signature_with_revocery_id)-1] // remove recovery id
	if verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID); !verified {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "verify failed")
	}

	var binding = types.Binding{
		Creator:      msg.Creator,
		EthAddress:   msg.EthAddress,
		EthSignature: msg.EthSignature,
		SignMessage:  msg.SignMessage,
	}

	k.SetBinding(
		ctx,
		binding,
	)
	return &types.MsgCreateBindingResponse{}, nil
}

func (k msgServer) UpdateBinding(goCtx context.Context, msg *types.MsgUpdateBinding) (*types.MsgUpdateBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetBinding(
		ctx,
		msg.EthAddress,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	// hash the message
	data := []byte(msg.SignMessage)
	hash := crypto.Keccak256Hash(data)
	var hash_bytes = hash.Bytes()

	//validate signature format
	decode_signature, err := hexutil.Decode(msg.EthSignature)
	if err != nil {
		// log.Fatalf("Failed to decode signature: %v", msg.Signature)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature")
	}

	// get pulic key from signature
	sigPublicKey, err := crypto.Ecrecover(hash_bytes, decode_signature) //recover publickey from signature and hash
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid signature or message")
	}
	// pubKey_encode := hexutil.Encode(sigPublicKey)

	// get address from public key
	pubEDCA, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "faild to unmarshal public key")
	}
	eth_address_from_pubkey := crypto.PubkeyToAddress(*pubEDCA)

	eth_address := common.HexToAddress(msg.EthAddress)
	if matches := bytes.Equal([]byte(eth_address_from_pubkey.Hex()), []byte(eth_address.Hex())); !matches {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fail to validate eth address, check signature and message are correct")
	}

	var binding = types.Binding{
		Creator:      msg.Creator,
		EthAddress:   msg.EthAddress,
		EthSignature: msg.EthSignature,
		SignMessage:  msg.SignMessage,
	}

	k.SetBinding(ctx, binding)

	return &types.MsgUpdateBindingResponse{}, nil
}

func (k msgServer) DeleteBinding(goCtx context.Context, msg *types.MsgDeleteBinding) (*types.MsgDeleteBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetBinding(
		ctx,
		msg.EthAddress,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBinding(
		ctx,
		msg.EthAddress,
	)

	return &types.MsgDeleteBindingResponse{}, nil
}
