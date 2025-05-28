package types

const (
	// ModuleName defines the module name
	ModuleName = "protocoladmin"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_protocoladmin"
)

var ParamsKey = []byte("p_protocoladmin")

func KeyPrefix(p string) []byte {
	return []byte(p)
}
