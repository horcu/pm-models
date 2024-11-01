package types

import c "github.com/horcu/mafia-models/types/characters"

type Game struct {
	Bin               string                      `json:"bin"`
	IsDaytime         bool                        `json:"is_daytime,omitempty"`
	ExplanationSeen   bool                        `json:"explanation_seen,omitempty"`
	FirstDayCompleted bool                        `json:"first_day_completed,omitempty"`
	Level             string                      `json:"level"`
	CurrentStep       string                      `json:"current_step"`
	GroupId           string                      `json:"group_id"`
	Status            string                      `json:"status"`
	ServerName        string                      `json:"server_name,omitempty"`
	ServerAddress     string                      `json:"server_ip,omitempty"`
	ServerPort        int                         `json:"server_port,omitempty"`
	StartTime         string                      `json:"start_time,omitempty"`
	EndTime           string                      `json:"end_time,omitempty"`
	Creator           *Player                     `json:"creator,omitempty"`
	NightCycles       int                         `json:"cycles,omitempty"`
	Invited           []string                    `json:"invited,omitempty"`
	AllVotesSubmitted bool                        `json:"all_votes_submitted,omitempty"`
	Steps             map[string]*Step            `json:"steps,omitempty"`
	Gamers            map[string]*Gamer           `json:"gamers,omitempty"`
	Characters        map[string]*c.GameCharacter `json:"characters,omitempty"`
	StepResults       map[string][]*Result        `json:"step_results,omitempty"` // map of playerId and the list of results
}
