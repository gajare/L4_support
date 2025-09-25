package handler

import (
	"L4_support/internal/models"
	"L4_support/internal/service"
	"L4_support/pkg/logger"
	"L4_support/pkg/metrics"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeeHandler(service *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	start := time.Now()

	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("POST", "/employees", "400").Inc()
		logger.Logger.Error("Invalid input", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.CreateEmployee(&employee); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("POST", "/employees", "500").Inc()
		logger.Logger.Error("Failed to create employee", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}

	metrics.HttpRequestsTotal.WithLabelValues("POST", "/employees", "201").Inc()
	metrics.HttpRequestDuration.WithLabelValues("POST", "/employees").Observe(time.Since(start).Seconds())
	logger.Logger.Info("Employee created successfully", zap.Uint("id", employee.ID))
	c.JSON(http.StatusCreated, employee)
}

func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	start := time.Now()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("GET", "/employees/:id", "400").Inc()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	employee, err := h.service.GetEmployee(uint(id))
	if err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("GET", "/employees/:id", "404").Inc()
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	metrics.HttpRequestsTotal.WithLabelValues("GET", "/employees/:id", "200").Inc()
	metrics.HttpRequestDuration.WithLabelValues("GET", "/employees/:id").Observe(time.Since(start).Seconds())
	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	start := time.Now()

	employees, err := h.service.GetAllEmployees()
	if err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("GET", "/employees", "500").Inc()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}

	metrics.HttpRequestsTotal.WithLabelValues("GET", "/employees", "200").Inc()
	metrics.HttpRequestDuration.WithLabelValues("GET", "/employees").Observe(time.Since(start).Seconds())
	c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	start := time.Now()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("PUT", "/employees/:id", "400").Inc()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("PUT", "/employees/:id", "400").Inc()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	employee.ID = uint(id)
	if err := h.service.UpdateEmployee(&employee); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("PUT", "/employees/:id", "500").Inc()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	metrics.HttpRequestsTotal.WithLabelValues("PUT", "/employees/:id", "200").Inc()
	metrics.HttpRequestDuration.WithLabelValues("PUT", "/employees/:id").Observe(time.Since(start).Seconds())
	logger.Logger.Info("Employee updated successfully", zap.Uint("id", employee.ID))
	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	start := time.Now()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("DELETE", "/employees/:id", "400").Inc()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteEmployee(uint(id)); err != nil {
		metrics.HttpRequestsTotal.WithLabelValues("DELETE", "/employees/:id", "500").Inc()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}

	metrics.HttpRequestsTotal.WithLabelValues("DELETE", "/employees/:id", "200").Inc()
	metrics.HttpRequestDuration.WithLabelValues("DELETE", "/employees/:id").Observe(time.Since(start).Seconds())
	logger.Logger.Info("Employee deleted successfully", zap.Int("id", id))
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
