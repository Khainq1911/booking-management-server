package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetUp(db *gorm.DB, e *echo.Echo) {
	publicRoutes := e.Group("")
	NewEmployeeRouter(db, publicRoutes)
	NewAuthRouter(db, publicRoutes)
	NewTyperoomRouter(db, publicRoutes)
	NewRoomRouter(db, publicRoutes)
	NewCustomerRouter(db, publicRoutes)
	NewBookingRouter(db, publicRoutes)
	NewServicesRouter(db, publicRoutes)
	NewPaymentRouter(db, publicRoutes)
	NewShiftsRouter(db, publicRoutes)
	NewSchedulerRouter(db, publicRoutes)
	NewSalaryRouter(db, publicRoutes)
}
