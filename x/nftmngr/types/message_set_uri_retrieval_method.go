package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetUriRetrievalMethod = "set_uri_retrieval_method"

var _ sdk.Msg = &MsgSetUriRetrievalMethod{}

func NewMsgSetUriRetrievalMethod(creator string, schemaCode string, newMethod int32) *MsgSetUriRetrievalMethod {
	return &MsgSetUriRetrievalMethod{
		Creator:    creator,
		SchemaCode: schemaCode,
		NewMethod:  newMethod,
	}
}

func (msg *MsgSetUriRetrievalMethod) Route() string {
	return RouterKey
}

func (msg *MsgSetUriRetrievalMethod) Type() string {
	return TypeMsgSetUriRetrievalMethod
}

func (msg *MsgSetUriRetrievalMethod) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetUriRetrievalMethod) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetUriRetrievalMethod) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
