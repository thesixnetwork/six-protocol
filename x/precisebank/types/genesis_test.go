package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

func TestGenesisState_Validate(t *testing.T) {
	addr1 := sdk.AccAddress([]byte("testaddress1________")).String()
	addr2 := sdk.AccAddress([]byte("testaddress2________")).String()

	tests := []struct {
		name    string
		gs      *types.GenesisState
		wantErr bool
	}{
		{
			name:    "default is valid",
			gs:      types.DefaultGenesisState(),
			wantErr: false,
		},
		{
			name: "valid with balances summing to conversion factor",
			gs: types.NewGenesisState(
				types.FractionalBalances{
					types.NewFractionalBalance(addr1, sdkmath.NewInt(500_000_000_000)),
					types.NewFractionalBalance(addr2, sdkmath.NewInt(500_000_000_000)),
				},
				sdkmath.ZeroInt(),
			),
			wantErr: false,
		},
		{
			name: "valid with remainder",
			gs: types.NewGenesisState(
				types.FractionalBalances{
					types.NewFractionalBalance(addr1, sdkmath.NewInt(400_000_000_000)),
				},
				sdkmath.NewInt(600_000_000_000),
			),
			wantErr: false,
		},
		{
			name: "nil remainder",
			gs: &types.GenesisState{
				Balances:  types.FractionalBalances{},
				Remainder: sdkmath.Int{},
			},
			wantErr: true,
		},
		{
			name: "negative remainder",
			gs: types.NewGenesisState(
				types.FractionalBalances{},
				sdkmath.NewInt(-1),
			),
			wantErr: true,
		},
		{
			name: "remainder exceeds conversion factor",
			gs: types.NewGenesisState(
				types.FractionalBalances{},
				types.ConversionFactor(),
			),
			wantErr: true,
		},
		{
			name: "sum + remainder not multiple of conversion factor",
			gs: types.NewGenesisState(
				types.FractionalBalances{
					types.NewFractionalBalance(addr1, sdkmath.NewInt(100)),
				},
				sdkmath.NewInt(200),
			),
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.gs.Validate()
			if tc.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestGenesisState_TotalAmountWithRemainder(t *testing.T) {
	addr1 := sdk.AccAddress([]byte("testaddress1________")).String()
	addr2 := sdk.AccAddress([]byte("testaddress2________")).String()

	gs := types.NewGenesisState(
		types.FractionalBalances{
			types.NewFractionalBalance(addr1, sdkmath.NewInt(300_000_000_000)),
			types.NewFractionalBalance(addr2, sdkmath.NewInt(400_000_000_000)),
		},
		sdkmath.NewInt(300_000_000_000),
	)

	total := gs.TotalAmountWithRemainder()
	expected := sdkmath.NewInt(1_000_000_000_000)
	if !total.Equal(expected) {
		t.Fatalf("expected total %s, got %s", expected, total)
	}
}

func TestDefaultGenesis(t *testing.T) {
	gs := types.DefaultGenesis()
	if gs == nil {
		t.Fatal("DefaultGenesis returned nil")
	}
	if len(gs.Balances) != 0 {
		t.Fatalf("expected empty balances, got %d", len(gs.Balances))
	}
	if !gs.Remainder.IsZero() {
		t.Fatalf("expected zero remainder, got %s", gs.Remainder)
	}
	if err := gs.Validate(); err != nil {
		t.Fatalf("default genesis should be valid: %v", err)
	}
}
