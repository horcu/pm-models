package types

type Level struct {
	Bin           string `json:"bin"`
	Value         string `json:"value"`
	AwardedPoints int    `json:"awarded_points"`
	Timestamp     string `json:"timestamp"`
}
