package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	employee "smartway-employee-app"
	"strconv"
)

func (h *Handler) createEmployee(c *gin.Context) {
	var input employee.Employee
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := checkEmployee(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Employee.CreateEmployee(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type deleteEmployeeInput struct {
	Id int `json:"id" binding:"required"`
}

func (h *Handler) deleteEmployee(c *gin.Context) {
	var input deleteEmployeeInput
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Employee.DeleteEmployee(input.Id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": input.Id,
	})
}

type getAllEmployeesInput struct {
	DepartmentName string `json:"departmentName"`
}

type getAllEmployeesResponse struct {
	Employees []employee.Employee `json:"employees"`
}

func (h *Handler) getEmployeesByCompanyId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input getAllEmployeesInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	employees, err := h.services.GetEmployeesByCompanyId(id, input.DepartmentName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllEmployeesResponse{
		Employees: employees,
	})
}

func (h *Handler) updateEmployee(c *gin.Context) {
	var input employee.UpdateEmployee

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = checkUpdatedEmployee(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.UpdateEmployee(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": input.Id,
	})
}
