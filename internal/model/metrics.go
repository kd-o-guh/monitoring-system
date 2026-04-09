package model

import "monitoring-system/internal/dto"

type Metric struct {
	Hostname string      `json:"hostname"`
	CPU      dto.Percent `json:"cpu"`
	Memory   dto.Percent `json:"memory"`
	Disk     dto.Percent `json:"disk"`
}
