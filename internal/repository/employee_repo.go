package repository

import (
	"L4_support/internal/models"
	"L4_support/pkg/logger"
	"L4_support/pkg/metrics"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(employee *models.Employee) error {
	metrics.DatabaseOperationsTotal.WithLabelValues("create", "employees").Inc()
	result := r.db.Create(employee)
	if result.Error != nil {
		logger.Logger.Error("Failed to create employee", zap.Error(result.Error))
	}
	return result.Error
}

func (r *EmployeeRepository) GetByID(id uint) (*models.Employee, error) {
	metrics.DatabaseOperationsTotal.WithLabelValues("read", "employees").Inc()
	var employee models.Employee
	result := r.db.First(&employee, id)
	if result.Error != nil {
		logger.Logger.Error("Failed to get employee by ID", zap.Error(result.Error))
	}
	return &employee, result.Error
}

func (r *EmployeeRepository) GetAll() ([]models.Employee, error) {
	metrics.DatabaseOperationsTotal.WithLabelValues("read_all", "employees").Inc()
	var employees []models.Employee
	result := r.db.Find(&employees)
	if result.Error != nil {
		logger.Logger.Error("Failed to get all employees", zap.Error(result.Error))
	}
	return employees, result.Error
}

func (r *EmployeeRepository) Update(employee *models.Employee) error {
	metrics.DatabaseOperationsTotal.WithLabelValues("update", "employees").Inc()
	result := r.db.Save(employee)
	if result.Error != nil {
		logger.Logger.Error("Failed to update employee", zap.Error(result.Error))
	}
	return result.Error
}

func (r *EmployeeRepository) Delete(id uint) error {
	metrics.DatabaseOperationsTotal.WithLabelValues("delete", "employees").Inc()
	result := r.db.Delete(&models.Employee{}, id)
	if result.Error != nil {
		logger.Logger.Error("Failed to delete employee", zap.Error(result.Error))
	}
	return result.Error
}
