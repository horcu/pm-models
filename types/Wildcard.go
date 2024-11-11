package types

type Wildcard struct {
	VillainCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &Wildcard{}
	i.Bin = "66d95d8c-600d-441f-a003-c52bc99df103"
	i.Abilities = []*Ability{
		{
			Bin:            "9c39409d-9e57-47c5-8da5-635f6c3a499e",
			Name:           "Retaliate",
			Character:      "Wildcard",
			Description:    "A prickly porcupine with a dangerous secret.",
			CycleUsedIndex: 0,
			Frequency:      "Once",
			TimesUsed:      0,
			Instructions:   "Take down anyone who votes to kill you",
			UIColor:        "#118AB2",
		},
		{
			Bin:            "f406631e-d868-4185-94e9-406d2b5ec89a",
			Name:           "Mark",
			Character:      "Wildcard",
			Description:    "Marks a player for elimination",
			CycleUsedIndex: 0,
			Frequency:      "every",
			TimesUsed:      0,
			Instructions:   "Select a player to mark",
			UIColor:        "#FFA3AF",
		},
	}
	i.Name = "Wildcard"
	i.ImageUrl = "/wildcard.png"
}

func (i *Wildcard) RetaliatedOn() string {
	return ""
}
