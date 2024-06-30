package handler

import (
	"github.com/gin-gonic/gin"
	"smartway-employee-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	employee := router.Group("/employee")
	{
		employee.POST("", h.createEmployee)
		employee.DELETE("", h.deleteEmployee)
		employee.PUT("", h.updateEmployee)
		employee.GET("/company/:id", h.getEmployeesByCompanyId)

	}

	return router
}
