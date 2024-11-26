package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetOriginChain = "set_origin_chain"

var _ sdk.Msg = &MsgSetOriginChain{}

func NewMsgSetOriginChain(creator string, schemaCode string, newOriginChain string) *MsgSetOriginChain {
	return &MsgSetOriginChain{
		Creator:        creator,
		SchemaCode:     schemaCode,
		NewOriginChain: newOriginChain,
	}
}

func (msg *MsgSetOriginChain) Route() string {
	return RouterKey
}

func (msg *MsgSetOriginChain) Type() string {
	return TypeMsgSetOriginChain
}

func (msg *MsgSetOriginChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetOriginChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetOriginChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
