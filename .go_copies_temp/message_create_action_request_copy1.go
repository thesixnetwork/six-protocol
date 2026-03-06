package types

import (
	errormod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateActionRequest = "create_action_request"

var _ sdk.Msg = &MsgCreateActionRequest{}

func NewMsgCreateActionRequest(creator string, vm string, base64ActionSignature string, requiredConfirm uint64) *MsgCreateActionRequest {
	return &MsgCreateActionRequest{
		Creator:               creator,
		Vm:                    vm,
		Base64ActionSignature: base64ActionSignature,
		RequiredConfirm:       requiredConfirm,
	}
}

func (msg *MsgCreateActionRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateActionRequest) Type() string {
	return TypeMsgCreateActionRequest
}

func (msg *MsgCreateActionRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateActionRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateActionRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
