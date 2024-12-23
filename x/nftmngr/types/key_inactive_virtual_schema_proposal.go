package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// InactiveVirtualSchemaProposalKeyPrefix is the prefix to retrieve all InactiveVirtualSchemaProposal
	InactiveVirtualSchemaProposalKeyPrefix = "InactiveVirtualSchemaProposal/value/"
)

// InactiveVirtualSchemaProposalKey returns the store key to retrieve a InactiveVirtualSchemaProposal from the index fields
func InactiveVirtualSchemaProposalKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
