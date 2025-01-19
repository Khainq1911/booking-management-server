package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoomRouter(db *gorm.DB, echo *echo.Group) {
	dbConn := repository.NewRoomStorage(db)
	rc := &controller.RoomController{
		Rc: dbConn,
	}
	echo.GET("/room", rc.ListRoom)
	echo.POST("/room", rc.AddRoom)
	echo.PUT("/room/:id", rc.UpdateRoom)
}
