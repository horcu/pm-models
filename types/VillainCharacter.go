package types

func init() {
	// set the  isInnocent property in the base class
	i := &VillainCharacter{}
	i.IsInnocent = false
}

type VillainCharacter struct {
	GameCharacter
}

func (i *VillainCharacter) LastKill() string {
	return ""
}
