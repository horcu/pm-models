package types

type Game struct {
	Bin               string               `json:"bin"`
	IsDaytime         bool                 `json:"is_daytime,omitempty"`
	ExplanationSeen   bool                 `json:"explanation_seen,omitempty"`
	FirstDayCompleted bool                 `json:"first_day_completed,omitempty"`
	CurrentStep       string               `json:"current_step"`
	Status            string               `json:"status"`
	Info              *ServerInfo          `json:"info,omitempty"`
	StartTime         string               `json:"start_time,omitempty"`
	EndTime           string               `json:"end_time,omitempty"`
	IsVoting          bool                 `json:"is_voting,omitempty"`
	IsResultStep      bool                 `json:"is_result_step,omitempty"`
	Creator           *Player              `json:"creator,omitempty"`
	NightCycles       int                  `json:"cycles,omitempty"`
	AllVotesSubmitted bool                 `json:"all_votes_submitted,omitempty"`
	Steps             map[string]*Step     `json:"steps,omitempty"`
	Gamers            map[string][]*Gamer  `json:"gamers,omitempty"`
	StepResults       map[string][]*Result `json:"result,omitempty"` // map of playerId and the list of results
}
