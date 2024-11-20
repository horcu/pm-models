package enums

import "strings"

type AbilityUseFrequency string

// Declare related constants for each ability starting with index 1
const (
	Once       AbilityUseFrequency = "once"
	EveryOther AbilityUseFrequency = "every_other"
	Every      AbilityUseFrequency = "every"
	Twice      AbilityUseFrequency = "twice"
)

func (a AbilityUseFrequency) String() string {
	return strings.ToLower(string(a))
}
