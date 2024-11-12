package types

type CharacterPack struct {
	Name       string                    `json:"name"`
	OriginDate string                    `json:"origin_date"`
	Creator    string                    `json:"creator"`
	Characters map[string]*GameCharacter `json:"characters,omitempty"`
}
