package types

type Investigator struct {
	InnocentCharacter
}

func init() {

	i := &Investigator{}
	i.Bin = "e2ab27bc-bf6e-4b13-a589-0fb47f22b8c2"
	i.Name = "Investigator"
	i.ImageUrl = "/investigator.png"
	i.SideId = 2

}

func (c *Investigator) LastInvestigated() string {
	return ""
}
