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

	if i.TypeId == 0 {
		return "Innocent"
	} else if i.TypeId == 1 {
		return "Villain"
	}
	return "Hybrid"
}
