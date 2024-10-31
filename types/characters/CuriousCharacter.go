package characters

type CuriousCharacter struct {
	InnocentCharacter
}

func (c *CuriousCharacter) LastGuess() string {
	return ""
}
