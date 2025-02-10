package domain

import (
	"context"
	"manage-system-server/model"
)

type ShiftStorage interface {
	AddShiftsRepo(ctx context.Context, payload model.Shift) error
	ListShiftsRepo(ctx context.Context) ([]model.Shift, error)
	UpdateShiftsRepo(ctx context.Context, shift_id int, payload model.UpdateShift) error
}
