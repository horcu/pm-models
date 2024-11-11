package types

type CorruptHealer struct {
	VillainCharacter
	IsCorrupted bool `json:"is_corrupt"`
}

func init() {
	// set the  isInnocent property in the base class
	i := &CorruptHealer{}
	i.Bin = "3b0415b4-e39e-462a-b005-18b9d07ba740"
	i.Abilities = []*Ability{
		{
			Bin:            "35ed7a60-1c7a-43d1-432a-364529fe65e6",
			Name:           "Healing Touch",
			Character:      "CorruptedHealingCharacter",
			Description:    "A deceitful snake with healing powers used for evil",
			CycleUsedIndex: 0,
			Frequency:      "Once",
			TimesUsed:      0,
			Instructions:   "Select a villain character to heal",
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
	i.Name = "Corrupt Healer"
	i.ImageUrl = "/healer.png"
	i.TypeId = 0
	i.IsCorrupted = true
}
