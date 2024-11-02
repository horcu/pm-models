package types

type MimicCharacter struct {
	InnocentCharacter
}

func (i *MimicCharacter) LastMimic() string {
	return ""
}

func (i *MimicCharacter) TimesUsed() string {
	return ""
}
