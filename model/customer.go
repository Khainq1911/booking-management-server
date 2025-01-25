package model

import "time"

type Customer struct {
	Id        int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"column:name"`
	Dob       time.Time `json:"dob" gorm:"column:dob"`
	Email     string    `json:"email" gorm:"column:email"`
	Phone     string    `json:"phone" gorm:"column:phone"`
	Address   string    `json:"address" gorm:"column:address"`
	ImgUrl    string    `json:"img_url" gorm:"column:img_url"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

