package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookingServiceController struct {
	BookingServiceStorage domain.BookingServiceStorage
}

func (controller *BookingServiceController) CreateBookingService(ctx echo.Context) error {
	req := []model.BookingService{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable to process request payload",
			Data:       nil,
		})
	}

	if err := controller.BookingServiceStorage.CreateBookingServiceRepo(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create booking service",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Booking service created successfully",
		Data:       nil, 
	})
}


func (controller *BookingServiceController) ListBookingService(ctx echo.Context) error {
	bookingIDParam := ctx.Param("booking_id")
	bookingID, err := strconv.Atoi(bookingIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid booking_id parameter",
			Data:       nil,
		})
	}

	
	data, err := controller.BookingServiceStorage.ListBookingServiceRepo(ctx.Request().Context(), bookingID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve booking services",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Booking services retrieved successfully",
		Data:       data,
	})
}
