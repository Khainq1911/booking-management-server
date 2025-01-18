package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewTyperoomRouter(db *gorm.DB, echo *echo.Group) {
	dbConn := repository.NewTyperoomStorage(db)
	trc := controller.TyperoomController{
		TrC: dbConn,
	}

	echo.GET("/typeroom", trc.ListTyperoom)
	echo.POST("/typeroom", trc.AddTyperoom)
	echo.PUT("/typeroom/:id", trc.UpdateTyperoom)
}
