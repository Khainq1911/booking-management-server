package controller

import (
	"manage-system-server/domain"
	"manage-system-server/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerController struct {
	CustomerStorage domain.CustomerStorage
}

func (controller *CustomerController) CreateCustomer(ctx echo.Context) error {
	var customer model.Customer

	if err := ctx.Bind(&customer); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable to process request payload",
			Data:       nil,
		})
	}

	if err := controller.CustomerStorage.CreateCustomerRepo(ctx.Request().Context(), customer); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to create customer",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Customer created successfully",
		Data:       customer,
	})
}

func (controller *CustomerController) ListCustomer(ctx echo.Context) error {
	customers, err := controller.CustomerStorage.ListCustomerRepo(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve customer list",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Customer list retrieved successfully",
		Data:       customers,
	})
}

func (controller *CustomerController) GetCustomerByPhone(ctx echo.Context) error {
	q := ctx.QueryParam("q")

	result, err := controller.CustomerStorage.FindCustomerByPhoneRepo(ctx.Request().Context(), q)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"exists": false,
		})
	}

	if result.Id == 0 {
		return ctx.JSON(http.StatusOK, echo.Map{
			"exists": false,
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"exists":      true,
		"customer_id": result.Id,
		"name":        result.Name,
		"email":       result.Email,
		"phone":       result.Phone,
	})
}
