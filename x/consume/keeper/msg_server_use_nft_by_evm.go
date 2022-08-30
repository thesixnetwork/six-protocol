package keeper

import (
	"context"
	// "fmt"
	// "strconv"
	// "time"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
	handler "github.com/thesixnetwork/six-protocol/handler"

	// "crypto/ecdsa"
	// "bytes"

	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/crypto"
)


type Message struct{
    Msg string
    Timestamp string
}

func (k msgServer) UseNftByEVM(goCtx context.Context, msg *types.MsgUseNftByEVM) (*types.MsgUseNftByEVMResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Check if the value already exists
	_, isFound := k.evmbindKeeper.GetBinding(
		ctx,
		msg.EthAddress,
	)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Eth address not found in binding")
	}

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	var raw_sign_msg Message
    json.Unmarshal([]byte(msg.SignMessage), &raw_sign_msg)

	msg_timestamp_string, err := handler.ValidateTimestamp(raw_sign_msg.Timestamp)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	_ ,err = handler.ValidateEVM(msg.SignMessage, msg.EthSignature, msg.EthAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
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
