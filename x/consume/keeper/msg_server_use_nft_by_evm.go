package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/consume/types"

	// "crypto/ecdsa"
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)


type Message struct{
    Msg string
    Timestamp string
}

func (k msgServer) UseNftByEVM(goCtx context.Context, msg *types.MsgUseNftByEVM) (*types.MsgUseNftByEVMResponse, error) {
	//chaeck creator is valid
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	var raw_sign_msg Message
    json.Unmarshal([]byte(msg.SignMessage), &raw_sign_msg)

	// hash the message
	sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(msg.SignMessage)), 10) + msg.SignMessage

	 // getting timestamp from message and transforming it to UTC
	time_stamp := raw_sign_msg.Timestamp
	_ = raw_sign_msg.Msg
	msg_timestamp_int, err := strconv.ParseInt(time_stamp[:len(time_stamp)-3], 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid message of timestamp")
	}
	msg_timestamp_epoch := time.Unix(msg_timestamp_int,0)

	msg_timestamp_string := msg_timestamp_epoch.Format("2006-01-02T15:04:05Z")
    _ ,error:= time.Parse(time.RFC3339,msg_timestamp_string)
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

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
		Timestamp: msg_timestamp_string,
	}
	_, err = k.UseNft(goCtx, &spend)
	if err != nil {
		return nil, err
	}

	return &types.MsgUseNftByEVMResponse{}, nil
}
