package types

type Hitman struct {
	VillainCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &Hitman{}
	i.Abilities = []*Ability{
		{
			Bin:            "f406631e-d868-4185-94e9-406d2b5ec89a",
			Name:           "Mark",
			Character:      "Hitman",
			Description:    "Marks a player for elimination",
			CycleUsedIndex: 0,
			Frequency:      "every",
			TimesUsed:      0,
			Instructions:   "Select a player to mark",
			UIColor:        "#FFA3AF",
		},
	}
	i.Name = "Hitman"
	i.ImageUrl = "/hitman.png"
	i.TypeId = 1
}

func (i *Hitman) LastKill() string {
	return ""
}
func (i *Hitman) LastVote() string {
	return ""
}
