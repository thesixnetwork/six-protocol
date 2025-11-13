package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
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
				AdminList: []types.Admin{
					{
						Group: "0",
						Admin: "0",
					},
					{
						Group: "1",
						Admin: "1",
					},
				},
				GroupList: []types.Group{
					{
						Name: "0",
					},
					{
						Name: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated admin",
			genState: &types.GenesisState{
				AdminList: []types.Admin{
					{
						Group: "0",
						Admin: "0",
					},
					{
						Group: "0",
						Admin: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated group",
			genState: &types.GenesisState{
				GroupList: []types.Group{
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
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
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
