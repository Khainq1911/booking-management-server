package repository

import (
	"context"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type SchedulerStorage struct {
	Sql *gorm.DB
}

func NewSchedulerStorage(db *gorm.DB) domain.SchedulerStorage {
	return &SchedulerStorage{
		Sql: db,
	}
}

func (db *SchedulerStorage) AddScheduler(ctx context.Context, payload []model.ShiftAssignment) error {
	result := db.Sql.Create(&payload)
	return result.Error
}

func (db *SchedulerStorage) UpdateScheduler(ctx context.Context, id int, payload model.UpdateShiftAssignment) error {
	result := db.Sql.Where("id = ?", id).Updates(&payload)
	return result.Error
}

func (db *SchedulerStorage) ListScheduler(ctx context.Context) ([]model.ShiftAssignment, error) {
	data := []model.ShiftAssignment{}
	result := db.Sql.Find(&data)
	return data, result.Error
}

func (db *SchedulerStorage) ListSchedulerByEmpId(ctx context.Context, empId int) ([]model.ShiftAssignment, error) {
	data := []model.ShiftAssignment{}
	result := db.Sql.Where("employee_id = ?", empId).Find(&data)
	return data, result.Error
}
