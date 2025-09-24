package handler

import (
	"L4_support/internal/models"
	"L4_support/internal/service"
	"L4_support/pkg/logger"
	"L4_support/pkg/metrics"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeehandler(service *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	start := time.Now()
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("POST", "/employees", "400").Inc()
		logger.Logger.Error("Invalide inpute", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.CreateEmployee(&employee); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("POST", "/employees", "500").Inc()
		logger.Logger.Error("Fail to create employee", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}

	metrics.HttpRequestsTotal.WithLabelValues("POST", "/employee", "201").Inc()
	metrics.HttpRequestDuration.WithLabelValues("POST", "/employees").Observe(time.Since(start).Seconds())

	logger.Logger.Info("Employee created successfully", zap.Uint("id", employee.ID))
	c.JSON(http.StatusCreated, employee)
}
