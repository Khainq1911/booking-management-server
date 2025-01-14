package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetUp(db *gorm.DB, e *echo.Echo) {
	publicRoutes := e.Group("")
	NewEmployeeRouter(db, publicRoutes)
	NewAuthRouter(db, publicRoutes)
}
