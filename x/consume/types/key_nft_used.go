package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NftUsedKeyPrefix is the prefix to retrieve all NftUsed
	NftUsedKeyPrefix = "NftUsed/value/"
)

// NftUsedKey returns the store key to retrieve a NftUsed from the index fields
func NftUsedKey(
	token string,
	cretor string,
) []byte {
	var key []byte

	tokenBytes := []byte(token)
	key = append(key, tokenBytes...)
	key = append(key, []byte("/")...)

	cretorBytes := []byte(cretor)
	key = append(key, cretorBytes...)
	key = append(key, []byte("/")...)

	return key
}
