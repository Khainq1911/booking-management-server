package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewSalaryRouter(db *gorm.DB, echo *echo.Group) {
	dbConn := repository.NewSalaryStorage(db)
	dbConn1 := repository.NewEmployeeStorage(db)
	controller := controller.SalaryController{
		EmployeeStorage: dbConn1,
		SalaryStorage: dbConn,
	}
	echo.POST("/payroll", controller.Add)
	echo.GET("/payroll", controller.List)
	echo.PUT("/payroll/:id", controller.Update)
}
