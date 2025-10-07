package types

func (p *Authorization) GetPermissionAddressByKey(key string) []string {
	// loop over p.Permissions
	for _, v := range p.Permissions {
		if v.Name == key {
			return v.Addresses
		}
	}
	return nil
}
