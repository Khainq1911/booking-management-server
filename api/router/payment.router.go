package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewPaymentRouter(db *gorm.DB, echo *echo.Group) {
	dbConn := repository.NewPaymentStorage(db)
	PaymentController := controller.PaymentController{
		PaymentStorage: dbConn,
	}
	echo.POST("payment", PaymentController.AddPayment)
	echo.PUT("payment/:booking_id/:room_id", PaymentController.UpdatePayment)
	echo.GET("payment/:booking_id", PaymentController.GetPayment)
}
