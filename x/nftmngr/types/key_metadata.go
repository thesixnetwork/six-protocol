package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MetadataCreatorKeyPrefix is the prefix to retrieve all MetadataCreator
	MetadataCreatorKeyPrefix = "MetadataCreator/value/"

	// NftCollectionKeyPrefix is the prefix to retrieve all NftCollection
	NftCollectionKeyPrefix = "NftCollection/value/"

	// NFTSchemaByContractKeyPrefix is the prefix to retrieve all NFTSchemaByContract
	NFTSchemaByContractKeyPrefix = "NFTSchemaByContract/value/"

	NftCollectionDataCountKey = "NftCollectionData-count-"
)

// MetadataCreatorKey returns the store key to retrieve a MetadataCreator from the index fields
func MetadataCreatorKey(
	nftSchemaCode string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	return key
}

// NftCollectionKey returns the store key to retrieve a NftCollection from the index fields
func NftCollectionKey(
	nftSchemaCode string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	return key
}

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
