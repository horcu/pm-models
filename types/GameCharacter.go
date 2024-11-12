package types

type GameCharacter struct {
	Bin         string              `json:"bin"`
	TypeId      int                 `json:"type_id"`
	IsInnocent  bool                `json:"is_innocent,omitempty"`
	Abilities   map[string]*Ability `json:"abilities,omitempty"`
	Name        string              `json:"name"`
	ImageUrl    string              `json:"image_url"`
	Description string              `json:"description"`
	Role        string              `json:"role"`
}

func (b *GameCharacter) LastVote() string {
	return ""
}

func (i *VillainCharacter) TimesUsed() string {
	return ""
}
