package types

type Result struct {
	Bin         string                  `json:"bin"`
	Index       int                     `json:"index"`
	EndTime     string                  `json:"end_time"`
	StepHistory map[string]*StepHistory `json:"decisions,omitempty"` //map of playerId and their stepHistory
	TimeStamp   string                  `json:"timestamp"`
}

func (result *Result) IsGamer(gamerId string) bool {
	if ok, _ := result.StepHistory[gamerId]; ok != nil {
		return true
	} else {
		return false
	}
}
