package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewServicesRouter(db *gorm.DB, echo *echo.Group) {
	dbConn := repository.NewServicesStorage(db)
	controller := controller.ServicesController{
		ServicesStorage: dbConn,
	}

	echo.POST("/service", controller.CreateItem)
	echo.PUT("/service/:id", controller.UpdateItem)
	echo.GET("/service", controller.ListItem)
}
