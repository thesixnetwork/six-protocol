package types

import (
	"strconv"
)

func (s *SchemaAttribute) CurrentValueToString() string {
	var attributeValue string
	if value, ok := s.GetCurrentValue().GetValue().(*SchemaAttributeValue_BooleanAttributeValue); ok {
		attributeValue = strconv.FormatBool(value.BooleanAttributeValue.Value)
	}

	if value, ok := s.GetCurrentValue().GetValue().(*SchemaAttributeValue_StringAttributeValue); ok {
		attributeValue = value.StringAttributeValue.Value
	}

	if value, ok := s.GetCurrentValue().GetValue().(*SchemaAttributeValue_NumberAttributeValue); ok {
		attributeValue = strconv.FormatUint(value.NumberAttributeValue.Value, 10)
	}

	if value, ok := s.GetCurrentValue().GetValue().(*SchemaAttributeValue_FloatAttributeValue); ok {
		attributeValue = strconv.FormatFloat(value.FloatAttributeValue.Value, 'f', -1, 64)
	}
	return attributeValue
}
