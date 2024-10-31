package types

type Metrics struct {
	LastHeal                 string   `json:"last_heal,omitempty"`
	LastVote                 string   `json:"last_vote,omitempty"`
	LastKill                 string   `json:"last_kill,omitempty"`
	LastPoisoning            string   `json:"last_poisoning,omitempty"`
	LastMimic                string   `json:"last_mimic,omitempty"`
	LastTrick                string   `json:"last_trick,omitempty"`
	Heals                    []string `json:"heals,omitempty"`
	TimesHealed              int      `json:"times_healed,omitempty"`
	Kills                    []string `json:"kills,omitempty"`
	Votes                    []string `json:"votes,omitempty"`
	Mimics                   []string `json:"mimics,omitempty"`
	TimesMimicked            int      `json:"times_mimicked,omitempty"`
	Tricks                   []string `json:"tricks,omitempty"`
	TimesTricked             int      `json:"times_tricked,omitempty"`
	Poisonings               []string `json:"poisonings,omitempty"`
	TimesPoisoned            int      `json:"times_poisoned,omitempty"`
	VotedAgainstBy           []string `json:"voted_against_by,omitempty"`
	VotedAgainstMostBy       string   `json:"voted_against_most_by,omitempty"`
	VoteSuccessRate          float64  `json:"vote_success_rate,omitempty"`
	NightsSurvived           int      `json:"nights_survived,omitempty"`
	SuccessfulInvestigations int      `json:"successful_investigations,omitempty"`
	TimesInvestigated        int      `json:"times_investigated,omitempty"`
}