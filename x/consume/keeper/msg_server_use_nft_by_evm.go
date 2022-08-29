package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/consume/types"

	// "crypto/ecdsa"
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k msgServer) UseNftByEVM(goCtx context.Context, msg *types.MsgUseNftByEVM) (*types.MsgUseNftByEVMResponse, error) {
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

	// get address from public key
	pubEDCA, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "faild to unmarshal public key")
	}
	eth_address_from_pubkey := crypto.PubkeyToAddress(*pubEDCA)

	eth_address := common.HexToAddress(msg.EthAddress)
	if matches := bytes.Equal([]byte(eth_address_from_pubkey.Hex()), []byte(eth_address.Hex())); !matches {
		var ret = fmt.Sprintf("eth_address_from_pubkey: %s ,eth_address %s", eth_address_from_pubkey.Hex(), eth_address.Hex())
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, ret)
	}

	signatureNoRecoverID := signature_with_revocery_id[:len(signature_with_revocery_id)-1] // remove recovery id
	if verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID); !verified {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "verify failed")
	}

	var spend = types.MsgUseNft{
		Creator: msg.Creator,
		Token:   msg.Token,
	}
	_, err = k.UseNft(goCtx, &spend)
	if err != nil {
		return nil, err
	}

	return &types.MsgUseNftByEVMResponse{}, nil
}
