package types

import (
	fmt "fmt"

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

func (c *CrossSchemaMetadata) ConvertNumberAttribute(srcSchemaName, srcAttributeName, dstSchemaName, dstAttributeName string, convertValue uint64) error {
	
	if c == nil {
		fmt.Println("CrossSchemaMetadata is nil")
        return sdkerrors.Wrap(ErrInvalidOperation, "CrossSchemaMetadata is nil")
    }

    if c.mapSchemaKey == nil {
		fmt.Println("mapSchemaKey is not initialized")
        return sdkerrors.Wrap(ErrInvalidOperation, "mapSchemaKey is not initialized")
    }

    if c.NftDataFunction == nil {
		fmt.Println("NftDataFunction is not set")
        return sdkerrors.Wrap(ErrInvalidOperation, "NftDataFunction is not set")
    }

	if err := c.validateSchemaName(srcSchemaName); err != nil {
		return err
	}

	if err := c.validateSchemaName(dstSchemaName); err != nil {
		return err
	}

	srcAttribute, ok := c.mapSchemaKey[srcSchemaName][srcAttributeName]
	if !ok {
		return sdkerrors.Wrap(ErrAttributeDoesNotExists, srcAttributeName)
	}

	srcNumberAttributeValue, ok := srcAttribute.AttributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue)
	if !ok {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, srcAttribute.AttributeValue.Name)
	}

	dstAttribute, ok := c.mapSchemaKey[dstSchemaName][dstAttributeName]
	if !ok {
		return sdkerrors.Wrap(ErrAttributeDoesNotExists, dstAttributeName)
	}

	dstNumberAttributeValue, ok := dstAttribute.AttributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue)
	if !ok {
		return sdkerrors.Wrap(ErrAttributeTypeNotMatch, dstAttribute.AttributeValue.Name)
	}

	srcNumberValue := srcNumberAttributeValue.NumberAttributeValue

	dstNumberValue := dstNumberAttributeValue.NumberAttributeValue

	var err error
	// Check if transfer value is sufficient
	if srcNumberValue.Value < convertValue {
		return sdkerrors.Wrap(ErrInsufficientValue, srcAttributeName)
	}

	// Deduct from source
	err = c.SetNumber(srcSchemaName, srcAttributeName, int64(srcNumberValue.Value - convertValue))
	if err != nil {
		return err
	}

	err = c.SetNumber(dstSchemaName, dstAttributeName, int64(dstNumberValue.Value + convertValue))
	if err != nil {
		return err
	}


	return nil
}
