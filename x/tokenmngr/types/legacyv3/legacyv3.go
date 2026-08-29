// Package legacyv3 registers the tokenmngr message types under their
// pre-v4 protobuf package name so that historical transactions can still be
// decoded and queried.
//
// Background: from v2.0.0 through v3.3.1 the tokenmngr proto package was
// `thesixnetwork.sixprotocol.tokenmngr`. In v4.0.0 it was renamed to
// `sixprotocol.tokenmngr`. The protobuf field layout of the messages did not
// change (with two exceptions noted below), but the package name is embedded
// in the `Any` type URL of every message stored inside a transaction. After
// the rename, the v4 InterfaceRegistry no longer knows how to resolve the old
// type URLs, so `sixd q tx <hash>` on a v2/v3 mint/burn tx fails with:
//
//	unable to resolve type URL /thesixnetwork.sixprotocol.tokenmngr.MsgMint
//
// This package defines thin shim types that report the OLD proto name via
// XXX_MessageName() while reusing the CURRENT wire marshaling, and registers
// them on the InterfaceRegistry. Because the wire bytes are identical, old
// transactions decode correctly again.
//
// Two messages changed field numbers between v3 and v4 (defaultMintee moved
// from field 2 to field 3), so MsgCreateOptions/MsgUpdateOptions carry their
// own v3 field layout instead of embedding the current type.
package legacyv3

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/jsonpb"
	proto "github.com/cosmos/gogoproto/proto"

	tokenmngrtypes "github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
)

// legacyPkg is the protobuf package name used for tokenmngr messages in
// v2.0.0 through v3.3.1.
const legacyPkg = "thesixnetwork.sixprotocol.tokenmngr."

// marshalJSONPB / unmarshalJSONPB let the embedded-alias types delegate JSON
// rendering to their current (tagged) type. Without this, gogoproto jsonpb
// walks only protobuf-tagged fields and renders the alias as `{}`, so
// `sixd q tx` would show an empty message body.
func marshalJSONPB(m *jsonpb.Marshaler, cur proto.Message) ([]byte, error) {
	s, err := m.MarshalToString(cur)
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func unmarshalJSONPB(u *jsonpb.Unmarshaler, data []byte, cur proto.Message) error {
	return u.Unmarshal(bytes.NewReader(data), cur)
}

// ---------------------------------------------------------------------------
// Wire-identical messages: embed the current type, override only the reported
// proto message name so the old type URL resolves. Marshal/Unmarshal/Size and
// all proto.Message methods are promoted from the embedded current type.
// ---------------------------------------------------------------------------

type MsgCreateToken struct{ tokenmngrtypes.MsgCreateToken }

func (*MsgCreateToken) XXX_MessageName() string { return legacyPkg + "MsgCreateToken" }
func (m *MsgCreateToken) MarshalJSONPB(mr *jsonpb.Marshaler) ([]byte, error) {
	return marshalJSONPB(mr, &m.MsgCreateToken)
}

func (m *MsgCreateToken) UnmarshalJSONPB(u *jsonpb.Unmarshaler, d []byte) error {
	return unmarshalJSONPB(u, d, &m.MsgCreateToken)
}

type MsgUpdateToken struct{ tokenmngrtypes.MsgUpdateToken }

func (*MsgUpdateToken) XXX_MessageName() string { return legacyPkg + "MsgUpdateToken" }
func (m *MsgUpdateToken) MarshalJSONPB(mr *jsonpb.Marshaler) ([]byte, error) {
	return marshalJSONPB(mr, &m.MsgUpdateToken)
}

func (m *MsgUpdateToken) UnmarshalJSONPB(u *jsonpb.Unmarshaler, d []byte) error {
	return unmarshalJSONPB(u, d, &m.MsgUpdateToken)
}

type MsgDeleteToken struct{ tokenmngrtypes.MsgDeleteToken }

func (*MsgDeleteToken) XXX_MessageName() string { return legacyPkg + "MsgDeleteToken" }
func (m *MsgDeleteToken) MarshalJSONPB(mr *jsonpb.Marshaler) ([]byte, error) {
	return marshalJSONPB(mr, &m.MsgDeleteToken)
}

func (m *MsgDeleteToken) UnmarshalJSONPB(u *jsonpb.Unmarshaler, d []byte) error {
	return unmarshalJSONPB(u, d, &m.MsgDeleteToken)
}

type MsgCreateMintperm struct {
	tokenmngrtypes.MsgCreateMintperm
}

func (*MsgCreateMintperm) XXX_MessageName() string { return legacyPkg + "MsgCreateMintperm" }
func (m *MsgCreateMintperm) MarshalJSONPB(mr *jsonpb.Marshaler) ([]byte, error) {
	return marshalJSONPB(mr, &m.MsgCreateMintperm)
}

func (m *MsgCreateMintperm) UnmarshalJSONPB(u *jsonpb.Unmarshaler, d []byte) error {
	return unmarshalJSONPB(u, d, &m.MsgCreateMintperm)
}

type MsgUpdateMintperm struct {
	tokenmngrtypes.MsgUpdateMintperm
}

func (*MsgUpdateMintperm) XXX_MessageName() string { return legacyPkg + "MsgUpdateMintperm" }
func (m *MsgUpdateMintperm) MarshalJSONPB(mr *jsonpb.Marshaler) ([]byte, error) {
	return marshalJSONPB(mr, &m.MsgUpdateMintperm)
}

func (m *MsgUpdateMintperm) UnmarshalJSONPB(u *jsonpb.Unmarshaler, d []byte) error {
	return unmarshalJSONPB(u, d, &m.MsgUpdateMintperm)
}

type MsgDeleteMintperm struct {
	tokenmngrtypes.MsgDeleteMintperm
}

func (*MsgDeleteMintperm) XXX_MessageName() string { return legacyPkg + "MsgDeleteMintperm" }
func (m *MsgDeleteMintperm) MarshalJSONPB(mr *jsonpb.Marshaler) ([]byte, error) {
	return marshalJSONPB(mr, &m.MsgDeleteMintperm)
}

func (m *MsgDeleteMintperm) UnmarshalJSONPB(u *jsonpb.Unmarshaler, d []byte) error {
	return unmarshalJSONPB(u, d, &m.MsgDeleteMintperm)
}

// MsgMint and MsgBurn live in mintburn.go: their amount field flipped from
// uint64 (v1) to Coin (v2.2+) under the same type URL, so they need a
// wire-type-aware Unmarshal rather than a simple embed-alias.

type MsgDeleteOptions struct {
	tokenmngrtypes.MsgDeleteOptions
}

func (*MsgDeleteOptions) XXX_MessageName() string { return legacyPkg + "MsgDeleteOptions" }
func (m *MsgDeleteOptions) MarshalJSONPB(mr *jsonpb.Marshaler) ([]byte, error) {
	return marshalJSONPB(mr, &m.MsgDeleteOptions)
}

func (m *MsgDeleteOptions) UnmarshalJSONPB(u *jsonpb.Unmarshaler, d []byte) error {
	return unmarshalJSONPB(u, d, &m.MsgDeleteOptions)
}

type MsgWrapToken struct{ tokenmngrtypes.MsgWrapToken }

func (*MsgWrapToken) XXX_MessageName() string { return legacyPkg + "MsgWrapToken" }
func (m *MsgWrapToken) MarshalJSONPB(mr *jsonpb.Marshaler) ([]byte, error) {
	return marshalJSONPB(mr, &m.MsgWrapToken)
}

func (m *MsgWrapToken) UnmarshalJSONPB(u *jsonpb.Unmarshaler, d []byte) error {
	return unmarshalJSONPB(u, d, &m.MsgWrapToken)
}

type MsgUnwrapToken struct{ tokenmngrtypes.MsgUnwrapToken }

func (*MsgUnwrapToken) XXX_MessageName() string { return legacyPkg + "MsgUnwrapToken" }
func (m *MsgUnwrapToken) MarshalJSONPB(mr *jsonpb.Marshaler) ([]byte, error) {
	return marshalJSONPB(mr, &m.MsgUnwrapToken)
}

func (m *MsgUnwrapToken) UnmarshalJSONPB(u *jsonpb.Unmarshaler, d []byte) error {
	return unmarshalJSONPB(u, d, &m.MsgUnwrapToken)
}

type MsgSendWrapToken struct {
	tokenmngrtypes.MsgSendWrapToken
}

func (*MsgSendWrapToken) XXX_MessageName() string { return legacyPkg + "MsgSendWrapToken" }
func (m *MsgSendWrapToken) MarshalJSONPB(mr *jsonpb.Marshaler) ([]byte, error) {
	return marshalJSONPB(mr, &m.MsgSendWrapToken)
}

func (m *MsgSendWrapToken) UnmarshalJSONPB(u *jsonpb.Unmarshaler, d []byte) error {
	return unmarshalJSONPB(u, d, &m.MsgSendWrapToken)
}

// RegisterInterfaces registers the legacy (v2/v3) tokenmngr message type URLs
// on the given registry so historical transactions can be decoded.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateToken{},
		&MsgUpdateToken{},
		&MsgDeleteToken{},
		&MsgCreateMintperm{},
		&MsgUpdateMintperm{},
		&MsgDeleteMintperm{},
		&MsgMint{},
		&MsgCreateOptions{},
		&MsgUpdateOptions{},
		&MsgDeleteOptions{},
		&MsgBurn{},
		&MsgWrapToken{},
		&MsgUnwrapToken{},
		&MsgSendWrapToken{},
	)
}
