package service

import (
	employee "smartway-employee-app"
	"smartway-employee-app/pkg/repository"
)

type Employee interface {
	CreateEmployee(employee employee.Employee) (int, error)
	DeleteEmployee(id int) error
	GetEmployeesByCompanyId(id int, departmentName string) ([]employee.Employee, error)
	UpdateEmployee(updatedEmployee employee.UpdateEmployee) error
}

type Service struct {
	Employee
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Employee: NewEmployeeService(repos.Employee),
	}
}
