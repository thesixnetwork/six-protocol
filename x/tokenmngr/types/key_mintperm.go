package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MintpermKeyPrefix is the prefix to retrieve all Mintperm
	MintpermKeyPrefix = "Mintperm/value/"
)

// MintpermKey returns the store key to retrieve a Mintperm from the index fields
func MintpermKey(
	token string,
	address string,
) []byte {
	var key []byte

	tokenBytes := []byte(token)
	key = append(key, tokenBytes...)
	key = append(key, []byte("/")...)

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
