package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GroupKeyPrefix is the prefix to retrieve all Group
	GroupKeyPrefix = "Group/value/"
)

// GroupKey returns the store key to retrieve a Group from the index fields
func GroupKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
