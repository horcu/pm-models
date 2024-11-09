package types

type Gamer struct {
	Bin         string           `json:"bin"`
	GameId      string           `json:"game_id"`
	CharacterId string           `json:"character_id"`
	Name        string           `json:"name"`
	Fate        *Fate            `json:"fate,omitempty"`
	ImageUrl    string           `json:"image_url"`
	IsAlive     bool             `json:"is_alive"`
	SeenSteps   map[int][]string `json:"seen_steps"`  // map of cycles and list of step bins
	VotedSteps  map[int][]string `json:"voted_steps"` // map of cycles and list of step bins
	Abilities   []*Ability       `json:"abilities,omitempty"`
	Metrics     *Metrics         `json:"metrics,omitempty"`
}
