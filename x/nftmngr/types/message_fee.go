package types

import (
	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetFeeConfig = "set_fee_config"

var _ sdk.Msg = &MsgSetFeeConfig{}

func NewMsgSetFeeConfig(creator string, feeConfig *FeeConfig) *MsgSetFeeConfig {
	return &MsgSetFeeConfig{
		Creator:   creator,
		FeeConfig: feeConfig,
	}
}

func (msg *MsgSetFeeConfig) Route() string {
	return RouterKey
}

func (msg *MsgSetFeeConfig) Type() string {
	return TypeMsgSetFeeConfig
}

func (msg *MsgSetFeeConfig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetFeeConfig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetFeeConfig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
