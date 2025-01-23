package domain

import (
	"context"
	"manage-system-server/model"
)

type CustomerStorage interface {
	CreateCustomerRepo(ctx context.Context, payload model.Customer) error
	ListCustomerRepo(ctx context.Context) ([]model.Customer, error)
	FindCustomerByPhoneRepo(ctx context.Context, phone string) (model.Customer, error)
}
