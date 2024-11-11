package types

type Mimicker struct {
	InnocentCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &Mimicker{}
	i.Bin = "1b0215b4-e39e-462a-ba05-18b2d07da714"
	i.Abilities = []*Ability{
		{
			Bin:            "f406631e-d868-4185-94e9-406d2b5ec89a",
			Name:           "Mark",
			Character:      "Mimicker",
			Description:    "Marks a player for elimination",
			CycleUsedIndex: 0,
			Frequency:      "every",
			TimesUsed:      0,
			Instructions:   "Select a player to mark",
			UIColor:        "#FFA3AF",
		},
		{
			Bin:            "f406631e-d868-4185-94e9-406d2b5ec89a",
			Name:           "Mark",
			Character:      "Mimicker",
			Description:    "Marks a player for elimination",
			CycleUsedIndex: 0,
			Frequency:      "every",
			TimesUsed:      0,
			Instructions:   "Select a player to mimic their abilities",
			UIColor:        "#FFA3AF",
		},
	}
	i.Name = "Mimicker"
	i.ImageUrl = "/mimicker.png"
}

func (i *Mimicker) LastMimic() string {
	return ""
}

func (i *Mimicker) TimesUsed() string {
	return ""
}
