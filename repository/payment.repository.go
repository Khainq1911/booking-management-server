package repository

import (
	"context"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type paymentStorage struct {
	Sql *gorm.DB
}

func NewPaymentStorage(db *gorm.DB) domain.PaymentStorage {
	return &paymentStorage{
		Sql: db,
	}
}

func (db *paymentStorage) AddPaymentRepo(ctx context.Context, payload model.Payment) error {
	err := db.Sql.Create(&payload).Error
	return err
}

func (db *paymentStorage) GetPaymentByBookingIdRepo(ctx context.Context, bookingId int) (model.Payment, error) {
	data := model.Payment{}
	err := db.Sql.Order("created_at DESC").Where("booking_id = ? and status = ?", bookingId, false).First(&data).Error
	return data, err
}

func (db *paymentStorage) UpdatePaymentRepo(ctx context.Context, booking_id int) error {
	err := db.Sql.Table("payments").Where("booking_id = ? and status = ?", booking_id, false).Update("status", true).Error
	return err
}

func (db *paymentStorage) UpdateBookingRepo(ctx context.Context, id int) error {
	err := db.Sql.Table("bookings").Where("id = ?", id).Update("is_paid", true).Error
	return err
}

func (db *paymentStorage) UpdateRoomRepo(ctx context.Context, id int) error {
	err := db.Sql.Table("rooms").Where("id = ?", id).Update("booking_status", false).Error
	return err
}