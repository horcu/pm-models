package types

type GameStats struct {
	GameIds    []string `json:"game_ids"`
	Roles      []string `json:"roles"`
	Wins       int      `json:"wins"`
	Losses     int      `json:"losses"`
	TotalGames int      `json:"total_games"`
}
