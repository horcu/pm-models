package types

type Healer struct {
	InnocentCharacter
	IsCorrupted bool `json:"is_corrupt"`
}

func init() {
	// set the  isInnocent property in the base class
	i := &Healer{}
	i.Bin = "02542022-9b7b-4fa2-a420-170c2fc1f054"
	i.Name = "Healer"
	i.ImageUrl = "/healer.png"
	i.TypeId = 1
	i.IsCorrupted = false

}

func (i *Healer) LastHealed() string {
	return ""
}

func (i *Healer) TimesUsed() string {
	return ""
}
