package types

type Trickster struct {
	VillainCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &Trickster{}
	i.Bin = "62542022-9b7b-4fa2-a420-172c2fc2f054"
	i.Abilities = []*Ability{
		{
			Bin:            "ae6d5dff-c252-4841-8b96-75f9efeb7983",
			Name:           "Trickery",
			Character:      "Trickster",
			Description:    "tricks a target character",
			CycleUsedIndex: 0,
			Frequency:      "Once",
			TimesUsed:      0,
			Instructions:   "Select a character to heal",
			UIColor:        "#00FF00",
		},
		{
			Bin:            "f406631e-d868-4185-94e9-406d2b5ec89a",
			Name:           "Mark",
			Character:      "Trickster",
			Description:    "Marks a player for elimination",
			CycleUsedIndex: 0,
			Frequency:      "every",
			TimesUsed:      0,
			Instructions:   "Select a player to mark",
			UIColor:        "#FFA3AF",
		},
	}
	i.Name = "Trickster"
	i.ImageUrl = "/trickster.png"
}

func (i *Trickster) LastTricked() string {
	return ""
}
