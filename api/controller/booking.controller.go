package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookingController struct {
	BookingStorage domain.BookingStorage
}

func (controller *BookingController) CreateBooking(ctx echo.Context) error {
	req := model.Booking{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable to process request payload",
			Data:       nil,
		})
	}

	if err := controller.BookingStorage.CreateBookingRepo(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create booking",
			Data:       nil,
		})
	}

	roomStatus := model.RoomStatusUpdate{
		BookingStatus: true,
	}

	if err := controller.BookingStorage.UpdateRoomStatus(ctx.Request().Context(), roomStatus, req.RoomId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update room status",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Booking created successfully",
		Data:       nil,
	})
}


func (controller *BookingController) ListBooking(ctx echo.Context) error {
	data, err := controller.BookingStorage.ListBookingRepo(ctx.Request().Context())
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
