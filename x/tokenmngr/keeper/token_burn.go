package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// SetTokenBurn set a specific tokenBurn in the store from its index
func (k Keeper) SetTokenBurn(ctx sdk.Context, tokenBurn types.TokenBurn) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenBurnKeyPrefix))
	b := k.cdc.MustMarshal(&tokenBurn)
	store.Set(types.TokenBurnKey(
		tokenBurn.Amount.Denom,
	), b)
}

// GetTokenBurn returns a tokenBurn from its index
func (k Keeper) GetTokenBurn(
	ctx sdk.Context,
	token string,
) (val types.TokenBurn, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenBurnKeyPrefix))

	b := store.Get(types.TokenBurnKey(
		token,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTokenBurn removes a tokenBurn from the store
func (k Keeper) RemoveTokenBurn(
	ctx sdk.Context,
	token string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenBurnKeyPrefix))
	store.Delete(types.TokenBurnKey(
		token,
	))
}

// GetAllTokenBurn returns all tokenBurn
func (k Keeper) GetAllTokenBurn(ctx sdk.Context) (list []types.TokenBurn) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenBurnKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokenBurn
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
