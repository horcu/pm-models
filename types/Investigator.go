package types

type Investigator struct {
	InnocentCharacter
}

func init() {

	i := &Investigator{}
	i.Bin = "e2ab27bc-bf6e-4b13-a589-0fb47f22b8c2"
	i.Abilities = []*Ability{
		{
			Bin:            "e588a971-1766-4166-8199-19686f56689a",
			Name:           "Investigate",
			Character:      "Investigator",
			Description:    "Investigate a player to see if they are a villain",
			CycleUsedIndex: 0,
			Frequency:      "Every",
			TimesUsed:      0,
			Instructions:   "Select a player to investigate",
			UIColor:        "#0000FF",
		},
		{
			Bin:            "f406631e-d868-4185-94e9-406d2b5ec89a",
			Name:           "Mark",
			Character:      "Investigator",
			Description:    "Marks a player for elimination",
			CycleUsedIndex: 0,
			Frequency:      "every",
			TimesUsed:      0,
			Instructions:   "Select a player to mark",
			UIColor:        "#FFA3AF",
		},
	}
	i.Name = "Investigator"
	i.ImageUrl = "/investigator.png"
	i.TypeId = 2

}

func (c *Investigator) LastInvestigated() string {
	return ""
}
