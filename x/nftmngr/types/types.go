package types

// Update first_mint_complete on metadata mint if first_mint_complete is false
const (
	KeyNFTStatusFirstMintComplete string = "first_mint_complete"
)

func (o *OnChainData) GetStatusByKey(key string) bool {
	// loop over o.Status
	for _, v := range o.Status {
		if v.StatusName == key {
			return v.StatusValue
		}
	}
	return false
}

// SetStatusByKey set status value by key
func (o *OnChainData) SetStatusByKey(key string, value bool) {
	// loop over o.Status
	found := false
	for _, v := range o.Status {
		if v.StatusName == key {
			v.StatusValue = value
			found = true
		}
	}
	if !found {
		o.Status = append(o.Status, &FlagStatus{
			StatusName:  key,
			StatusValue: value,
		})
	}
}
