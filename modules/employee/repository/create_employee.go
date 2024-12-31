package repository

import (
	"context"
	"manage-system-server/modules/employee/model"

	"gorm.io/gorm"
)

type EmployeeStorage interface {
	CreateEmployeeRepo(ctx context.Context, newEmployee model.Employee) error
	ListEmployeeRepo(ctx context.Context) ([]model.Employee, error)
}

type employeeStorage struct {
	sql *gorm.DB
}

func NewMySQLStorage(db *gorm.DB) EmployeeStorage {
	return &employeeStorage{sql: db}
}

func (db *employeeStorage) CreateEmployeeRepo(ctx context.Context, newEmployee model.Employee) error {
	result := db.sql.Select("Username", "Email", "Phone", "Password", "Position", "HiredDate", "Dob", "Salary").Create(&newEmployee)
	return result.Error
}
