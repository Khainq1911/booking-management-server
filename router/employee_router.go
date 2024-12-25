package router

import (
	"manage-system-server/modules/employee/handler"

	"github.com/labstack/echo/v4"
)

type EmployeeApi struct {
	Echo            *echo.Echo
	EmployeeHandler handler.EmployeeRepo
}

func (api *EmployeeApi) SetupRouter() {
	api.Echo.POST("/employee", api.EmployeeHandler.CreateEmployee)
}
