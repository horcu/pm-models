package types

type Gamer struct {
	Bin         string     `json:"bin"`
	GameId      string     `json:"game_id"`
	CharacterId string     `json:"character_id"`
	Name        string     `json:"name"`
	ImageUrl    string     `json:"image_url"`
	IsAlive     bool       `json:"is_alive"`
	Abilities   []*Ability `json:"abilities"`
	Metrics     *Metrics   `json:"metrics"`
}
