package types

type Player struct {
	UserName     string                 `json:"user_name"`
	Bin          string                 `json:"bin"`
	Photo        string                 `json:"photo"`
	Status       string                 `json:"status"` //inAGame, afk, available
	Privacy      string                 `json:"privacy"`
	Invitations  map[string]*Invitation `json:"invitations,omitempty"`
	GroupIds     []string               `json:"group_ids,omitempty"`
	GameIds      []string               `json:"game_ids,omitempty"`
	Achievements *Achievements          `json:"achievements,omitempt"`
}
