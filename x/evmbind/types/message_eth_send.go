package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEthSend = "eth_send"

var _ sdk.Msg = &MsgEthSend{}

func NewMsgEthSend(creator string, fromEth string, toEth string, amount string) *MsgEthSend {
	return &MsgEthSend{
		Creator: creator,
		FromEth: fromEth,
		ToEth:   toEth,
		Amount:  amount,
	}
}

func (msg *MsgEthSend) Route() string {
	return RouterKey
}

func (msg *MsgEthSend) Type() string {
	return TypeMsgEthSend
}

func (msg *MsgEthSend) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEthSend) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEthSend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
