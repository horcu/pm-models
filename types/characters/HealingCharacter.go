package characters

type HealingCharacter struct {
	InnocentCharacter
}

func (i *HealingCharacter) LastHealed() string {
	return ""
}

func (i *HealingCharacter) TimesUsed() string {
	return ""
}
