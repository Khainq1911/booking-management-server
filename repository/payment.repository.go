package repository

import (
	"manage-system-server/domain"

	"gorm.io/gorm"
)

type paymentStorage struct {
	Sql *gorm.DB
}

func NewPaymentStorage (db *gorm.DB) domain.PaymentStorage{
	return &paymentStorage{
		Sql: db,
	}
}

 