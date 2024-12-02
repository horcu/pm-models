package types

type Saboteur struct {
	VillainCharacter
}

func (i *Saboteur) LastBlocked(g *Game) string {
	return ""
}
