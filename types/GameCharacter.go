package types

type GameCharacter struct {
	Bin           string              `json:"bin"`
	SideId        int                 `json:"type_id,omitempty"`
	IsInnocent    bool                `json:"is_innocent,omitempty"`
	Abilities     map[string]*Ability `json:"abilities,omitempty"`
	Name          string              `json:"name,omitempty"`
	ImageUrl      string              `json:"image_url,omitempty"`
	Description   string              `json:"description,omitempty"`
	Role          string              `json:"role,omitempty"`
	CharacterType interface{}         `json:"character_type,omitempty"`
}

func init() {
}

func (i *GameCharacter) LastVote() string {
	return ""
}

func (i *GameCharacter) TimesUsed() string {
	return ""
}
