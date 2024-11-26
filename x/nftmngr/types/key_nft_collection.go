package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NftCollectionKeyPrefix is the prefix to retrieve all NftCollection
	NftCollectionKeyPrefix = "NftCollection/value/"
)

// NftCollectionKey returns the store key to retrieve a NftCollection from the index fields
func NftCollectionKey(
	nftSchemaCode string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	return key
}
