package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MetadataCreatorKeyPrefix is the prefix to retrieve all MetadataCreator
	MetadataCreatorKeyPrefix = "MetadataCreator/value/"
)

// MetadataCreatorKey returns the store key to retrieve a MetadataCreator from the index fields
func MetadataCreatorKey(
	nftSchemaCode string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	return key
}
