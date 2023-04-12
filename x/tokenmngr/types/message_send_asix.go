package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendAsix = "send_asix"

var _ sdk.Msg = &MsgSendAsix{}

func NewMsgSendAsix(creator string, ethAddress string, amount sdk.Coin) *MsgSendAsix {
	return &MsgSendAsix{
		Creator:    creator,
		EthAddress: ethAddress,
		Amount:     amount,
	}
}

func (msg *MsgSendAsix) Route() string {
	return RouterKey
}

func (msg *MsgSendAsix) Type() string {
	return TypeMsgSendAsix
}

func (msg *MsgSendAsix) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendAsix) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendAsix) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
