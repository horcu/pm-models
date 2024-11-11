package types

type ServerInfo struct {
	Name      string `json:"name,omitempty"`
	Address   string `json:"address,omitempty"`
	Port      int    `json:"port,omitempty"`
	CreatorId string `json:"creator_id,omitempty"`
	GroupId   string `json:"group_id,omitempty"`
}
