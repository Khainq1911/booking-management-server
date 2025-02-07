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
		Select("rooms.*, typerooms.name AS typeroom_name, typerooms.description AS typeroom_description, typerooms.hourly_price, typerooms.daily_price, typerooms.max_capacity, typerooms.standard_capacity").
		Joins("LEFT JOIN typerooms ON rooms.type_id = typerooms.id").
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

func (db *roomStorage) QueryRoomRepo(ctx context.Context, query string) ([]model.Room, error) {
	var rooms []model.Room
	result := db.Sql.WithContext(ctx).
		Where("LOWER(name) LIKE LOWER(?)", "%"+query+"%").
		Find(&rooms)

	if result.Error != nil {
		return nil, result.Error
	}
	return rooms, nil
}
