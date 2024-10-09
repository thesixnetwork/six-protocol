package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NFTSchemaByContractKeyPrefix is the prefix to retrieve all NFTSchemaByContract
	NFTSchemaByContractKeyPrefix = "NFTSchemaByContract/value/"
)

// NFTSchemaByContractKey returns the store key to retrieve a NFTSchemaByContract from the index fields
func NFTSchemaByContractKey(
	originContractAddress string,
) []byte {
	var key []byte

	originContractAddressBytes := []byte(originContractAddress)
	key = append(key, originContractAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}
