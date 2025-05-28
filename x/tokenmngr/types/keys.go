package types

const (
	// ModuleName defines the module name
	ModuleName = "tokenmngr"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tokenmngr"
)

var ParamsKey = []byte("p_tokenmngr")

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	OptionsKey   = "Options/value/"
	BurnKey      = "Burn/value/"
	BurnCountKey = "Burn/count/"
)
