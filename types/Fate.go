package types

type Fate struct {
	Bin          string `json:"bin,omitempty"`
	AbilityBin   string `json:"ability,omitempty"`
	TimeStamp    string `json:"timestamp,omitempty"`
	isEliminated bool   `json:"is_eliminated,omitempty"`
}
