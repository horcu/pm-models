package types

type Leaderboard struct {
	Players      []*Player          `json:"players"`
	BestVillain  map[string]*Player `json:"best_villain"`
	BestInnocent map[string]*Player `json:"best_innocent"`
	MostWins     *Player            `json:"most_wins"`
}
