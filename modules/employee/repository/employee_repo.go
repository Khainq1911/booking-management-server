package repository

import (
	"context"
	"manage-system-server/modules/employee/model"

	"gorm.io/gorm"
)

type EmployeeStorage interface {
	CreateEmployeeRepo(ctx context.Context, newEmployee model.Employee) error
	ListEmployeeRepo(ctx context.Context) ([]model.Employee, error)
	UpdateEmployRepo(ctx context.Context, id int, updateData model.UpdateEmploy) error
}

type employeeStorage struct {
	sql *gorm.DB
}

func NewMySQLStorage(db *gorm.DB) EmployeeStorage {
	return &employeeStorage{sql: db}
}
