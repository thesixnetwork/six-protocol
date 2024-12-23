package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetInactiveVirtualSchemaProposal set a specific inactiveVirtualSchemaProposal in the store from its index
func (k Keeper) SetInactiveVirtualSchemaProposal(ctx sdk.Context, inactiveVirtualSchemaProposal types.InactiveVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&inactiveVirtualSchemaProposal)
	store.Set(types.InactiveVirtualSchemaProposalKey(
		inactiveVirtualSchemaProposal.Id,
	), b)
}

// GetInactiveVirtualSchemaProposal returns a inactiveVirtualSchemaProposal from its index
func (k Keeper) GetInactiveVirtualSchemaProposal(
	ctx sdk.Context,
	index string,
) (val types.InactiveVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.InactiveVirtualSchemaProposalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInactiveVirtualSchemaProposal removes a inactiveVirtualSchemaProposal from the store
func (k Keeper) RemoveInactiveVirtualSchemaProposal(
	ctx sdk.Context,
	index string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	store.Delete(types.InactiveVirtualSchemaProposalKey(
		index,
	))
}

// GetAllInactiveVirtualSchemaProposal returns all inactiveVirtualSchemaProposal
func (k Keeper) GetAllInactiveVirtualSchemaProposal(ctx sdk.Context) (list []types.InactiveVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InactiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
