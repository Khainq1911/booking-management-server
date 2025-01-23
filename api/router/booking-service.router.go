package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewBookingServiceRouter(db *gorm.DB, echo *echo.Group) {
	dbConn := repository.NewBookingServiceStorage(db)
	controller := controller.BookingServiceController{
		BookingServiceStorage: dbConn,
	}
	echo.GET("/booking-service/:booking_id", controller.ListBookingService)
	echo.POST("/booking-service", controller.CreateBookingService)
}
