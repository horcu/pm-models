package types

type Group struct {
	Bin       string             `json:"bin"`
	Creator   *Player            `json:"creator"`
	Members   map[string]*Player `json:"members,omitempty"`
	GroupName string             `json:"group_name"`
	Capacity  int                `json:"capacity"`
	Status    string             `json:"status"`
}
