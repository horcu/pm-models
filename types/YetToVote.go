package types

type YetToVote struct {
	Bin     string `json:"bin"`
	GameBin string `json:"game_bin"`
	StepBin string `json:"step_bin"`
	Gamer   Gamer  `json:"gamer"`
}
