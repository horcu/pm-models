package types

type StepHistory struct {
	Bin            string      `json:"bin"`
	Stamp          string      `json:"stamp"` //millis timestamp same as key for step history entry
	StepBin        string      `json:"step_bin"`
	GameBin        string      `json:"game_bin"`
	VotedAgainstBy []string    `json:"voted_against_by"`
	VoteAction     *VoteAction `json:"actions,omitempty"`
}
