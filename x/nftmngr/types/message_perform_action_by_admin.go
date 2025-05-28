package types

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPerformActionByAdmin{}

const TypeMsgPerformActionByAdmin = "perform_action_by_admin"

func NewMsgPerformActionByAdmin(creator, nftSchemaCode, tokenId, action, actionPrams, refId string) *MsgPerformActionByAdmin {
	// string to json object of  []*ActionParameter
	var actionPrams_ []*ActionParameter
	err := json.Unmarshal([]byte(actionPrams), &actionPrams_)
	if err != nil {
		panic(err)
	}

	return &MsgPerformActionByAdmin{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		TokenId:       tokenId,
		Action:        action,
		Parameters:    actionPrams_,
		RefId:         refId,
	}
}

func (msg *MsgPerformActionByAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgPerformActionByAdmin) Route() string {
	return RouterKey
}

func (msg *MsgPerformActionByAdmin) Type() string {
	return TypeMsgPerformActionByAdmin
}

func (msg *MsgPerformActionByAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPerformActionByAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
