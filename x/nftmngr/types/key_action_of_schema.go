package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActionOfSchemaKeyPrefix is the prefix to retrieve all ActionOfSchema
	ActionOfSchemaKeyPrefix = "ActionOfSchema/value/"
)

// ActionOfSchemaKey returns the store key to retrieve a ActionOfSchema from the index fields
func ActionOfSchemaKey(
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
