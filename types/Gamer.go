package types

import "github.com/horcu/mafia-models"

type Gamer struct {
	Bin         string        `json:"bin"`
	GameId      string        `json:"game_id"`
	CharacterId string        `json:"character_id"`
	Name        string        `json:"name"`
	ImageUrl    string        `json:"image_url"`
	IsAlive     bool          `json:"is_alive"`
	Abilities   []*v1.Ability `json:"abilities"`
	Metrics     *v1.Metrics   `json:"metrics"`
}
