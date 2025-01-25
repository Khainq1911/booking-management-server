package model

import "time"

type Service struct {
	Id        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Quantity  int       `json:"quantity" gorm:"column:quantity"`
	Price     float32   `json:"price" gorm:"column:price"`
	ImgUrl    string    `json:"img_url" gorm:"column:img_url"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

type UpdateService struct {
	Quantity  int       `json:"quantity" gorm:"column:quantity"`
	Price     int       `json:"price" gorm:"column:price"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (UpdateService) TableName() string { return "services" }
