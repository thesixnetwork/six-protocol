package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActiveVirtualSchemaProposalKeyPrefix is the prefix to retrieve all ActiveVirtualSchemaProposal
	ActiveVirtualSchemaProposalKeyPrefix = "ActiveVirtualSchemaProposal/value/"

	// VirtualSchemaProposalKeyPrefix is the prefix to retrieve all VirtualSchemaProposal
	VirtualSchemaProposalKeyPrefix = "VirtualSchemaProposal/value/"

	// AttributeOfSchemaKeyPrefix is the prefix to retrieve all AttributeOfSchema
	AttributeOfSchemaKeyPrefix = "AttributeOfSchema/value/"

	// DisableVirtualSchemaKeyPrefix is the prefix to retrieve all DisableVirtualSchema
	DisableVirtualSchemaProposalKeyPrefix = "DisableVirtualSchemaProposal/value/"

	// InactiveVirtualSchemaProposalKeyPrefix is the prefix to retrieve all InactiveVirtualSchemaProposal
	InactiveVirtualSchemaProposalKeyPrefix = "InactiveVirtualSchemaProposal/value/"

	// ActiveDislabeVirtualSchemaProposalKeyPrefix is the prefix to retrieve all ActiveDislabeVirtualSchemaProposal
	ActiveDislabeVirtualSchemaProposalKeyPrefix = "ActiveDislabeVirtualSchemaProposal/value/"

	// InactiveDisableVirtualSchemaProposalKeyPrefix is the prefix to retrieve all InactiveDisableVirtualSchemaProposal
	InactiveDisableVirtualSchemaProposalKeyPrefix = "InactiveDisableVirtualSchemaProposal/value/"
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

// AttributeOfSchemaKey returns the store key to retrieve a AttributeOfSchema from the index fields
func AttributeOfSchemaKey(
	nftSchemaCode string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	return key
}

// DisableVirtualSchemaKey returns the store key to retrieve a DisableVirtualSchema from the index fields
func DisableVirtualSchemaKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
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

// ActiveDislabeVirtualSchemaProposalKey returns the store key to retrieve a ActiveDislabeVirtualSchemaProposal from the index fields
func ActiveDislabeVirtualSchemaProposalKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

// InactiveDisableVirtualSchemaProposalKey returns the store key to retrieve a InactiveDisableVirtualSchemaProposal from the index fields
func InactiveDisableVirtualSchemaProposalKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
