package types

import "github.com/horcu/pm-models/enums"

type Media struct {
	Bin       string          `json:"bin"`
	Type      enums.MediaType `json:"type,omitempty"`
	Icon      string          `json:"icon,omitempty"`
	Text      string          `json:"text,omitempty"`
	Path      string          `json:"path,omitempty"` // either audio or video path
	TimeStamp string          `json:"time_stamp,omitempty"`
}
