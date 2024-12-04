package types

type Message struct {
	Bin       string `json:"bin,omitempty"`
	Target    string `json:"target,omitempty"`
	Source    string `json:"source,omitempty"` // either the game bin or the bin of a gamer
	Payload   *Media `json:"payload,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}
