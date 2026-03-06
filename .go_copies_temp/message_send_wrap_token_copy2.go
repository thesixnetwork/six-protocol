package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendWrapToken{}

func NewMsgSendWrapToken(creator string, ethAddress string, amount sdk.Coin) *MsgSendWrapToken {
	return &MsgSendWrapToken{
		Creator:    creator,
		EthAddress: ethAddress,
		Amount:     amount,
	}
}

func (msg *MsgSendWrapToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
