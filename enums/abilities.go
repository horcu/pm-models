package enums

import "strings"

type AllowedAbilities string

// Declare related constants for each weekday starting with index 1
const (
	Vote        AllowedAbilities = "vote"
	Mimic                        = "mimic"
	Hide                         = "hide"
	Retaliate                    = "retaliate"
	Investigate                  = "investigate"
	Poison                       = "poison"
	Heal                         = "heal"
	Mark                         = "mark"
	Block                        = "block"
	Direct                       = "direct"
	Trick                        = "trick"
	Kill                         = "kill"
)

func (a AllowedAbilities) String() string {
	return strings.ToLower(string(a))
}
