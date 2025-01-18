package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TyperoomController struct {
	TrC domain.TyperoomStorage
}

func (u *TyperoomController) AddTyperoom(ctx echo.Context) error {
	req := model.CreateTyperoom{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable to process request payload",
			Data:       nil,
		})
	}

	if err := u.TrC.AddTyperoomRepo(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error occurred while adding the typeroom",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Typeroom has been successfully created",
		Data:       req,
	})
}

func (u *TyperoomController) UpdateTyperoom(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID provided",
			Data:       nil,
		})
	}

	req := model.CreateTyperoom{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable to process request payload",
			Data:       nil,
		})
	}

	rowsAffected, err := u.TrC.UpdateTyperoomRepo(ctx.Request().Context(), id, req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error occurred while updating the typeroom",
			Data:       nil,
		})
	}

	if rowsAffected == 0 {
		return ctx.JSON(http.StatusNotFound, model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "No typeroom found to update",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Typeroom updated successfully",
		Data:       nil,
	})
}

func (u *TyperoomController) ListTyperoom(ctx echo.Context) error {
	res, err := u.TrC.ListTyperoomRepo(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error occurred while retrieving the typeroom list",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Typeroom list retrieved successfully",
		Data:       res,
	})
}
