package types

func NewVirtualSchemaRegistryRequest(code string, sharedAttributes []string) *VirtualSchemaRegistryRequest {
	return &VirtualSchemaRegistryRequest{
		NftSchemaCode:    code,
		SharedAttributes: sharedAttributes,
	}
}

func (req VirtualSchemaRegistryRequest) ConvertRequestToVirtualRegistry() *VirtualSchemaRegistry {
	return &VirtualSchemaRegistry{
		NftSchemaCode:    req.NftSchemaCode,
		SharedAttributes: req.SharedAttributes,
		Status:           RegistryStatus_PENDING,
	}
}
