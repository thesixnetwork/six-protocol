package types

import (
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
	virtualSchemaRegistry []VirtualSchemaRegistryRequest,
	actions []Action,
	enable bool,
) *MsgProposalVirtualSchema {
	actionsPointer := make([]*Action, len(actions))
	for i, action := range actions {
		actionsPointer[i] = &action
	}

	registryPointer := make([]*VirtualSchemaRegistryRequest, len(virtualSchemaRegistry))
	for i, registry := range virtualSchemaRegistry {
		registryPointer[i] = &registry
	}
	return &MsgProposalVirtualSchema{
		Creator:              creator,
		VirtualNftSchemaCode: schemacode,
		ProposalType:         proposalType,
		Registry:             registryPointer,
		Actions:              actionsPointer,
		Enable:               enable,
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
