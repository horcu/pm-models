package types

type Witch struct {
	InnocentCharacter
	IsCorrupted bool
}

func (i *Witch) LastHealed() string {
	return ""
}

func (i *Witch) TimesUsed() string {
	return ""
}
