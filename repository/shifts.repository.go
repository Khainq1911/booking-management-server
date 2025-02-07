package repository

import (
	"context"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type ShiftStorage struct {
	Sql *gorm.DB
}

func NewShiftStorage(db *gorm.DB) domain.ShiftStorage {
	return &ShiftStorage{
		Sql: db,
	}
}

func (db *ShiftStorage) AddShiftsRepo(ctx context.Context, payload []model.Shift) error {
	result := db.Sql.Create(&payload)
	return result.Error
}

func (db *ShiftStorage) ListShiftsRepo(ctx context.Context) ([]model.Shift, error) {
	res := []model.Shift{}
	result := db.Sql.Table("shifts").
	Select(`shifts.*, employees.username as employee_name`).
	Joins("join employees on employees.id = shifts.employee_id").
	Find(&res)
	return res, result.Error
}

func (db *ShiftStorage) ListShiftsById(ctx context.Context, employeeId int) ([]model.Shift, error) {
	res := []model.Shift{}
	result := db.Sql.Where("employee_id = ?", employeeId).Find(&res)
	return res, result.Error
}

func (db *ShiftStorage) UpdateShiftsRepo(ctx context.Context, shift_id int, payload model.UpdateShift) error {
	result := db.Sql.Where("id = ?", shift_id).Updates(&payload)
	return result.Error
}
