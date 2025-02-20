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

func (db *roomStorage) UpdateRoomRepo(ctx context.Context, id int, payload model.RoomAction) error {
	result := db.Sql.Where("id = ?", id).Updates(&payload)
	return result.Error
}

func (db *roomStorage) QueryRoomRepo(ctx context.Context, query string, roomStatus string) ([]model.RoomWithType, error) {
	var rooms []model.RoomWithType
	result := db.Sql.WithContext(ctx).Table("rooms").
		Select("rooms.*, typerooms.name AS typeroom_name, typerooms.description AS typeroom_description, typerooms.hourly_price, typerooms.daily_price, typerooms.max_capacity, typerooms.standard_capacity").
		Joins("LEFT JOIN typerooms ON rooms.type_id = typerooms.id").
		Where("LOWER(rooms.name) LIKE LOWER(?)", "%"+query+"%")

	if roomStatus != "" {
		result = result.Where("rooms.booking_status = ?", roomStatus)
	}

	if err := result.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}
