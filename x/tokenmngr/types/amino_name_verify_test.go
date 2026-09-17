package types_test

import (
	"context"
	"strings"
	"testing"

	msgv1 "cosmossdk.io/api/cosmos/msg/v1"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txsigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	gogoproto "github.com/cosmos/gogoproto/proto"
	protov2 "google.golang.org/protobuf/proto"

	amino "cosmossdk.io/api/amino"

	_ "github.com/thesixnetwork/six-protocol/v4/x/nftadmin/types"
	_ "github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
	_ "github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"
	_ "github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
)

// Verifies that SIGN_MODE_LEGACY_AMINO_JSON sign bytes now carry the classic
// amino names (e.g. "tokenmngr/Mint") instead of type-URL fallbacks.
func TestAminoJSONSignBytesUseLegacyNames(t *testing.T) {
	reg := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(reg)
	cdc := codec.NewProtoCodec(reg)
	txCfg := authtx.NewTxConfig(cdc, authtx.DefaultSignModes)

	msg := &types.MsgMint{
		Creator: "6x1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp0xhjg",
		Amount:  sdk.NewInt64Coin("usix", 1),
	}

	b := txCfg.NewTxBuilder()
	if err := b.SetMsgs(msg); err != nil {
		t.Fatal(err)
	}

	signBz, err := authsigning.GetSignBytesAdapter(
		context.Background(),
		txCfg.SignModeHandler(),
		txsigning.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
		authsigning.SignerData{
			Address:       "6x1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqp0xhjg",
			ChainID:       "sixnet",
			AccountNumber: 1,
			Sequence:      1,
		},
		b.GetTx(),
	)
	if err != nil {
		t.Fatal(err)
	}

	got := string(signBz)
	if !strings.Contains(got, `"type":"tokenmngr/Mint"`) {
		t.Fatalf("sign bytes do not contain legacy amino name, got: %s", got)
	}
	if strings.Contains(got, "sixprotocol.tokenmngr.MsgMint") {
		t.Fatalf("sign bytes still contain type URL fallback: %s", got)
	}
}

// Every tx message with a cosmos.msg.v1.signer option must also declare an
// amino.name, otherwise SIGN_MODE_LEGACY_AMINO_JSON falls back to the type URL
// and standard amino clients fail signature verification.
func TestAllSignableMsgsHaveAminoName(t *testing.T) {
	modules := []string{"tokenmngr", "nftmngr", "nftoracle", "nftadmin", "protocoladmin"}
	for _, mod := range modules {
		fd, err := gogoproto.HybridResolver.FindFileByPath("sixprotocol/" + mod + "/tx.proto")
		if err != nil {
			t.Fatalf("%s: %v", mod, err)
		}
		msgs := fd.Messages()
		for i := 0; i < msgs.Len(); i++ {
			md := msgs.Get(i)
			opts := md.Options()
			if !protov2.HasExtension(opts, msgv1.E_Signer) {
				continue
			}
			// Create/Update/DeleteVirtualAction are not wired into the Msg
			// service nor the amino codec; skip them like codec.go does.
			name := string(md.Name())
			if name == "MsgCreateVirtualAction" || name == "MsgUpdateVirtualAction" || name == "MsgDeleteVirtualAction" {
				continue
			}
			if !protov2.HasExtension(opts, amino.E_Name) {
				t.Errorf("%s: %s has a signer but no amino.name option", mod, md.Name())
			}
		}
	}
}
