package service

import (
	"monitoring-system/internal/dto"
	"monitoring-system/internal/model"
	"monitoring-system/internal/repository"
)

type MetricsService struct {
	repo *repository.MetricsRepository
}

func NewMetricsService(r *repository.MetricsRepository) *MetricsService {
	return &MetricsService{repo: r}
}

func (s *MetricsService) CreateMetric(metric model.Metric) (map[string]interface{}, error) {

	if metric.Hostname == "" {
		return nil, ErrHostnameRequired
	}

	err := s.repo.Save(metric)
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"message": "metric saved",
	}

	return response, nil
}

func (s *MetricsService) GetMetrics() ([]model.Metric, error) {
	metrics, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var response []model.Metric

	for _, m := range metrics {

		resp := model.Metric{
			Hostname: m.Hostname,
			CPU:      dto.Percent(m.CPU),
			Memory:   dto.Percent(m.Memory),
			Disk:     dto.Percent(m.Disk),
		}

		response = append(response, resp)
	}

	return response, nil
}
