package types

type Wildcard struct {
	VillainCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &Wildcard{}
	i.Bin = "66d95d8c-600d-441f-a003-c52bc99df103"
	i.Name = "Wildcard"
	i.ImageUrl = "/wildcard.png"
}

func (i *Wildcard) RetaliatedOn() string {
	return ""
}
