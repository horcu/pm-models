package types

type Doctor struct {
	Healer
	IsCorrupted bool `json:"is_corrupt"`
}

func init() {
	// set the  isInnocent property in the base class
	i := &Doctor{}
	i.Bin = "3b7e89b5-92e7-42c9-8675-8df8b58ba961"
	i.Name = "Doctor"
	i.ImageUrl = "/doctor.png"
	i.SideId = 1
	i.IsCorrupted = false

}

func (i *Doctor) LastHealed() string {
	return ""
}

func (i *Doctor) TimesUsed() string {
	return ""
}
