package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

const crossSchemaCode = "schemaA"

// newCrossMeta builds a CrossSchemaMetadata with one schema whose attributes are
// all "chain" (settable).
func newCrossMeta(onchain ...*types.NftAttributeValue) *types.CrossSchemaMetadata {
	schema := &types.NFTSchema{Code: crossSchemaCode}
	token := &types.NftData{TokenId: "t1", OnchainAttributes: onchain}
	return types.NewCrossSchemaMetadata(
		[]*types.NFTSchema{schema},
		[]*types.NftData{token},
		types.CrossSchemaAttributeOverriding{crossSchemaCode: types.AttributeOverriding_CHAIN},
		types.CrossSchemaGlobalAttributes{},
	)
}

// --- new SetBoolean parity ---

func TestCrossSetBooleanSuccess(t *testing.T) {
	c := newCrossMeta(boolAttr("flag", false))

	require.NoError(t, c.SetBoolean(crossSchemaCode, "flag", true))

	got, err := c.MustGetBool(crossSchemaCode, "flag")
	require.NoError(t, err)
	require.True(t, got)

	cl := c.GetChangeList(crossSchemaCode)
	require.Len(t, cl, 1)
	require.Equal(t, "flag", cl[0].Key)
	require.Equal(t, "false", cl[0].PreviousValue)
	require.Equal(t, "true", cl[0].NewValue)
}

func TestCrossSetBooleanTypeMismatch(t *testing.T) {
	c := newCrossMeta(numberAttr("flag", 1))
	require.ErrorIs(t, c.SetBoolean(crossSchemaCode, "flag", true), types.ErrAttributeTypeNotMatch)
}

func TestCrossSetBooleanMissingKey(t *testing.T) {
	c := newCrossMeta()
	require.ErrorIs(t, c.SetBoolean(crossSchemaCode, "flag", true), types.ErrAttributeNotFoundForAction)
}

func TestCrossSetBooleanInvalidSchema(t *testing.T) {
	c := newCrossMeta(boolAttr("flag", false))
	require.ErrorIs(t, c.SetBoolean("unknown", "flag", true), types.ErrSchemaNotFound)
}

// --- #4: GetTokenURI / GetImage guard invalid schema codes ---

func TestCrossGetTokenURIInvalidSchemaPanics(t *testing.T) {
	c := newCrossMeta()
	require.Panics(t, func() { c.GetTokenURI("unknown") })
}

func TestCrossGetImageInvalidSchemaPanics(t *testing.T) {
	c := newCrossMeta()
	require.Panics(t, func() { c.GetImage("unknown") })
}

func TestCrossSetImageInvalidSchemaPanics(t *testing.T) {
	c := newCrossMeta()
	require.Panics(t, func() { c.SetImage("unknown", "ipfs://x") })
}

func TestCrossGetImageValid(t *testing.T) {
	c := newCrossMeta()
	require.NotPanics(t, func() { c.SetImage(crossSchemaCode, "ipfs://img") })
	require.Equal(t, "ipfs://img", c.GetImage(crossSchemaCode))
}

// --- #3: cross GetSubString ---

func TestCrossGetSubString(t *testing.T) {
	c := newCrossMeta(stringAttr("name", "helloworld"))
	require.Equal(t, "hello", c.GetSubString(crossSchemaCode, "name", 0, 5))
	require.Equal(t, "helloworld", c.GetSubString(crossSchemaCode, "name", 0, -1))
}

func TestCrossGetSubStringMissingKeyPanics(t *testing.T) {
	c := newCrossMeta()
	require.Panics(t, func() { c.GetSubString(crossSchemaCode, "nope", 0, 1) })
}
