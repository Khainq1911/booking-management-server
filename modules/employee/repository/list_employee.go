package repository

import (
	"context"
	"manage-system-server/modules/employee/model"
)

func (db *employeeStorage) ListEmployeeRepo(ctx context.Context) ([]model.Employee, error) {
	employees := []model.Employee{}
	retult := db.sql.Find(&employees)
	return employees, retult.Error
}
