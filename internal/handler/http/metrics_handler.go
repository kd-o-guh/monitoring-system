package http

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"monitoring-system/internal/dto"
	"monitoring-system/internal/model"
	"monitoring-system/internal/service"

	"github.com/gin-gonic/gin"
)

type MetricsHandler struct {
	service *service.MetricsService
}

type createMetricRequest struct {
	Hostname string   `json:"hostname" binding:"required"`
	CPU      *float64 `json:"cpu" binding:"required"`
	Memory   *float64 `json:"memory" binding:"required"`
	Disk     *float64 `json:"disk" binding:"required"`
}

func NewMetricsHandler(s *service.MetricsService) *MetricsHandler {
	return &MetricsHandler{service: s}
}

func (h *MetricsHandler) CreateMetric(c *gin.Context) {
	var request createMetricRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	metric := model.Metric{
		Hostname: request.Hostname,
		CPU:      dto.Percent(*request.CPU),
		Memory:   dto.Percent(*request.Memory),
		Disk:     dto.Percent(*request.Disk),
	}

	response, err := h.service.CreateMetric(metric)
	if err != nil {
		if errors.Is(err, service.ErrHostnameRequired) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, service.ErrMetricOutOfRange) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("erro ao salvar métrica: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error saving metrics"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *MetricsHandler) GetMetrics(c *gin.Context) {
	limit, err := queryInt(c, "limit", 100)
	if err != nil || limit < 1 || limit > 500 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be between 1 and 500"})
		return
	}
	offset, err := queryInt(c, "offset", 0)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "offset must be a non-negative integer"})
		return
	}

	metrics, err := h.service.GetMetrics(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching metrics"})
		return
	}

	c.JSON(http.StatusOK, metrics)
}

func queryInt(c *gin.Context, key string, fallback int) (int, error) {
	value := c.Query(key)
	if value == "" {
		return fallback, nil
	}
	return strconv.Atoi(value)
}
