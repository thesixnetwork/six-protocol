package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

func TestConversionFactor(t *testing.T) {
	cf := types.ConversionFactor()
	expected := sdkmath.NewInt(1_000_000_000_000)
	if !cf.Equal(expected) {
		t.Fatalf("expected conversion factor %s, got %s", expected, cf)
	}
}

func TestConversionFactor_Immutable(t *testing.T) {
	cf1 := types.ConversionFactor()
	cf1.AddRaw(1)
	cf2 := types.ConversionFactor()
	expected := sdkmath.NewInt(1_000_000_000_000)
	if !cf2.Equal(expected) {
		t.Fatalf("ConversionFactor should be immutable, but got modified: %s", cf2)
	}
}

func TestNewFractionalBalance(t *testing.T) {
	addr := sdk.AccAddress([]byte("testaddress1________")).String()
	amount := sdkmath.NewInt(100)
	fb := types.NewFractionalBalance(addr, amount)

	if fb.Address != addr {
		t.Fatalf("expected address %s, got %s", addr, fb.Address)
	}
	if !fb.Amount.Equal(amount) {
		t.Fatalf("expected amount %s, got %s", amount, fb.Amount)
	}
}

func TestFractionalBalance_Validate(t *testing.T) {
	validAddr := sdk.AccAddress([]byte("testaddress1________")).String()

	tests := []struct {
		name    string
		fb      types.FractionalBalance
		wantErr bool
	}{
		{
			name:    "valid",
			fb:      types.NewFractionalBalance(validAddr, sdkmath.NewInt(100)),
			wantErr: false,
		},
		{
			name:    "invalid address",
			fb:      types.NewFractionalBalance("invalid", sdkmath.NewInt(100)),
			wantErr: true,
		},
		{
			name:    "zero amount",
			fb:      types.NewFractionalBalance(validAddr, sdkmath.ZeroInt()),
			wantErr: true,
		},
		{
			name:    "negative amount",
			fb:      types.NewFractionalBalance(validAddr, sdkmath.NewInt(-1)),
			wantErr: true,
		},
		{
			name:    "amount at conversion factor",
			fb:      types.NewFractionalBalance(validAddr, types.ConversionFactor()),
			wantErr: true,
		},
		{
			name:    "amount just below conversion factor",
			fb:      types.NewFractionalBalance(validAddr, types.ConversionFactor().SubRaw(1)),
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.fb.Validate()
			if tc.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestValidateFractionalAmount(t *testing.T) {
	tests := []struct {
		name    string
		amt     sdkmath.Int
		wantErr bool
	}{
		{
			name:    "valid",
			amt:     sdkmath.NewInt(500_000_000_000),
			wantErr: false,
		},
		{
			name:    "nil",
			amt:     sdkmath.Int{},
			wantErr: true,
		},
		{
			name:    "zero",
			amt:     sdkmath.ZeroInt(),
			wantErr: true,
		},
		{
			name:    "negative",
			amt:     sdkmath.NewInt(-1),
			wantErr: true,
		},
		{
			name:    "at max",
			amt:     sdkmath.NewInt(999_999_999_999),
			wantErr: false,
		},
		{
			name:    "exceeds max",
			amt:     sdkmath.NewInt(1_000_000_000_000),
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateFractionalAmount(tc.amt)
			if tc.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
