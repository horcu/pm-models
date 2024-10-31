package types

type Vote struct {
	Bin       string `json:"bin"`
	StepBin   string `json:"step_bin"`
	Target    string `json:"target"`
	Source    string `json:"source"`
	TimeStamp string `json:"time_stamp"`
	VotedBy   string `json:"voted_by"`
	Ability   string `json:"ability"` // what ability to use on the target
}
