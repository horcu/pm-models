package types

type CorruptHealer struct {
	VillainCharacter
	IsCorrupted bool `json:"is_corrupt"`
}

func init() {
	// set the  isInnocent property in the base class
	i := &CorruptHealer{}
	i.Bin = "3b0415b4-e39e-462a-b005-18b9d07ba740"
	i.Name = "Corrupt Healer"
	i.ImageUrl = "/healer.png"
	i.TypeId = 0
	i.IsCorrupted = true
}
