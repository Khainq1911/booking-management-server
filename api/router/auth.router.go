package router

import (
	"manage-system-server/api/controller"
	"manage-system-server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewAuthRouter(db *gorm.DB, group *echo.Group) {
	dbConn := repository.NewAuthStorage(db)
	ac := controller.AuthController{
		Ac: dbConn,
	}
	group.POST("/auth/login", ac.Login)
}
