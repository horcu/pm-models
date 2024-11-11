package types

type Healer struct {
	InnocentCharacter
	IsCorrupted bool `json:"is_corrupt"`
}

func init() {
	// set the  isInnocent property in the base class
	i := &Healer{}
	i.Bin = "02542022-9b7b-4fa2-a420-170c2fc1f054"
	i.Abilities = []*Ability{
		{
			Bin:            "35ed7a60-1c7a-43d1-432a-364529fe65e6",
			Name:           "Healing Touch",
			Character:      "Healer",
			Description:    "Heals a target character",
			CycleUsedIndex: 0,
			Frequency:      "Once",
			TimesUsed:      0,
			Instructions:   "Select a character to heal",
			UIColor:        "#00FF00",
		},
		{
			Bin:            "f406631e-d868-4185-94e9-406d2b5ec89a",
			Name:           "Mark",
			Character:      "Healer",
			Description:    "Marks a player for elimination",
			CycleUsedIndex: 0,
			Frequency:      "every",
			TimesUsed:      0,
			Instructions:   "Select a player to mark",
			UIColor:        "#FFA3AF",
		},
	}
	i.Name = "Healer"
	i.ImageUrl = "/healer.png"
	i.TypeId = 1
	i.IsCorrupted = false

}

func (i *Healer) LastHealed() string {
	return ""
}

func (i *Healer) TimesUsed() string {
	return ""
}
