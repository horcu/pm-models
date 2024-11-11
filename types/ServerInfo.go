package types

type ServerInfo struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Port    int    `json:"port,omitempty"`
}
