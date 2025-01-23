package domain

import (
	"context"
	"manage-system-server/model"
)

type BookingServiceStorage interface {
	CreateBookingServiceRepo(ctx context.Context, payload []model.BookingService) error
	ListBookingServiceRepo(ctx context.Context, bookingId int) ([]model.BookingService, error)
}
