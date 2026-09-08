package types

import (
	errormod "cosmossdk.io/errors"
)

func (m *Metadata) TransferNumber(attributeName string, targetTokenId string, transferValue uint64) error {
	// Check if attribute exists in m.MapAllKey
	if _, ok := m.MapAllKey[attributeName]; !ok {
		return errormod.Wrap(ErrAttributeDoesNotExists, attributeName)
	}

	attri := m.MapAllKey[attributeName]

	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue); !ok {
		// Number
		return errormod.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
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
		return errormod.Wrap(ErrInsufficientValue, attributeName)
	}

	// Locate and type-check the target attribute BEFORE debiting the source, so
	// we never decrement the source unless the credit is guaranteed to land.
	targetIndex := -1
	for i, targetAttri := range targetNftData.OnchainAttributes {
		if targetAttri.Name == attributeName {
			targetIndex = i
			break
		}
	}
	if targetIndex == -1 {
		return errormod.Wrap(ErrAttributeDoesNotExists, attributeName)
	}
	targetAttri := targetNftData.OnchainAttributes[targetIndex]
	if _, ok := targetAttri.GetValue().(*NftAttributeValue_NumberAttributeValue); !ok {
		return errormod.Wrap(ErrAttributeTypeNotMatch, targetAttri.Name)
	}

	// decrease transferValue from source (surface the error instead of dropping it)
	if err := m.SetNumber(attributeName, int64(numberValue.Value-transferValue)); err != nil {
		return err
	}

	// increase transferValue on target
	targetNftData.OnchainAttributes[targetIndex] = &NftAttributeValue{
		Name: attri.AttributeValue.Name,
		Value: &NftAttributeValue_NumberAttributeValue{
			NumberAttributeValue: &NumberAttributeValue{
				Value: targetAttri.GetNumberAttributeValue().Value + transferValue,
			},
		},
	}
	// check if exists m.OtherUpdatedTokenDatas map
	if _, ok := m.OtherUpdatedTokenDatas[targetTokenId]; !ok {
		m.OtherUpdatedTokenDatas[targetTokenId] = targetNftData
	}

	return nil
}

func (m *Metadata) TransferFloat(attributeName string, targetTokenId string, transferValue float64) error {
	// Check if attribute exists in m.MapAllKey
	if _, ok := m.MapAllKey[attributeName]; !ok {
		return errormod.Wrap(ErrAttributeDoesNotExists, attributeName)
	}

	attri := m.MapAllKey[attributeName]

	if _, ok := attri.AttributeValue.GetValue().(*NftAttributeValue_FloatAttributeValue); !ok {
		// Float
		return errormod.Wrap(ErrAttributeTypeNotMatch, attri.AttributeValue.Name)
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
		return errormod.Wrap(ErrInsufficientValue, attributeName)
	}

	// Locate and type-check the target attribute BEFORE debiting the source, so
	// we never decrement the source unless the credit is guaranteed to land.
	targetIndex := -1
	for i, targetAttri := range targetNftData.OnchainAttributes {
		if targetAttri.Name == attributeName {
			targetIndex = i
			break
		}
	}
	if targetIndex == -1 {
		return errormod.Wrap(ErrAttributeDoesNotExists, attributeName)
	}
	targetAttri := targetNftData.OnchainAttributes[targetIndex]
	if _, ok := targetAttri.GetValue().(*NftAttributeValue_FloatAttributeValue); !ok {
		return errormod.Wrap(ErrAttributeTypeNotMatch, targetAttri.Name)
	}

	// decrease transferValue from source (surface the error instead of dropping it)
	if err := m.SetFloat(attributeName, floatValue.Value-transferValue); err != nil {
		return err
	}

	// increase transferValue on target
	targetNftData.OnchainAttributes[targetIndex] = &NftAttributeValue{
		Name: attri.AttributeValue.Name,
		Value: &NftAttributeValue_FloatAttributeValue{
			FloatAttributeValue: &FloatAttributeValue{
				Value: targetAttri.GetFloatAttributeValue().Value + transferValue,
			},
		},
	}
	// check if exists m.OtherUpdatedTokenDatas map
	if _, ok := m.OtherUpdatedTokenDatas[targetTokenId]; !ok {
		m.OtherUpdatedTokenDatas[targetTokenId] = targetNftData
	}

	return nil
}

func (c *CrossSchemaMetadata) validateState() error {
	if c == nil {
		return errormod.Wrap(ErrInvalidOperation, "CrossSchemaMetadata is nil")
	}
	if c.mapSchemaKey == nil {
		return errormod.Wrap(ErrInvalidOperation, "mapSchemaKey is not initialized")
	}
	// if c.NftDataFunction == nil {
	//     return errormod.Wrap(ErrInvalidOperation, "NftDataFunction is not set")
	// }
	return nil
}

func (c *CrossSchemaMetadata) validateNumberAttribute(attr *MetadataAttribute, attrName string) (*NumberAttributeValue, error) {
	numberAttr, ok := attr.AttributeValue.GetValue().(*NftAttributeValue_NumberAttributeValue)
	if !ok {
		return nil, errormod.Wrapf(ErrAttributeTypeNotMatch, "attribute %s is not a number", attrName)
	}
	return numberAttr.NumberAttributeValue, nil
}

func (c *CrossSchemaMetadata) ConvertNumberAttribute(srcSchemaName, srcAttributeName, dstSchemaName, dstAttributeName string, convertValue uint64) error {
	// Validate metadata state
	if err := c.validateState(); err != nil {
		return err
	}

	// Get and validate attributes
	srcAttribute, err := c.getAttribute(srcSchemaName, srcAttributeName)
	if err != nil {
		return errormod.Wrapf(err, "source attribute %s", srcAttributeName)
	}

	dstAttribute, err := c.getAttribute(dstSchemaName, dstAttributeName)
	if err != nil {
		return errormod.Wrapf(err, "destination attribute %s", dstAttributeName)
	}

	// Validate number attributes
	srcNumberValue, err := c.validateNumberAttribute(srcAttribute, srcAttributeName)
	if err != nil {
		return err
	}

	dstNumberValue, err := c.validateNumberAttribute(dstAttribute, dstAttributeName)
	if err != nil {
		return err
	}

	// Validate sufficient balance
	if srcNumberValue.Value < convertValue {
		return errormod.Wrapf(ErrInsufficientValue,
			"insufficient balance in %s: has %d, need %d",
			srcAttributeName, srcNumberValue.Value, convertValue)
	}

	// Perform transfer
	if err := c.SetNumber(srcSchemaName, srcAttributeName,
		int64(srcNumberValue.Value-convertValue)); err != nil {
		return errormod.Wrap(err, "failed to update source value")
	}

	if err := c.SetNumber(dstSchemaName, dstAttributeName,
		int64(dstNumberValue.Value+convertValue)); err != nil {
		return errormod.Wrap(err, "failed to update destination value")
	}

	return nil
}
