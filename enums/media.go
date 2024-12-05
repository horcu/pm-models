package enums

import "strings"

type MediaType string

// Declare related constants for each ability starting with index 1
const (
	Audio     MediaType = "audio"
	Video     MediaType = "video"
	Text      MediaType = "text"
	GamerVote MediaType = "vote"
)

func (s MediaType) String() string {
	return strings.ToLower(string(s))
}
