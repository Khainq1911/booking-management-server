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

func (db *bookingStorage) GetBookingByRoomId(ctx context.Context, roomId int) (model.BookingWithCustomerAndRoom, error) {
	data := model.BookingWithCustomerAndRoom{}

	err := db.Sql.Table("bookings").
		Select(`
		bookings.*, 
		customers.name AS customer_name, customers.dob AS customer_dob, 
		customers.email AS customer_email, customers.phone AS customer_phone, customers.address AS customer_address,
		rooms.id as room_id, rooms.name AS room_name, rooms.description AS room_description, rooms.floor AS room_floor, rooms.img_url,
		typerooms.name AS typeroom_name, typerooms.description AS typeroom_description, 
		typerooms.hourly_price, typerooms.daily_price, typerooms.max_capacity, typerooms.standard_capacity
		`).
		Joins("JOIN customers ON customers.id = bookings.guest_id").
		Joins("JOIN rooms ON rooms.id = bookings.room_id").
		Joins("LEFT JOIN typerooms ON rooms.type_id = typerooms.id").
		Where("bookings.is_paid = ? AND bookings.room_id = ?", false, roomId).
		First(&data).Error
	if err != nil {
		return model.BookingWithCustomerAndRoom{}, err
	}

	return data, nil
}
