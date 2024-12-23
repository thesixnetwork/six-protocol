package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetActiveVirtualSchemaProposal set a specific activeVirtualSchemaProposal in the store from its index
func (k Keeper) SetActiveVirtualSchemaProposal(ctx sdk.Context, activeVirtualSchemaProposal types.ActiveVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&activeVirtualSchemaProposal)
	store.Set(types.ActiveVirtualSchemaProposalKey(
		activeVirtualSchemaProposal.Id,
	), b)
}

// GetActiveVirtualSchemaProposal returns a activeVirtualSchemaProposal from its index
func (k Keeper) GetActiveVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) (val types.ActiveVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.ActiveVirtualSchemaProposalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActiveVirtualSchemaProposal removes a activeVirtualSchemaProposal from the store
func (k Keeper) RemoveActiveVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))
	store.Delete(types.ActiveVirtualSchemaProposalKey(
		index,
	))
}

// GetAllActiveVirtualSchemaProposal returns all activeVirtualSchemaProposal
func (k Keeper) GetAllActiveVirtualSchemaProposal(ctx sdk.Context) (list []types.ActiveVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
