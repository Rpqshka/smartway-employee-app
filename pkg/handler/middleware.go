package handler

import (
	"errors"
	"regexp"
	employee "smartway-employee-app"
)

func checkEmployee(employee employee.Employee) error {

	err := checkName(employee.Name)
	if err != nil {
		return err
	}

	if err = checkSurname(employee.Surname); err != nil {
		return err
	}

	if err = checkPhone(employee.Phone); err != nil {
		return err
	}

	if err = checkPassport(employee.Passport.Number); err != nil {
		return err
	}

	if employee.CompanyId < 1 {
		return errors.New("incorrect Company ID")
	}

	if err = checkPhone(employee.Department.Phone); err != nil {
		return err
	}

	return nil
}

func checkName(name string) error {
	regex := regexp.MustCompile(`^[A-Z][a-zA-Z'-]*$`)

	if !regex.MatchString(name) {
		return errors.New("name must start with a capital letter and contain only letters, apostrophes, and hyphens")
	}
	return nil
}

func checkSurname(surname string) error {
	regex := regexp.MustCompile(`^[A-Z][a-zA-Z'-]*$`)

	if !regex.MatchString(surname) {
		return errors.New("surname must start with a capital letter and contain only letters, apostrophes, and hyphens")
	}
	return nil
}

func checkPhone(phone string) error {
	if phone[0] == '+' {
		return errors.New("phone number must start with a '8'")
	}
	if len(phone) != 11 {
		return errors.New("phone number must be exactly 11 digits long")
	}
	return nil
}

func checkPassport(number string) error {
	regex := regexp.MustCompile(`^\d{10}$`)
	if !regex.MatchString(number) {
		return errors.New("passport number must contain only 10 digits")
	}
	return nil
}

func checkUpdatedEmployee(updatedEmployee employee.UpdateEmployee) error {
	if updatedEmployee.Name != "" {
		if err := checkName(updatedEmployee.Name); err != nil {
			return err
		}
	}
	if updatedEmployee.Surname != "" {
		if err := checkSurname(updatedEmployee.Surname); err != nil {
			return err
		}
	}
	if updatedEmployee.Phone != "" {
		if err := checkPhone(updatedEmployee.Phone); err != nil {
			return err
		}
	}

	if updatedEmployee.Passport.Number != "" {
		if err := checkPassport(updatedEmployee.Passport.Number); err != nil {
			return err
		}
	}

	if updatedEmployee.CompanyId < 0 {
		return errors.New("incorrect Company ID")
	}

	if updatedEmployee.Department.Phone != "" {
		if err := checkPhone(updatedEmployee.Department.Phone); err != nil {
			return err
		}
	}

	return nil
}
