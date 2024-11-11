package types

type HybridCharacter struct {
	InnocentCharacter
	VillainCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &HybridCharacter{}
	i.IsInnocent = false
}

func (i *HybridCharacter) Side() {

	if i.TypeId == 3 {
		return
	}
}
