package repository

import (
	"context"
	"fmt"
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

func (db *ShiftStorage) AddShiftsRepo(ctx context.Context, payload model.Shift) error {
	result := db.Sql.Create(&payload)
	fmt.Println(result.Error)
	return result.Error
}

func (db *ShiftStorage) ListShiftsRepo(ctx context.Context) ([]model.Shift, error) {
	res := []model.Shift{}
	result := db.Sql.Table("shifts").Find(&res)
	return res, result.Error
}


func (db *ShiftStorage) UpdateShiftsRepo(ctx context.Context, shift_id int, payload model.UpdateShift) error {
	result := db.Sql.Where("id = ?", shift_id).Updates(&payload)
	return result.Error
}
