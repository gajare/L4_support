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

// why we this function ?
func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(employee *models.Employee) error {
	metrics.DatabaseOperationsTotal.WithLabelValues("create", "employee").Inc()
	result := r.db.Create(employee)
	if result.Error != nil {
		logger.Logger.Error("Fail to create employee", zap.Error(result.Error))
	}
	return result.Error
}

func (r *EmployeeRepository) GetByID(id uint) (*models.Employee, error) {
	metrics.DatabaseOperationsTotal.WithLabelValues("read", "employee").Inc()

	var employee models.Employee
	result := r.db.First(&employee, id)
	if result.Error != nil {
		logger.Logger.Error("fail to get employee by id", zap.Error(result.Error))
	}
	return &employee, result.Error
}

func (r *EmployeeRepository) GetAll() ([]models.Employee, error) {
	metrics.DatabaseOperationsTotal.WithLabelValues("read_all", "employeess").Inc()

	var employees []models.Employee

	result := r.db.Find(&employees)
	if result.Error != nil {
		logger.Logger.Error("Fail to get all employees", zap.Error(result.Error))
	}
	return employees, result.Error
}

func (r *EmployeeRepository) Update(employee models.Employee) error {
	metrics.DatabaseOperationsTotal.WithLabelValues("update", "employee").Inc()

	result := r.db.Save(employee)
	if result.Error != nil {
		logger.Logger.Error("Fail to update employee", zap.Error(result.Error))
	}
	return result.Error
}

func (r *EmployeeRepository) Delete(id uint) error {
	metrics.DatabaseOperationsTotal.WithLabelValues("delete", "employee").Inc()
	result := r.db.Delete(&models.Employee{}, id)
	if result.Error != nil {
		logger.Logger.Error("Failt to delete employee", zap.Error(result.Error))
	}
	return result.Error
}
