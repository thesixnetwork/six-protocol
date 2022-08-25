package keeper

import (
	"context"

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

	var spend = types.MsgUseNft{
		Creator: msg.Creator,
		Token:   msg.Token,
	}
	k.UseNft(goCtx,&spend)
	return &types.MsgUseNftByEVMResponse{}, nil
}
