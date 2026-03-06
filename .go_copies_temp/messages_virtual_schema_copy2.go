package types

import (
	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgProposalVirtualSchema = "virtual_schema_proposal"
)

var _ sdk.Msg = &MsgProposalVirtualSchema{}

func NewMsgProposalVirtualSchema(
	creator string,
	schemacode string,
	proposalType ProposalType,
	virtualSchemaRegistry []*VirtualSchemaRegistryRequest,
	actions []*Action,
	executors []string,
	enable bool,
) *MsgProposalVirtualSchema {
	return &MsgProposalVirtualSchema{
		Creator:              creator,
		VirtualNftSchemaCode: schemacode,
		ProposalType:         proposalType,
		Registry:             virtualSchemaRegistry,
		Actions:              actions,
		Enable:               enable,
		Executors:            executors,
	}
}

func (msg *MsgProposalVirtualSchema) Route() string {
	return RouterKey
}

func (msg *MsgProposalVirtualSchema) Type() string {
	return TypeMsgProposalVirtualSchema
}

func (msg *MsgProposalVirtualSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgProposalVirtualSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposalVirtualSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
