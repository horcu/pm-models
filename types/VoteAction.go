package types

type VoteAction struct {
	Bin     string            `json:"bin"`
	StepBin string            `json:"step_bin"`
	GameBin string            `json:"game_bin"`
	Vote    Vote              `json:"vote,omitempty"`
	Media   map[string]*Media `json:"media,omitempty"` // map of string (timestamp millis) and message
}
