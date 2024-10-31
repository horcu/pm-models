package types

type Achievements struct {
	Levels     []*Level     `json:"levels"`
	Characters []*Character `json:"characters"`
	Stats      []*GameStats `json:"stats"`
}
