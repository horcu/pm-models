package types

type Enforcer struct {
	VillainCharacter
}

func (i *Enforcer) LastKill() string {
	return ""
}
func (i *Enforcer) LastMark() string {
	return ""
}
