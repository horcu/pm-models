package types

type Saboteur struct {
	VillainCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &Saboteur{}
	i.Bin = "123e4567-e89b-12d3-a456-426614174000"
	i.Abilities = []*Ability{
		{
			Bin:            "b67d33c0-f307-46ef-a1ca-5091cb4f1a92",
			Name:           "Block",
			Character:      "Saboteur",
			Description:    "Saboteur who blocks the healer",
			CycleUsedIndex: 0,
			Frequency:      "Once",
			TimesUsed:      0,
			Instructions:   "Block the witch from healing or poisoning someone",
			UIColor:        "#00FF00",
		},
		{
			Bin:            "f406631e-d868-4185-94e9-406d2b5ec89a",
			Name:           "Mark",
			Character:      "Saboteur",
			Description:    "Marks a player for elimination",
			CycleUsedIndex: 0,
			Frequency:      "every",
			TimesUsed:      0,
			Instructions:   "Select a player to mark",
			UIColor:        "#FFA3AF",
		},
	}
	i.Name = "Saboteur"
	i.ImageUrl = "/saboteur.png"
}

func (i *Trickster) LastBlocked() string {
	return ""
}
