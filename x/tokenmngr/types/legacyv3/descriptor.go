package legacyv3

import (
	"bytes"
	"compress/gzip"

	protov2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// The Cosmos SDK tx decoder runs codec/unknownproto.RejectUnknownFields over
// every message inside a tx BEFORE any Go-level Unmarshal. That check does not
// deserialize; it walks the raw wire bytes and, for each field tag, asserts the
// on-wire wire-type is compatible with the field type declared in the message's
// protobuf Descriptor() (see canEncodeType / the `checks` table in
// codec/unknownproto/unknown_fields.go).
//
// For MsgMint/MsgBurn this is exactly where historical txs died: the current
// generated descriptor declares field 2 (`amount`) as a Coin message
// (length-delimited, wire type 2), but the v1-era on-chain bytes carry a bare
// uint64 (varint, wire type 0). unknownproto rejects it with
//
//	Mismatched "*legacyv3.MsgMint": {TagNum: 2, GotWireType: "varint" != WantWireType: "bytes"}
//
// before mintBurn.Unmarshal ever runs.
//
// The fix: hand these legacy types a Descriptor() that declares field 2 as
// `uint64`. A uint64 field is packable, so unknownproto's `checks` table accepts
// BOTH wire type 0 (a single varint — the v1 layout) AND wire type 2 (a packed
// run — indistinguishable at the wire-type level from the v2/v3 Coin bytes).
// Because uint64 is scalar, unknownproto also does NOT recurse into the field's
// content, so the v2/v3 Coin bytes pass through untouched. The real
// disambiguation still happens afterwards in mintBurn.Unmarshal, which branches
// on the actual wire type to recover either the uint64 or the Coin.
//
// Field layout described here (the v1 superset — every era is a subset):
//
//	1: creator string
//	2: amount  uint64   // v1 uint64 OR v2/v3 Coin, both wire-compatible
//	3: token   string   // v1 only
//
// The message name in the descriptor is irrelevant to unknownproto (it keys
// purely on field number + type), so MsgMint and MsgBurn share one descriptor.
var mintBurnFileDescriptorGz = buildMintBurnDescriptor()

func buildMintBurnDescriptor() []byte {
	fd := &descriptorpb.FileDescriptorProto{
		Name:    protov2.String("thesixnetwork/sixprotocol/tokenmngr/legacy_mintburn.proto"),
		Package: protov2.String("thesixnetwork.sixprotocol.tokenmngr"),
		Syntax:  protov2.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: protov2.String("LegacyMintBurn"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{
						Name:   protov2.String("creator"),
						Number: protov2.Int32(1),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					},
					{
						Name:   protov2.String("amount"),
						Number: protov2.Int32(2),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_UINT64.Enum(),
					},
					{
						Name:   protov2.String("token"),
						Number: protov2.Int32(3),
						Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
						Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
					},
				},
			},
		},
	}

	blob, err := protov2.Marshal(fd)
	if err != nil {
		panic("legacyv3: failed to marshal mint/burn file descriptor: " + err.Error())
	}

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(blob); err != nil {
		panic("legacyv3: failed to gzip mint/burn file descriptor: " + err.Error())
	}
	if err := zw.Close(); err != nil {
		panic("legacyv3: failed to close gzip writer: " + err.Error())
	}
	return buf.Bytes()
}

// Descriptor satisfies the descriptorIface that unknownproto requires. The
// returned gzipped FileDescriptorProto declares field 2 as uint64 so both the
// v1 (varint) and v2/v3 (Coin) wire layouts survive the strict unknown-field
// check. MsgMint and MsgBurn both embed mintBurn, so both inherit this.
func (*mintBurn) Descriptor() ([]byte, []int) { return mintBurnFileDescriptorGz, []int{0} }
