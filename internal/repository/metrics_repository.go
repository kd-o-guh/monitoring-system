package repository

import (
	"database/sql"

	"monitoring-system/internal/model"
)

type MetricsRepository struct {
	db *sql.DB
}

func NewMetricsRepository(db *sql.DB) *MetricsRepository {
	return &MetricsRepository{db: db}
}

func (r *MetricsRepository) Save(metric model.Metric) error {
	query := `
		INSERT INTO metrics (hostname, cpu, memory, disk)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(query, metric.Hostname, metric.CPU, metric.Memory, metric.Disk)
	return err
}

func (r *MetricsRepository) FindAll() ([]model.Metric, error) {
	query := `SELECT hostname, cpu, memory, disk FROM metrics`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metrics []model.Metric

	for rows.Next() {
		var m model.Metric

		err := rows.Scan(&m.Hostname, &m.CPU, &m.Memory, &m.Disk)
		if err != nil {
			return nil, err
		}

		metrics = append(metrics, m)
	}

	return metrics, nil
}
