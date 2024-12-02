package types

type Healer struct {
	HybridCharacter
	IsCorrupted bool
}

func (i *Healer) LastHealed() string {
	return ""
}

func (i *Healer) TimesUsed() string {
	return ""
}
