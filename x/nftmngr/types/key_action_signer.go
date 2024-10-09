package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActionSignerKeyPrefix is the prefix to retrieve all ActionSigner
	ActionSignerKeyPrefix = "ActionSigner/value/"
)

// ActionSignerKey returns the store key to retrieve a ActionSigner from the index fields
func ActionSignerKey(
	actorAddress string,
) []byte {
	var key []byte

	actorAddressBytes := []byte(actorAddress)
	key = append(key, actorAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}
