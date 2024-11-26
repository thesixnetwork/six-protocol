package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// OrganizationKeyPrefix is the prefix to retrieve all Organization
	OrganizationKeyPrefix = "Organization/value/"
)

// OrganizationKey returns the store key to retrieve a Organization from the index fields
func OrganizationKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
