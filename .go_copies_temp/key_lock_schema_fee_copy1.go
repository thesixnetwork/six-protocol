package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LockSchemaFeeKeyPrefix is the prefix to retrieve all LockSchemaFee
	LockSchemaFeeKeyPrefix = "LockSchemaFee/value/"
)

// LockSchemaFeeKey returns the store key to retrieve a LockSchemaFee from the index fields
func LockSchemaFeeKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
