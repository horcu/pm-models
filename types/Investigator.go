package types

type Investigator struct {
	InnocentCharacter
}

func (c *Investigator) LastInvestigated() string {
	return ""
}
