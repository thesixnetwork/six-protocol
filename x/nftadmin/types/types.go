package types

// GetPermissionByKey
func (p *Permissions) GetPermissionAddressByKey(key string) *AddressList {
	// loop over p.Permissions
	for _, v := range p.Permissions {
		if v.Name == key {
			return v.Addresses
		}
	}
	return nil
}
