package model

type Metric struct {
	Hostname string  `json:"hostname"`
	CPU      float64 `json:"cpu"`
	Memory   float64 `json:"memory"`
	Disk     float64 `json:"disk"`
}
