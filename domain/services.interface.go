package domain

import (
	"context"
	"manage-system-server/model"
)

type ServicesStorage interface {
	CreateItemRepo(ctx context.Context, payload []model.Service) error
	UpdateItemRepo(ctx context.Context, payload model.UpdateService, service_id int) (int64, error)
	ListItemRepo(ctx context.Context) ([]model.Service, error)
}
