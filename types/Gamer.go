package types

import (
	"github.com/horcu/pm-models/enums"
)

type Gamer struct {
	Bin         string              `json:"bin"`
	GameId      string              `json:"game_id"`
	CharacterId string              `json:"character_id"`
	Name        string              `json:"name"`
	Fate        *Fate               `json:"fate,omitempty"`
	ImageUrl    string              `json:"image_url"`
	IsAlive     bool                `json:"is_alive"`
	SeenSteps   []string            `json:"seen_steps"`  //list of step bins
	VotedSteps  []string            `json:"voted_steps"` // map of cycles and list of step bins
	Abilities   map[string]*Ability `json:"abilities,omitempty"`
	Metrics     *Metrics            `json:"metrics,omitempty"`
	State       enums.State         `json:"state,omitempty"`
}
