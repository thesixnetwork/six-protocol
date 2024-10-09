package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetBaseUri = "set_base_uri"

var _ sdk.Msg = &MsgSetBaseUri{}

func NewMsgSetBaseUri(creator string, code string, newBaseUri string) *MsgSetBaseUri {
	return &MsgSetBaseUri{
		Creator:    creator,
		Code:       code,
		NewBaseUri: newBaseUri,
	}
}

func (msg *MsgSetBaseUri) Route() string {
	return RouterKey
}

func (msg *MsgSetBaseUri) Type() string {
	return TypeMsgSetBaseUri
}

func (msg *MsgSetBaseUri) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetBaseUri) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetBaseUri) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
