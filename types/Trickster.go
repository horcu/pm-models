package types

type Trickster struct {
	VillainCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &Trickster{}
	i.Bin = "62542022-9b7b-4fa2-a420-172c2fc2f054"
	i.Name = "Trickster"
	i.ImageUrl = "/trickster.png"
}

func (i *Trickster) LastTricked() string {
	return ""
}
