package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetNFTFeeConfig set nFTFeeConfig in the store
func (k Keeper) SetNFTFeeConfig(ctx sdk.Context, nFTFeeConfig types.NFTFeeConfig) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeConfigKey))
	b := k.cdc.MustMarshal(&nFTFeeConfig)
	store.Set([]byte{0}, b)
}

// GetNFTFeeConfig returns nFTFeeConfig
func (k Keeper) GetNFTFeeConfig(ctx sdk.Context) (val types.NFTFeeConfig, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeConfigKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFTFeeConfig removes nFTFeeConfig from the store
func (k Keeper) RemoveNFTFeeConfig(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeConfigKey))
	store.Delete([]byte{0})
}
