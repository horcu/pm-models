package types

type Hitman struct {
	VillainCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &Hitman{}
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
