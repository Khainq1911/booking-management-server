package repository

import (
	"context"
	"fmt"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type typeroomStorage struct {
	Sql *gorm.DB
}

func NewTyperoomStorage(db *gorm.DB) domain.TyperoomStorage {
	return &typeroomStorage{
		Sql: db,
	}
}
func (db *typeroomStorage) ListTyperoomRepo(ctx context.Context) ([]model.Typeroom, error) {
	data := []model.Typeroom{}
	result := db.Sql.Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no records found")
	}

	return data, nil
}

func (db *typeroomStorage) AddTyperoomRepo(ctx context.Context, createData model.CreateTyperoom) error {
	if err := db.Sql.Create(&createData).Error; err != nil {
		return err
	}
	return nil
}

func (db *typeroomStorage) UpdateTyperoomRepo(ctx context.Context, id int, updateData model.CreateTyperoom) (int64, error) {
	result := db.Sql.Where("id = ?", id).Updates(&updateData)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
