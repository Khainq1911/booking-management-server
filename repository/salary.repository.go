package repository

import (
	"context"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type SalaryStorage struct {
	Sql *gorm.DB
}

func NewSalaryStorage(db *gorm.DB) domain.SalaryStorage {
	return &SalaryStorage{
		Sql: db,
	}
}

func (db *SalaryStorage) Add(ctx context.Context, payload model.Salary) error {
	result := db.Sql.Create(&payload)
	return result.Error
}

func (db *SalaryStorage) List(ctx context.Context) ([]model.ListSalary, error) {
	data := []model.ListSalary{}
	result := db.Sql.Table("payrolls").
	Select("payrolls.*, employees.username as employee_name").
	Joins("left join employees on employees.id = payrolls.employee_id").Find(&data)
	return data, result.Error
}

func (db *SalaryStorage) Update(ctx context.Context, id int, payload model.UpdateSalary) error {
	result := db.Sql.Where("id = ?", id).Updates(&payload)
	return result.Error
}
