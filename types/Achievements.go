package types

type Achievements struct {
	Levels     []*Level         `json:"levels"`
	Characters []*GameCharacter `json:"characters"`
	Stats      []*GameStats     `json:"stats"`
}
