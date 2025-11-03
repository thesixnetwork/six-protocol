package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NftDataKeyPrefix is the prefix to retrieve all NftData
	NftDataKeyPrefix = "NftData/value/"
)

// NftDataKey returns the store key to retrieve a NftData from the index fields
func NftDataKey(
	nftSchemaCode string,
	tokenId string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	tokenIdBytes := []byte(tokenId)
	key = append(key, tokenIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
