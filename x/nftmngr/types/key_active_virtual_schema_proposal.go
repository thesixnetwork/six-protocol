package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActiveVirtualSchemaProposalKeyPrefix is the prefix to retrieve all ActiveVirtualSchemaProposal
	ActiveVirtualSchemaProposalKeyPrefix = "ActiveVirtualSchemaProposal/value/"
)

// ActiveVirtualSchemaProposalKey returns the store key to retrieve a ActiveVirtualSchemaProposal from the index fields
func ActiveVirtualSchemaProposalKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
