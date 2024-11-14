package types

type StepSequence struct {
	SDuration   string `json:"s_duration"`
	CurrentStep *Step  `json:"current_step"`
	Game        *Game  `json:"game"`
}
