package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetOriginContract = "set_origin_contract"

var _ sdk.Msg = &MsgSetOriginContract{}

func NewMsgSetOriginContract(creator string, schemaCode string, newContractAddress string) *MsgSetOriginContract {
	return &MsgSetOriginContract{
		Creator:            creator,
		SchemaCode:         schemaCode,
		NewContractAddress: newContractAddress,
	}
}

func (msg *MsgSetOriginContract) Route() string {
	return RouterKey
}

func (msg *MsgSetOriginContract) Type() string {
	return TypeMsgSetOriginContract
}

func (msg *MsgSetOriginContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetOriginContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetOriginContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
