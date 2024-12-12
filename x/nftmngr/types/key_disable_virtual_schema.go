package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DisableVirtualSchemaKeyPrefix is the prefix to retrieve all DisableVirtualSchema
	DisableVirtualSchemaKeyPrefix = "DisableVirtualSchema/value/"
)

// DisableVirtualSchemaKey returns the store key to retrieve a DisableVirtualSchema from the index fields
func DisableVirtualSchemaKey(
	nftSchemaCode string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	return key
}
