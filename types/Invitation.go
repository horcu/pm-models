package types

type Invitation struct {
	Bin        string `json:"bin"`
	GameGroup  string `json:"game_group"`
	CreatorId  string `json:"creator_id"`
	Status     string `json:"status"`
	Invitation string `json:"invitation"` // type - game/group/message etc.
	Message    string `json:"message"`
	Time       string `json:"time"`
	GameId     string `json:"game_id"`
	Accepted   bool   `json:"accepted"`
	Declined   bool   `json:"declined"`
}
