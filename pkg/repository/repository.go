package repository

import (
	"github.com/jmoiron/sqlx"
	employee "smartway-employee-app"
)

type Employee interface {
	CreateEmployee(employee employee.Employee) (int, error)
	DeleteEmployee(id int) error
	GetEmployeesByCompanyId(id int, departmentName string) ([]employee.Employee, error)
	UpdateEmployee(updatedEmployee employee.UpdateEmployee) error
}

type Repository struct {
	Employee
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Employee: NewEmployeePostgres(db),
	}
}
