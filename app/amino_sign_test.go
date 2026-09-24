package app_test

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdksigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"

	// The CLI links the protov2 (pulsar) staking descriptors from the
	// cosmossdk.io/api fork; once registered they shadow the gogo descriptors
	// in the hybrid resolver the amino-json handler uses. Import them here so
	// the test exercises the same path Ledger signing does.
	_ "cosmossdk.io/api/cosmos/staking/v1beta1"
)

// TestAminoJSONSignEditValidatorForkFields guards Ledger signing of the
// staking fork's MsgEditValidator fields. SIGN_MODE_LEGACY_AMINO_JSON (the
// only mode the Cosmos Ledger app supports) walks the message with the
// protov2 descriptors from the cosmossdk.io/api fork — if that module tag
// lags the proto (as api v0.7.6-six-1 did), signing fails with
// `{TagNum: 8, WireType:"bytes"}: unknown protobuf field` even though
// simulation and SIGN_MODE_DIRECT work.
func TestAminoJSONSignEditValidatorForkFields(t *testing.T) {
	registry := codectypes.NewInterfaceRegistry()
	stakingtypes.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)
	txConfig := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)

	maxLicense := math.NewInt(7000)
	increment := math.NewInt(10_000_000_000)
	minDelegation := math.NewInt(1)

	msg := &stakingtypes.MsgEditValidator{
		Description:         stakingtypes.NewDescription("[do-not-modify]", "[do-not-modify]", "[do-not-modify]", "[do-not-modify]", "[do-not-modify]"),
		ValidatorAddress:    sdk.ValAddress([]byte("test_validator_addr_")).String(),
		MaxLicense:          &maxLicense,
		DelegationIncrement: &increment,
		Mode:                "license",
		MinDelegation:       &minDelegation,
	}

	txb := txConfig.NewTxBuilder()
	require.NoError(t, txb.SetMsgs(msg))

	signerData := authsigning.SignerData{
		ChainID:       "fivenet",
		AccountNumber: 1,
		Sequence:      1,
		Address:       sdk.AccAddress([]byte("test_validator_addr_")).String(),
	}

	_, err := authsigning.GetSignBytesAdapter(
		context.Background(),
		txConfig.SignModeHandler(),
		sdksigning.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
		signerData,
		txb.GetTx(),
	)
	require.NoError(t, err,
		"amino-json sign bytes for MsgEditValidator with fork fields (delegation_increment=8, mode=9, min_delegation=10) — cosmossdk.io/api fork descriptors are stale if this fails")
}
