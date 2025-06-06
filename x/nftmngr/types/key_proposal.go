package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VirtualSchemaProposalKeyPrefix is the prefix to retrieve all VirtualSchemaProposal
	VirtualSchemaProposalKeyPrefix = "VirtualSchemaProposal/value/"

	// AttributeOfSchemaKeyPrefix is the prefix to retrieve all AttributeOfSchema
	AttributeOfSchemaKeyPrefix = "AttributeOfSchema/value/"

	// InactiveVirtualSchemaProposalKeyPrefix is the prefix to retrieve all InactiveVirtualSchemaProposal
	InactiveVirtualSchemaProposalKeyPrefix = "InactiveVirtualSchemaProposal/value/"

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
