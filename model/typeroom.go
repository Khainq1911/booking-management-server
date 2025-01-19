package model

import "time"

type Typeroom struct {
	Id               int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name             string    `json:"name" gorm:"column:name"`
	Description      string    `json:"description" gorm:"column:description"`
	HourlyPrice      float32   `json:"hourly_price" gorm:"column:hourly_price"`
	DailyPrice       float32   `json:"daily_price" gorm:"column:daily_price"`
	MaxCapacity      int       `json:"max_capacity" gorm:"column:max_capacity"`
	StandardCapacity int       `json:"standard_capacity" gorm:"column:standard_capacity"`
	ImgUrl           string    `json:"img_url" gorm:"column:img_url"`
	CreatedAt        time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

type CreateTyperoom struct {
	Name             string    `json:"name" gorm:"column:name"`
	Description      string    `json:"description" gorm:"column:description"`
	HourlyPrice      float32   `json:"hourly_price" gorm:"column:hourly_price"`
	DailyPrice       float32   `json:"daily_price" gorm:"column:daily_price"`
	MaxCapacity      int       `json:"max_capacity" gorm:"column:max_capacity"`
	StandardCapacity int       `json:"standard_capacity" gorm:"column:standard_capacity"`
	ImgUrl           string    `json:"img_url" gorm:"column:img_url"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (CreateTyperoom) TableName() string { return "typerooms" }
