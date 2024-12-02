package enums

import "strings"

type Characters string

// Declare related constants for each ability starting with index 1
const (
	Mastermind   Characters = "mastermind"
	Enforcer     Characters = "enforcer"
	Doppelganger Characters = "doppelganger"
	Healer       Characters = "healer"
	Witch        Characters = "witch"
	Investigator Characters = "investigator"
	Coward       Characters = "coward"
	Doctor       Characters = "doctor"
	Saboteur     Characters = "saboteur"
	Wildcard     Characters = "wildcard"
	Sly          Characters = "sly"
	Lookout      Characters = "lookout"
	Moderator    Characters = "moderator"
)

func (a Characters) String() string {
	return strings.ToLower(string(a))
}
