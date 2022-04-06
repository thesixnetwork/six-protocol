package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateGroup{}, "protocoladmin/CreateGroup", nil)
	cdc.RegisterConcrete(&MsgUpdateGroup{}, "protocoladmin/UpdateGroup", nil)
	cdc.RegisterConcrete(&MsgDeleteGroup{}, "protocoladmin/DeleteGroup", nil)
	cdc.RegisterConcrete(&MsgAddAdminToGroup{}, "protocoladmin/AddAdminToGroup", nil)
	cdc.RegisterConcrete(&MsgRemoveAdminFromGroup{}, "protocoladmin/RemoveAdminFromGroup", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGroup{},
		&MsgUpdateGroup{},
		&MsgDeleteGroup{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddAdminToGroup{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveAdminFromGroup{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
