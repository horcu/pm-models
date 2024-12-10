package types

type Events struct {
	IsVoting          bool `json:"is_voting,omitempty"`
	IsGamePlay        bool `json:"is_game_play,omitempty"`
	IsInTransition    bool `json:"is_in_transition,omitempty"`
	IsInDiscussion    bool `json:"is_in_discussion,omitempty"`
	IsInDefend        bool `json:"is_in_defend,omitempty"`
	IsCharacterStep   bool `json:"is_character_step,omitempty"`
	IsGameOver        bool `json:"is_game_over,omitempty"`
	IsIntroductions   bool `json:"is_introductions,omitempty"`
	IsResultStep      bool `json:"is_result_step,omitempty"`
	AllVotesSubmitted bool `json:"all_votes_submitted,omitempty"`
	IsDaytime         bool `json:"is_daytime,omitempty"`
	ExplanationSeen   bool `json:"explanation_seen,omitempty"`
	FirstDayCompleted bool `json:"first_day_completed,omitempty"`
}
