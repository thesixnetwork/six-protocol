package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TokenBurnKeyPrefix is the prefix to retrieve all TokenBurn
	TokenBurnKeyPrefix = "TokenBurn/value/"
)

// TokenBurnKey returns the store key to retrieve a TokenBurn from the index fields
func TokenBurnKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
