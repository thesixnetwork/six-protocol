package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AttributeOfSchemaKeyPrefix is the prefix to retrieve all AttributeOfSchema
	AttributeOfSchemaKeyPrefix = "AttributeOfSchema/value/"
)

// AttributeOfSchemaKey returns the store key to retrieve a AttributeOfSchema from the index fields
func AttributeOfSchemaKey(
	nftSchemaCode string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	return key
}
