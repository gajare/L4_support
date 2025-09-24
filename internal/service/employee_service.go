package service

import (
	"L4_support/internal/models"
	"L4_support/internal/repository"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeRepository(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) error {
	return s.repo.Create(employee)
}

func (s *EmployeeService) GetEmployee(id uint) (*models.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *EmployeeService) GetEmployees() ([]models.Employee, error) {
	return s.repo.GetAll()
}

func (s *EmployeeService) UpdateEmployee(employee *models.Employee) error {
	return s.repo.Update(*employee)
}

func (s *EmployeeService) Delete(id uint) error {
	return s.repo.Delete(id)
}
