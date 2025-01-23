package repository

import (
	"context"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type bookingStorage struct {
	Sql *gorm.DB
}

func NewBookingStorage(db *gorm.DB) domain.BookingStorage {
	return &bookingStorage{
		Sql: db,
	}
}

func (db *bookingStorage) CreateBookingRepo(ctx context.Context, payload model.Booking) error {
	result := db.Sql.Create(&payload)
	return result.Error
}

func (db *bookingStorage) ListBookingRepo(ctx context.Context) ([]model.Booking, error) {
	data := []model.Booking{}
	result := db.Sql.Find(&data)
	return data, result.Error
}

func (db *bookingStorage) UpdateRoomStatus(ctx context.Context, payload model.RoomStatusUpdate, roomId int) error {
	result := db.Sql.Where("id = ?", roomId).Updates(&payload)
	return result.Error
}