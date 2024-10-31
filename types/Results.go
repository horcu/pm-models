package types

type Result struct {
	Bin       string            `json:"bin"`
	StepBin   string            `json:"step_bin"`
	GameBin   string            `json:"game_bin"`
	GamerId   string            `json:"gamer_id,omitempty"`
	TimeStamp string            `json:"timestamp"`
	Vote      Vote              `json:"vote,omitempty"`
	Media     map[string]*Media `json:"media,omitempty"` // map of string (timestamp millis) and message
}

func (result *Result) IsGamer(gamerId string) bool {
	if result.GamerId == gamerId {
		return true
	} else {
		return false
	}
}
