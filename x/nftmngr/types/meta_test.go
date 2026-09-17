package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

// --- attribute constructors ---

func numberAttr(name string, v uint64) *types.NftAttributeValue {
	return &types.NftAttributeValue{
		Name: name,
		Value: &types.NftAttributeValue_NumberAttributeValue{
			NumberAttributeValue: &types.NumberAttributeValue{Value: v},
		},
	}
}

func floatAttr(name string, v float64) *types.NftAttributeValue {
	return &types.NftAttributeValue{
		Name: name,
		Value: &types.NftAttributeValue_FloatAttributeValue{
			FloatAttributeValue: &types.FloatAttributeValue{Value: v},
		},
	}
}

func stringAttr(name, v string) *types.NftAttributeValue {
	return &types.NftAttributeValue{
		Name: name,
		Value: &types.NftAttributeValue_StringAttributeValue{
			StringAttributeValue: &types.StringAttributeValue{Value: v},
		},
	}
}

func boolAttr(name string, v bool) *types.NftAttributeValue {
	return &types.NftAttributeValue{
		Name: name,
		Value: &types.NftAttributeValue_BooleanAttributeValue{
			BooleanAttributeValue: &types.BooleanAttributeValue{Value: v},
		},
	}
}

// newChainMeta builds a Metadata whose attributes are all "chain" (settable).
func newChainMeta(tokenID string, onchain ...*types.NftAttributeValue) *types.Metadata {
	token := &types.NftData{TokenId: tokenID, OnchainAttributes: onchain}
	return types.NewMetadata(&types.NFTSchema{}, token, types.AttributeOverriding_CHAIN, nil)
}

// --- #2: setters return an error (not panic) on a missing key ---

func TestSettersReturnErrorOnMissingKey(t *testing.T) {
	meta := newChainMeta("t1") // no attributes

	require.NotPanics(t, func() {
		require.ErrorIs(t, meta.SetNumber("nope", 1), types.ErrAttributeNotFoundForAction)
		require.ErrorIs(t, meta.SetString("nope", "x"), types.ErrAttributeNotFoundForAction)
		require.ErrorIs(t, meta.SetFloat("nope", 1.5), types.ErrAttributeNotFoundForAction)
		require.ErrorIs(t, meta.SetBoolean("nope", true), types.ErrAttributeNotFoundForAction)
	})
}

func TestSettersTypeMismatchReturnsError(t *testing.T) {
	meta := newChainMeta("t1", numberAttr("level", 1))
	// "level" is a number; setting it as a string must error, not corrupt it.
	require.ErrorIs(t, meta.SetString("level", "x"), types.ErrAttributeTypeNotMatch)
}

func TestSetNumberUpdatesChainAttribute(t *testing.T) {
	meta := newChainMeta("t1", numberAttr("level", 10))

	require.NoError(t, meta.SetNumber("level", 42))

	got, err := meta.MustGetNumber("level")
	require.NoError(t, err)
	require.Equal(t, int64(42), got)
	require.Len(t, meta.ChangeList, 1)
	require.Equal(t, "level", meta.ChangeList[0].Key)
	require.Equal(t, "10", meta.ChangeList[0].PreviousValue)
	require.Equal(t, "42", meta.ChangeList[0].NewValue)
}

// --- #1: TransferNumber value integrity ---

func TestTransferNumberSuccess(t *testing.T) {
	source := newChainMeta("src", numberAttr("power", 100))
	target := &types.NftData{TokenId: "dst", OnchainAttributes: []*types.NftAttributeValue{numberAttr("power", 5)}}
	source.SetGetNFTFunction(func(string) (*types.NftData, error) { return target, nil })

	require.NoError(t, source.TransferNumber("power", "dst", 30))

	// source debited
	srcVal, err := source.MustGetNumber("power")
	require.NoError(t, err)
	require.Equal(t, int64(70), srcVal)

	// target credited and registered for persistence
	updated, ok := source.OtherUpdatedTokenDatas["dst"]
	require.True(t, ok, "target must be recorded in OtherUpdatedTokenDatas")
	require.Equal(t, uint64(35), updated.OnchainAttributes[0].GetNumberAttributeValue().Value)
}

// The core value-loss fix: if the target lacks the attribute, the transfer must
// fail WITHOUT having debited the source.
func TestTransferNumberTargetMissingAttributeDoesNotDebitSource(t *testing.T) {
	source := newChainMeta("src", numberAttr("power", 100))
	target := &types.NftData{TokenId: "dst", OnchainAttributes: []*types.NftAttributeValue{numberAttr("other", 5)}}
	source.SetGetNFTFunction(func(string) (*types.NftData, error) { return target, nil })

	err := source.TransferNumber("power", "dst", 30)
	require.ErrorIs(t, err, types.ErrAttributeDoesNotExists)

	// source untouched
	srcVal, gErr := source.MustGetNumber("power")
	require.NoError(t, gErr)
	require.Equal(t, int64(100), srcVal)
	require.Empty(t, source.OtherUpdatedTokenDatas, "nothing should be staged for persistence")
}

// The inflation fix: an origin-sourced attribute cannot be debited (SetNumber
// rejects it), so the target must NOT be credited either.
func TestTransferNumberOriginSourceDoesNotInflateTarget(t *testing.T) {
	// place "power" in OriginAttributes -> From == "origin" -> not settable.
	token := &types.NftData{
		TokenId:          "src",
		OriginAttributes: []*types.NftAttributeValue{numberAttr("power", 100)},
	}
	source := types.NewMetadata(&types.NFTSchema{}, token, types.AttributeOverriding_ORIGIN, nil)
	target := &types.NftData{TokenId: "dst", OnchainAttributes: []*types.NftAttributeValue{numberAttr("power", 5)}}
	source.SetGetNFTFunction(func(string) (*types.NftData, error) { return target, nil })

	err := source.TransferNumber("power", "dst", 30)
	require.Error(t, err, "origin attribute must not be transferable")

	// target not credited
	require.Empty(t, source.OtherUpdatedTokenDatas)
	require.Equal(t, uint64(5), target.OnchainAttributes[0].GetNumberAttributeValue().Value)
}

func TestTransferNumberInsufficientValue(t *testing.T) {
	source := newChainMeta("src", numberAttr("power", 100))
	target := &types.NftData{TokenId: "dst", OnchainAttributes: []*types.NftAttributeValue{numberAttr("power", 5)}}
	source.SetGetNFTFunction(func(string) (*types.NftData, error) { return target, nil })

	require.ErrorIs(t, source.TransferNumber("power", "dst", 200), types.ErrInsufficientValue)

	srcVal, _ := source.MustGetNumber("power")
	require.Equal(t, int64(100), srcVal)
}

func TestTransferFloatSuccess(t *testing.T) {
	source := newChainMeta("src", floatAttr("hp", 10.5))
	target := &types.NftData{TokenId: "dst", OnchainAttributes: []*types.NftAttributeValue{floatAttr("hp", 1.0)}}
	source.SetGetNFTFunction(func(string) (*types.NftData, error) { return target, nil })

	require.NoError(t, source.TransferFloat("hp", "dst", 2.5))

	srcVal, err := source.MustGetFloat("hp")
	require.NoError(t, err)
	require.InDelta(t, 8.0, srcVal, 1e-9)
	require.InDelta(t, 3.5, target.OnchainAttributes[0].GetFloatAttributeValue().Value, 1e-9)
}

// --- #3: GetSubString ---

func TestGetSubString(t *testing.T) {
	meta := newChainMeta("t1", stringAttr("name", "helloworld"))

	require.Equal(t, "hello", meta.GetSubString("name", 0, 5))
	require.Equal(t, "", meta.GetSubString("name", 3, 3))
	// negative end normalizes to length (existing convention: -1 -> len)
	require.Equal(t, "helloworld", meta.GetSubString("name", 0, -1))
}

func TestGetSubStringEndBeyondLengthPanics(t *testing.T) {
	meta := newChainMeta("t1", stringAttr("name", "abc"))
	require.Panics(t, func() { meta.GetSubString("name", 0, 99) })
}

// error is now checked first: a missing key panics with the not-found error
// rather than being used before the error check.
func TestGetSubStringMissingKeyPanics(t *testing.T) {
	meta := newChainMeta("t1")
	require.Panics(t, func() { meta.GetSubString("nope", 0, 1) })
}
