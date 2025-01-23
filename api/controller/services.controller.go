package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ServicesController struct {
	ServicesStorage domain.ServicesStorage
}

func (controller *ServicesController) CreateItem(ctx echo.Context) error {
	req := []model.Service{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable to process request payload",
			Data:       nil,
		})
	}

	if err := controller.ServicesStorage.CreateItemRepo(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create service item",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Service item created successfully",
		Data:       nil,
	})
}

func (controller *ServicesController) UpdateItem(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid service ID",
			Data:       nil,
		})
	}
	req := model.UpdateService{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable to process request payload",
			Data:       nil,
		})
	}

	rowsAffected, err := controller.ServicesStorage.UpdateItemRepo(ctx.Request().Context(), req, id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update service item",
			Data:       nil,
		})
	}

	if rowsAffected == 0 {
		return ctx.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "No service item found to update",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Service item updated successfully",
		Data:       nil,
	})
}

func (controller *ServicesController) ListItem(ctx echo.Context) error {
	res, err := controller.ServicesStorage.ListItemRepo(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve service items",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Service items retrieved successfully",
		Data:       res,
	})
}
