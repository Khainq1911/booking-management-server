package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SchedulerController struct {
	SchedulerStorage domain.SchedulerStorage
}

func (controller *SchedulerController) AddScheduler(ctx echo.Context) error {
	req := []model.ShiftAssignment{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,

			Message: "Invalid request format. Please check your input data.",

			Data: nil,
		})
	}

	if err := controller.SchedulerStorage.AddScheduler(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,

			Message: "Failed to add scheduler. Please try again later.",

			Data: nil,
		})
	}
	return ctx.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,

		Message: "Scheduler created successfully.",

		Data: nil,
	})
}

func (controller *SchedulerController) UpdateScheduler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,

			Message: "Invalid scheduler ID. Please provide a valid ID.",

			Data: nil,
		})
	}

	req := model.UpdateShiftAssignment{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,

			Message: "Invalid request format. Please check your input data.",

			Data: nil,
		})
	}

	if err := controller.SchedulerStorage.UpdateScheduler(ctx.Request().Context(), id, req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,

			Message: "Failed to update scheduler. Please try again later.",

			Data: nil,
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,

		Message: "Scheduler updated successfully.",

		Data: nil,
	})
}

func (controller *SchedulerController) ListScheduler(ctx echo.Context) error {
	data, err := controller.SchedulerStorage.ListScheduler(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,

			Message: "Failed to retrieve scheduler list. Please try again later.",

			Data: nil,
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,

		Message: "Scheduler list retrieved successfully.",

		Data: data,
	})
}

func (controller *SchedulerController) ListSchedulerByEmpId(ctx echo.Context) error {
	EmployeeId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,

			Message: "Invalid employee ID. Please provide a valid ID.",

			Data: nil,
		})
	}

	data, err := controller.SchedulerStorage.ListSchedulerByEmpId(ctx.Request().Context(), EmployeeId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,

			Message: "Failed to retrieve scheduler list for the employee. Please try again later.",

			Data: nil,
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,

		Message: "Scheduler list for the employee retrieved successfully.",

		Data: data,
	})
}
