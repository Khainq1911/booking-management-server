package domain

import (
	"context"
	"manage-system-server/model"
)

type SchedulerStorage interface {
	AddScheduler(ctx context.Context, payload []model.ShiftAssignment) error
	UpdateScheduler(ctx context.Context, id int, payload model.UpdateShiftAssignment) error
	ListScheduler(ctx context.Context) ([]model.ShiftAssignment, error)
	ListSchedulerByEmpId(ctx context.Context, empId int) ([]model.ShiftAssignment, error)
}
