package types

type Feed struct {
	Bin      string                `json:"bin,omitempty"`
	GameBin  string                `json:"game_bin,omitempty"`
	StepBin  string                `json:"step_bin,omitempty"`
	Messages map[string][]*Message `json:"messages,omitempty"`
	// map of timestamp and list of messages
}
