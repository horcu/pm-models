package types

import c "github.com/horcu/mafia-models/types/characters"

type Achievements struct {
	Levels     []*Level           `json:"levels"`
	Characters []*c.GameCharacter `json:"characters"`
	Stats      []*GameStats       `json:"stats"`
}
