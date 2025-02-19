package controller

import (
	"manage-system-server/domain"
	"manage-system-server/email"
	"manage-system-server/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SalaryController struct {
	SalaryStorage   domain.SalaryStorage
	EmployeeStorage domain.EmployeeStorage
}

func (sc SalaryController) Add(ctx echo.Context) error {
	var req model.Salary

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Invalid request payload",
			Data:       nil,
		})
	}

	if err := sc.SalaryStorage.Add(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to add salary record",
			Data:       nil,
		})
	}

	employee, err := sc.EmployeeStorage.GetEmployeeById(ctx.Request().Context(), req.EmployeeId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve employee information",
			Data:       nil,
		})
	}

	message := email.GeneratePayrollMessage(
		employee.Username, req.EmployeeId, req.PayPeriodStart, req.PayPeriodEnd,
		req.BaseSalary, req.Overtime, req.Bonus, req.Deductions,
		req.NetSalary, req.PaymentDate, req.Notes,
	)

	if err := email.SendMail(employee.Email, message); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to send payroll email",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Salary record added and email sent successfully",
		Data:       nil,
	})
}

func (sc SalaryController) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid employee ID",
			Data:       nil,
		})
	}

	var req model.UpdateSalary

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Invalid request payload",
			Data:       nil,
		})
	}

	if err := sc.SalaryStorage.Update(ctx.Request().Context(), id, req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update salary record",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Salary record updated successfully",
		Data:       nil,
	})
}

func (sc *SalaryController) List(ctx echo.Context) error {
	result, err := sc.SalaryStorage.List(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update salary record",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Salary record updated successfully",
		Data:       result,
	})
}
