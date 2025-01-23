package repository

import (
	"context"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type BookingServiceStorage struct {
	Sql *gorm.DB
}

func NewBookingServiceStorage(db *gorm.DB) domain.BookingServiceStorage {
	return &BookingServiceStorage{
		Sql: db,
	}
}

func (db *BookingServiceStorage) CreateBookingServiceRepo(ctx context.Context, payload []model.BookingService) error {
	result := db.Sql.Create(&payload)
	return result.Error
}

func (db *BookingServiceStorage) ListBookingServiceRepo(ctx context.Context, bookingId int) ([]model.BookingService, error) {
	data := []model.BookingService{}
	result := db.Sql.Where("booking_id = ?", bookingId).Find(&data)
	return data, result.Error
}
