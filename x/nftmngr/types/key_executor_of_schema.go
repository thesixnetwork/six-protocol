package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ExecutorOfSchemaKeyPrefix is the prefix to retrieve all ExecutorOfSchema
	ExecutorOfSchemaKeyPrefix = "ExecutorOfSchema/value/"
)

// ExecutorOfSchemaKey returns the store key to retrieve a ExecutorOfSchema from the index fields
func ExecutorOfSchemaKey(
	nftSchemaCode string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	return key
}
