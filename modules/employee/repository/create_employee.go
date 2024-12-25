package repository

import (
	"context"
	"manage-system-server/modules/employee/model"

	"gorm.io/gorm"
)

type EmployeeStorage interface {
	CreateEmployeeRepo(ctx context.Context, newEmployee model.Employee) error
}

type postgresqlStorage struct {
	db *gorm.DB
}

func NewMySQLStorage(db *gorm.DB) EmployeeStorage {
	return &postgresqlStorage{db: db}
}

func (db *postgresqlStorage) CreateEmployeeRepo(ctx context.Context, newEmployee model.Employee) error {
	result := db.db.Create(&newEmployee)
	return result.Error
}
