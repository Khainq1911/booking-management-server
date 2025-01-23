package repository

import (
	"context"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type ServicesStorage struct {
	Sql *gorm.DB
}

func NewServicesStorage(db *gorm.DB) domain.ServicesStorage {
	return &ServicesStorage{
		Sql: db,
	}
}

func (db *ServicesStorage) CreateItemRepo(ctx context.Context, payload []model.Service) error {
	err := db.Sql.Create(&payload).Error
	return err
}

func (db *ServicesStorage) UpdateItemRepo(ctx context.Context, payload model.UpdateService, service_id int) (int64, error) {
	result := db.Sql.Where("id = ?", service_id).Updates(&payload)
	return result.RowsAffected, result.Error
}

func (db *ServicesStorage) ListItemRepo(ctx context.Context) ([]model.Service, error) {
	data := []model.Service{}
	result := db.Sql.Find(&data)
	return data, result.Error
}
