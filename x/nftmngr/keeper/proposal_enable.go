package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetEnableVirtualSchemaProposal set a specific enableVirtualSchemaProposal in the store from its index
func (k Keeper) SetEnableVirtualSchemaProposal(ctx sdk.Context, enableVirtualSchemaProposal types.EnableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnableVirtualSchemaProposalKeyPrefix))
	b := k.cdc.MustMarshal(&enableVirtualSchemaProposal)
	store.Set(types.EnableVirtualSchemaProposalKey(
		enableVirtualSchemaProposal.Id,
	), b)
}

// GetEnableVirtualSchemaProposal returns a enableVirtualSchemaProposal from its index
func (k Keeper) GetEnableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) (val types.EnableVirtualSchemaProposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnableVirtualSchemaProposalKeyPrefix))

	b := store.Get(types.EnableVirtualSchemaProposalKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEnableVirtualSchemaProposal removes a enableVirtualSchemaProposal from the store
func (k Keeper) RemoveEnableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnableVirtualSchemaProposalKeyPrefix))
	store.Delete(types.EnableVirtualSchemaProposalKey(
		index,
	))
}

// GetAllEnableVirtualSchemaProposal returns all enableVirtualSchemaProposal
func (k Keeper) GetAllEnableVirtualSchemaProposal(ctx sdk.Context) (list []types.EnableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EnableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EnableVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// func (k Keeper) IterateActiveProposalEnableVirtualSchema(ctx sdk.Context, endTime time.Time, cb func(proposal types.EnableVirtualSchemaProposal) (stop bool)) {
// 	iterator := k.ActiveProposalDisableVirtualSchemaQueryIterator(ctx, endTime)
// 	defer iterator.Close()

// 	for ; iterator.Valid(); iterator.Next() {
// 		var val types.ActiveVirtualSchemaProposal
// 		k.cdc.MustUnmarshal(iterator.Value(), &val)

// 		proposal, found := k.GetDisableVirtualSchemaProposal(ctx, val.Id)
// 		if !found {
// 			panic(fmt.Sprintf("proposal %d does not exist", &val.Id))
// 		}

// 		if cb(proposal) {
// 			break
// 		}
// 	}
// }
