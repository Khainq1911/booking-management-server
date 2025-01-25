package domain

import (
	"context"
	"manage-system-server/model"
)

type BookingStorage interface {
	CreateBookingRepo(ctx context.Context, payload model.Booking) error
	ListBookingRepo(ctx context.Context) ([]model.Booking, error)
	UpdateRoomStatus(ctx context.Context, payload model.RoomStatusUpdate, roomId int) error
	GetBookingByRoomId(ctx context.Context, roomId int) (model.BookingWithCustomerAndRoom, error)
}
