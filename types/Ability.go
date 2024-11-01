package types

type Ability struct {
	Bin            string `json:"bin"`
	Name           string `json:"name"`
	Character      string `json:"character"`
	Description    string `json:"description"`
	CycleUsedIndex int    `json:"cycle_used_index"`
	Frequency      string `json:"frequency"`
	TimesUsed      int    `json:"times_used,omitempty"`
	Instructions   string `json:"instructions,omitempty"`
}
