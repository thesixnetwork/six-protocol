package utils

import (
	"os"
	"strconv"
	"strings"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	errormod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type (
	VirtualSchemaRegistryRequestJSON []VirtualSchemaRegistryJSON

	VirtualSchemaRegistryJSON struct {
		VirtualSchemaCode string `json:"code"`
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

	FeeDistribution struct {
		Method  string `json:"method"`
		Portion string `json:"portion"`
	}

	FeeConfig struct {
		FeeAmount        string            `json:"fee_amount"`
		FeeDistributions []FeeDistribution `json:"fee_distributions"`
	}
)

func ParseProposalFile(cdc *codec.LegacyAmino, proposalFile string) (types.VirtualSchemaProposalRequest, error) {
	request := types.VirtualSchemaProposalRequest{}

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

func ParseFeeConfigJSON(cdc *codec.LegacyAmino, newFeeConfigFile string) (*types.FeeConfig, error) {
	newFeeConfig := types.FeeConfig{}

	feeconfigInput := FeeConfig{}

	feeconfigContent, err := os.ReadFile(newFeeConfigFile)
	if err != nil {
		return nil, err
	}

	if err := cdc.UnmarshalJSON(feeconfigContent, &feeconfigInput); err != nil {
		return nil, err
	}

	// convert input to nftmngrtypes.FeeConfig
	newFeeConfig.FeeAmount = feeconfigInput.FeeAmount
	for _, feeDis := range feeconfigInput.FeeDistributions {
		methodLower := strings.ToLower(feeDis.Method)
		portionFloat, err := strconv.ParseFloat(feeDis.Portion, 32)
		if err != nil {
			return nil, err
		}

		switch methodLower {
		case "burn":
			newFeeConfig.FeeDistributions = append(newFeeConfig.FeeDistributions, &types.FeeDistribution{
				Method:  types.FeeDistributionMethod_BURN,
				Portion: float32(portionFloat),
			})
		case "reward_pool":
			newFeeConfig.FeeDistributions = append(newFeeConfig.FeeDistributions, &types.FeeDistribution{
				Method:  types.FeeDistributionMethod_REWARD_POOL,
				Portion: float32(portionFloat),
			})
		case "transfer":
			newFeeConfig.FeeDistributions = append(newFeeConfig.FeeDistributions, &types.FeeDistribution{
				Method:  types.FeeDistributionMethod_TRANSFER,
				Portion: float32(portionFloat),
			})
		default:
			return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "invalid subject. Use burn, reward_pool, or transfer")
		}
	}
	return &newFeeConfig, nil
}
