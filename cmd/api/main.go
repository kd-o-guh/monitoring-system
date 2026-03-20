package main

import (
	"log"
	"net/http"

	"monitoring-system/internal/config"
	handler "monitoring-system/internal/handler/http"
	"monitoring-system/internal/repository"
	"monitoring-system/internal/service"
)

func main() {
	mux := http.NewServeMux()

	db := config.ConnectDB()

	repo := repository.NewMetricsRepository(db)
	service := service.NewMetricsService(repo)
	metricsHandler := handler.NewMetricsHandler(service)

	mux.HandleFunc("/health", handler.HealthHandler)

	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			metricsHandler.GetMetrics(w, r)
			return
		}
		if r.Method == http.MethodPost {
			metricsHandler.CreateMetric(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	log.Println("Servidor rodando na porta :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
