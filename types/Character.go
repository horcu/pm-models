package types

type Character struct {
	Bin         string   `json:"bin"`
	TypeId      int      `json:"type_id"`
	Name        string   `json:"name"`
	ImageUrl    string   `json:"image_url"`
	Description string   `json:"description"`
	Abilities   []string `json:"abilities,omitempty"`
	Role        string   `json:"role"`
}
