package keeper

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (m *Metadata) TransferNumber(attributeName string, targetTokenId string, transferValue uint64) error {
	// Check if attribute exists in m.MapAllKey
	if _, ok := m.MapAllKey[attributeName]; !ok {
		return sdkerrors.Wrap(types.ErrAttributeDoesNotExists, attributeName)
	}

	attri := m.MapAllKey[attributeName]

	if _, ok := attri.AttributeValue.GetValue().(*types.NftAttributeValue_NumberAttributeValue); !ok {
		// Number
		return sdkerrors.Wrap(types.ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}

	numberValue := attri.AttributeValue.GetValue().(*types.NftAttributeValue_NumberAttributeValue).NumberAttributeValue
	// check if exists in m.OtherUpdatedTokenDatas
	var targetNftData *types.NftData
	if _, ok := m.OtherUpdatedTokenDatas[targetTokenId]; ok {
		targetNftData = m.OtherUpdatedTokenDatas[targetTokenId]
	} else {
		var err error
		// Get target NFTData
		targetNftData, err = m.NftDataFunction(targetTokenId)
		if err != nil {
			return err
		}
	}
	// check if numberValue.Value > transferValue
	if numberValue.Value < transferValue {
		return sdkerrors.Wrap(types.ErrInsufficientValue, attributeName)
	}
	// decrease transferValue
	m.SetNumber(attributeName, int64(numberValue.Value-transferValue))
	// increase transferValu
	// loop over targetNftData.OnchainAttributes to find attributeName
	for i, targetAttri := range targetNftData.OnchainAttributes {
		if targetAttri.Name == attributeName {
			newAttributeValue := &types.NftAttributeValue{
				Name: attri.AttributeValue.Name,
				Value: &types.NftAttributeValue_NumberAttributeValue{
					NumberAttributeValue: &types.NumberAttributeValue{
						Value: uint64(targetAttri.GetNumberAttributeValue().Value + transferValue),
					},
				},
			}
			targetNftData.OnchainAttributes[i] = newAttributeValue
			// check if exists m.OtherUpdatedTokenDatas map
			if _, ok := m.OtherUpdatedTokenDatas[targetTokenId]; !ok {
				m.OtherUpdatedTokenDatas[targetTokenId] = targetNftData
			}
			break
		}
	}

	return nil
}
