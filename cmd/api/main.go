package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"monitoring-system/internal/config"
	handler "monitoring-system/internal/handler/http"
	"monitoring-system/internal/repository"
	"monitoring-system/internal/service"
	"monitoring-system/migrations"
)

func main() {
	r := gin.Default()

	db, err := config.ConnectDB(context.Background())
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}
	defer db.Close()
	if err := migrations.Apply(context.Background(), db); err != nil {
		log.Fatal("Erro ao aplicar migrações:", err)
	}

	repo := repository.NewMetricsRepository(db)
	service := service.NewMetricsService(repo)
	metricsHandler := handler.NewMetricsHandler(service)

	r.GET("/health", handler.HealthHandler(db))

	r.GET("/metrics", metricsHandler.GetMetrics)
	r.POST("/metrics", metricsHandler.CreateMetric)

	log.Println("Servidor rodando na porta :8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
