package handler

import (
	"manage-system-server/modules/employee/model"
	"manage-system-server/modules/employee/repository"
	"manage-system-server/security"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EmployeeRepo struct {
	Repo repository.EmployeeStorage
}

func (repo *EmployeeRepo) CreateEmployee(ctx echo.Context) error {
	req := model.Employee{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
			Data:       nil,
		})
	}

	hashPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to hash password",
			Data:       nil,
		})
	}
	req.Password = hashPassword

	if err := repo.Repo.CreateEmployeeRepo(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create employee",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Employee created successfully",
		Data:       req,
	})
}
