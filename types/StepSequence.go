package types

type StepSequence struct {
	SDuration string `json:"s_duration"`
	CStep     Step   `json:"c_step"`
	GameId    string `json:"game_id"`
}
