package types

type DoppleGanger struct {
	HybridCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &DoppleGanger{}
	i.Bin = "3b0415b4-e39e-462a-b005-18b9d07ba740"
	i.Abilities = []*Ability{
		{
			Bin:            "ae6d5dff-c252-4841-8b96-75f9efeb7983",
			Name:           "Trickery",
			Character:      "DoppleGanger",
			Description:    "tricks a target character",
			CycleUsedIndex: 0,
			Frequency:      "Once",
			TimesUsed:      0,
			Instructions:   "A shape shifting Chamillion who blends in with the innocent but causes confusion",
			UIColor:        "#00FF00",
		},
		{
			Bin:            "f406631e-d868-4185-94e9-406d2b5ec89a",
			Name:           "Mark",
			Character:      "DoppleGanger",
			Description:    "Marks a player for elimination",
			CycleUsedIndex: 0,
			Frequency:      "every",
			TimesUsed:      0,
			Instructions:   "Select a player to mark",
			UIColor:        "#FFA3AF",
		},
	}
	i.Name = "DoppleGanger"
	i.ImageUrl = "/doppleganger.png"
}

func (i *DoppleGanger) LastTricked() string {
	return ""
}
