package types

import (
	// "encoding/json"
	// fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPerformMultiTokenAction = "perform_multi_token_action"

var _ sdk.Msg = &MsgPerformMultiTokenAction{}

func NewMsgPerformMultiTokenAction(creator string, nftSchemaCode string, tokenId []string, action []string, refId string, parameters string) *MsgPerformMultiTokenAction {
	// // string to json object of  []*ActionParameter
	// var actionPrams_ []*ActionParameter
	// err := json.Unmarshal([]byte(parameters), &actionPrams_)
	// if err != nil {
	// 	fmt.Println("Error in NewMsgPerformActionByAdmin: ", err)
	// }

	return &MsgPerformMultiTokenAction{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		TokenId:       tokenId,
		Action:        action,
		RefId:         refId,
		Parameters:    parameters,
	}
}

func (msg *MsgPerformMultiTokenAction) Route() string {
	return RouterKey
}

func (msg *MsgPerformMultiTokenAction) Type() string {
	return TypeMsgPerformMultiTokenAction
}

func (msg *MsgPerformMultiTokenAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPerformMultiTokenAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPerformMultiTokenAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
