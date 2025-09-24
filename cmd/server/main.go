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
	// initialize logger
	if err := logger.InitLogger(); err != nil {
		log.Fatal("Fail to initialize logger")
	}
	defer logger.Sync()

	//database connection

	dsn := "host=postgres user=postgres password=postgres dbname=l4_support port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Fatal("Failed ti connect to database", zap.Error(err))
	}

	//auto migration
	if err := db.AutoMigrate(&models.Employee{}); err != nil {
		logger.Logger.Fatal("Fail to migrate database", zap.Error(err))
	}

	//initialize componets

	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeehandler(employeeService)

	//gin router
	r := gin.Default()

	//Routers
	api := r.Group("/api/v1")
	{
		api.POST("/employees", employeeHandler.CreateEmployee)
		api.GET("/employees", employeeHandler.GetAllEmployee)
		api.GET("/employees/:id", employeeHandler.GetEmployee)
		api.PUT("/employees/:id", employeeHandler.UpdateEmployee)
		api.DELETE("/employees/:id", employeeHandler.DeleteEmployee)
		api.GET("/metrics", handler.MetricsHandler())
	}

	//Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "timestamp": time.Now()})
	})

	logger.Logger.Info("Starting L4 Support Service on :8080")
	if err := r.Run(":8080"); err != nil {
		logger.Logger.Fatal("failed to start server", zap.Error(err))
	}
}
