package types

type Doctor struct {
	InnocentCharacter
}

func (i *Doctor) LastHealed() string {
	return ""
}

func (i *Doctor) TimesUsed() string {
	return ""
}
