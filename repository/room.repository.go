package repository

import (
	"context"
	"fmt"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type roomStorage struct {
	Sql *gorm.DB
}

func NewRoomStorage(db *gorm.DB) domain.RoomStorage {
	return &roomStorage{
		Sql: db,
	}
}


func (db *roomStorage) ListRoomWithType(ctx context.Context) ([]model.RoomWithType, error) {
	data := []model.RoomWithType{}

	result := db.Sql.
		Table("rooms").
		Select("rooms.*, room_types.name AS typeroom_name, room_types.description AS typeroom_description, room_types.hourly_price, room_types.daily_price, room_types.max_capacity, room_types.standard_capacity").
		Joins("LEFT JOIN room_types ON rooms.type_id = room_types.id").
		Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no rooms found with their type information")
	}

	return data, nil
}

func (db *roomStorage) AddRoomRepo(ctx context.Context, payload model.RoomAction) error {
	if err := db.Sql.Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

func (db *roomStorage) UpdateRoomRepo(ctx context.Context, id int, payload model.RoomAction) (int64, error) {
	result := db.Sql.Where("id = ?", id).Updates(&payload)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
