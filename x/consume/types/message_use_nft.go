package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUseNft = "use_nft"

var _ sdk.Msg = &MsgUseNft{}

func NewMsgUseNft(creator string, token string) *MsgUseNft {
	return &MsgUseNft{
		Creator: creator,
		Token:   token,
	}
}

func (msg *MsgUseNft) Route() string {
	return RouterKey
}

func (msg *MsgUseNft) Type() string {
	return TypeMsgUseNft
}

func (msg *MsgUseNft) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUseNft) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUseNft) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
