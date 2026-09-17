package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

func numberAttr(name string, v uint64) *types.NftAttributeValue {
	return &types.NftAttributeValue{
		Name: name,
		Value: &types.NftAttributeValue_NumberAttributeValue{
			NumberAttributeValue: &types.NumberAttributeValue{Value: v},
		},
	}
}

func newChainMeta(onchain ...*types.NftAttributeValue) *types.Metadata {
	token := &types.NftData{TokenId: "t1", OnchainAttributes: onchain}
	return types.NewMetadata(&types.NFTSchema{}, token, types.AttributeOverriding_CHAIN, nil)
}

// levelUpAction sets level to 10 when it is below 10.
func levelUpAction() *types.Action {
	return &types.Action{
		Name: "level_up",
		When: `meta.GetNumber("level") < 10`,
		Then: []string{`meta.SetNumber("level", 10);`},
	}
}

// ProcessAction compiles and runs a rule that mutates the metadata.
func TestProcessActionExecutesRule(t *testing.T) {
	meta := newChainMeta(numberAttr("level", 1))

	require.NoError(t, keeper.ProcessAction(meta, levelUpAction(), nil))

	got, err := meta.MustGetNumber("level")
	require.NoError(t, err)
	require.Equal(t, int64(10), got)
	require.Len(t, meta.ChangeList, 1)
}

// Running the same action twice must produce identical results — this exercises
// the compiled-rule cache path (second run reuses the parsed knowledge base).
func TestProcessActionCachedRuleRepeatable(t *testing.T) {
	action := levelUpAction()

	for i := range 3 {
		meta := newChainMeta(numberAttr("level", 1))
		require.NoError(t, keeper.ProcessAction(meta, action, nil))
		got, err := meta.MustGetNumber("level")
		require.NoError(t, err)
		require.Equalf(t, int64(10), got, "run %d", i)
	}
}

// GetNumber on a missing attribute panics inside the rule; ProcessAction's
// deferred recover must turn that into a returned error rather than crashing.
func TestProcessActionRecoversPanicAsError(t *testing.T) {
	meta := newChainMeta(numberAttr("level", 1))
	action := &types.Action{
		Name: "bad_action",
		When: `meta.GetNumber("missing") < 10`,
		Then: []string{`meta.SetNumber("level", 5);`},
	}

	var err error
	require.NotPanics(t, func() { err = keeper.ProcessAction(meta, action, nil) })
	require.Error(t, err)
}
