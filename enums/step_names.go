package enums

import "strings"

type StepNames string

// Declare related constants for each ability starting with index 1
const (
	Intro       StepNames = "introductions"
	Discussions StepNames = "discussions"
	Gameplay    StepNames = "gameplay"
	GameOver    StepNames = "gameover"
	Transition  StepNames = "transition"
	Defend      StepNames = "defend"
	Eliminate   StepNames = "elimination"
	Result      StepNames = "result"
)

func (a StepNames) String() string {
	return strings.ToLower(string(a))
}
