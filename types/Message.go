package types

type Message struct {
	Bin       string `json:"bin,omitempty"`
	Type      string `json:"type,omitempty"`
	Target    string `json:"target,omitempty"`
	Source    string `json:"source,omitempty"` // either the game bin or the bin of a gamer
	payload   *Media `json:"payload,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}
