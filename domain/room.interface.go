package domain

import (
	"context"
	"manage-system-server/model"
)

type RoomStorage interface {
	ListRoomWithType(ctx context.Context) ([]model.RoomWithType, error)
	AddRoomRepo(ctx context.Context, payload model.RoomAction) error
	UpdateRoomRepo(ctx context.Context, id int, payload model.RoomAction) (int64, error)
	QueryRoomRepo(ctx context.Context, query string) ([]model.Room, error) 
}
