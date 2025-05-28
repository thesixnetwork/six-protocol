package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NftschemaKeyPrefix is the prefix to retrieve all Nftschema
	NftschemaKeyPrefix = "Nftschema/value/"
)

// NftschemaKey returns the store key to retrieve a Nftschema from the index fields
func NftschemaKey(
	code string,
) []byte {
	var key []byte

	codeBytes := []byte(code)
	key = append(key, codeBytes...)
	key = append(key, []byte("/")...)

	return key
}
