package domain

import (
	"context"
	"manage-system-server/model"
)

type EmployeeStorage interface {
	GetEmployeeById(ctx context.Context, id int) (model.Employee, error)
	CreateEmployeeRepo(ctx context.Context, newEmployee model.Employee) error
	ListEmployeeRepo(ctx context.Context) ([]model.Employee, error)
	QueryEmployeeRepo(ctx context.Context, query string) ([]model.Employee, error)
	UpdateEmployRepo(ctx context.Context, id int, updateData model.UpdateEmploy) error
}
