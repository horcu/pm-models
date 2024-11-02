package types

type TrickyCharacter struct {
	VillainCharacter
}

func (t *TrickyCharacter) LastTricked() string {
	return ""
}

func (i *TrickyCharacter) TimesUsed() string {
	return ""
}
