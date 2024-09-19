package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

const TypeMsgMint = "mint"

var _ sdk.Msg = &MsgMintLegacy{}

func NewMsgMintLagacy(creator string, amount uint64, token string) *MsgMintLegacy {
	return &MsgMintLegacy{
		Creator: creator,
		Amount:  amount,
		Token:   token,
	}
}

func (msg *MsgMintLegacy ) Route() string {
	return RouterKey
}

func (msg *MsgMintLegacy ) Type() string {
	return TypeMsgMint
}

func (msg *MsgMintLegacy ) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintLegacy ) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintLegacy ) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}