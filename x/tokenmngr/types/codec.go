package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

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
		&MsgCreateOptions{},
		&MsgUpdateOptions{},
		&MsgDeleteOptions{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMint{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurn{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWrapToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnwrapToken{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendWrapToken{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
