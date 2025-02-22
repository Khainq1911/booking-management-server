package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewEmployeeRouter(db *gorm.DB, echo *echo.Group) {
	dbConn := repository.NewEmployeeStorage(db)
	ec := &controller.EmployeeController{
		Repo: dbConn,
	}
	echo.GET("/employee", ec.ListEmployee)
	echo.GET("/employee/search", ec.QueryEmployee)
	echo.POST("/employee", ec.CreateEmployee)
	echo.PUT("/employee/:id", ec.UpdateEmployee)
}
