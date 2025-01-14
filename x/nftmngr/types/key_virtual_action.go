package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VirtualKeyPrefix is the prefix to retrieve all Virtual
	VirtualActionKeyPrefix = "VirtualAction/value/"

	// VirtualSchemaKeyPrefix is the prefix to retrieve all VirtualSchema
	VirtualSchemaKeyPrefix = "VirtualSchema/value/"
)

// VirtualKey returns the store key to retrieve a Virtual from the index fields
func VirtualActionKey(
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

// VirtualSchemaKey returns the store key to retrieve a VirtualSchema from the index fields
func VirtualSchemaKey(
	nftSchemaCode string,
) []byte {
	var key []byte

	indexBytes := []byte(nftSchemaCode)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
