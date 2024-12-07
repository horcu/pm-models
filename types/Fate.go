package types

type Fate struct {
	Bin             string `json:"bin,omitempty"`
	AbilityBin      string `json:"ability,omitempty"`
	TimeStamp       string `json:"timestamp,omitempty"`
	IsEliminated    bool   `json:"is_eliminated,omitempty"`
	CycleFateSet    string `json:"cycle_fate_set,omitempty"`
	CycleEliminated bool   `json:"cycle_eliminated,omitempty"`
}
