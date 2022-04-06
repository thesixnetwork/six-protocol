package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// SetMintdes set mintdes in the store
func (k Keeper) SetMintdes(ctx sdk.Context, mintdes string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintdesKey))
	store.Set([]byte{0}, []byte(mintdes))
}

// GetMintdes returns mintdes
func (k Keeper) GetMintdes(ctx sdk.Context) (val string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintdesKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val
	}

	return string(val)
}

// RemoveMintdes removes mintdes from the store
func (k Keeper) RemoveMintdes(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MintdesKey))
	store.Delete([]byte{0})
}
