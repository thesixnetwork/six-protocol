package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnwrapToken = "unwrap_token"

var _ sdk.Msg = &MsgUnwrapToken{}

func NewMsgUnwrapToken(creator string, amount sdk.Coin, receiver string) *MsgUnwrapToken {
	return &MsgUnwrapToken{
		Creator:  creator,
		Amount:   amount,
		Receiver: receiver,
	}
}

func (msg *MsgUnwrapToken) Route() string {
	return RouterKey
}

func (msg *MsgUnwrapToken) Type() string {
	return TypeMsgUnwrapToken
}

func (msg *MsgUnwrapToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnwrapToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnwrapToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
