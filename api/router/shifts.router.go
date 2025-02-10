package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewShiftsRouter(db *gorm.DB, echo *echo.Group) {
	dbConn := repository.NewShiftStorage(db)
	controller := controller.ShiftsController{
		ShiftStorage: dbConn,
	}
	echo.POST("/shifts", controller.AddShifts)
	echo.GET("/shifts/list", controller.ListShifts)
	echo.PUT("/shifts/:id", controller.UpdateShift)
}
