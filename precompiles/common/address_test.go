package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/ethermint/utils"
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

	assert.NotNil(t, from, "Address should not be nil")
	assert.Equal(t, address, from.String(), "Address strings should match")
}

func TestAccAddressFromEthCommon(t *testing.T) {
	commonAddress := "0x3fab184622dc19b6109349b94811493bf2a45362"
	address := common.HexToAddress(commonAddress)
	bech32Address := utils.EthToCosmosAddr(address)

	assert.NotNil(t, bech32Address, "Bech32 address should not be nil")
	assert.Equal(t, "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq", bech32Address.String(), "Bech32 address strings should match")
}
