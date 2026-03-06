package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AdminKeyPrefix is the prefix to retrieve all Admin
	AdminKeyPrefix = "Admin/value/"
)

// AdminKey returns the store key to retrieve a Admin from the index fields
func AdminKey(
	group string,
	admin string,
) []byte {
	var key []byte

	groupBytes := []byte(group)
	key = append(key, groupBytes...)
	key = append(key, []byte("/")...)

	adminBytes := []byte(admin)
	key = append(key, adminBytes...)
	key = append(key, []byte("/")...)

	return key
}
