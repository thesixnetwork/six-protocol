package common

import (
	"fmt"
	"testing"

	"github.com/evmos/ethermint/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
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

func TestAccAddressFromEthCommon(t *testing.T) {
	commonAddress := "0x3fab184622dc19b6109349b94811493bf2a45362"
	address := common.HexToAddress(commonAddress)
  bech32Address := utils.EthToCosmosAddr(address)

	fmt.Printf("######## ADDRESS BYTE: %v\n", bech32Address)
	fmt.Printf("######## ADDRESS String: %v\n", bech32Address.String())
}
