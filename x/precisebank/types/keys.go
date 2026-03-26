package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "precisebank"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for precisebank
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_precisebank"
)

var (
	// Key prefixes for store
	FractionalBalancePrefix = []byte{0x01}
)

var (
	// Keys for store that are not prefixed
	RemainderBalanceKey = []byte{0x02}
)

// FractionalBalanceKey returns a KV store key for an address's fractional balance.
func FractionalBalanceKey(address sdk.AccAddress) []byte {
	return address.Bytes()
}
