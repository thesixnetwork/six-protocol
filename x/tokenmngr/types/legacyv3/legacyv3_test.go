package legacyv3_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"

	tokenmngrtypes "github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types/legacyv3"
)

const oldMintURL = "/thesixnetwork.sixprotocol.tokenmngr.MsgMint"
const oldCreateOptionsURL = "/thesixnetwork.sixprotocol.tokenmngr.MsgCreateOptions"

func newRegistry() codectypes.InterfaceRegistry {
	reg := codectypes.NewInterfaceRegistry()
	tokenmngrtypes.RegisterInterfaces(reg)
	legacyv3.RegisterInterfaces(reg)
	return reg
}

// v2/v3-era mint: amount is a Coin (field 2, wire type 2). Must resolve under
// the old type URL and render the Coin.
func TestLegacyMintCoinEra(t *testing.T) {
	reg := newRegistry()
	cdc := codec.NewProtoCodec(reg)

	orig := &tokenmngrtypes.MsgMint{
		Creator: "6x1abcdef",
		Amount:  sdk.NewCoin("usix", sdkmath.NewInt(1_000_000)),
	}
	val, err := orig.Marshal()
	require.NoError(t, err)

	got := unpackMint(t, reg, val)
	require.False(t, got.IsLegacy)
	require.Equal(t, orig.Creator, got.Creator)
	require.Equal(t, orig.Amount, got.Amount)

	out, err := cdc.MarshalJSON(got)
	require.NoError(t, err)
	t.Logf("Coin-era MsgMint JSON: %s", out)
	require.Contains(t, string(out), "6x1abcdef")
	require.Contains(t, string(out), "1000000")
	require.Contains(t, string(out), "usix")
}

// v1-era mint: amount is a bare uint64 (field 2, wire type 0) plus a token
// string at field 3. This is the layout in the production tx that failed with
// "GotWireType varint != WantWireType bytes". It must now decode.
func TestLegacyMintUint64Era(t *testing.T) {
	reg := newRegistry()
	cdc := codec.NewProtoCodec(reg)

	// Hand-build v1 wire bytes: creator=1(string), amount=2(uint64), token=3(string).
	val := buildV1MintBurn("6x1creator", 1_000_000, "usix")

	got := unpackMint(t, reg, val)
	require.True(t, got.IsLegacy, "field-2 varint must be detected as v1 layout")
	require.Equal(t, "6x1creator", got.Creator)
	require.Equal(t, uint64(1_000_000), got.LegacyAmount)
	require.Equal(t, "usix", got.Token)

	// Verify the full Any display path (what `sixd q tx` emits): the @type
	// wrapper plus the original fields.
	anyMsg := &codectypes.Any{TypeUrl: oldMintURL, Value: val}
	anyJSON, err := cdc.MarshalJSON(anyMsg)
	require.NoError(t, err)
	t.Logf("Any JSON: %s", anyJSON)
	require.Contains(t, string(anyJSON), oldMintURL)
	require.Contains(t, string(anyJSON), "6x1creator")

	// Round-trips back to the exact original bytes.
	remarshaled, err := got.Marshal()
	require.NoError(t, err)
	require.Equal(t, val, remarshaled)

	out, err := cdc.MarshalJSON(got)
	require.NoError(t, err)
	t.Logf("v1 MsgMint JSON: %s", out)
	require.Contains(t, string(out), "6x1creator")
	require.Contains(t, string(out), "1000000")
	require.Contains(t, string(out), "usix")
}

func unpackMint(t *testing.T, reg codectypes.InterfaceRegistry, val []byte) *legacyv3.MsgMint {
	t.Helper()
	anyMsg := &codectypes.Any{TypeUrl: oldMintURL, Value: val}
	var msg sdk.Msg
	require.NoError(t, reg.UnpackAny(anyMsg, &msg))
	got, ok := msg.(*legacyv3.MsgMint)
	require.True(t, ok, "resolved to %T", msg)
	return got
}

// buildV1MintBurn encodes the v1 MsgMint/MsgBurn wire layout.
func buildV1MintBurn(creator string, amount uint64, token string) []byte {
	var b []byte
	// field 1: creator (string)
	b = append(b, 0x0a, byte(len(creator)))
	b = append(b, creator...)
	// field 2: amount (uint64 varint)
	b = append(b, 0x10)
	for amount >= 0x80 {
		b = append(b, byte(amount)|0x80)
		amount >>= 7
	}
	b = append(b, byte(amount))
	// field 3: token (string)
	b = append(b, 0x1a, byte(len(token)))
	b = append(b, token...)
	return b
}

// MsgCreateOptions moved defaultMintee from field 2 (v3) to field 3 (v4).
// The legacy type must read field 2; the current type would drop it.
func TestLegacyCreateOptionsFieldTwo(t *testing.T) {
	reg := newRegistry()

	legacyMsg := &legacyv3.MsgCreateOptions{Creator: "6x1creator", DefaultMintee: "6x1mintee"}
	val, err := legacyMsg.Marshal()
	require.NoError(t, err)

	// Byte layout sanity: field1(0x0a) + field2(0x12), defaultMintee at field 2.
	require.Equal(t, byte(0x0a), val[0])
	require.Contains(t, string(val), "6x1mintee")

	// Round-trip through the registry under the old type URL.
	anyMsg := &codectypes.Any{TypeUrl: oldCreateOptionsURL, Value: val}
	var msg sdk.Msg
	require.NoError(t, reg.UnpackAny(anyMsg, &msg))
	got := msg.(*legacyv3.MsgCreateOptions)
	require.Equal(t, "6x1creator", got.Creator)
	require.Equal(t, "6x1mintee", got.DefaultMintee)

	// Prove the regression the legacy type avoids: the CURRENT type, which
	// expects defaultMintee at field 3, silently loses it when reading v3 bytes.
	cur := &tokenmngrtypes.MsgCreateOptions{}
	require.NoError(t, cur.Unmarshal(val))
	require.Equal(t, "6x1creator", cur.Creator)
	require.Empty(t, cur.DefaultMintee, "current type drops v3 defaultMintee (field 2)")
}

// Guards the XXX_MessageName wiring used to derive the old type URLs.
func TestLegacyMessageNames(t *testing.T) {
	require.Equal(t, "thesixnetwork.sixprotocol.tokenmngr.MsgMint", proto.MessageName(&legacyv3.MsgMint{}))
	require.Equal(t, "thesixnetwork.sixprotocol.tokenmngr.MsgBurn", proto.MessageName(&legacyv3.MsgBurn{}))
	require.Equal(t, "thesixnetwork.sixprotocol.tokenmngr.MsgCreateOptions", proto.MessageName(&legacyv3.MsgCreateOptions{}))
}
