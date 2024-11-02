package types

type CharacterPack struct {
	Name       string           `json:"name"`
	OriginDate string           `json:"origin_date"`
	Creator    string           `json:"creator"`
	Characters []*GameCharacter `json:"characters,omitempty"`
}
