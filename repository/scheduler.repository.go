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


func (db *SchedulerStorage) ListScheduler(ctx context.Context) ([]model.GetScheduler, error) {
	data := []model.GetScheduler{}
	result := db.Sql.
	Table("shift_assignments").Select("shift_assignments.*, employees.username, employees.phone, shifts.start_time, shifts.end_time").
	Joins("left join shifts on shifts.id = shift_assignments.shift_id left join employees on employees.id = shift_assignments.employee_id").
	Find(&data)
	return data, result.Error
}

func (db *SchedulerStorage) ListSchedulerByEmpId(ctx context.Context, empId int) ([]model.GetScheduler, error) {
	data := []model.GetScheduler{}
	result := db.Sql.Where("employee_id = ?", empId).
	Select("shift_assignments.*, employees.username, employees.phone, shifts.start_time, shifts.end_time").
	Joins("left join shifts on shifts.id = shift_assignments.shift_id left join employees on employees.id = shift_assignments.employee_id").
	Find(&data)

	return data, result.Error
}
