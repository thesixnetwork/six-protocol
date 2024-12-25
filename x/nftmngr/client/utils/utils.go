package utils

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
)

type (
	VirtualSchemaRegistryRequestJSON []VirtualSchemaRegistryJSON

	VirtualSchemaRegistryJSON struct {
		Code             string   `json:"code"`
		SharedAttributes []string `json:"sharedAttributes"`
	}

	ActionParameter struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	TokenIds struct {
		Schema  string `json:"schema"`
		TokenId string `json:"tokenId"`
	}

	VirtualSchemaActionJSON struct {
		Code     string            `json:"code"`
		TokenIds []TokenIds        `json:"tokenIds"`
		Action   string            `json:"action"`
		Params   []ActionParameter `json:"params"`
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

func NewVirtualSchemaActionJSON(code string, tokenIds []TokenIds, actionName string, params []ActionParameter) VirtualSchemaActionJSON {
	return VirtualSchemaActionJSON{
		Code:     code,
		TokenIds: tokenIds,
		Action:   actionName,
		Params:   params,
	}
}

func ParseActionJSON(cdc *codec.LegacyAmino, proposalFile string) (VirtualSchemaActionJSON, error) {
	request := VirtualSchemaActionJSON{}
	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return request, err
	}
	if err := cdc.UnmarshalJSON(contents, &request); err != nil {
		return request, err
	}
	return request, nil
}
