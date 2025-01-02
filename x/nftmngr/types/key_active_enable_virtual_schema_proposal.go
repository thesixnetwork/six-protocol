package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActiveEnableVirtualSchemaProposalKeyPrefix is the prefix to retrieve all ActiveEnableVirtualSchemaProposal
	ActiveEnableVirtualSchemaProposalKeyPrefix = "ActiveEnableVirtualSchemaProposal/value/"
)

// ActiveEnableVirtualSchemaProposalKey returns the store key to retrieve a ActiveEnableVirtualSchemaProposal from the index fields
func ActiveEnableVirtualSchemaProposalKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
