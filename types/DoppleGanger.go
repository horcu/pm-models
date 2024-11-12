package types

type DoppleGanger struct {
	HybridCharacter
}

func init() {
	// set the  isInnocent property in the base class
	i := &DoppleGanger{}
	i.Bin = "3b0415b4-e39e-462a-b005-18b9d07ba740"
	i.Name = "DoppleGanger"
	i.ImageUrl = "/doppleganger.png"
}

func (i *DoppleGanger) LastTricked() string {
	return ""
}
