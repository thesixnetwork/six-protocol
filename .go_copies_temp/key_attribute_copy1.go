package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SchemaAttributeKeyPrefix is the prefix to retrieve all SchemaAttribute
	SchemaAttributeKeyPrefix = "SchemaAttribute/value/"
)

// SchemaAttributeKey returns the store key to retrieve a SchemaAttribute from the index fields
func SchemaAttributeKey(
	nftSchemaCode string,
	name string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
