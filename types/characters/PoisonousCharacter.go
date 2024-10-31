package characters

type PoisonousCharacter struct {
	VillainCharacter
}

func (i *PoisonousCharacter) LastPoisoned() string {
	return ""
}

func (i *PoisonousCharacter) IsCorrupt() bool {
	return false
}

func (i *PoisonousCharacter) TimesUsed() string {
	return ""
}
