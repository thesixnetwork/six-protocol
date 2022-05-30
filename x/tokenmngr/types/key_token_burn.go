package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TokenBurnKeyPrefix is the prefix to retrieve all TokenBurn
	TokenBurnKeyPrefix = "TokenBurn/value/"
)

// TokenBurnKey returns the store key to retrieve a TokenBurn from the index fields
func TokenBurnKey(
	token string,
) []byte {
	var key []byte

	tokenBytes := []byte(token)
	key = append(key, tokenBytes...)
	key = append(key, []byte("/")...)

	return key
}
