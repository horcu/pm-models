package enums

import "strings"

type AllowedAbilities string

// Declare related constants for each ability starting with index 1
const (
	Vote        AllowedAbilities = "vote"
	Mimic       AllowedAbilities = "mimic"
	Hide        AllowedAbilities = "hide"
	Retaliate   AllowedAbilities = "retaliate"
	Investigate AllowedAbilities = "investigate"
	Poison      AllowedAbilities = "poison"
	Heal        AllowedAbilities = "heal"
	Mark        AllowedAbilities = "mark"
	Block       AllowedAbilities = "block"
	Direct      AllowedAbilities = "direct"
	Trick       AllowedAbilities = "trick"
	Kill        AllowedAbilities = "kill"
)

func (a AllowedAbilities) String() string {
	return strings.ToLower(string(a))
}
