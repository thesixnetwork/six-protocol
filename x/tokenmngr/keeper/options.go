package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// SetOptions set options in the store
func (k Keeper) SetOptions(ctx sdk.Context, options types.Options) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OptionsKey))
	b := k.cdc.MustMarshal(&options)
	store.Set([]byte{0}, b)
}

// GetOptions returns options
func (k Keeper) GetOptions(ctx sdk.Context) (val types.Options, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OptionsKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOptions removes options from the store
func (k Keeper) RemoveOptions(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OptionsKey))
	store.Delete([]byte{0})
}
