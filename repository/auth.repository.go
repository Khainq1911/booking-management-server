package repository

import (
	"context"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type authStorage struct {
	Db *gorm.DB
}

func NewAuthStorage(db *gorm.DB) domain.AuthStorage {
	return &authStorage{
		Db: db,
	}
}

func (db *authStorage) CheckLogin(ctx context.Context, username string) (model.Employee, error) {
	employee := model.Employee{}
	result := db.Db.Where("email = ? or phone = ?", username, username).First(&employee)
	if result.Error != nil {
		return employee, result.Error
	}
	return employee, nil
}
