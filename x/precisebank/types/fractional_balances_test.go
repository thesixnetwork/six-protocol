package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

func TestFractionalBalances_Validate(t *testing.T) {
	addr1 := sdk.AccAddress([]byte("testaddress1________")).String()
	addr2 := sdk.AccAddress([]byte("testaddress2________")).String()

	tests := []struct {
		name    string
		fbs     types.FractionalBalances
		wantErr bool
	}{
		{
			name:    "empty is valid",
			fbs:     types.FractionalBalances{},
			wantErr: false,
		},
		{
			name: "valid single",
			fbs: types.FractionalBalances{
				types.NewFractionalBalance(addr1, sdkmath.NewInt(100)),
			},
			wantErr: false,
		},
		{
			name: "valid multiple",
			fbs: types.FractionalBalances{
				types.NewFractionalBalance(addr1, sdkmath.NewInt(100)),
				types.NewFractionalBalance(addr2, sdkmath.NewInt(200)),
			},
			wantErr: false,
		},
		{
			name: "invalid balance in slice",
			fbs: types.FractionalBalances{
				types.NewFractionalBalance("invalid", sdkmath.NewInt(100)),
			},
			wantErr: true,
		},
		{
			name: "duplicate addresses",
			fbs: types.FractionalBalances{
				types.NewFractionalBalance(addr1, sdkmath.NewInt(100)),
				types.NewFractionalBalance(addr1, sdkmath.NewInt(200)),
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.fbs.Validate()
			if tc.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestFractionalBalances_SumAmount(t *testing.T) {
	addr1 := sdk.AccAddress([]byte("testaddress1________")).String()
	addr2 := sdk.AccAddress([]byte("testaddress2________")).String()

	tests := []struct {
		name     string
		fbs      types.FractionalBalances
		expected sdkmath.Int
	}{
		{
			name:     "empty",
			fbs:      types.FractionalBalances{},
			expected: sdkmath.ZeroInt(),
		},
		{
			name: "single balance",
			fbs: types.FractionalBalances{
				types.NewFractionalBalance(addr1, sdkmath.NewInt(500)),
			},
			expected: sdkmath.NewInt(500),
		},
		{
			name: "multiple balances",
			fbs: types.FractionalBalances{
				types.NewFractionalBalance(addr1, sdkmath.NewInt(300)),
				types.NewFractionalBalance(addr2, sdkmath.NewInt(700)),
			},
			expected: sdkmath.NewInt(1000),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sum := tc.fbs.SumAmount()
			if !sum.Equal(tc.expected) {
				t.Fatalf("expected sum %s, got %s", tc.expected, sum)
			}
		})
	}
}
