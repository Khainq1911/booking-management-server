package repository

import (
	"context"
	"manage-system-server/modules/employee/model"
)

func (db *employeeStorage) CreateEmployeeRepo(ctx context.Context, newEmployee model.Employee) error {
	result := db.sql.Select("Username", "Email", "Phone", "Password", "Position", "HiredDate", "Dob", "Salary").Create(&newEmployee)
	return result.Error
}
