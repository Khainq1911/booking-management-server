package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ShiftsController struct {
	ShiftStorage domain.ShiftStorage
}

func (controller *ShiftsController) AddShifts(ctx echo.Context) error {
	var req []model.Shift

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Invalid request payload",
			Data:       nil,
		})
	}

	if err := controller.ShiftStorage.AddShiftsRepo(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to add shifts",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Shifts added successfully",
		Data:       nil,
	})
}

func (controller *ShiftsController) ListShifts(ctx echo.Context) error {
	res, err := controller.ShiftStorage.ListShiftsRepo(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve shifts",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Shifts retrieved successfully",
		Data:       res,
	})
}

func (controller *ShiftsController) ListShiftsById(ctx echo.Context) error {
	employeeId, err := strconv.Atoi(ctx.QueryParam("employee_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid employee ID provided",
			Data:       nil,
		})
	}

	res, err := controller.ShiftStorage.ListShiftsById(ctx.Request().Context(), employeeId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve shifts for the specified employee",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Shifts retrieved successfully",
		Data:       res,
	})
}

func (controller *ShiftsController) UpdateShift(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid employee ID provided",
			Data:       nil,
		})
	}

	var req model.UpdateShift
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Invalid request payload",
			Data:       nil,
		})
	}

	if err := controller.ShiftStorage.UpdateShiftsRepo(ctx.Request().Context(), id, req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update shift",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Shift updated successfully",
		Data:       nil,
	})
}