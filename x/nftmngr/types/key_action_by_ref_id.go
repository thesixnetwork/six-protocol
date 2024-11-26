package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActionByRefIdKeyPrefix is the prefix to retrieve all ActionByRefId
	ActionByRefIdKeyPrefix = "ActionByRefId/value/"
)

// ActionByRefIdKey returns the store key to retrieve a ActionByRefId from the index fields
func ActionByRefIdKey(
	refId string,
) []byte {
	var key []byte

	refIdBytes := []byte(refId)
	key = append(key, refIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
