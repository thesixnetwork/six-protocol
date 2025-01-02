package types

import (
	"strconv"
)

func (a *NftAttributeValue) AttributeValueToString() string {
	var attributeValue string

	if value, ok := a.GetValue().(*NftAttributeValue_BooleanAttributeValue); ok {
		attributeValue = strconv.FormatBool(value.BooleanAttributeValue.Value)
	}

	if value, ok := a.GetValue().(*NftAttributeValue_StringAttributeValue); ok {
		attributeValue = value.StringAttributeValue.Value
	}

	if value, ok := a.GetValue().(*NftAttributeValue_NumberAttributeValue); ok {
		attributeValue = strconv.FormatUint(value.NumberAttributeValue.Value, 10)
	}

	if value, ok := a.GetValue().(*NftAttributeValue_FloatAttributeValue); ok {
		attributeValue = strconv.FormatFloat(value.FloatAttributeValue.Value, 'f', -1, 64)
	}

	return attributeValue
}
