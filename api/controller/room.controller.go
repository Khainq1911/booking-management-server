package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomController struct {
	Rc domain.RoomStorage
}

func (u *RoomController) AddRoom(ctx echo.Context) error {
	req := model.RoomAction{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable to process request payload",
			Data:       nil,
		})
	}

	if err := u.Rc.AddRoomRepo(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error occurred while adding the room",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Room has been successfully created",
		Data:       req,
	})
}

func (u *RoomController) UpdateRoom(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID provided",
			Data:       nil,
		})
	}

	req := model.RoomAction{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable to process request payload",
			Data:       nil,
		})
	}

	if err := u.Rc.UpdateRoomRepo(ctx.Request().Context(), id, req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error occurred while updating the room",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Room updated successfully",
		Data:       nil,
	})
}

func (u *RoomController) ListRoom(ctx echo.Context) error {
	res, err := u.Rc.ListRoomWithType(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error occurred while retrieving the room list",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Room list retrieved successfully",
		Data:       res,
	})
}

func (controller *RoomController) QueryRoom(ctx echo.Context) error {
	query := ctx.QueryParam("q")
	status := ctx.QueryParam("booking_status")

	data, err := controller.Rc.QueryRoomRepo(ctx.Request().Context(), query, status)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error occurred while retrieving the room list",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Room updated successfully",
		Data:       data,
	})
}
