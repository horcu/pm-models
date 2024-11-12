package types

type Ability struct {
	Bin             string `json:"bin"`
	Name            string `json:"name"`
	Character       string `json:"character"`
	Description     string `json:"description"`
	CyclesUsedIndex []int  `json:"cycle_used_index,omitempty"`
	Frequency       string `json:"frequency"`
	TimesUsed       int    `json:"times_used,omitempty"`
	Instructions    string `json:"instructions,omitempty"`
	UIColor         string `json:"ui_color,omitempty"`
}
