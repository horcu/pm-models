package types

type Step struct {
	Bin               string                    `json:"bin"`
	NextStep          string                    `json:"next_step"`
	StepType          string                    `json:"step_type"`
	Duration          string                    `json:"duration"`
	Command           string                    `json:"command"`
	ImageURL          string                    `json:"img_url,omitempty"`
	ShowTimer         bool                      `json:"show_timer"`
	Text              string                    `json:"text"`
	Characters        map[string]*GameCharacter `json:"characters,omitempty"`
	StepIndex         int                       `json:"step_index"`
	StartTime         string                    `json:"start_time"`
	RequiresVote      bool                      `json:"requires_vote"`
	VoteType          string                    `json:"vote_type"`
	VillainVoteCount  int                       `json:"villain_vote_count"`
	InnocentVoteCount int                       `json:"innocent_vote_count"`
	EndTime           string                    `json:"end_time"`
	Result            map[string][]*Result      `json:"result,omitempty"` // map of player bin and list of results
	Allowed           []string                  `json:"allowed"`
	SubSteps          map[string]*Step          `json:"sub_steps,omitempty"`
	CanVoteHere       []string                  `json:"can_vote_here"`
	YetToVote         []string                  `json:"ytv,omitempty"`
}

func (s *Step) GetNext(g *Game) *Step {

	return g.Steps[s.NextStep]
}
