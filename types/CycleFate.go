package types

type CycleFate struct {
	Bin        string           `json:"bin,omitempty`          // the cycle
	GamersFate map[string]*Fate `json:"gamer_fates,omitempty"` // the fate of all gamers in that cycle
}
