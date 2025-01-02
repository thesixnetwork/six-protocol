package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)
func init() {
	// Set the prefix for addresses
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("6x", "6xpub")
	config.Seal()
}

func TestAccAddressFromBech32(t *testing.T) {
	address := "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq"
	from, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("######## ADDRESS BYTE: %v\n", from)
	fmt.Printf("######## ADDRESS String: %v\n", from.String())
}
