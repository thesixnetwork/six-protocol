package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VirtualSchemaProposalKeyPrefix is the prefix to retrieve all VirtualSchemaProposal
	VirtualSchemaProposalKeyPrefix = "VirtualSchemaProposal/value/"
)

// VirtualSchemaProposalKey returns the store key to retrieve a VirtualSchemaProposal from the index fields
func VirtualSchemaProposalKey(
	id string,
) []byte {
	var key []byte

	idBytes := []byte(id)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}
