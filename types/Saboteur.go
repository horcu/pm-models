package types

type Saboteur struct {
	VillainCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &Saboteur{}
	i.Bin = "123e4567-e89b-12d3-a456-426614174000"
	i.Name = "Saboteur"
	i.ImageUrl = "/saboteur.png"
}

func (i *Trickster) LastBlocked() string {
	return ""
}
