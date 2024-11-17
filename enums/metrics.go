package enums

import "strings"

type MetricsValues string

// Declare related constants for each metric value starting with index 1
const (
	LastHeal                 MetricsValues = "last_heal"
	LastVote                 MetricsValues = "last_vote"
	LastKill                 MetricsValues = "last_kill"
	LastPoisoning            MetricsValues = "last_poisoning"
	LastMimic                MetricsValues = "last_mimic"
	LastTrick                MetricsValues = "last_trick"
	LastInvestigated         MetricsValues = "last_investigated"
	Heals                    MetricsValues = "heals"
	TimesHealed              MetricsValues = "times_healed"
	Kills                    MetricsValues = "kills"
	Votes                    MetricsValues = "votes"
	Mimics                   MetricsValues = "mimics"
	TimesUsed                MetricsValues = "times_used"
	TimesBlocked             MetricsValues = "times_blocked"
	HealedBy                 MetricsValues = "healed_by"
	KilledBy                 MetricsValues = "killed_by"
	MimickedBy               MetricsValues = "mimicked_by"
	TrickedBy                MetricsValues = "tricked_by"
	PoisonedBy               MetricsValues = "poisoned_by"
	InvestigatedBy           MetricsValues = "investigated_by"
	Investigations           MetricsValues = "investigations"
	TimesMimicked            MetricsValues = "times_mimicked"
	TimesMarked              MetricsValues = "times_marked_against"
	Tricks                   MetricsValues = "tricks"
	TimesTricked             MetricsValues = "times_tricked"
	Poisonings               MetricsValues = "poisonings"
	TimesPoisoned            MetricsValues = "times_poisoned"
	MarkedBy                 MetricsValues = "marked_by"
	MarkedMostBy             MetricsValues = "marked_most_by"
	KilledByBy               MetricsValues = "killed_by"
	VoteSuccessRate          MetricsValues = "vote_success_rate"
	NightsSurvived           MetricsValues = "nights_survived"
	SuccessfulInvestigations MetricsValues = "successful_investigations"
	TimesInvestigated        MetricsValues = "times_investigated"
)

func (a MetricsValues) String() string {
	return strings.ToLower(string(a))
}
