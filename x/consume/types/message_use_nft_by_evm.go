package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUseNftByEVM = "use_nft_by_evm"

var _ sdk.Msg = &MsgUseNftByEVM{}

func NewMsgUseNftByEVM(creator string, token string, ethAddress string, ethSignature string, signMessage string) *MsgUseNftByEVM {
	return &MsgUseNftByEVM{
		Creator:      creator,
		Token:        token,
		EthAddress:   ethAddress,
		EthSignature: ethSignature,
		SignMessage:  signMessage,
	}
}

func (msg *MsgUseNftByEVM) Route() string {
	return RouterKey
}

func (msg *MsgUseNftByEVM) Type() string {
	return TypeMsgUseNftByEVM
}

func (msg *MsgUseNftByEVM) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUseNftByEVM) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUseNftByEVM) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
