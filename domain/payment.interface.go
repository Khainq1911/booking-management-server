package domain

import (
	"context"
	"manage-system-server/model"
)

type PaymentStorage interface {
	AddPaymentRepo(ctx context.Context, payload model.Payment) error
	UpdatePaymentRepo(ctx context.Context, booking_id int) error
	UpdateBookingRepo(ctx context.Context, id int) error
	GetPaymentByBookingIdRepo(ctx context.Context, bookingId int) (model.Payment, error)
	UpdateRoomRepo(ctx context.Context, id int) error
}
