package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (m *Metadata) TransferNumber(attributeName string, targetTokenId string, transferValue uint64) error {
	// Check if attribute exists in m.MapAllKey
	if _, ok := m.MapAllKey[attributeName]; !ok {
		return sdkerrors.Wrap(ErrAttributeDoesNotExists, attributeName)
	}

	attri := m.MapAllKey[attributeName]

	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue); !ok {
		// Number
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}

	numberValue := attri.AttributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue).NumberAttributeValue
	// check if exists in m.OtherUpdatedTokenDatas
	var targetNftData *NftData
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
		return sdkerrors.Wrap(ErrInsufficientValue, attributeName)
	}
	// decrease transferValue
	m.SetNumber(attributeName, int64(numberValue.Value-transferValue))
	// increase transferValu
	// loop over targetNftData.OnchainAttributes to find attributeName
	for i, targetAttri := range targetNftData.OnchainAttributes {
		if targetAttri.Name == attributeName {
			newAttributeValue := &NftAttributeValue{
				Name: attri.AttributeValue.Name,
				Value: &NftAttributeValue_NumberAttributeValue{
					NumberAttributeValue: &NumberAttributeValue{
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

func (m *Metadata) TransferFloat(attributeName string, targetTokenId string, transferValue float64) error {
	// Check if attribute exists in m.MapAllKey
	if _, ok := m.MapAllKey[attributeName]; !ok {
		return sdkerrors.Wrap(ErrAttributeDoesNotExists, attributeName)
	}

	attri := m.MapAllKey[attributeName]

	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_FloatAttributeValue); !ok {
		// Float
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
	}

	floatValue := attri.AttributeValue.GetValue().(*NftAttributeValue_FloatAttributeValue).FloatAttributeValue
	// check if exists in m.OtherUpdatedTokenDatas
	var targetNftData *NftData
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
	// check if floatValue.Value > transferValue
	if floatValue.Value < transferValue {
		return sdkerrors.Wrap(ErrInsufficientValue, attributeName)
	}
	// decrease transferValue
	m.SetFloat(attributeName, floatValue.Value-transferValue)
	// increase transferValu
	// loop over targetNftData.OnchainAttributes to find attributeName
	for i, targetAttri := range targetNftData.OnchainAttributes {
		if targetAttri.Name == attributeName {
			newAttributeValue := &NftAttributeValue{
				Name: attri.AttributeValue.Name,
				Value: &NftAttributeValue_FloatAttributeValue{
					FloatAttributeValue: &FloatAttributeValue{
						Value: targetAttri.GetFloatAttributeValue().Value + transferValue,
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