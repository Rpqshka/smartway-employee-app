package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	employee "smartway-employee-app"
	"strings"
)

type EmployeePostgres struct {
	db *sqlx.DB
}

func NewEmployeePostgres(db *sqlx.DB) *EmployeePostgres {
	return &EmployeePostgres{db: db}
}

func (r *EmployeePostgres) CreateEmployee(employee employee.Employee) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	selectQuery := fmt.Sprintf(`SELECT id FROM %s WHERE passport_number = $1 FOR UPDATE`, employeesTable)
	err = tx.QueryRow(selectQuery, employee.Passport.Number).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return 0, err
	}
	if id != 0 {
		tx.Rollback()
		return 0, fmt.Errorf("employee with passport number %s already exists with id %d", employee.Passport.Number, id)
	}

	insertQuery := fmt.Sprintf(`
		INSERT INTO %s 
		(name, surname, employee_phone, company_id, passport_type, passport_number, department_name, department_phone)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`, employeesTable)

	row := tx.QueryRow(insertQuery, employee.Name, employee.Surname, employee.Phone, employee.CompanyId,
		employee.Passport.Type, employee.Passport.Number, employee.Department.Name, employee.Department.Phone)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}

func (r *EmployeePostgres) DeleteEmployee(id int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", employeesTable)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(deleteQuery, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if int(rowsAffected) == 0 {
		return errors.New("incorrect employee id")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *EmployeePostgres) GetEmployeesByCompanyId(id int, departmentName string) ([]employee.Employee, error) {
	var employees []employee.Employee
	var selectQuery string
	var args []interface{}

	selectQuery = fmt.Sprintf(`
		SELECT id, name, surname, employee_phone, company_id,
		       passport_type, passport_number,
		       department_name, department_phone
		FROM %s
		WHERE company_id = $1`, employeesTable)

	args = append(args, id)

	if departmentName != "" {
		selectQuery += ` AND department_name = $2`
		args = append(args, departmentName)
	}

	rows, err := r.db.Query(selectQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var emp employee.Employee
		err := rows.Scan(&emp.Id, &emp.Name, &emp.Surname, &emp.Phone, &emp.CompanyId,
			&emp.Passport.Type, &emp.Passport.Number, &emp.Department.Name, &emp.Department.Phone)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	if len(employees) == 0 {
		return nil, errors.New(fmt.Sprintf("no employees found for company_id %d with department_name %s",
			id, departmentName))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *EmployeePostgres) UpdateEmployee(updatedEmployee employee.UpdateEmployee) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if updatedEmployee.Passport.Number != "" {
		var id int

		selectQuery := fmt.Sprintf(`SELECT id FROM %s WHERE passport_number = $1 FOR UPDATE`, employeesTable)
		err = tx.QueryRow(selectQuery, updatedEmployee.Passport.Number).Scan(&id)
		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return err
		}
		if id != 0 {
			tx.Rollback()
			return fmt.Errorf("employee with passport number %s already exists with id %d", updatedEmployee.Passport.Number, id)
		}
	}

	setFields := map[string]interface{}{
		"name":             updatedEmployee.Name,
		"surname":          updatedEmployee.Surname,
		"employee_phone":   updatedEmployee.Phone,
		"company_id":       updatedEmployee.CompanyId,
		"passport_type":    updatedEmployee.Passport.Type,
		"passport_number":  updatedEmployee.Passport.Number,
		"department_name":  updatedEmployee.Department.Name,
		"department_phone": updatedEmployee.Department.Phone,
	}

	var setValues []string
	args := []interface{}{}
	args = append(args, updatedEmployee.Id)
	argID := 2

	for key, value := range setFields {
		if value == "" || (key == "company_id" && value.(int) == 0) {
			continue
		}
		setValues = append(setValues, fmt.Sprintf("%s=$%d", key, argID))
		args = append(args, value)
		argID++
	}

	if len(setValues) == 0 {
		return errors.New("0 columns updated")
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`
		UPDATE %s
		SET %s
		WHERE id = $1`, employeesTable, setQuery)

	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
