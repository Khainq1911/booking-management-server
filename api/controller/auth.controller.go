package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"
	"manage-system-server/security"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Ac domain.AuthStorage
}

func (a *AuthController) Login(ctx echo.Context) error {
	req := model.Login{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
			Data:       nil,
		})
	}

	empData, err := a.Ac.CheckLogin(ctx.Request().Context(), req.Username)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid username or password",
			Data:       nil,
		})
	}

	if !security.CheckPasswordHash(req.Password, empData.Password) {
		return ctx.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid username or password",
			Data:       nil,
		})
	}

	accessToken, err := security.GenerateToken(empData.Username, empData.Role, empData.Id)
	if err != nil {
		ctx.Logger().Errorf("Token generation failed: %v", err)
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"accessToken": accessToken,
	})
}
