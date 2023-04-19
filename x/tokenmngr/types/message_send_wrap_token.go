package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendWrapToken = "send_wrap_token"

var _ sdk.Msg = &MsgSendWrapToken{}

func NewMsgSendWrapToken(creator string, ethAddress string, amount sdk.Coin) *MsgSendWrapToken {
	return &MsgSendWrapToken{
		Creator:    creator,
		EthAddress: ethAddress,
		Amount:     amount,
	}
}

func (msg *MsgSendWrapToken) Route() string {
	return RouterKey
}

func (msg *MsgSendWrapToken) Type() string {
	return TypeMsgSendWrapToken
}

func (msg *MsgSendWrapToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendWrapToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendWrapToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
