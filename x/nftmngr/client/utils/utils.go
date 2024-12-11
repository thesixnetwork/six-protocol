package utils

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
)

type (
	VirtualSchemaRegistryRequestJSON []VirtualSchemaRegistryJSON

	VirtualSchemaRegistryJSON struct {
		Code             string   `json:"code" yaml:"code"`
		SharedAttributes []string `json:"sharedAttributes" yaml:"sharedAttributes"`
	}
)

func NewVirtualSchemaRegistryRequestJSON(code string, sharedAttributes []string) VirtualSchemaRegistryRequestJSON {
	return VirtualSchemaRegistryRequestJSON{
		{
			Code:             code,
			SharedAttributes: sharedAttributes,
		},
	}
}

func ParseVirtualSchemaRegistryRequestJSON(cdc *codec.LegacyAmino, proposalFile string) (VirtualSchemaRegistryRequestJSON, error) {
	request := VirtualSchemaRegistryRequestJSON{}

	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return request, err
	}

	if err := cdc.UnmarshalJSON(contents, &request); err != nil {
		return request, err
	}

	return request, nil
}