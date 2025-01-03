package enums

import "strings"

type State string

// Declare related constants for each ability starting with index 1
const (
	Marked            State = "marked"
	Healed            State = "healed"
	Poisoned          State = "poisoned"
	RetaliatedAgainst State = "retaliatedAgainst"
	Hidden            State = "hidden"
	Tricked           State = "tricked"
	Mimicked          State = "mimicked"
	Killed            State = "killed"
	Investigated      State = "investigated"
	None              State = "none"
)

func (s State) String() string {
	return strings.ToLower(string(s))
}
