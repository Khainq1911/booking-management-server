package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentStorage domain.PaymentStorage
}

func (controller *PaymentController) AddPayment(ctx echo.Context) error {
	var payload model.Payment

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Invalid request payload",
			Data:       nil,
		})
	}

	if err := controller.PaymentStorage.AddPaymentRepo(ctx.Request().Context(), payload); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create the payment",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Payment created successfully",
		Data:       nil,
	})
}

func (controller *PaymentController) ListPayment(ctx echo.Context) error {
	data, err := controller.PaymentStorage.ListPayment(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get the payment",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Payment is gotten successfully",
		Data:       data,
	})
}

func (controller *PaymentController) GetPayment(ctx echo.Context) error {
	booking_id, err := strconv.Atoi(ctx.Param("booking_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid payment ID",
			Data:       nil,
		})
	}

	data, err := controller.PaymentStorage.GetPaymentByBookingIdRepo(ctx.Request().Context(), booking_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get the payment",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Payment is gotten successfully",
		Data:       data,
	})
}

func (controller *PaymentController) UpdatePayment(ctx echo.Context) error {
	booking_id, err := strconv.Atoi(ctx.Param("booking_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid payment ID",
			Data:       nil,
		})
	}

	room_id, err := strconv.Atoi(ctx.Param("room_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid room ID",
			Data:       nil,
		})
	}

	if err := controller.PaymentStorage.UpdatePaymentRepo(ctx.Request().Context(), booking_id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update the payment",
			Data:       nil,
		})
	}

	if err := controller.PaymentStorage.UpdateBookingRepo(ctx.Request().Context(), booking_id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update the booking",
			Data:       nil,
		})
	}

	if err := controller.PaymentStorage.UpdateRoomRepo(ctx.Request().Context(), room_id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update the room status",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Payment and associated booking updated successfully",
		Data:       nil,
	})
}
