package handler

import (
	"manage-system-server/modules/employee/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (e *EmployeeRepo) ListEmployee(ctx echo.Context) error {
	data, err := e.Repo.ListEmployeeRepo(ctx.Request().Context())
	if err != nil {

		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve employee list",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Employee list retrieved successfully",
		Data:       data,
	})
}
