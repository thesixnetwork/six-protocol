package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetConverterParams = "set_converter_params"

var _ sdk.Msg = &MsgSetConverterParams{}

func NewMsgSetConverterParams(creator string, contractAddress string, eventName string, eventTuple string, abi string) *MsgSetConverterParams {
	return &MsgSetConverterParams{
		Creator:         creator,
		ContractAddress: contractAddress,
		EventName:       eventName,
		EventTuple:      eventTuple,
		Abi:             abi,
	}
}

func (msg *MsgSetConverterParams) Route() string {
	return RouterKey
}

func (msg *MsgSetConverterParams) Type() string {
	return TypeMsgSetConverterParams
}

func (msg *MsgSetConverterParams) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetConverterParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetConverterParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
