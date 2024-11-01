package types

import c "github.com/horcu/mafia-models/types/characters"

type CharacterPack struct {
	Name       string             `json:"name"`
	OriginDate string             `json:"origin_date"`
	Creator    string             `json:"creator"`
	Characters []*c.GameCharacter `json:"characters,omitempty"`
}
