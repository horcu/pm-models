package types

type CuriousCharacter struct {
	InnocentCharacter
}

func (c *CuriousCharacter) LastInvestigated() string {
	return ""
}
