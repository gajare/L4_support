package main

import (
	"L4_support/internal/handler"
	"L4_support/internal/models"
	"L4_support/internal/repository"
	"L4_support/internal/service"
	"L4_support/pkg/logger"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize logger
	if err := logger.InitLogger(); err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer logger.Sync()

	// Database connection
	dsn := "host=postgres user=postgres password=password dbname=l4_support port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	// Auto migrate
	if err := db.AutoMigrate(&models.Employee{}); err != nil {
		logger.Logger.Fatal("Failed to migrate database", zap.Error(err))
	}

	// Initialize components
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	// Create Gin router
	r := gin.Default()

	// Routes
	api := r.Group("/api/v1")
	{
		api.POST("/employees", employeeHandler.CreateEmployee)
		api.GET("/employees", employeeHandler.GetAllEmployees)
		api.GET("/employees/:id", employeeHandler.GetEmployee)
		api.PUT("/employees/:id", employeeHandler.UpdateEmployee)
		api.DELETE("/employees/:id", employeeHandler.DeleteEmployee)
		api.GET("/metrics", handler.MetricsHandler())
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "timestamp": time.Now()})
	})

	logger.Logger.Info("Starting L4 Support service on :8080")
	if err := r.Run(":8080"); err != nil {
		logger.Logger.Fatal("Failed to start server", zap.Error(err))
	}
}
