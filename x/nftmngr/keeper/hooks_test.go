package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func TestKeeper_VirtualSchemaHook(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cdc                codec.BinaryCodec
		storeKey           sdk.StoreKey
		memKey             sdk.StoreKey
		ps                 params.Subspace
		nftadminKeeper     types.NftadminKeeper
    accountKeeper      types.AccountKeeper
		bankKeeper         types.BankKeeper
		stakingKeeper      types.StakingKeeper
		distributionKeeper types.DistributionKeeper
		govKeeper          types.GovKeeper
		// Named input parameters for target function.
		ctx                   sdk.Context
		virtualSchemaProposal types.VirtualSchemaProposal
		want                  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := keeper.NewKeeper(tt.cdc, tt.storeKey, tt.memKey, tt.ps, tt.nftadminKeeper, tt.accountKeeper, tt.bankKeeper, tt.stakingKeeper, tt.distributionKeeper, tt.govKeeper)
			got := k.VirtualSchemaHook(tt.ctx, tt.virtualSchemaProposal)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("VirtualSchemaHook() = %v, want %v", got, tt.want)
			}
		})
	}
}
