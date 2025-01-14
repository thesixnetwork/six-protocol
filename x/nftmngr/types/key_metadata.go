package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MetadataCreatorKeyPrefix is the prefix to retrieve all MetadataCreator
	MetadataCreatorKeyPrefix = "MetadataCreator/value/"

	// NftCollectionKeyPrefix is the prefix to retrieve all NftCollection
	NftCollectionKeyPrefix = "NftCollection/value/"

	// NftDataKeyPrefix is the prefix to retrieve all NftData
	NftDataKeyPrefix = "NftData/value/"

	// NFTSchemaByContractKeyPrefix is the prefix to retrieve all NFTSchemaByContract
	NFTSchemaByContractKeyPrefix = "NFTSchemaByContract/value/"
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
