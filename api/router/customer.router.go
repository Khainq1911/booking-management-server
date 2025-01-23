package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewCustomerRouter(db *gorm.DB, echo *echo.Group) {
	dbConn := repository.NewCustomerStorage(db)
	CustomerController := controller.CustomerController{
		CustomerStorage: dbConn,
	}
	echo.GET("/customer", CustomerController.ListCustomer)
	echo.POST("/customer", CustomerController.CreateCustomer)
	echo.GET("/customer/search", CustomerController.GetCustomerByPhone)
}
