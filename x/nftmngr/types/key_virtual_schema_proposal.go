package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VirtualSchemaProposalKeyPrefix is the prefix to retrieve all VirtualSchemaProposal
	VirtualSchemaProposalKeyPrefix = "VirtualSchemaProposal/value/"
)

// VirtualSchemaProposalKey returns the store key to retrieve a VirtualSchemaProposal from the index fields
func VirtualSchemaProposalKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
