package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEnableContractConverter = "enable_contract_converter"

var _ sdk.Msg = &MsgEnableContractConverter{}

func NewMsgEnableContractConverter(creator string, enable bool) *MsgEnableContractConverter {
	return &MsgEnableContractConverter{
		Creator: creator,
		Enable:  enable,
	}
}

func (msg *MsgEnableContractConverter) Route() string {
	return RouterKey
}

func (msg *MsgEnableContractConverter) Type() string {
	return TypeMsgEnableContractConverter
}

func (msg *MsgEnableContractConverter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnableContractConverter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableContractConverter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
