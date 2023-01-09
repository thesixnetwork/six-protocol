package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

const TOKEN_ADMIN = "token.admin"

// SetToken set a specific token in the store from its index
func (k Keeper) SetToken(ctx sdk.Context, token types.Token) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKeyPrefix))
	b := k.cdc.MustMarshal(&token)
	store.Set(types.TokenKey(
		token.Name,
	), b)
}

// GetToken returns a token from its index
func (k Keeper) GetToken(
	ctx sdk.Context,
	name string,

) (val types.Token, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKeyPrefix))

	b := store.Get(types.TokenKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetToken returns a token from its index
func (k Keeper) GetTokenV202(
	ctx sdk.Context,
	name string,

) (val types.TokenV202, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKeyPrefix))

	b := store.Get(types.TokenKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveToken removes a token from the store
func (k Keeper) RemoveToken(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKeyPrefix))
	store.Delete(types.TokenKey(
		name,
	))
}

// GetAllToken returns all token
func (k Keeper) GetAllToken(ctx sdk.Context) (list []types.Token) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllToken returns all token
func (k Keeper) GetAllTokenV202(ctx sdk.Context) (list []types.TokenV202) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokenV202
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
