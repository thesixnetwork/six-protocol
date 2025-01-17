package types

func NewVirtualSchemaRegistryRequest(code string, sharedAttributes []string) *VirtualSchemaRegistryRequest {
	return &VirtualSchemaRegistryRequest{
		NftSchemaCode: code,
	}
}

func (req VirtualSchemaRegistryRequest) ConvertRequestToVirtualRegistry() *VirtualSchemaRegistry {
	return &VirtualSchemaRegistry{
		NftSchemaCode: req.NftSchemaCode,
		Decision:      RegistryStatus_PENDING,
	}
}
