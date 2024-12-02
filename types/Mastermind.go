package types

type Mastermind struct {
	VillainCharacter
}

func (i *Mastermind) LastKill() string {
	return ""
}
func (i *Mastermind) LastVote() string {
	return ""
}
