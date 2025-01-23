package repository

import (
	"context"
	"errors"
	"manage-system-server/domain"
	"manage-system-server/model"

	"gorm.io/gorm"
)

type customerStorage struct {
	Sql *gorm.DB
}

func NewCustomerStorage(db *gorm.DB) domain.CustomerStorage {
	return &customerStorage{
		Sql: db,
	}
}

func (db *customerStorage) CreateCustomerRepo(ctx context.Context, payload model.Customer) error {
	result := db.Sql.Create(&payload)
	return result.Error

}

func (db *customerStorage) ListCustomerRepo(ctx context.Context) ([]model.Customer, error) {
	data := []model.Customer{}
	result := db.Sql.Find(&data)
	return data, result.Error
}

func (db *customerStorage) FindCustomerByPhoneRepo(ctx context.Context, phone string) (model.Customer, error) {
	var customer model.Customer

	// Tìm kiếm khách hàng bằng số điện thoại
	result := db.Sql.Where("phone = ?", phone).First(&customer)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Không tìm thấy khách hàng
			return customer, nil // Trả về customer rỗng và không có lỗi
		}
		// Lỗi khác từ GORM hoặc database
		return model.Customer{}, result.Error
	}

	// Trả về khách hàng nếu tìm thấy
	return customer, nil
}
