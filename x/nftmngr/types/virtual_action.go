package types

func (v *VirtualAction) ToAction() *Action {
	return &Action{
		Name: v.Name,
		Desc: v.Desc,
		When: v.When,
		Then: v.Then,
		Disable: v.Disable,
		AllowedActioner: v.AllowedActioner,
		Params: v.Params,
	}
}