package types

type HybridCharacter struct {
	GameCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &HybridCharacter{}
	i.IsInnocent = false
}

func (i *HybridCharacter) Side() string {
	return "Hybrid"
}
