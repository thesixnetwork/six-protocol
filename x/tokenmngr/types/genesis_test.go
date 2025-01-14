package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				PortId: types.PortID,
				TokenList: []types.Token{
					{
						Name: "0",
					},
					{
						Name: "1",
					},
				},
				MintpermList: []types.Mintperm{
					{
						Token:   "0",
						Address: "0",
					},
					{
						Token:   "1",
						Address: "1",
					},
				},
				Options: &types.Options{
					DefaultMintee: "39",
				},
				TokenBurnList: []types.TokenBurn{
					{
						Amount: sdk.NewCoin("test", sdk.NewInt(int64(1))),
					},
					{
						Amount: sdk.NewCoin("test", sdk.NewInt(int64(1))),
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated token",
			genState: &types.GenesisState{
				TokenList: []types.Token{
					{
						Name: "0",
					},
					{
						Name: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated mintperm",
			genState: &types.GenesisState{
				MintpermList: []types.Mintperm{
					{
						Token:   "0",
						Address: "0",
					},
					{
						Token:   "0",
						Address: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated tokenBurn",
			genState: &types.GenesisState{
				TokenBurnList: []types.TokenBurn{
					{
						Amount: sdk.NewCoin("test", sdk.NewInt(int64(1))),
					},
					{
						Amount: sdk.NewCoin("test", sdk.NewInt(int64(1))),
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
