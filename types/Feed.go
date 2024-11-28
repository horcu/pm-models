package types

type Feed struct {
	Bin      string              `json:"bin,omitempty"`
	GameBin  string              `json:"game_bin,omitempty"`
	Messages map[string]*Message `json:"messages,omitempty"` // map of gamer bin and list of messages
}
