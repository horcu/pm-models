package types

func init() {
	// set the  isInnocent property in the base class
	i := &GameCharacter{}
	i.IsInnocent = true
	i.TypeId = 0
}

type InnocentCharacter struct {
	GameCharacter
}
