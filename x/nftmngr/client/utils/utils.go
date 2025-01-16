package utils

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

type (
	VirtualSchemaRegistryRequestJSON []VirtualSchemaRegistryJSON

	VirtualSchemaRegistryJSON struct {
		VirtualSchemaCode string `json:"code"`
	}

	VirtualSchemaRequest struct {
		VirtualSchemaCode     string         `json:"virtualSchemaCode"`
		Actions               []types.Action `json:"actions"`
		VirtualSchemaRegistry []string       `json:"virtualSchemaRegistry"`
		Enable                bool           `json:"enable"`
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

func ParseProposalFile(cdc *codec.LegacyAmino, proposalFile string) (VirtualSchemaRequest, error) {
	request := VirtualSchemaRequest{}

	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return request, err
	}

	if err := cdc.UnmarshalJSON(contents, &request); err != nil {
		return request, err
	}

	return request, nil
}

func NewVirtualSchemaRegistryRequestJSON(code string, sharedAttributes []string) VirtualSchemaRegistryRequestJSON {
	return VirtualSchemaRegistryRequestJSON{
		{
			VirtualSchemaCode: code,
		},
	}
}

// DEPRECATED
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

func ParseActionJSON(cdc *codec.LegacyAmino, actionFile string) (VirtualSchemaActionJSON, error) {
	request := VirtualSchemaActionJSON{}
	contents, err := os.ReadFile(actionFile)
	if err != nil {
		return request, err
	}
	if err := cdc.UnmarshalJSON(contents, &request); err != nil {
		return request, err
	}
	return request, nil
}

func ParseNewVirtualActionJSON(cdc *codec.LegacyAmino, newActionFile string) ([]*types.Action, error) {
	newAction := []*types.Action{}

	contents, err := os.ReadFile(newActionFile)
	if err != nil {
		return newAction, err
	}

	if err := cdc.UnmarshalJSON(contents, &newAction); err != nil {
		return newAction, err
	}

	return newAction, err
}
