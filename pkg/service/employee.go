package service

import (
	employee "smartway-employee-app"
	"smartway-employee-app/pkg/repository"
)

type EmployeeService struct {
	repo repository.Employee
}

func NewEmployeeService(repo repository.Employee) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(employee employee.Employee) (int, error) {
	return s.repo.CreateEmployee(employee)
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	return s.repo.DeleteEmployee(id)
}

func (s *EmployeeService) GetEmployeesByCompanyId(id int, departmentName string) ([]employee.Employee, error) {
	return s.repo.GetEmployeesByCompanyId(id, departmentName)
}

func (s *EmployeeService) UpdateEmployee(updatedEmployee employee.UpdateEmployee) error {
	return s.repo.UpdateEmployee(updatedEmployee)
}
