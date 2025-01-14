package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActionExecutorKeyPrefix is the prefix to retrieve all ActionExecutor
	ActionExecutorKeyPrefix = "ActionExecutor/value/"

	// ExecutorOfSchemaKeyPrefix is the prefix to retrieve all ExecutorOfSchema
	ExecutorOfSchemaKeyPrefix = "ExecutorOfSchema/value/"
)

// ActionExecutorKey returns the store key to retrieve a ActionExecutor from the index fields
func ActionExecutorKey(
	nftSchemaCode string,
	executorAddress string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	executorAddressBytes := []byte(executorAddress)
	key = append(key, executorAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// ExecutorOfSchemaKey returns the store key to retrieve a ExecutorOfSchema from the index fields
func ExecutorOfSchemaKey(
	nftSchemaCode string,
) []byte {
	var key []byte

	nftSchemaCodeBytes := []byte(nftSchemaCode)
	key = append(key, nftSchemaCodeBytes...)
	key = append(key, []byte("/")...)

	return key
}
