package common_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/precompiles/common"
	"github.com/thesixnetwork/six-protocol/utils"
	"github.com/stretchr/testify/require"
)

var largeAmt, _ = sdk.NewIntFromString("1000000000000000000000000000000000000000")

func TestNewCoinsResponse(t *testing.T) {
	testCases := []struct {
		amount sdk.Int
	}{
		{amount: sdk.NewInt(1)},
		{amount: largeAmt},
	}

	for _, tc := range testCases {
		coin := sdk.NewCoin(utils.BaseDenom, tc.amount)
		coins := sdk.NewCoins(coin)
		res := common.NewCoinsResponse(coins)
		require.Equal(t, 1, len(res))
		require.Equal(t, tc.amount.BigInt(), res[0].Amount)
	}
}

func TestNewDecCoinsResponse(t *testing.T) {
	testCases := []struct {
		amount sdk.Int
	}{
		{amount: sdk.NewInt(1)},
		{amount: largeAmt},
	}

	for _, tc := range testCases {
		coin := sdk.NewDecCoin(utils.BaseDenom, tc.amount)
		coins := sdk.NewDecCoins(coin)
		res := common.NewDecCoinsResponse(coins)
		require.Equal(t, 1, len(res))
		require.Equal(t, tc.amount.BigInt(), res[0].Amount)
	}
}
