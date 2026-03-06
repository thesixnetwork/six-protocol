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

func (p *Authorization) SetPermissionAddressByKey(key string, listAddress []string) {
	// Handle empty address list by removing the permission entirely
	if len(listAddress) == 0 {
		p.RemovePermissionByKey(key)
		return
	}

	// loop over p.Permissions and update the matching one
	for _, v := range p.Permissions {
		if v.Name == key {
			v.Addresses = listAddress
			return
		}
	}
}

// RemovePermissionByKey removes a permission entry entirely when no addresses remain
func (p *Authorization) RemovePermissionByKey(key string) {
	for i, v := range p.Permissions {
		if v.Name == key {
			// Remove the permission at index i
			p.Permissions = append(p.Permissions[:i], p.Permissions[i+1:]...)
			return
		}
	}
}
