package types

const (
	// ModuleName defines the module name
	ModuleName = "nftmngr"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_nftmngr"
)

const (
	KeyMintPermissionOnlySystem = "system"
	KeyMintPermissionAll        = "all"
)

var ParamsKey = []byte("p_nftmngr")

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// ! Fixed from prefix appent([]byte{0x1} []byte(nftSchemaCode)) because it will return 0x1nftSchemaCode
func CollectionkeyPrefix(nftSchemaCode string) []byte {
	var key []byte
	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)
	return key
}
