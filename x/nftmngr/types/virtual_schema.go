package types


func NewVirtualSchemaRegistryRequest(code string, sharedAttributes []string) *VirtualSchemaRegistryRequest{
	return &VirtualSchemaRegistryRequest{
		NftSchemaCode: code,
		SharedAttributes: sharedAttributes,
	}
}