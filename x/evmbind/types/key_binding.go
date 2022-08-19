package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BindingKeyPrefix is the prefix to retrieve all Binding
	BindingKeyPrefix = "Binding/value/"
)

// BindingKey returns the store key to retrieve a Binding from the index fields
func BindingKey(
	ethAddress string,
) []byte {
	var key []byte

	ethAddressBytes := []byte(ethAddress)
	key = append(key, ethAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}
