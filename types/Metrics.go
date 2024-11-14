package types

type Metrics struct {
	Values map[string]interface{} `json:"values"` // enum metricsValues are keys
}
