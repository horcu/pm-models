package characters

type VillainCharacter struct {
	BaseCharacter
}

func (v *VillainCharacter) LastKill() string {
	return ""
}
