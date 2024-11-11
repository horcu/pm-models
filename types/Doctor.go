package types

type Doctor struct {
	Healer
	IsCorrupted bool `json:"is_corrupt"`
}

func init() {
	// set the  isInnocent property in the base class
	i := &Doctor{}
	i.Bin = "3b7e89b5-92e7-42c9-8675-8df8b58ba961"
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
			UIColor:        "#FB8500",
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
	i.Name = "Doctor"
	i.ImageUrl = "/doctor.png"
	i.TypeId = 1
	i.IsCorrupted = false

}

func (i *Doctor) LastHealed() string {
	return ""
}

func (i *Doctor) TimesUsed() string {
	return ""
}
