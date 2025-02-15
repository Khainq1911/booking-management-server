package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewSchedulerRouter(db *gorm.DB, echo *echo.Group) {
	dbConn := repository.NewSchedulerStorage(db)
	controller := controller.SchedulerController{
		SchedulerStorage: dbConn,
	}
	echo.POST("/scheduler", controller.AddScheduler)
	echo.PUT("/scheduler/update/:id", controller.UpdateScheduler)
	echo.GET("/scheduler", controller.ListScheduler)
	echo.GET("/scheduler/:id", controller.ListSchedulerByEmpId)
}
