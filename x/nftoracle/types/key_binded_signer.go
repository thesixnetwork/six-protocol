package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BindedSignerKeyPrefix is the prefix to retrieve all BindedSigner
	BindedSignerKeyPrefix = "BindedSigner/value/"
)

// BindedSignerKey returns the store key to retrieve a BindedSigner from the index fields
func BindedSignerKey(
	ownerAddress string,
) []byte {
	var key []byte

	ownerAddressBytes := []byte(ownerAddress)
	key = append(key, ownerAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}
