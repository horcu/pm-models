package types

type Mastermind struct {
	Character
}

func (m *Mastermind) LastKill() string {
	return ""
}
func (m *Mastermind) LastVote() string {
	return ""
}

func (m *Mastermind) Abilities() []string {
	return m.Character.Abilities
}
