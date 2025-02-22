package repository

import (
	"context"
	"fmt"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type employeeStorage struct {
	sql *gorm.DB
}

func NewEmployeeStorage(db *gorm.DB) domain.EmployeeStorage {
	return &employeeStorage{sql: db}
}

func (db *employeeStorage) CreateEmployeeRepo(ctx context.Context, newEmployee model.Employee) error {
	result := db.sql.Create(&newEmployee)
	return result.Error
}

func (db *employeeStorage) ListEmployeeRepo(ctx context.Context) ([]model.Employee, error) {
	employees := []model.Employee{}
	retult := db.sql.Find(&employees)
	return employees, retult.Error
}

func (db *employeeStorage) UpdateEmployRepo(ctx context.Context, id int, updateData model.UpdateEmploy) error {

	result := db.sql.Model(&model.Employee{}).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows updated, employee with id %d not found", id)
	}
	return nil
}

func (db *employeeStorage) QueryEmployeeRepo(ctx context.Context, query string) ([]model.Employee, error) {
    var employee []model.Employee
    result := db.sql.
        Where("username ILIKE ?", "%"+query+"%").
        Or("phone ILIKE ?", "%"+query+"%").
        Find(&employee)
    return employee, result.Error
}

func (db *employeeStorage) GetEmployeeById(ctx context.Context, id int) (model.Employee, error) {
	employee := model.Employee{}
	retult := db.sql.Where("id=?", id).First(&employee)
	return employee, retult.Error
}
