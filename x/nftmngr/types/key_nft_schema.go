package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NFTSchemaKeyPrefix is the prefix to retrieve all NFTSchema
	NFTSchemaKeyPrefix = "NFTSchema/value/"
)

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
