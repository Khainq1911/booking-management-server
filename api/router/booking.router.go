package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewBookingRouter(db *gorm.DB, echo *echo.Group) {
	dbConn := repository.NewBookingStorage(db)
	controller := controller.BookingController{
		BookingStorage: dbConn,
	}
	echo.GET("/booking", controller.ListBooking)
	echo.GET("/booking/:room_id", controller.GetBooking)
	echo.POST("/booking", controller.CreateBooking)
}
