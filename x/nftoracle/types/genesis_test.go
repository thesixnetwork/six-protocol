package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
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

				MintRequestList: []types.MintRequest{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				MintRequestCount: 2,
				ActionRequestList: []types.ActionOracleRequest{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ActionRequestCount: 2,
				CollectionOwnerRequestList: []types.CollectionOwnerRequest{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				CollectionOwnerRequestCount: 2,
				OracleConfig:                &types.OracleConfig{},
				ActionSignerList: []types.ActionSigner{
					{
						ActorAddress: "0",
					},
					{
						ActorAddress: "1",
					},
				},
				BindedSignerList: []types.BindedSigner{
					{
						OwnerAddress: "0",
					},
					{
						OwnerAddress: "1",
					},
				},
				ActionSignerConfigList: []types.ActionSignerConfig{
					{
						Chain: "0",
					},
					{
						Chain: "1",
					},
				},
				SyncActionSignerList: []types.SyncActionSigner{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				SyncActionSignerCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated mintRequest",
			genState: &types.GenesisState{
				MintRequestList: []types.MintRequest{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid mintRequest count",
			genState: &types.GenesisState{
				MintRequestList: []types.MintRequest{
					{
						Id: 1,
					},
				},
				MintRequestCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated actionRequest",
			genState: &types.GenesisState{
				ActionRequestList: []types.ActionOracleRequest{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid actionRequest count",
			genState: &types.GenesisState{
				ActionRequestList: []types.ActionOracleRequest{
					{
						Id: 1,
					},
				},
				ActionRequestCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated collectionOwnerRequest",
			genState: &types.GenesisState{
				CollectionOwnerRequestList: []types.CollectionOwnerRequest{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid collectionOwnerRequest count",
			genState: &types.GenesisState{
				CollectionOwnerRequestList: []types.CollectionOwnerRequest{
					{
						Id: 1,
					},
				},
				CollectionOwnerRequestCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated actionSigner",
			genState: &types.GenesisState{
				ActionSignerList: []types.ActionSigner{
					{
						ActorAddress: "0",
					},
					{
						ActorAddress: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated bindedSigner",
			genState: &types.GenesisState{
				BindedSignerList: []types.BindedSigner{
					{
						OwnerAddress: "0",
					},
					{
						OwnerAddress: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated actionSignerConfig",
			genState: &types.GenesisState{
				ActionSignerConfigList: []types.ActionSignerConfig{
					{
						Chain: "0",
					},
					{
						Chain: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated syncActionSigner",
			genState: &types.GenesisState{
				SyncActionSignerList: []types.SyncActionSigner{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid syncActionSigner count",
			genState: &types.GenesisState{
				SyncActionSignerList: []types.SyncActionSigner{
					{
						Id: 1,
					},
				},
				SyncActionSignerCount: 0,
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
