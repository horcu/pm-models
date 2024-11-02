package types

type VillainCharacter struct {
	BaseCharacter
}

func (v *VillainCharacter) LastKill() string {
	return ""
}
