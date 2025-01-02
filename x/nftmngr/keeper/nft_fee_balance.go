package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetNFTFeeBalance set nFTFeeBalance in the store
func (k Keeper) SetNFTFeeBalance(ctx sdk.Context, nFTFeeBalance types.NFTFeeBalance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeBalanceKey))
	b := k.cdc.MustMarshal(&nFTFeeBalance)
	store.Set([]byte{0}, b)
}

// GetNFTFeeBalance returns nFTFeeBalance
func (k Keeper) GetNFTFeeBalance(ctx sdk.Context) (val types.NFTFeeBalance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeBalanceKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFTFeeBalance removes nFTFeeBalance from the store
func (k Keeper) RemoveNFTFeeBalance(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeBalanceKey))
	store.Delete([]byte{0})
}
