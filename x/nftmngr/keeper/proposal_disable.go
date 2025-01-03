package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// RemoveInactiveDisableVirtualSchemaProposal removes a inactiveDisableVirtualSchemaProposal from the store
func (k Keeper) RemoveInactiveDisableVirtualSchemaProposal(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveDisableVirtualSchemaProposalKeyPrefix))
	store.Delete(types.InactiveDisableVirtualSchemaProposalKey(
		index,
	))
}

// GetAllInactiveDisableVirtualSchemaProposal returns all inactiveDisableVirtualSchemaProposal
func (k Keeper) GetAllInactiveDisableVirtualSchemaProposal(ctx sdk.Context) (list []types.InactiveDisableVirtualSchemaProposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InactiveDisableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InactiveDisableVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) DisableVirtualSchemaIterateActiveProposal(ctx sdk.Context, endTime time.Time, cb func(proposal types.DisableVirtualSchemaProposal) (stop bool)) {
	iterator := k.DisableVirtualSchemaActiveProposalQueryIterator(ctx, endTime)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ActiveVirtualSchemaProposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		proposal, found := k.GetDisableVirtualSchemaProposal(ctx, val.Id)
		if !found {
			panic(fmt.Sprintf("proposal %d does not exist", &val.Id))
		}

		if cb(proposal) {
			break
		}
	}
}

func (k Keeper) DisableVirtualSchemaActiveProposalQueryIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ActiveDisableVirtualSchemaProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	return iterator
}
