package domain

import (
	"context"
	"manage-system-server/model"
)

type EmployeeStorage interface {
	CreateEmployeeRepo(ctx context.Context, newEmployee model.Employee) error
	ListEmployeeRepo(ctx context.Context) ([]model.Employee, error)
	GetEmployeeById(ctx context.Context, id int) (model.Employee, error)
	UpdateEmployRepo(ctx context.Context, id int, updateData model.UpdateEmploy) error
}
