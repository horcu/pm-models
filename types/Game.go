package types

type Game struct {
	Bin             string                    `json:"bin"`
	GroupId         string                    `json:"group_id,omitempty"`
	GameOverStepBin string                    `json:"gameover_step_bin,omitempty"`
	FirstStepBin    string                    `json:"first_step_bin,omitempty"`
	GameplayBin     string                    `json:"gameplay_bin,omitempty"`
	CurrentStep     string                    `json:"current_step"`
	NightCycles     int                       `json:"cycles,omitempty"`
	Status          string                    `json:"status"`
	Creator         *Player                   `json:"creator,omitempty"`
	Info            *ServerInfo               `json:"info,omitempty"`
	Sync            *TimeSync                 `json:"sync,omitempty"`
	GameEvents      Events                    `json:"events,omitempty"`
	Messages        map[string]*Message       `json:"messages,omitempty"`
	MainSuspects    map[string][]*Gamer       `json:"main_suspects,omitempty"`
	Steps           map[string]*Step          `json:"steps,omitempty"`
	Characters      map[string]*GameCharacter `json:"characters,omitempty"`
	Gamers          map[string]*Gamer         `json:"gamers,omitempty"`
	StepResults     map[string][]*Result      `json:"result,omitempty"`     // map of playerId and the list of results
	CycleFate       map[string]*CycleFate     `json:"cycle_fate,omitempty"` // map of user and the cycle fate
	Winners         []*Gamer                  `json:"winners,omitempty"`
}
