package handler

import (
	"manage-system-server/modules/employee/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (e *EmployeeRepo) UpdateEmployee(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid employee ID",
			Data:       nil,
		})
	}


	req := model.UpdateEmploy{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
			Data:       nil,
		})
	}

	if err := e.Repo.UpdateEmployRepo(ctx.Request().Context(), id, req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update employee",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Employee updated successfully",
		Data:       nil,
	})
}
