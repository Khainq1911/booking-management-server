package domain

import (
	"context"
	"manage-system-server/model"
)

type SalaryStorage interface {
	Add(ctx context.Context, payload model.Salary) error
	Update(ctx context.Context, id int, payload model.UpdateSalary) error
	List(ctx context.Context) ([]model.ListSalary, error)
}
