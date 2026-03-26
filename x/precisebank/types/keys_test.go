package types_test

import (
	"testing"

	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

func TestConstants(t *testing.T) {
	if types.IntegerCoinDenom != "usix" {
		t.Fatalf("expected IntegerCoinDenom to be 'usix', got '%s'", types.IntegerCoinDenom)
	}
	if types.ExtendedCoinDenom != "asix" {
		t.Fatalf("expected ExtendedCoinDenom to be 'asix', got '%s'", types.ExtendedCoinDenom)
	}
}

func TestModuleName(t *testing.T) {
	if types.ModuleName != "precisebank" {
		t.Fatalf("expected ModuleName to be 'precisebank', got '%s'", types.ModuleName)
	}
	if types.StoreKey != "precisebank" {
		t.Fatalf("expected StoreKey to be 'precisebank', got '%s'", types.StoreKey)
	}
}

func TestKeyPrefixes(t *testing.T) {
	if len(types.FractionalBalancePrefix) != 1 || types.FractionalBalancePrefix[0] != 0x01 {
		t.Fatal("FractionalBalancePrefix should be [0x01]")
	}
	if len(types.RemainderBalanceKey) != 1 || types.RemainderBalanceKey[0] != 0x02 {
		t.Fatal("RemainderBalanceKey should be [0x02]")
	}
}
