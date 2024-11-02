package types

type Step struct {
	Bin               string                    `json:"bin"`
	NextStep          string                    `json:"next_step"`
	StepType          string                    `json:"step_type"`
	Duration          string                    `json:"duration"`
	Command           string                    `json:"command"`
	ShowTimer         bool                      `json:"show_timer"`
	CharPhotoUrl      string                    `json:"char_photo_url"`
	Characters        map[string]*GameCharacter `json:"characters"`
	StepIndex         int                       `json:"step_index"`
	StartTime         string                    `json:"start_time"`
	RequiresVote      bool                      `json:"requires_vote"`
	VoteType          string                    `json:"vote_type"`
	VillainVoteCount  int                       `json:"villain_vote_count"`
	InnocentVoteCount int                       `json:"innocent_vote_count"`
	EndTime           string                    `json:"end_time"`
	Result            map[int]*Result           `json:"result"` // map of game cycles and results
	Allowed           []string                  `json:"allowed"`
	SubSteps          map[string]*Step          `json:"sub_steps,omitempty"`
}
