package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"

	"manage-system-server/security"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EmployeeController struct {
	Repo domain.EmployeeStorage
}

func (repo *EmployeeController) CreateEmployee(ctx echo.Context) error {
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

func (e *EmployeeController) ListEmployee(ctx echo.Context) error {
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

func (e *EmployeeController) UpdateEmployee(ctx echo.Context) error {
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

func(controller *EmployeeController) QueryEmployee(ctx echo.Context) error{
	q := ctx.QueryParam("q")

	data, err := controller.Repo.QueryEmployeeRepo(ctx.Request().Context(), q)
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
