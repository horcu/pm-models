package types

type Mimicker struct {
	InnocentCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &Mimicker{}
	i.Bin = "1b0215b4-e39e-462a-ba05-18b2d07da714"
	i.Name = "Mimicker"
	i.ImageUrl = "/mimicker.png"
}

func (i *Mimicker) LastMimic() string {
	return ""
}

func (i *Mimicker) TimesUsed() string {
	return ""
}
