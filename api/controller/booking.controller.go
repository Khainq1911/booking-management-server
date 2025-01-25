package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


type BookingController struct {
	BookingStorage domain.BookingStorage
}

func (controller *BookingController) CreateBooking(ctx echo.Context) error {
	var bookingRequest model.Booking

	if err := ctx.Bind(&bookingRequest); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Invalid request payload",
			Data:       nil,
		})
	}

	if err := controller.BookingStorage.CreateBookingRepo(ctx.Request().Context(), bookingRequest); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create the booking",
			Data:       nil,
		})
	}

	roomStatus := model.RoomStatusUpdate{
		BookingStatus: true,
	}

	if err := controller.BookingStorage.UpdateRoomStatus(ctx.Request().Context(), roomStatus, bookingRequest.RoomId); err != nil {
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
	bookings, err := controller.BookingStorage.ListBookingRepo(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve booking list",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Booking list retrieved successfully",
		Data:       bookings,
	})
}

func (controller *BookingController) GetBooking(ctx echo.Context) error {
	roomIdParam := ctx.Param("room_id")
	roomId, err := strconv.Atoi(roomIdParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid room_id parameter",
			Data:       nil,
		})
	}

	bookingDetails, err := controller.BookingStorage.GetBookingByRoomId(ctx.Request().Context(), roomId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve booking details",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Booking details retrieved successfully",
		Data:       bookingDetails,
	})
}
