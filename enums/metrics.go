package enums

import "strings"

type MetricsValues string

// Declare related constants for each weekday starting with index 1
const (
	LastHeal                 MetricsValues = "last_heal"
	LastVote                               = "last_vote"
	LastKill                               = "last_kill"
	LastPoisoning                          = "last_poisoning"
	LastMimic                              = "last_mimic"
	LastTrick                              = "last_trick"
	LastInvestigated                       = "last_investigated"
	Heals                                  = "heals"
	TimesHealed                            = "times_healed"
	Kills                                  = "kills"
	Votes                                  = "votes"
	Mimics                                 = "mimics"
	TimesUsed                              = "times_used"
	TimesBlocked                           = "times_blocked"
	HealedBy                               = "healed_by"
	KilledBy                               = "killed_by"
	MimickedBy                             = "mimicked_by"
	TrickedBy                              = "tricked_by"
	PoisonedBy                             = "poisoned_by"
	InvestigatedBy                         = "investigated_by"
	Investigations                         = "investigations"
	TimesMimicked                          = "times_mimicked"
	TimesVotedAgainst                      = "times_votes_against"
	Tricks                                 = "tricks"
	TimesTricked                           = "times_tricked"
	Poisonings                             = "poisonings"
	TimesPoisoned                          = "times_poisoned"
	VotedAgainstBy                         = "voted_against_by"
	VotedAgainstMostBy                     = "voted_against_most_by"
	VoteSuccessRate                        = "vote_success_rate"
	NightsSurvived                         = "nights_survived"
	SuccessfulInvestigations               = "successful_investigations"
	TimesInvestigated                      = "times_investigated"
)

func (a MetricsValues) String() string {
	return strings.ToLower(string(a))
}
