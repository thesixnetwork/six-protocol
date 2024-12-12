package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetVirtualSchemaProposal set a specific virtualSchemaProposal in the store from its index
func (k Keeper) SetVirtualSchemaProposal(ctx sdk.Context, virtualSchemaProposal types.VirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&virtualSchemaProposal)
	store.Set(types.VirtualSchemaProposalKey(
		virtualSchemaProposal.Id,
	), b)
}

// GetVirtualSchemaProposal returns a virtualSchemaProposal from its index
func (k Keeper) GetVirtualSchemaProposal(
	ctx sdk.Context,
	id string,

) (val types.VirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))

	b := store.Get(types.VirtualSchemaProposalKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVirtualSchemaProposal removes a virtualSchemaProposal from the store
func (k Keeper) RemoveVirtualSchemaProposal(
	ctx sdk.Context,
	id string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))
	store.Delete(types.VirtualSchemaProposalKey(
		id,
	))
}

// GetAllVirtualSchemaProposal returns all virtualSchemaProposal
func (k Keeper) GetAllVirtualSchemaProposal(ctx sdk.Context) (list []types.VirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
