package legacy

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateToken{}, "tokenmngr/CreateToken", nil)
	cdc.RegisterConcrete(&MsgUpdateToken{}, "tokenmngr/UpdateToken", nil)
	cdc.RegisterConcrete(&MsgDeleteToken{}, "tokenmngr/DeleteToken", nil)
	cdc.RegisterConcrete(&MsgCreateMintperm{}, "tokenmngr/CreateMintperm", nil)
	cdc.RegisterConcrete(&MsgUpdateMintperm{}, "tokenmngr/UpdateMintperm", nil)
	cdc.RegisterConcrete(&MsgDeleteMintperm{}, "tokenmngr/DeleteMintperm", nil)
	cdc.RegisterConcrete(&MsgMint{}, "tokenmngr/Mint", nil)
	cdc.RegisterConcrete(&MsgCreateOptions{}, "tokenmngr/CreateOptions", nil)
	cdc.RegisterConcrete(&MsgUpdateOptions{}, "tokenmngr/UpdateOptions", nil)
	cdc.RegisterConcrete(&MsgDeleteOptions{}, "tokenmngr/DeleteOptions", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "tokenmngr/Burn", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateToken{},
		&MsgUpdateToken{},
		&MsgDeleteToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMintperm{},
		&MsgUpdateMintperm{},
		&MsgDeleteMintperm{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMint{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOptions{},
		&MsgUpdateOptions{},
		&MsgDeleteOptions{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurn{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
