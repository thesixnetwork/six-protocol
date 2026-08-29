package legacyv3_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/codec/unknownproto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/stretchr/testify/require"

	tokenmngrtypes "github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
)

// The tx decoder (x/auth/tx/decoder.go) runs unknownproto.RejectUnknownFields
// over the whole TxBody before any Go Unmarshal, resolving each message Any via
// the InterfaceRegistry and validating its wire types against the resolved
// type's Descriptor(). This is the exact gate that rejected historical mints
// with `Mismatched ... GotWireType: "varint" != WantWireType: "bytes"`.
//
// These tests wrap a legacy MsgMint (both eras) in a real TxBody and push it
// through that same gate to prove the legacy descriptor lets both survive.

func txBodyBytesWithMintValue(t *testing.T, val []byte) []byte {
	t.Helper()
	body := &txtypes.TxBody{
		Messages: []*codectypes.Any{{TypeUrl: oldMintURL, Value: val}},
	}
	bz, err := body.Marshal()
	require.NoError(t, err)
	return bz
}

// v1-era mint: amount is a bare uint64 (field 2, wire type 0). This is the
// production tx that failed. It must now pass unknownproto.
func TestUnknownProtoAcceptsV1Mint(t *testing.T) {
	reg := newRegistry()
	val := buildV1MintBurn("6x1creator", 1_000_000, "usix")

	bodyBz := txBodyBytesWithMintValue(t, val)

	_, err := unknownproto.RejectUnknownFields(bodyBz, &txtypes.TxBody{}, true, reg)
	require.NoError(t, err, "v1 uint64 mint must pass the strict tx-decoder gate")
}

// v2/v3-era mint: amount is a Coin (field 2, wire type 2 / length-delimited).
// The uint64 descriptor accepts wire type 2 too (uint64 is packable), so this
// must also pass.
func TestUnknownProtoAcceptsCoinMint(t *testing.T) {
	reg := newRegistry()
	orig := &tokenmngrtypes.MsgMint{
		Creator: "6x1abcdef",
		Amount:  sdk.NewCoin("usix", sdkmath.NewInt(1_000_000)),
	}
	val, err := orig.Marshal()
	require.NoError(t, err)

	bodyBz := txBodyBytesWithMintValue(t, val)

	_, err = unknownproto.RejectUnknownFields(bodyBz, &txtypes.TxBody{}, true, reg)
	require.NoError(t, err, "v2/v3 Coin mint must pass the strict tx-decoder gate")
}
