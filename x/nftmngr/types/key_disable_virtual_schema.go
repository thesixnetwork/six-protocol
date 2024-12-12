package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DisableVirtualSchemaKeyPrefix is the prefix to retrieve all DisableVirtualSchema
	DisableVirtualSchemaKeyPrefix = "DisableVirtualSchema/value/"
)

// DisableVirtualSchemaKey returns the store key to retrieve a DisableVirtualSchema from the index fields
func DisableVirtualSchemaKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
