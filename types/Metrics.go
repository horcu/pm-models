package types

type Metrics struct {
	Values map[string][]string `json:"values,omitempty"`  // enum metricsValues are keys
	Single map[string]string   `json:"singles,omitempty"` // enum metricsValues are keys
	Sums   map[string]int      `json:"sums,omitempty"`    // enum metricsValues are keys
}
