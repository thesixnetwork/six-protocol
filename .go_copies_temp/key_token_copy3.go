package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TokenKeyPrefix is the prefix to retrieve all Token
	TokenKeyPrefix = "Token/value/"
)

// TokenKey returns the store key to retrieve a Token from the index fields
func TokenKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
