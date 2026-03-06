package types

import (
	errormod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitActionResponse = "submit_action_response"

var _ sdk.Msg = &MsgSubmitActionResponse{}

func NewMsgSubmitActionResponse(creator string, actionRequestID uint64, base64NftData string) *MsgSubmitActionResponse {
	return &MsgSubmitActionResponse{
		Creator:         creator,
		ActionRequestID: actionRequestID,
		Base64NftData:   base64NftData,
	}
}

func (msg *MsgSubmitActionResponse) Route() string {
	return RouterKey
}

func (msg *MsgSubmitActionResponse) Type() string {
	return TypeMsgSubmitActionResponse
}

func (msg *MsgSubmitActionResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitActionResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitActionResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
