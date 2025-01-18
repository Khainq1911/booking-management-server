package domain

import (
	"context"
	"manage-system-server/model"
)

type TyperoomStorage interface {
	AddTyperoomRepo(ctx context.Context, createData model.CreateTyperoom) error
	UpdateTyperoomRepo(ctx context.Context, id int, updateData model.CreateTyperoom) (int64, error)
	ListTyperoomRepo(ctx context.Context) ([]model.Typeroom, error)
}
