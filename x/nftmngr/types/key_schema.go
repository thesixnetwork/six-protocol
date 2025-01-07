package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActionOfSchemaKeyPrefix is the prefix to retrieve all ActionOfSchema
	ActionOfSchemaKeyPrefix = "ActionOfSchema/value/"

	// NFTSchemaKeyPrefix is the prefix to retrieve all NFTSchema
	NFTSchemaKeyPrefix = "NFTSchema/value/"
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

// NFTSchemaKey returns the store key to retrieve a NFTSchema from the index fields
func NFTSchemaKey(
	code string,
) []byte {
	var key []byte

	codeBytes := []byte(code)
	key = append(key, codeBytes...)
	key = append(key, []byte("/")...)

	return key
}

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