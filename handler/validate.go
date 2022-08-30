package handler

import (
	"strconv"
	"time"
	"fmt"
	
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"bytes"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func ValidateEVM(message string, signature string, eth_address string) (_signature []byte,err error) {
    sign_msg := "\x19Ethereum Signed Message:\n" + strconv.FormatInt(int64(len(message)), 10) + message

	data := []byte(sign_msg)
	hash := crypto.Keccak256Hash(data)
	var hash_bytes = hash.Bytes()

	//validate signature format
	decode_signature, err := hexutil.Decode(signature)
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

	common_eth_address := common.HexToAddress(eth_address)
	if matches := bytes.Equal([]byte(eth_address_from_pubkey.Hex()), []byte(common_eth_address.Hex())); !matches {
		var ret = fmt.Sprintf("eth_address_from_pubkey: %s ,eth_address %s", eth_address_from_pubkey.Hex(), common_eth_address.Hex())
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, ret)
	}

	signatureNoRecoverID := signature_with_revocery_id[:len(signature_with_revocery_id)-1] // remove recovery id
	if verified := crypto.VerifySignature(sigPublicKey, hash.Bytes(), signatureNoRecoverID); !verified {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "verify failed")
	}
	
	return signatureNoRecoverID, nil
}

func ValidateTimestamp(timestamp string) (utc_time_to_string string, err error) {

	msg_timestamp_int, err := strconv.ParseInt(timestamp[:len(timestamp)-3], 10, 64)
	if err != nil {
		return "" , sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid message of timestamp")
	}
	msg_timestamp_epoch := time.Unix(msg_timestamp_int,0)

	msg_timestamp_string := msg_timestamp_epoch.Format("2006-01-02T15:04:05Z")
    _ ,error:= time.Parse(time.RFC3339,msg_timestamp_string)
	if error != nil {
		return "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

    return msg_timestamp_string, nil
}